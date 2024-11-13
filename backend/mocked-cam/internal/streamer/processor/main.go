package processor

import (
	"context"
	"fmt"
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/internal/rabbitmq"
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/internal/s3"
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/internal/streamer/config"
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/pkg/randq"
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/pkg/waiter"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"
	"time"
)

type Processor struct {
	opts    config.StreamOptions
	storage *s3.CarStorage
	logger  *logan.Entry

	publisher *rabbitmq.Publisher
	queueName string

	enterq *randq.RandomizedQueue[string]
	exitq  *randq.RandomizedQueue[string]
}

func New(opts config.StreamOptions, storage *s3.CarStorage, publisher *rabbitmq.Publisher) *Processor {
	return &Processor{
		opts:      opts,
		storage:   storage,
		publisher: publisher,
	}
}

func (e *Processor) WithQueues(enterq, exitq *randq.RandomizedQueue[string]) *Processor {
	e.enterq = enterq
	e.exitq = exitq

	return e
}

func (e *Processor) WithRabbitQueue(rabbitQ string) *Processor {
	e.queueName = rabbitQ

	return e
}

func (e *Processor) WithLogger(logger *logan.Entry) *Processor {
	e.logger = logger

	return e
}

func (e *Processor) Stream(ctx context.Context) {
	e.logger.Info(fmt.Sprintf("sleeping for %s", e.opts.InitDelay))
	<-time.After(e.opts.InitDelay)

	e.logger.Info("starting to stream cars")

	for {
		select {
		case <-ctx.Done():
			e.logger.Info("streaming finished")
			return
		case <-waiter.WaitRandom(e.opts.MinDelay, e.opts.MaxDelay):
			car, err := e.enterq.Dequeue()
			if err != nil {
				if errors.Is(err, randq.ErrEmptyQueue) {
					e.logger.Debug("no cars to process")
					continue
				}
				e.logger.Error(errors.Wrap(err, "failed to dequeue car"))
				continue
			}

			exits, err := e.storage.Exists(car)
			if err != nil {
				e.logger.Error(errors.Wrap(err, "failed to check car existence"))
				continue
			}

			if !exits {
				e.logger.Debug(fmt.Sprintf("car %s does not exist in storage, omitting", car))
				continue
			}

			if err = e.publisher.PublishCar(e.queueName, car); err != nil {
				e.logger.Debug(errors.Wrap(err, "failed to publish car"))
				continue
			}

			e.logger.Info(fmt.Sprintf("car %s processed", car))
			e.exitq.Enqueue(car)
		}
	}
}

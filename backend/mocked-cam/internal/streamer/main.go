package streamer

import (
	"context"
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/internal/rabbitmq"
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/internal/s3"
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/internal/streamer/config"
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/internal/streamer/initializer"
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/internal/streamer/processor"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"
	"sync"
)

type Streamer struct {
	logger    *logan.Entry
	storage   *s3.CarStorage
	publisher *rabbitmq.Publisher
	cfg       config.StreamerConfig
}

func New(cfg config.StreamerConfig, logger *logan.Entry, storage *s3.CarStorage, publisher *rabbitmq.Publisher) *Streamer {
	return &Streamer{
		logger:    logger,
		storage:   storage,
		publisher: publisher,
		cfg:       cfg,
	}
}

func (s *Streamer) Stream(ctx context.Context) {
	s.logger.Info("initializing queues...")

	cars, err := s.storage.ListCars()
	if err != nil {
		panic(errors.Wrap(err, "failed to get available cars to stream"))
	}

	enterq, exitq, err := initializer.InitializeQueues(s.cfg.DumpFile, cars)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize queues"))
	}

	s.logger.Info("starting streaming...")
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		processor.
			New(s.cfg.EnteredOptions, s.storage, s.publisher).
			WithQueues(enterq, exitq).
			WithRabbitQueue(s.publisher.EnterQueueName()).
			WithLogger(s.logger.WithField("component", "enterer")).
			Stream(ctx)
	}()
	go func() {
		defer wg.Done()
		processor.
			New(s.cfg.ExitedOptions, s.storage, s.publisher).
			WithQueues(exitq, enterq).
			WithRabbitQueue(s.publisher.ExitQueueName()).
			WithLogger(s.logger.WithField("component", "exiter")).
			Stream(ctx)
	}()

	wg.Wait()
	if err = initializer.DumpExitQueue(s.cfg.DumpFile, exitq); err != nil {
		panic(errors.Wrap(err, "failed to dump exit queue"))
	}

	s.logger.Info("streaming finished")
}

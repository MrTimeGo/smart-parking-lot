package streamer

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
)

type Streamer struct {
	logger   *logan.Entry
	dumpFile string
}

func New() *Streamer {
	return &Streamer{}
}

func (s *Streamer) Stream(ctx context.Context) {
	s.logger.Info("initializing queues...")
	defer s.logger.Info("streaming stopped")
}

package cmd

import (
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/internal/config"
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/internal/rabbitmq"
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/internal/s3"
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/internal/streamer"
	"github.com/spf13/cobra"
	"gitlab.com/distributed_lab/kit/kv"
	"os/signal"
	"syscall"
)

var (
	runServiceCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the mocked camera service",
		Run: func(cmd *cobra.Command, args []string) {
			var (
				cfg         = config.New(kv.MustFromEnv())
				storage     = s3.New(cfg.S3Config())
				publisher   = rabbitmq.New(cfg.AmqpConfig())
				streamerCfg = cfg.StreamerConfig()
				logger      = cfg.Log().WithField("service", "streamer")
			)

			ctx, cancel := signal.NotifyContext(cmd.Context(), syscall.SIGTERM, syscall.SIGINT)
			defer cancel()

			streamer.New(streamerCfg, logger, storage, publisher).Stream(ctx)
		},
	}
)

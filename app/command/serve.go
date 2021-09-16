package command

import (
	"os"

	servercomp "buttress.io/app/component/server"
	"buttress.io/app/config"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

// Serve
type Serve struct {
	Verbose bool `short:"v" default:"false" help:"Enable verbose mode."`
}

// Run Starts the app.
func (c *Serve) Run() error {
	// Init the config verbose value.
	config.Verbose = c.Verbose

	log := config.Log.Named("serve")

	// Print the process id of current running instance.
	log.Info("starting...", zap.Int("pid", os.Getpid()))

	// Run!
	fx.New(
		fx.WithLogger(func() fxevent.Logger {
			if config.Verbose {
				return &fxevent.ZapLogger{
					Logger: config.Log.Named("fx"),
				}
			}

			return fxevent.NopLogger
		}),
		servercomp.RpcComp,
	).Run()

	return nil
}

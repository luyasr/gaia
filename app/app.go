package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/ioc"
	"github.com/luyasr/gaia/log"
	"github.com/luyasr/gaia/log/zerolog"
	"github.com/luyasr/gaia/transport"
	"golang.org/x/sync/errgroup"
)

type App struct {
	servers         []transport.Server
	ctx             context.Context
	shutdownTimeout time.Duration
	sigs            []os.Signal
	log             *log.Helper
}

// Option is the option of App
type Option func(*App)

// Server adds transport.Server to App
func Server(s ...transport.Server) Option {
	return func(a *App) {
		a.servers = append(a.servers, s...)
	}
}

// ShutdownTimeout adds shutdown timeout to App
func ShutdownTimeout(shutdownTimeout time.Duration) Option {
	return func(a *App) {
		a.shutdownTimeout = shutdownTimeout
	}
}

// Signal adds os.Signal to App
func Signal(sigs []os.Signal) Option {
	return func(a *App) {
		a.sigs = sigs
	}
}

func New(opt ...Option) *App {
	app := &App{
		ctx:             context.Background(),
		shutdownTimeout: 10 * time.Second,
		sigs:            []os.Signal{os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT},
		log:             log.NewHelper(zerolog.DefaultLogger),
	}

	for _, o := range opt {
		o(app)
	}

	return app
}

func (a *App) Run() error {
	eg, ctx := errgroup.WithContext(a.ctx)
	wg := sync.WaitGroup{}

	// run servers
	for _, svr := range a.servers {
		svr := svr
		eg.Go(func() error {
			<-ctx.Done() // wait for signal
			shutdownCtx, cancel := context.WithTimeout(ctx, a.shutdownTimeout)
			defer cancel()
			return svr.Shutdown(shutdownCtx)
		})

		wg.Add(1)
		eg.Go(func() error {
			wg.Done()
			return svr.Run()
		})
	}

	wg.Wait()

	// handle signals
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, a.sigs...)
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return nil
		case s := <-ch:
			a.log.Infof("receive signal %s, shutdown server", s)
			return context.Canceled
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	// close ioc container
	if err := ioc.Container.Close(); err != nil {
		return err
	}	

	return nil
}

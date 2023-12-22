package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/log"
	"github.com/luyasr/gaia/log/zerolog"
	"github.com/luyasr/gaia/transport"
	"golang.org/x/sync/errgroup"
)

type App struct {
	servers         []transport.Server
	shutdownTimeout time.Duration
	sigs            []os.Signal
	ctx             context.Context
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
		shutdownTimeout: 10 * time.Second,
		sigs:            []os.Signal{os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT},
		ctx:             context.Background(),
		log:             log.NewHelper(zerolog.New(zerolog.DefaultLogger)),
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
		wg.Add(1)
		svr := svr
		eg.Go(func() error {
			defer wg.Done()
			return svr.Run()
		})
	}

	// handle signals
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, a.sigs...)
	eg.Go(func() error {
		s := <-ch
		a.log.Infof("receive signal %s, shutdown server", s)
		return a.Shutdown(ctx)
	})

	wg.Wait()

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	eg, c := errgroup.WithContext(ctx)

	for _, svr := range a.servers {
		svr := svr
		eg.Go(func() error {
			shutdownCtx, cancel := context.WithTimeout(c, a.shutdownTimeout)
			defer cancel()

			return svr.Shutdown(shutdownCtx)
		})
	}

	return eg.Wait()
}

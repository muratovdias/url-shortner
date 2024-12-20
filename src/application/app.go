package application

import (
	"context"
	"fmt"
	"github.com/muratovdias/url-shortner/src/config"
	"github.com/muratovdias/url-shortner/src/databases"
	"github.com/muratovdias/url-shortner/src/databases/drivers"
	http2 "github.com/muratovdias/url-shortner/src/server/http"
	"github.com/muratovdias/url-shortner/src/service"
	"github.com/muratovdias/url-shortner/src/service/shortner"
	"log"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type Application interface {
	Run()
	Exit()
}

type application struct {
	log    *slog.Logger
	cfg    *config.Configuration
	ds     drivers.DataStore
	srv    *http.Server
	ctx    context.Context
	cancel context.CancelFunc
}

func (a *application) Run() {
	errChan := make(chan error, 1)

	go func() {
		a.log.Info("starting HTTP server", "addr", a.srv.Addr)

		if err := a.srv.ListenAndServe(); err != nil {
			errChan <- fmt.Errorf("HTTP server error: %w", err)
		}
		close(errChan)
	}()

	select {
	case err := <-errChan:
		if err != nil {
			log.Fatal(err)
		}
	case <-a.ctx.Done():
		a.log.Info("shutting down application")
	}

	a.Exit()
}

func (a *application) Exit() {
	defer a.cancel()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	a.log.Info("shutting down HTTP server")
	if err := a.srv.Shutdown(shutdownCtx); err != nil {
		a.log.Error("failed to shutdown HTTP server", "error", err)
	}

	a.log.Info("closing datastore connections")
	if err := a.ds.Close(shutdownCtx); err != nil {
		a.log.Error("failed to close datastore", "error", err)
	}

	a.log.Info("application shutdown completed")
}

func Init() (Application, error) {
	var err error
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	// init config
	cfg, err := config.FromEnvs()
	if err != nil {
		return nil, err
	}

	app := &application{
		ctx:    ctx,
		cancel: cancel,
		cfg:    cfg,
		log:    setupLogger(cfg.Env),
	}

	// init database
	app.ds, err = databases.New(cfg.DataStore)
	if err != nil {
		return nil, err
	}

	if err = app.ds.Connect(); err != nil {
		return nil, err
	}

	// init urlShortener
	urlShortener := shortner.NewUrlShortener(app.ds.UrlShortenerRepo(), app.log)

	// init service
	serv := service.NewService(urlShortener)

	//init router
	router := http2.NewRouterImpl(serv, app.ds)

	// init http server
	app.srv = &http.Server{
		Addr:         cfg.Address,
		Handler:      router.Routes(),
		ReadTimeout:  cfg.HttpServer.Timeout,
		WriteTimeout: cfg.HttpServer.Timeout,
	}

	return app, nil
}

package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/halyph/go-service-blueprint/pkg/handler"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	gr, ctx := errgroup.WithContext(ctx)
	gr.Go(func() error {
		for {
			log.Println("I am alive")

			select {
			case <-time.After(time.Minute):
				continue
			case <-ctx.Done():
				return nil
			}
		}
	})

	gr.Go(func() error {
		userHandler := handler.NewUserHandler()
		mux := http.NewServeMux()
		mux.HandleFunc("/users", userHandler.ListUsers)
		mux.HandleFunc("/users/", userHandler.GetUser)
		mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte("ok"))
		})

		server := &http.Server{
			Addr:              ":8080",
			ReadHeaderTimeout: 10 * time.Second,
			Handler:           mux,
		}

		go func() {
			<-ctx.Done()
			_ = server.Shutdown(context.Background())
		}()

		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	})

	return gr.Wait()
}

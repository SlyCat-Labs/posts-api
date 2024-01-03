package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/jdashel/posts-api/internal/infra/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure cancellation, even for early errors

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		<-ch
		cancel()
		fmt.Println("\nShutting down gracefully...")
	}()

	err := http.StartServer(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	<-ctx.Done() // Wait for cancellation
	fmt.Println("Server stopped")
}

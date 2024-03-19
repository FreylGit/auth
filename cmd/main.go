package main

import (
	"context"
	"github.com/FreylGit/auth/internal/app"
	"log"
)

func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed settings:%w", err.Error())
	}
	err = a.Run()
	if err != nil {
		panic("Failed run")
	}
}

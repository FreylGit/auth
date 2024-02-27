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
		log.Fatal("failed settings")
	}
	err = a.Run()
	if err != nil {
		panic("Failed run")
	}
}

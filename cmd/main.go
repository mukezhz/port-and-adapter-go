package main

import (
	"myapp/internal/framework"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		framework.Module,
		fx.Invoke(framework.Start),
	).Run()
}

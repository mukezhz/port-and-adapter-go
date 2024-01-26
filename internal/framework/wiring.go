package framework

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func Start(lifecycle fx.Lifecycle, router *gin.Engine) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go router.Run()
			return nil
		},
	})
}

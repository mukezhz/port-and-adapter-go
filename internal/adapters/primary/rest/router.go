package rest

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Handler interface {
	RegisterRoutes(router *gin.Engine)
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Handler)),
		fx.ResultTags(`group:"handlers"`),
	)
}

func NewRouter(handlers []Handler) *gin.Engine {
	router := gin.Default()
	for _, h := range handlers {
		h.RegisterRoutes(router)
	}
	return router
}

var Module = fx.Options(
	fx.Provide(
		AsRoute(NewProcessHandler),
		fx.Annotate(
			NewRouter,
			fx.ParamTags(`group:"handlers"`),
		),
	),
)

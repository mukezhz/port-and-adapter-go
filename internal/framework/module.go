package framework

import (
	"myapp/internal/adapters/primary/rest"

	"go.uber.org/fx"
)

var Module = fx.Options(
	rest.Module,
	ProcessModule,
)

package framework

import (
	"myapp/internal/adapters/secondary/persistent"
	"myapp/internal/domain"
	"myapp/internal/domain/interactors"
	incomingports "myapp/internal/ports/incoming_ports"
	outgoingports "myapp/internal/ports/outgoing_ports"

	"go.uber.org/fx"
)

var ProcessModule = fx.Options(
	fx.Provide(
		fx.Annotate(
			persistent.NewMockRepository,
			fx.As(new(outgoingports.ProcessCreator)),
			fx.As(new(outgoingports.ProcessFinder)),
		),

		fx.Annotate(
			interactors.NewProcessInteractor,
			fx.As(new(incomingports.ProcessUseCase)),
		),

		domain.NewProcessService,
	),
)

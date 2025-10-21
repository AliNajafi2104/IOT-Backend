package http

import "go.uber.org/fx"

var Module = fx.Module("http",
	fx.Provide(
		NewHandler,
	),
	fx.Invoke(RegisterHandlers),
	fx.Invoke(NewHTTPServer),
)

package main

import (
	"net/http"

	"github.com/corioders/gokit/web"
	"github.com/corioders/gokit/web/middleware"
	"github.com/corioders/qr-games.server/app/routes/hello"
	"github.com/corioders/qr-games.server/foundation"
)

func newRouter(app *foundation.Application) (http.Handler, error) {
	routerLogger := app.GetLogger().Child("Router")
	_ = routerLogger

	router := web.NewRouter(app.GetLogger(),
		middleware.Errors(app.GetLogger()),
		middleware.Compression(),
		middleware.Cors("*"),
	)

	helloController, err := hello.NewController()
	if err != nil {
		return nil, err
	}

	router.Handle(http.MethodGet, "/", helloController.Hello)

	return router, nil
}

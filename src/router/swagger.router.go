package router

import (
	"template-app/cmd/log"
	_ "template-app/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Swagger struct {
	Log    *log.Logger
	Router *mux.Router
}

func (i *Swagger) Build() {
	i.Log.Router("Build swagger routes starting...")

	i.Router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(httpSwagger.URL("swagger/doc.json")))
	i.Log.Router("Build swagger routes finished")
}

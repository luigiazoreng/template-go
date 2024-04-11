package router

import (
	"template-app/cmd/log"
	"template-app/src/handler"
	"template-app/src/middleware"
	types "template-app/src/structs"

	"github.com/gorilla/mux"
)

type Product struct {
	Log     *log.Logger
	Router  *mux.Router
	Handler *handler.Product
}

func (i *Product) Build() {
	i.Log.Router("Build ticket routes starting...")
	i.POST()
	i.PATCH()
	i.GET()
	i.DELETE()
	i.Log.Router("Build ticket routes finished")
}

func (i *Product) POST() {
	i.Log.Router("Building POST method handler...")
	postRouter := i.Router.Methods("POST").Subrouter()
	postRouter.Use(middleware.ValidateBody(&types.Product_POST_Body{}))
	postRouter.HandleFunc("/{rest:.*}", i.Handler.POST)
}
func (i *Product) PATCH() {
	i.Log.Router("Building PATCH method handler...")
	postRouter := i.Router.Methods("PATCH").Subrouter()
	postRouter.Use(middleware.ValidateBody(&types.Product_PATCH_Body{}))
	postRouter.HandleFunc("/{rest:.*}", i.Handler.PATCH)
}

func (i *Product) GET() {
	i.Log.Router("Building GET method handler...")
	getRouter := i.Router.Methods("GET").Subrouter()
	getRouter.Use(middleware.ValidateQuery(&types.Product_GET_Query{}))
	getRouter.HandleFunc("/{rest:.*}", i.Handler.GET)
}

func (i *Product) DELETE() {
	i.Log.Router("Building DELETE method handler...")
	deleteRouter := i.Router.Methods("DELETE").Subrouter()
	deleteRouter.Use(middleware.ValidateQuery(&types.Product_DELETE_Query{}))
	deleteRouter.Use(middleware.ValidateBody(&types.Product_DELETE_Body{}))
	deleteRouter.HandleFunc("/{rest:.*}", i.Handler.DELETE)
}

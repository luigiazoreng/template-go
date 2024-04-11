package handler

import (
	"net/http"
	"template-app/cmd/log"
)

type Product struct {
	Log *log.Logger
}

// @Summary Post Product
// @Description this endpoint is used for...
// @Tags Product
// @Produce json
// @Param product body types.Product_POST_Body true "Provide product details in the request body"
// @Success 200 {object} types.Response "Ok"
// @Failure 403 "Unauthorized, check HTTP Authorization header"
// @Router /product [post]
func (r *Product) POST(res http.ResponseWriter, req *http.Request) {
	// Log the ticket post method
	r.Log.Product("Post method")
}

// @Summary Patch Product
// @Description this endpoint is used for...
// @Tags Product
// @Produce json
// @Param product body types.Product_PATCH_Body true "Provide product details in the request body"
// @Success 200 {object} types.Response "Ok"
// @Failure 403 "Unauthorized, check HTTP Authorization header"
// @Router /product [patch]
func (r *Product) PATCH(res http.ResponseWriter, req *http.Request) {
	r.Log.Product("Patch method")

}

// @Summary Get Product
// @Description This endpoint is used to...
// @Tags Product
// @Produce json
// @Success 200 {object} types.Response "Ok"
// @Failure 403 "Unauthorized, check HTTP Authorization header"
// @Router /product [get]
func (r *Product) GET(res http.ResponseWriter, req *http.Request) {
	r.Log.Product("Get method")
}

// @Summary Delete Product
// @Description This endpoint is used to delete product
// @Tags Product
// @Produce json
// @Success 200 {object} types.Response "Ok"
// @Failure 403 "Unauthorized, check HTTP Authorization header"
// @Router /product [delete]
func (r *Product) DELETE(res http.ResponseWriter, req *http.Request) {
	r.Log.Product("Delete method")
}

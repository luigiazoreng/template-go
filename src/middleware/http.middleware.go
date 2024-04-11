package middleware

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"reflect"
	"template-app/src/utils"
)

/*
The function first retrieves the HTTP method from the request.
It then checks if there is a handler for this method in the handlers map.
If there is, it calls this handler with the response writer and the request as parameters.
If there isn't, it logs an error message and sends an
HTTP error response with the status code 405 (Method Not Allowed).
*/
func FowardRequest(req *http.Request, res http.ResponseWriter, handlers map[string]http.HandlerFunc, log *log.Logger) {
	method := req.Method
	if handler, ok := handlers[method]; ok {
		handler(res, req)
	} else {
		log.Println("Unsupported HTTP method:", method)
		http.Error(res, "Unsupported HTTP method", http.StatusMethodNotAllowed)
	}
}

func ValidateBody(structInstance interface{}) func(http.Handler) http.Handler {
	structType := reflect.TypeOf(structInstance)

	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			// Create a new instance of the struct type for each request
			newStructInstance := reflect.New(structType).Interface()
			bodyBytes, err := io.ReadAll(req.Body)
			if err != nil {
				http.Error(res, "Failed to read the request body", http.StatusInternalServerError)
				return
			}
			// Create a new reader from the byte slice and assign it back to req.Body
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			if err := utils.DecodeIO(bytes.NewBuffer(bodyBytes), newStructInstance); err != nil {
				http.Error(res, "Invalid request body", http.StatusBadRequest)
				return
			}

			if err := utils.Validate(newStructInstance); err != nil {
				http.Error(res, "Validation error: "+err.Error(), http.StatusBadRequest)
				return
			}

			ctx := context.WithValue(req.Context(), structType, newStructInstance)
			reqContext := req.WithContext(ctx)
			next.ServeHTTP(res, reqContext)
		})
	}
}

func ValidateQuery(structInstance interface{}) func(http.Handler) http.Handler {
	structType := reflect.TypeOf(structInstance)

	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			// Create a new instance of the struct type for each request
			newStructInstance := reflect.New(structType).Interface()

			// Decode query parameters into the new struct instance
			queryValues := req.URL.Query()
			if err := utils.DecodeQueryParams(queryValues, newStructInstance); err != nil {
				http.Error(res, "Invalid query parameters", http.StatusBadRequest)
				return
			}

			if validator, ok := newStructInstance.(interface{ Validate() error }); ok {
				if err := validator.Validate(); err != nil {
					http.Error(res, "Validation error: "+err.Error(), http.StatusBadRequest)
					return
				}
			}

			ctx := context.WithValue(req.Context(), structType, newStructInstance)
			reqContext := req.WithContext(ctx)
			next.ServeHTTP(res, reqContext)
		})
	}
}

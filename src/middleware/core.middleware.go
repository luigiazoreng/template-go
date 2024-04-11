package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"template-app/cmd/log"
	"time"

	"gopkg.in/yaml.v2"
)

// LoggerMiddleware is a middleware that logs information about the incoming request.
func LogRequest(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			start := time.Now()
			defer func() {
				logger.Middleware("Request: %s %s | Duration: %s", req.Method, req.URL.Path, time.Since(start))
			}()
			next.ServeHTTP(res, req)
		})
	}
}

/*
Is a middleware function for HTTP handlers in Go.
It takes two parameters: next which is the next HTTP handler function to be executed, and log which is a logger instance.
The function returns another function that matches the signature of an HTTP handler function.
This returned function is responsible for performing authentication logic before the next HTTP handler is executed.
*/
func AuthMiddleware(next http.HandlerFunc, log *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Perform authentication logic here
		token := r.Header.Get("Authorization")
		expectedToken1 := os.Getenv("")
		expectedToken2 := os.Getenv("")
		expectedToken3 := os.Getenv("")
		expectedToken4 := os.Getenv("")

		if token == "" {
			log.Middleware("Authentication failed: Authorization token is missing")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		} else if token != expectedToken1 && token != expectedToken2 && token != expectedToken3 && token != expectedToken4 {
			log.Middleware("Autenticação falhou: Token inválido")
			http.Error(w, "Não autorizado", http.StatusUnauthorized)
			return
		}
		log.Middleware("Authentication passed")
		next(w, r)
	}
}

// RecoverMiddleware recovers from panics and logs the error.
func RecoverMiddleware(next http.HandlerFunc, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Middleware("Panic recovered: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next(w, r)
	}
}

/*This functionis used to read a YAML file and unmarshal its content into a given target structure.*/
func ReadYAMLFile(filePath string, targetStruct interface{}) error {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading YAML file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, targetStruct)
	if err != nil {
		return fmt.Errorf("error unmarshalling YAML: %v", err)
	}

	return nil
}

/*Is used to send a HTTP response with a specific status code and response body.*/
func Response(res http.ResponseWriter, statusCode int, response any) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	json.NewEncoder(res).Encode(response)
}

func ExtractResponseBody(r *http.Response, l *log.Logger) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		l.Println(err)
		return nil, err
	}

	return body, nil
}

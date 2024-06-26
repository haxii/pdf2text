package rest_api

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/haxii/pdf2text/rest_api/service"
)

func configureFlags(api *service.Pdf2textAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *service.Pdf2textAPI) http.Handler {
	api.ServeError = errors.ServeError
	api.UseSwaggerUI()
	api.BinConsumer = runtime.ByteStreamConsumer()
	api.TxtProducer = runtime.TextProducer()

	api.PDF2TextHandler = service.PDF2TextHandlerFunc(
		func(params service.PDF2TextParams) middleware.Responder {
			text, err := pdf2Text(params)
			if err != nil {
				middleware.Error(http.StatusInternalServerError, err.Error())
			}
			return service.NewPDF2TextOK().WithPayload(string(text))
		},
	)
	api.PreServerShutdown = func() {}
	api.ServerShutdown = func() {}
	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

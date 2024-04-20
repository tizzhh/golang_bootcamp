// This file is safe to edit. Once it exists it will not be overwritten

package restapi

/*
#include "../cow.c"
*/
import "C"

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"unsafe"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"candyShop/restapi/operations"
)

const (
	COOL_ESKIMO         string = "CE"
	APRICOT_AARDVARK    string = "AA"
	NATURAL_TIGER       string = "NT"
	DAZZLING_ELDERBERRY string = "DE"
	YELLOW_RAMBUTAN     string = "YR"
)

var ShopPrices = map[string]int64{
	COOL_ESKIMO:         10,
	APRICOT_AARDVARK:    15,
	NATURAL_TIGER:       17,
	DAZZLING_ELDERBERRY: 21,
	YELLOW_RAMBUTAN:     23,
}

//go:generate swagger generate server --target ../../ex01 --name CandyServer --spec ../swagger_spec.yaml --principal interface{}

func configureFlags(api *operations.CandyServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CandyServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.BuyCandyHandler = operations.BuyCandyHandlerFunc(func(params operations.BuyCandyParams) middleware.Responder {
		if _, ok := ShopPrices[*params.Order.CandyType]; !ok || *params.Order.CandyCount < 1 {
			return operations.NewBuyCandyBadRequest().WithPayload(&operations.BuyCandyBadRequestBody{Error: "uknown candy type or count is < 1"})
		}

		reqMoney := ShopPrices[*params.Order.CandyType] * *params.Order.CandyCount
		if *params.Order.Money < reqMoney {
			notEnoughMoney := fmt.Sprintf("You need %d more money!", reqMoney-*params.Order.Money)
			return operations.NewBuyCandyPaymentRequired().WithPayload(&operations.BuyCandyPaymentRequiredBody{Error: notEnoughMoney})
		}

		cstr := C.CString("Thanks")
		defer C.free(unsafe.Pointer(cstr))
		cow := C.ask_cow(cstr)
		return operations.NewBuyCandyCreated().WithPayload(&operations.BuyCandyCreatedBody{Thanks: C.GoString(cow), Change: *params.Order.Money - reqMoney})
	})

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

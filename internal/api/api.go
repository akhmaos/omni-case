package api

import (
	"net/http"
	"strings"

	"github.com/akhmaos/omni-case/internal/api/restapi/restapi"
	"github.com/akhmaos/omni-case/internal/api/restapi/restapi/operations"
	"github.com/akhmaos/omni-case/internal/api/restapi/restapi/operations/process"
	"github.com/akhmaos/omni-case/internal/api/restapi/restapi/operations/standard"
	"github.com/akhmaos/omni-case/internal/usecase"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sebest/xff"
)

type service struct {
	client usecase.Client
}

// NewServer - return openapi server implementation
func NewServer(client usecase.Client, host string, port int) (*restapi.Server, error) {
	svc := &service{
		client: client,
	}

	serverHost := host
	serverPort := port

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, err
	}

	api := operations.NewOnmiAPI(swaggerSpec)

	api.StandardHealthCheckHandler = standard.HealthCheckHandlerFunc(healthCheck)
	api.ProcessPostProcessItemsHandler = process.PostProcessItemsHandlerFunc(svc.ProcessItems)

	server := restapi.NewServer(api)

	server.EnabledListeners = []string{"http"}
	server.Host = serverHost
	server.Port = serverPort

	// The middleware executes before anything.
	globalMiddlewares := func(handler http.Handler) http.Handler {
		xffmw, _ := xff.Default()
		swaggerUIOpts := middleware.SwaggerUIOpts{
			BasePath: "/",
			Path:     "/onmi/docs",
		}

		return xffmw.Handler(recovery(
			middleware.Spec(
				serverHost,
				restapi.FlatSwaggerJSON,
				middleware.SwaggerUI(swaggerUIOpts, handler)),
		))
	}

	server.SetHandler(globalMiddlewares(api.Serve(nil)))

	return server, nil
}

func splitCommaSeparatedStr(commaSeparated string) (result []string) {
	for _, item := range strings.Split(commaSeparated, ",") {
		item = strings.TrimSpace(item)
		if item != "" {
			result = append(result, item)
		}
	}
	return
}

func healthCheck(params standard.HealthCheckParams, profile interface{}) middleware.Responder {
	return standard.NewHealthCheckOK().WithPayload(&standard.HealthCheckOKBody{Ok: true})
}

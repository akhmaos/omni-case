package api

import (
	"context"

	"github.com/akhmaos/omni-case/internal/api/restapi/models"
	"github.com/akhmaos/omni-case/internal/api/restapi/restapi/operations/process"
	externalSrv "github.com/akhmaos/omni-case/internal/service"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

func (s *service) ProcessItems(params process.PostProcessItemsParams) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	items := convertAPIItemsToAppItems(params.Body.Items)

	err := s.client.ProcessItems(ctx, items)

	switch {
	default:
		return process.NewPostProcessItemsDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case err == nil:
		return process.NewPostProcessItemsOK().WithPayload(&process.PostProcessItemsOKBody{Ok: true})
	}

}

func convertAPIItemsToAppItems(items []*models.Item) []externalSrv.Item {
	appItems := make([]externalSrv.Item, len(items))

	for i := range items {
		appItems[i] = convertAPIItemToAppItem(*items[i])
	}

	return appItems
}

func convertAPIItemToAppItem(item models.Item) externalSrv.Item {
	return externalSrv.Item{
		Title: *item.Title,
		Key:   *item.Key,
	}
}

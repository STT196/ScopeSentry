package assets

import (
	iphandler "github.com/STT196/ScopeSentry/internal/api/handlers/assets/ip"
	"github.com/STT196/ScopeSentry/internal/api/routes/common"
	"github.com/STT196/ScopeSentry/internal/models"
)

func registerIPRoutes() models.RouteGroup {
	return models.RouteGroup{
		Path: "/ip",
		Routes: []models.Route{
			{
				Method:      "POST",
				Path:        "",
				Handler:     iphandler.GetIPAssets,
				Middlewares: common.WithAuth(),
			},
		},
	}
}

// assets-------------------------------------
// @file      : subdomain.go
// @author    : Autumn
// @contact   : rainy-autumn@outlook.com
// @time      : 2025/5/13 23:32
// -------------------------------------------

package assets

import (
	"github.com/STT196/ScopeSentry/internal/api/handlers/assets/subdomain"
	"github.com/STT196/ScopeSentry/internal/api/routes/common"
	"github.com/STT196/ScopeSentry/internal/models"
)

func registerSubDomainRoutes() models.RouteGroup {
	return models.RouteGroup{
		Path: "/subdomain",
		Routes: []models.Route{
			{
				Method:      "POST",
				Path:        "",
				Handler:     subdomain.GetSubdomains,
				Middlewares: common.WithAuth(),
			},
			{
				Method:      "POST",
				Path:        "/taker",
				Handler:     subdomain.GetSubdomainTakerData,
				Middlewares: common.WithAuth(),
			},
		},
	}
}

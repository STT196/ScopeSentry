// assets-------------------------------------
// @file      : url.go
// @author    : Autumn
// @contact   : rainy-autumn@outlook.com
// @time      : 2025/5/14 23:11
// -------------------------------------------

package assets

import (
	"github.com/STT196/ScopeSentry/internal/api/handlers/assets/url"
	"github.com/STT196/ScopeSentry/internal/api/routes/common"
	"github.com/STT196/ScopeSentry/internal/models"
)

func registerUrlRoutes() models.RouteGroup {
	return models.RouteGroup{
		Path: "/url",
		Routes: []models.Route{
			{
				Method:      "POST",
				Path:        "",
				Handler:     url.GetURLs,
				Middlewares: common.WithAuth(),
			},
		},
	}
}

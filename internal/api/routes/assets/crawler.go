// assets-------------------------------------
// @file      : crawler.go
// @author    : Autumn
// @contact   : rainy-autumn@outlook.com
// @time      : 2025/5/14 23:33
// -------------------------------------------

package assets

import (
	"github.com/STT196/ScopeSentry/internal/api/handlers/assets/crawler"
	"github.com/STT196/ScopeSentry/internal/api/routes/common"
	"github.com/STT196/ScopeSentry/internal/models"
)

func registerCrawlerRoutes() models.RouteGroup {
	return models.RouteGroup{
		Path: "/crawler",
		Routes: []models.Route{
			{
				Method:      "POST",
				Path:        "",
				Handler:     crawler.GetCrawlers,
				Middlewares: common.WithAuth(),
			},
		},
	}
}

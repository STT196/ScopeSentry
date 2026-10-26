// assets-------------------------------------
// @file      : dirscan.go
// @author    : Autumn
// @contact   : rainy-autumn@outlook.com
// @time      : 2025/5/17 15:20
// -------------------------------------------

package assets

import (
	"github.com/STT196/ScopeSentry/internal/api/handlers/assets/dirscan"
	"github.com/STT196/ScopeSentry/internal/api/routes/common"
	"github.com/STT196/ScopeSentry/internal/models"
)

func registerDirscanRoutes() models.RouteGroup {
	return models.RouteGroup{
		Path: "/dirscan",
		Routes: []models.Route{
			{
				Method:      "POST",
				Path:        "",
				Handler:     dirscan.List,
				Middlewares: common.WithAuth(),
			},
		},
	}
}

// assets-------------------------------------
// @file      : mp.go
// @author    : Autumn
// @contact   : rainy-autumn@outlook.com
// @time      : 2025/5/14 22:56
// -------------------------------------------

package assets

import (
	"github.com/STT196/ScopeSentry/internal/api/handlers/assets/mp"
	"github.com/STT196/ScopeSentry/internal/api/routes/common"
	"github.com/STT196/ScopeSentry/internal/models"
)

func registerMpRoutes() models.RouteGroup {
	return models.RouteGroup{
		Path: "/mp",
		Routes: []models.Route{
			{
				Method:      "POST",
				Path:        "",
				Handler:     mp.GetMPData,
				Middlewares: common.WithAuth(),
			},
		},
	}
}

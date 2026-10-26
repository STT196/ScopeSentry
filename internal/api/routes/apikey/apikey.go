package apikey

import (
	"github.com/STT196/ScopeSentry/internal/api/handlers/apikey"
	"github.com/STT196/ScopeSentry/internal/api/routes/common"
	"github.com/STT196/ScopeSentry/internal/models"
	"github.com/gin-gonic/gin"
)

func RegisterApiKeyRoutes(api *gin.RouterGroup) {
	routes := models.RouteGroup{
		Path: "/apikey",
		Routes: []models.Route{
			{
				Method:      "GET",
				Path:        "/list",
				Handler:     apikey.List,
				Middlewares: common.WithAuth(),
			},
			{
				Method:      "POST",
				Path:        "/create",
				Handler:     apikey.Create,
				Middlewares: common.WithAuth(),
			},
			{
				Method:      "POST",
				Path:        "/delete",
				Handler:     apikey.Delete,
				Middlewares: common.WithAuth(),
			},
		},
	}
	common.RegisterRouteGroup(api, routes)
}

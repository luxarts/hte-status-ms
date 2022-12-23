package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luxarts/jsend-go"
	"hte-status-ms/internal/controller"
	"hte-status-ms/internal/defines"
	"hte-status-ms/internal/repository"
	"hte-status-ms/internal/service"
	"net/http"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {
	// DB connectors, rest clients, and other stuff init

	// Repositories init
	repo := repository.NewStatusRepository()

	// Services init
	svc := service.NewStatusService(repo)

	// Controllers init
	ctrl := controller.NewStatusController(svc)

	// Endpoints
	r.PUT(defines.EndpointUpdateStatus, ctrl.Update)

	// Health check endpoint
	r.GET(defines.EndpointPing, healthCheck)
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, jsend.NewSuccess("pong"))
}

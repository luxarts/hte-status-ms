package router

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/luxarts/jsend-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"hte-status-ms/internal/controller"
	"hte-status-ms/internal/defines"
	"hte-status-ms/internal/repository"
	"hte-status-ms/internal/service"
	"log"
	"net/http"
	"os"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {
	// DB connectors, rest clients, and other stuff init
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv(defines.EnvMongoDBURI)))
	if err != nil {
		log.Fatalln(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln(err)
	}
	// Repositories init
	repo := repository.NewStatusRepository(client)

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

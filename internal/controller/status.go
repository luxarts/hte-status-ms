package controller

import (
	"github.com/gin-gonic/gin"
	"hte-status-ms/internal/domain"
	"hte-status-ms/internal/service"
	"net/http"
)

type StatusController interface {
	Update(ctx *gin.Context)
}
type statusController struct {
	svc service.StatusService
}

func NewStatusController(svc service.StatusService) StatusController {
	return &statusController{svc: svc}
}

func (c *statusController) Update(ctx *gin.Context) {
	var p domain.Status
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid payload"})
		return
	}
	if !p.IsValid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid payload"})
		return
	}
	s, err := c.svc.Update(&p)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, s)
}

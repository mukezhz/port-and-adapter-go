package rest

import (
	"myapp/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProcessHandler struct {
	service *domain.ProcessService
}

func NewProcessHandler(service *domain.ProcessService) *ProcessHandler {
	h := ProcessHandler{service: service}
	return &h
}

func (h *ProcessHandler) HandleRequest(c *gin.Context) {
	data := c.Query("data")
	result := h.service.Find(data)
	h.service.Create(data)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func (h *ProcessHandler) RegisterRoutes(router *gin.Engine) {
	userGroup := router.Group("/process")
	userGroup.GET("/", h.HandleRequest)
}

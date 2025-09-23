package handler

import (
	"net/http"

	"github.com/ZakSlinin/licey-maps-backend/point-service/internal/model"
	"github.com/ZakSlinin/licey-maps-backend/point-service/internal/service"
	"github.com/gin-gonic/gin"
)

type PointHandler struct {
	service *service.PointService
}

func NewPointHandler(service *service.PointService) *PointHandler {
	return &PointHandler{service: service}
}

func (h *PointHandler) CreatePoint(c *gin.Context) {
	var point model.Point
	if err := c.Bind(&point); err != nil {
		return
	}

	createdPoint, err := h.service.CreatePoint(c.Request.Context(), point)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"point": createdPoint})
}

func (h *PointHandler) GetPointByPointId(c *gin.Context) {
	pointId := c.Query("pointId")

	getPoint, err := h.service.GetPointByPointId(c.Request.Context(), pointId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"point": getPoint})
}

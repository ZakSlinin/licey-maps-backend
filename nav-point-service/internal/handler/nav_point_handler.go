package handler

import (
	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/model"
	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NavPointHandler struct {
	service *service.NavPointService
}

func NewNavPointHandler(service *service.NavPointService) *NavPointHandler {
	return &NavPointHandler{
		service: service,
	}
}

func (h *NavPointHandler) CreateNavPoint(c *gin.Context) {
	var navPoint model.NavPoint
	if err := c.Bind(&navPoint); err != nil {
		return
	}

	createdNavPoint, err := h.service.CreateNavPoint(c.Request.Context(), navPoint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdNavPoint)
	return
}

func (h *NavPointHandler) GetNavPointByNavPointId(c *gin.Context) {
	navPointId := c.Param("navPointId")

	getNavPoint, err := h.service.GetNavPointByNavPointId(c.Request.Context(), navPointId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getNavPoint)
}

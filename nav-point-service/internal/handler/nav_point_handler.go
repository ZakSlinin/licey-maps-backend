package handler

import (
	"net/http"
	"strconv"

	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/model"
	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/service"
	"github.com/gin-gonic/gin"
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
}

func (h *NavPointHandler) GetNavPointByNavPointId(c *gin.Context) {
	navPointId := c.Query("navPointId")

	getNavPoint, err := h.service.GetNavPointByNavPointId(c.Request.Context(), navPointId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getNavPoint)
}

// запрос на поиск маршрута
type RouteRequest struct {
	StartPointID int `json:"start_point_id" binding:"required"`
	EndPointID   int `json:"end_point_id" binding:"required"`
}

// ответ с найденным маршрутом
type RouteResponse struct {
	Path  []int    `json:"path"`
	Rooms []string `json:"rooms"`
}

func (h *NavPointHandler) FindRoute(c *gin.Context) {
	var req RouteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	path, rooms, err := h.service.FindRouteBetweenPoints(c.Request.Context(), req.StartPointID, req.EndPointID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := RouteResponse{
		Path:  path,
		Rooms: rooms,
	}

	c.JSON(http.StatusOK, response)
}

func (h *NavPointHandler) FindRouteByQuery(c *gin.Context) {
	startPointIDStr := c.Query("start_point_id")
	endPointIDStr := c.Query("end_point_id")

	if startPointIDStr == "" || endPointIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_point_id and end_point_id are required"})
		return
	}

	startPointID, err := strconv.Atoi(startPointIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_point_id"})
		return
	}

	endPointID, err := strconv.Atoi(endPointIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_point_id"})
		return
	}

	path, rooms, err := h.service.FindRouteBetweenPoints(c.Request.Context(), startPointID, endPointID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := RouteResponse{
		Path:  path,
		Rooms: rooms,
	}

	c.JSON(http.StatusOK, response)
}

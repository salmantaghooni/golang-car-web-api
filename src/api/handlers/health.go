package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/salmantaghooni/golang-car-web-api/src/api/helper"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("working", true, 0))
	return
}
 
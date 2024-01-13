package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
}

func NewTesthHandler() *TestHandler {
	return &TestHandler{}

}

func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"resault": "test",
	})
}

func (h *TestHandler) Users(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"resault": "users",
	})
}

func (h *TestHandler) UserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"type":    "userbyid",
		"resault": id,
	})
}

func (h *TestHandler) UserByUsername(c *gin.Context) {
	userName := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"type":    "username",
		"resault": userName,
	})
}

func (h *TestHandler) AccountByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"type":    "account",
		"resault": id,
	})
}

func (h *TestHandler) AddUser(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"resault": "add user",
	})
}

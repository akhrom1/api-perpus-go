package handlers

import (
	"net/http"
	"strconv"

	"api-perpus-go/internal/models"
	"api-perpus-go/internal/services"

	"github.com/gin-gonic/gin"
)

type memberReq struct {
	Name   string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
	Phone  string `form:"phone" json:"phone"`
}

func CreateMember(c *gin.Context) {
	var req memberReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	data, err := services.CreateMember(req.Name, req.Address, req.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, data)
}

func ListMembers(c *gin.Context) {
	data, err := services.ListMembers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := services.GetMember(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func UpdateMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req memberReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	m := &models.Member{
		Name:   req.Name,
		Address: req.Address,
		Phone:  req.Phone,
	}

	if err := services.UpdateMember(id, m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func DeleteMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := services.DeleteMember(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}




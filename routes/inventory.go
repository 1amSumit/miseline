package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sumit.com/mise-link/model"
)

func addInventory(c *gin.Context) {
	var inventory model.Inventory

	err := c.ShouldBind(&inventory)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err,
			"meesage": "An error occured",
		})

		return
	}

	inventory.Save()

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "inventory created successfully",
		"data":    inventory,
	})
}

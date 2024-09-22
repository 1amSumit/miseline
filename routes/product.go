package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sumit.com/mise-link/model"
	"sumit.com/mise-link/utils"
)

func addProduct(c *gin.Context) {
	var product model.Product

	inventory_id := c.Query("inventoryId")

	err := c.ShouldBind(&product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err,
			"message": "An error occured",
		})
	}
	product.InventoryId, err = strconv.ParseInt(inventory_id, 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err,
			"message": "An error occured",
		})
		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err,
			"message": "An error occured",
		})
		return
	}

	url, err := utils.UploadImage(file)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err,
			"message": "An error occured",
		})
		return
	}

	product.Image = url

	product.Save()

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "product created successfully",
		"data":    product,
	})

}

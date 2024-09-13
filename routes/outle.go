package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"sumit.com/mise-link/model"
	"sumit.com/mise-link/utils"
)

func createOutlet(c *gin.Context) {
	var outlet model.Outlet

	err := c.ShouldBind(&outlet)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = utils.SaveImage(file, c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	url := utils.UploadImage(file)

	fmt.Println(url)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Outlet created successfully",
		"data":    outlet,
	})
}

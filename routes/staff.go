package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sumit.com/mise-link/model"
	"sumit.com/mise-link/utils"
)

func createStaff(c *gin.Context) {
	var staff model.Staff

	err := c.ShouldBind(&staff)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := utils.UploadImage(file)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	outletId := c.Query("outletId")
	staff.OutletId, err = strconv.ParseInt(outletId, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	staff.Image = url

	err = staff.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error ocured",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "staff created successfully",
		"data":    staff,
	})
}

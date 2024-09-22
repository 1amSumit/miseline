package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sumit.com/mise-link/model"
	"sumit.com/mise-link/utils"
)

func createOutlet(c *gin.Context) {
	var outlet model.Outlet

	user_id := c.GetInt64("user_id")

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

	url, err := utils.UploadImage(file)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	outlet.Image = url
	outlet.UserId = int64(user_id)

	outlet.Save()

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Outlet created successfully",
		"data":    outlet,
	})
}

func getOutletsByUserId(c *gin.Context) {
	var outlets []model.Outlet

	user_id, _, boolErr := utils.IsLoggedIn(c)

	if !boolErr {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not logged in",
		})

		return
	}

	outlets, err := model.GetOutletsByUserId(user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to retrieve outlets",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, outlets)

}

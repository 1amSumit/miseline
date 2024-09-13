package utils

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func SaveImage(file *multipart.FileHeader, c *gin.Context) error {
	err := c.SaveUploadedFile(file, "./uploads/"+file.Filename)
	if err != nil {
		return err
	}

	return nil
}

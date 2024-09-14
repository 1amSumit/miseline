package utils

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

func SaveImage(file *multipart.FileHeader, c *gin.Context) error {
	err := c.SaveUploadedFile(file, "./uploads/"+file.Filename)
	if err != nil {
		return err
	}

	return nil
}

func credentials() (*cloudinary.Cloudinary, context.Context, error) {

	cld, err := cloudinary.New()
	if err != nil {
		return nil, nil, err
	}

	cld.Config.URL.Secure = true
	ctx := context.Background()
	return cld, ctx, nil
}

func UploadImage(file *multipart.FileHeader) (string, error) {
	cld, ctx, err := credentials()
	if err != nil {
		return "", err

	}

	src, err := file.Open()
	if err != nil {
		return "", err

	}
	defer src.Close()
	resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{
		PublicID:       "shopOutlet",
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
	})
	if err != nil {
		return "", err

	}

	if resp.SecureURL == "" {
		return "", err

	}

	return resp.SecureURL, nil
}

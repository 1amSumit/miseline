package utils

import (
	"context"
	"fmt"
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

func UploadImage(file *multipart.FileHeader) string {

	cld, ctx, err := credentials()
	if err != nil {
		fmt.Println("Error initializing Cloudinary:", err)
		return ""
	}

	src, err := file.Open()
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer src.Close()
	resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{
		PublicID:       "shopOutlet",
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
	})
	if err != nil {
		fmt.Println("Error uploading image:", err)
		return ""
	}

	fmt.Printf("Full Cloudinary response: %+v\n", resp)

	if resp.SecureURL == "" {
		fmt.Println("SecureURL is missing in the response")
		return ""
	}

	fmt.Println("Uploaded image Secure URL:", resp.SecureURL)
	return resp.URL
}

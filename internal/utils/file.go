package utils

import (
	"final_project/internal/helper"
	"fmt"
	"log"

	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

var imageDir = "assets/public/images"

var fileUploadDir = filepath.Join(helper.ProjectDirectory, imageDir)

func HandleUploadFile(ctx *fiber.Ctx) error {
	file, _ := ctx.FormFile("photo")
	if file != nil {
		if !checkFileType(file) {
			return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, "invalid file type", nil, fiber.StatusBadRequest)
		}
		fileName := replaceFileName(file.Filename)
		if err := ctx.SaveFile(file, fmt.Sprintf("./%s/store/%s", imageDir, fileName)); err != nil {
			log.Println("Failed to save file.")
			return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, err.Error(), nil, fiber.StatusBadRequest)
		}
		ctx.Locals("photoUrl", fmt.Sprintf("store/%s", fileName))
	} else {
		ctx.Locals("photoUrl", "")
	}
	return ctx.Next()
}

func HandleMultipleFile(ctx *fiber.Ctx) error {
	fmt.Println("triggered")
	form, err := ctx.MultipartForm()
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, err.Error(), nil, fiber.StatusBadRequest)
	}
	var photoUrls []string
	for _, photo := range form.File["photos"] {
		if !checkFileType(photo) {
			return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, "invalid file type", nil, fiber.StatusBadRequest)
		}
		fileName := replaceFileName(photo.Filename)
		if err := ctx.SaveFile(photo, fmt.Sprintf("./%s/products/%s", imageDir, fileName)); err != nil {
			log.Println("Failed to save file.")
			return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, err.Error(), nil, fiber.StatusBadRequest)
		}
		photoUrls = append(photoUrls, fmt.Sprintf("products/%s", fileName))
	}
	ctx.Locals("photoUrls", photoUrls)
	return ctx.Next()
}

func replaceFileName(filename string) string {
	newName := strconv.FormatInt(time.Now().Unix(), 10)
	return newName + "-" + filename
}

func RemoveUnusedPhoto(filename string) error {
	err := os.Remove(filepath.Join(fileUploadDir, filename))
	if err != nil {
		return err
	}
	return nil
}

func checkFileType(file *multipart.FileHeader) bool {
	return strings.Split(file.Header["Content-Type"][0], "/")[0] == "image"
}

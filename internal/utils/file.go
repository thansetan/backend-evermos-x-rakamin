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
		if err := ctx.SaveFile(file, fmt.Sprintf("./%s/%s", imageDir, fileName)); err != nil {
			log.Println("Failed to save file.")
			return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, err.Error(), nil, fiber.StatusBadRequest)
		}

		ctx.Locals("photoUrl", fileName)
	} else {
		ctx.Locals("photoUrl", "")
	}
	return ctx.Next()
}

func replaceFileName(filename string) string {
	newName := strconv.FormatInt(time.Now().Unix(), 10)
	return strings.Replace(filename, strings.Split(filename, ".")[0], newName, 1)
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

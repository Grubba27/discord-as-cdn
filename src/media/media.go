package media

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"os"
)

var Path string

func check(file *multipart.FileHeader) error {
	if file.Size > 10*1024*1024 {
		fmt.Println("File is too big")
		return fmt.Errorf("file is too big")
	}

	fileMimeType := file.Header.Get("Content-Type")

	allowedMimesList := []string{"image/jpeg", "image/pjpeg", "image/png", "image/gif"}
	allowedMimes := map[string]bool{}
	for _, mime := range allowedMimesList {
		allowedMimes[mime] = true
	}
	if !allowedMimes[fileMimeType] {
		fmt.Println("File is not an image")
		return fmt.Errorf("file is not an image")
	}
	return nil
}

func ToOSFile(ctx *fiber.Ctx, file *multipart.FileHeader) (*os.File, error) {
	if err := check(file); err != nil {
		return nil, err
	}

	Path = fmt.Sprintf("./%s", file.Filename)
	if err := ctx.SaveFile(file, Path); err != nil {
		e := fmt.Errorf("was not able to save file")
		return nil, e
	}

	osFile, _ := os.Open(Path)
	return osFile, nil
}

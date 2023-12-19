package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		err = c.SaveFile(file, filepath.Join("./images", file.Filename))
		if err != nil {
			return err
		}

		return c.SendString("File successfully uploaded")
	})

	app.Get("/download/:filename", func(c *fiber.Ctx) error {
		filename := c.Params("filename")

		file, err := os.Open(filepath.Join("./images", filename))
		if err != nil {
			return err
		}
		defer file.Close()

		// os.FileInfo("./downloads", os.ModePerm)
		os.MkdirAll("./downloads", os.ModePerm)

		out, err := os.Create(filepath.Join("./downloads", filename))
		if err != nil {
			return err
		}
		defer out.Close()

		// ok, err = io.Copy(out, filename)
		_, err = io.Copy(out, file)
		if err != nil {
			return err
		}

		return c.SendString("File successfully downloaded")
	})

	app.Listen(":3000")
}

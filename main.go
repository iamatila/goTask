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

		// Save the file to the 'images' directory
		err = c.SaveFile(file, filepath.Join("./images", file.Filename))
		if err != nil {
			return err
		}

		return c.SendString("File successfully uploaded")
	})

	app.Get("/download/:filename", func(c *fiber.Ctx) error {
		filename := c.Params("filename")

		// Open the file
		file, err := os.Open(filepath.Join("./images", filename))
		if err != nil {
			return err
		}
		defer file.Close()

		// Create the 'downloads' directory if it doesn't exist
		os.MkdirAll("./downloads", os.ModePerm)

		// Create a new file in the 'downloads' directory
		out, err := os.Create(filepath.Join("./downloads", filename))
		if err != nil {
			return err
		}
		defer out.Close()

		// Copy the file
		_, err = io.Copy(out, file)
		if err != nil {
			return err
		}

		return c.SendString("File successfully downloaded")
	})

	app.Listen(":3000")
}

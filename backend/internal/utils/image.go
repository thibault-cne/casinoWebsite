package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
)

func ImageToWebp(fileHeader *multipart.FileHeader, filename string) error {
	// Open file
	file, err := fileHeader.Open()
	temp := fmt.Sprintf("./images/temp/%s", fileHeader.Filename)
	filePath := fmt.Sprintf("./images/webp/%s.webp", filename)

	if err != nil {
		return err
	}
	defer file.Close()

	buffer, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	// Write temp file
	os.WriteFile(temp, buffer, 0644)

	// Convert temp file to webp with ffmpeg
	// We add -y to overwrite the file
	// We wait for cmd to finish
	cmd := exec.Command("ffmpeg", "-i", temp, filePath, "-y")
	cmd.Start()
	err = cmd.Wait()

	if err != nil {
		return err
	}

	// Delete temp file
	err = os.Remove(temp)

	if err != nil {
		return err
	}

	return nil
}

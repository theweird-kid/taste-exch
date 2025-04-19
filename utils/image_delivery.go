package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const imgbbAPI = "https://api.imgbb.com/1/upload"

// ImgBBResponse represents the response structure from ImgBB API
type ImgBBResponse struct {
	Data struct {
		URL string `json:"url"`
	} `json:"data"`
	Success bool `json:"success"`
}

// UploadImage uploads an image to ImgBB and returns the URL of the uploaded image
func UploadImage(filePath string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("failed to load .env file: %w", err)
	}

	apiKey := os.Getenv("IMGBB_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("IMGBB_KEY is not set in the environment")
	}

	// Open the image file
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a buffer to hold the multipart form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the API key to the form
	if err := writer.WriteField("key", apiKey); err != nil {
		return "", fmt.Errorf("failed to write API key to form: %w", err)
	}

	// Add the image file to the form
	part, err := writer.CreateFormFile("image", filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %w", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return "", fmt.Errorf("failed to copy file content: %w", err)
	}

	// Close the writer to finalize the form
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %w", err)
	}

	// Make the HTTP POST request
	req, err := http.NewRequest("POST", imgbbAPI, body)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Read and parse the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to upload image: %s", string(respBody))
	}

	var imgbbResp ImgBBResponse
	if err := json.Unmarshal(respBody, &imgbbResp); err != nil {
		return "", fmt.Errorf("failed to parse response JSON: %w", err)
	}

	if !imgbbResp.Success {
		return "", fmt.Errorf("image upload failed: %s", string(respBody))
	}

	// Return the image URL
	return imgbbResp.Data.URL, nil
}

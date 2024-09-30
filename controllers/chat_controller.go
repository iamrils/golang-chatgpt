package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// Request structure for incoming requests
type Request struct {
	Prompt string `json:"prompt" binding:"required"`
}

// Response structure for the API response
type Response struct {
	Choices []Choice `json:"choices"`
}

// Choice structure to hold response choices
type Choice struct {
	Message Message `json:"message"`
}

// Message structure to hold message content
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// GetUsers retrieves the list of users
func GenerateChat(c *gin.Context) {
	var req Request

	// Bind JSON request to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a client
	client := resty.New()

	// Prepare the request to OpenAI API
	apiKey := os.Getenv("OPENAI_API_KEY")
	response, err := client.R().
		SetHeader("Authorization", "Bearer "+apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model": "gpt-3.5-turbo", // Use gpt-3.5-turbo model
			"messages": []map[string]string{
				{"role": "user", "content": req.Prompt},
			},
			"max_tokens": 100,
		}).
		Post("https://api.openai.com/v1/chat/completions") // Use the chat/completions endpoint

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get response from OpenAI"})
		return
	}

	// Check if the request was successful
	if response.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": "OpenAI API request failed", "details": response.String()})
		return
	}

	// Parse the response
	var gptResponse Response
	err = json.Unmarshal(response.Body(), &gptResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	// Return the content of the first message choice
	c.JSON(http.StatusOK, gin.H{"response": gptResponse.Choices[0].Message.Content})
}

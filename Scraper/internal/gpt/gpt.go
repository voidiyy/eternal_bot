package gpt

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
)

type Client struct {
	apiKey  string
	history map[int64][]map[string]string
	mu      sync.Mutex
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		history: make(map[int64][]map[string]string),
	}
}

func (c *Client) ProcessMessage(chatID int64, message string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.history[chatID] = append(c.history[chatID], map[string]string{
		"role":    "user",
		"content": message,
	})

	reqBody := map[string]interface{}{
		"model":    "gpt-4",
		"messages": c.history[chatID],
	}
	reqBytes, _ := json.Marshal(reqBody) // Отримуємо відповідь від GPT і додаємо її до контексту

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	//gpt resp
	response := result["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	c.history[chatID] = append(c.history[chatID], map[string]string{
		"role":    "assistant",
		"content": response,
	})

	return response, nil
}

func (c *Client) ClearHistory(chatID int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.history[chatID] = []map[string]string{}
}

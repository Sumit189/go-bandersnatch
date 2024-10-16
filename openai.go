package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const apiKey = "OPENAI_API_KEY"

func generateDynamicResponse(memory *Memory, choice string) (string, map[string]string, string, error) {
	prompt := fmt.Sprintf("You chose to go %s. What happens next? Provide a short story and two choices, choices should not be included in the story", choice)
	chatHistory := []map[string]string{
		{"role": "system", "content": "You are a helpful assistant that generates bandersnatch-style stories."},
	}
	for _, message := range memory.GetMemory() {
		chatHistory = append(chatHistory, map[string]string{"role": message.Role, "content": message.Content})
	}

	messagePayload := append(chatHistory, map[string]string{"role": "user", "content": prompt})
	client := &http.Client{}
	reqBody, err := json.Marshal(map[string]interface{}{
		"model":    "gpt-4o-2024-08-06",
		"messages": messagePayload,
		"response_format": map[string]interface{}{
			"type": "json_schema",
			"json_schema": map[string]interface{}{
				"name": "story_response",
				"schema": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"story": map[string]interface{}{
							"type": "string",
						},
						"1": map[string]interface{}{
							"type": "string",
						},
						"2": map[string]interface{}{
							"type": "string",
						},
					},
					"required": []string{"story", "1", "2"},
				},
			},
		},
	})

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", nil, "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", nil, "", err
	}

	choicesInterface, ok := result["choices"].([]interface{})
	if !ok || len(choicesInterface) == 0 {
		return "", nil, "", fmt.Errorf("no choices found")
	}

	firstChoice := choicesInterface[0].(map[string]interface{})
	message := firstChoice["message"].(map[string]interface{})
	content := message["content"].(string)

	var rawResponse map[string]interface{}
	if err := json.Unmarshal([]byte(content), &rawResponse); err != nil {
		return "", nil, "", err
	}

	choices := make(map[string]string)
	choices["1"] = rawResponse["1"].(string)
	choices["2"] = rawResponse["2"].(string)

	return rawResponse["story"].(string), choices, choice, nil
}

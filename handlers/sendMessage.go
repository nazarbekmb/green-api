package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"green-api/models"
	"io/ioutil"
	"net/http"
)

func SendMessageHandler(w http.ResponseWriter, r *http.Request, idInstance, apiTokenInstance, apiUrl string, chatID, message string) (string, error) {
	url := fmt.Sprintf("%s/waInstance%s/sendMessage/%s", apiUrl, idInstance, apiTokenInstance)

	requestBody := models.SendMessageRequest{
		ChatID:  chatID,
		Message: message,
	}
	payload, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error creating JSON payload: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	return string(body), nil
}

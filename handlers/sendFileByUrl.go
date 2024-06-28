package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"green-api/models"
	"io/ioutil"
	"net/http"
)

func SendFileByUrlHandler(w http.ResponseWriter, r *http.Request, idInstance, apiTokenInstance, apiUrl, chatID, urlFile, fileName, caption string) (string, error) {
	url := fmt.Sprintf("%s/waInstance%s/sendFileByUrl/%s", apiUrl, idInstance, apiTokenInstance)

	requestBody := models.SendFileRequest{
		ChatID:   chatID,
		URLFile:  urlFile,
		FileName: fileName,
		Caption:  caption,
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

	if resp.StatusCode == http.StatusRequestEntityTooLarge {
		return "", fmt.Errorf("File from URL exceeded max upload size. Size: XXXXmb Limit: 100mb Url: %s", urlFile)
	} else if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	return string(body), nil
}

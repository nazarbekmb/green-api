package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetStateInstanceHandler(w http.ResponseWriter, r *http.Request, idInstance, apiTokenInstance, apiUrl string) (string, error) {
	url := fmt.Sprintf("%s/waInstance%s/getStateInstance/%s", apiUrl, idInstance, apiTokenInstance)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("Error creating request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return "", fmt.Errorf("401 Unauthorized: Проверьте корректность указания apiTokenInstance, partnerToken")
	} else if resp.StatusCode == http.StatusForbidden {
		return "", fmt.Errorf("403 Forbidden: Проверьте корректность указания idInstance")
	} else if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response: %v", err)
	}

	return string(body), nil
}

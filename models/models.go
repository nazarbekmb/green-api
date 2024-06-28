package models

type ResponseData struct {
	Response         string
	IDInstance       string
	APITokenInstance string
	ErrorMessage     string
}

type SendFileRequest struct {
	ChatID   string `json:"chatId"`
	URLFile  string `json:"urlFile"`
	FileName string `json:"fileName"`
	Caption  string `json:"caption"`
}

type SendMessageRequest struct {
	ChatID  string `json:"chatId"`
	Message string `json:"message"`
}

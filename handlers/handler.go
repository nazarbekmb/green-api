package handlers

import (
	"fmt"
	"green-api/models"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var data models.ResponseData

	if idInstanceCookie, err := r.Cookie("idInstance"); err == nil {
		data.IDInstance = idInstanceCookie.Value
	}
	if apiTokenInstanceCookie, err := r.Cookie("apiTokenInstance"); err == nil {
		data.APITokenInstance = apiTokenInstanceCookie.Value
	}

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Unable to load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
	} else if r.Method == http.MethodPost {
		idInstance := r.FormValue("idInstance")
		apiTokenInstance := r.FormValue("apiTokenInstance")
		http.SetCookie(w, &http.Cookie{
			Name:  "idInstance",
			Value: idInstance,
			Path:  "/",
		})
		http.SetCookie(w, &http.Cookie{
			Name:  "apiTokenInstance",
			Value: apiTokenInstance,
			Path:  "/",
		})

		chatId := r.FormValue("chatId")
		message := r.FormValue("message")
		fileChatId := r.FormValue("fileChatId")
		fileUrl := r.FormValue("fileUrl")
		fileName := r.FormValue("fileName")
		caption := r.FormValue("caption")
		var apiUrl string

		if len(idInstance) >= 4 {
			apiUrl = fmt.Sprintf("https://%s.api.greenapi.com", idInstance[:4])
		}
		action := r.FormValue("action")
		var response string
		var err error
		switch action {
		case "getSettings":
			response, err = GetSettingsHandler(w, r, idInstance, apiTokenInstance, apiUrl)
		case "getStateInstance":
			response, err = GetStateInstanceHandler(w, r, idInstance, apiTokenInstance, apiUrl)
		case "sendMessage":
			response, err = SendMessageHandler(w, r, idInstance, apiTokenInstance, apiUrl, chatId, message)
		case "sendFileByUrl":
			response, err = SendFileByUrlHandler(w, r, idInstance, apiTokenInstance, apiUrl, fileChatId, fileUrl, fileName, caption)
		default:
			http.Error(w, "Unknown action", http.StatusBadRequest)
			return
		}
		if err != nil {
			data.ErrorMessage = err.Error()
		} else {

			data.Response = response
		}
		data.IDInstance = idInstance
		data.APITokenInstance = apiTokenInstance

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Unable to load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
	} else {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Unable to load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
	}
}

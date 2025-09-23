package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Mathis-Pain/Forum/utils"
)

var CreatTopicHtml = template.Must(template.New("create-topic.html").Funcs(funcMap).ParseFiles(
	"templates/create-topic.html",
	"templates/login.html",
	"templates/header.html",
	"templates/categorie.html",
	"templates/initpage.html",
))

func CreateTopicHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			utils.InternalServError(w)
			return
		}
		// Verification username et password non nul
		topicTitle := r.FormValue("title")
		topicDescription := r.FormValue("description")
		if topicTitle == "" || topicDescription == "" {
			http.Error(w, "Tous les champs sont requis", http.StatusBadRequest)
			return
		}
	}

	err := CreatTopicHtml.Execute(w, nil)
	if err != nil {
		log.Printf("<create-topic-handler.go> Could not execute template <create-topic.html>: %v\n", err)
		utils.NotFoundHandler(w)

	}
}

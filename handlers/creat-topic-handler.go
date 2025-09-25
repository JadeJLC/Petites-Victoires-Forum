package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Mathis-Pain/Forum/utils"
	"github.com/Mathis-Pain/Forum/utils/postactions"
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
		topicName := r.FormValue("title")
		message := r.FormValue("message")
		stringcatID := r.FormValue("category_id")
		catID, err := strconv.Atoi(stringcatID)
		if err != nil {
			fmt.Println("Erreur de conversion dans creat-topic-handler catID:", err)
			return
		}
		if topicName == "" || message == "" {
			http.Error(w, "Tous les champs sont requis", http.StatusBadRequest)
			return
		}

		db, err := sql.Open("sqlite3", "./data/forum.db")
		if err != nil {
			log.Printf("<cathandler.go> Could not open database : %v\n", err)
			return
		}
		defer db.Close()

		_, userID, _ := utils.GetUserNameAndIDByCookie(r, db)
		postactions.CreateNewtopic(userID, catID, topicName, message)

		// Redirection vers la page catégorie
		http.Redirect(w, r, fmt.Sprintf("/topic?id=%d", catID), http.StatusSeeOther)
		return
	}

	err := CreatTopicHtml.Execute(w, nil)
	if err != nil {
		log.Printf("<create-topic-handler.go> Could not execute template <create-topic.html>: %v\n", err)
		utils.NotFoundHandler(w)

	}
}

package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/Mathis-Pain/Forum/handlers/subhandlers"
	"github.com/Mathis-Pain/Forum/models"
	"github.com/Mathis-Pain/Forum/utils"
	admin "github.com/Mathis-Pain/Forum/utils/adminfuncs"
)

var funcShort = template.FuncMap{
	"preview": utils.Preview,
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./data/forum.db")
	if err != nil {
		log.Print("<profilhandler.go> Erreur à l'ouverture de la base de données :", err)
		utils.InternalServError(w)
		return
	}
	defer db.Close()

	categories, currentUser, err := subhandlers.BuildHeader(r, w, db)
	if err != nil {
		log.Printf("<cathandler.go> Erreur dans la construction du header : %v\n", err)
		utils.InternalServError(w)
		return
	}

	isAdmin, err := admin.CheckIfAdmin(currentUser.Username)
	if !isAdmin && err == nil {
		log.Print("Tentative d'accès non autorisé au panneau d'administration.")
		utils.UnauthorizedError(w)
		return
	} else if err != nil {
		if err == sql.ErrNoRows {
			log.Print("Tentative d'accès non autorisé au panneau d'administration.")
			utils.UnauthorizedError(w)
			return
		}
		log.Print("<adminhandler.go> Erreur dans la vérification des accréditations :", err)
		utils.InternalServError(w)
		return
	}

	parts := strings.Split(r.URL.Path, "/")

	categories, topics, err := admin.GetAllTopics(categories, db)
	if err != nil {
		log.Print("Erreur dans la récupération des sujets", err)
		utils.InternalServError(w)
		return
	}
	lastmonthpost, stats, users, err := admin.GetStats(topics)
	if err != nil {
		log.Print("Erreur dans la récupération des statistiques", err)
		utils.InternalServError(w)
		return
	}
	if len(parts) == 3 && parts[1] == "admin" {
		adminHome(categories, topics, stats, users, w, currentUser, lastmonthpost)
	} else if len(parts) > 3 {
		fmt.Println(parts[2])
		switch parts[1] {
		case "userlist":
			adminUsers(users, w)
		case "catlist":
			adminCategories(categories, w)
		case "topiclist":
			adminTopics(topics, w)
		case "seeposts":
			adminPost(lastmonthpost, w)
		}
	}
}

func adminPost(lastmonthpost []models.LastPost, w http.ResponseWriter) {
	data := struct {
		PageName  string
		LastMonth []models.LastPost
	}{
		PageName:  "Messages du dernier mois",
		LastMonth: lastmonthpost,
	}

	pageToLoad, err := template.ParseFiles("templates/all-posts.html", "templates/header.html", "templates/initpage.html")
	if err != nil {
		log.Printf("<adminhandler.go> Erreur dans la génération du template adminPost : %v", err)
		utils.InternalServError(w)
		return
	}

	err = pageToLoad.Execute(w, data)
	if err != nil {
		utils.InternalServError(w)
		return
	}
}

func adminTopics(topics []models.Topic, w http.ResponseWriter) {
	data := struct {
		PageName string
		Topics   []models.Topic
	}{
		PageName: "Administration des sujets",
		Topics:   topics,
	}

	pageToLoad, err := template.ParseFiles("templates/all-topics.html", "templates/header.html", "templates/initpage.html")
	if err != nil {
		log.Printf("<adminhandler.go> Erreur dans la génération du template adminCategories : %v", err)
		utils.InternalServError(w)
		return
	}

	err = pageToLoad.Execute(w, data)
	if err != nil {
		utils.InternalServError(w)
		return
	}
}

func adminCategories(categories []models.Category, w http.ResponseWriter) {
	data := struct {
		PageName   string
		Categories []models.Category
	}{
		PageName:   "Administration des catégories",
		Categories: categories,
	}

	pageToLoad, err := template.ParseFiles("templates/all-categories.html", "templates/header.html", "templates/initpage.html")
	if err != nil {
		log.Printf("<adminhandler.go> Erreur dans la génération du template adminCategories : %v", err)
		utils.InternalServError(w)
		return
	}

	err = pageToLoad.Execute(w, data)
	if err != nil {
		utils.InternalServError(w)
		return
	}
}

func adminUsers(users []models.User, w http.ResponseWriter) {
	data := struct {
		PageName string
		Users    []models.User
	}{
		PageName: "Administrer les utilisateurs",
		Users:    users,
	}

	pageToLoad, err := template.ParseFiles("templates/all-users.html", "templates/header.html", "templates/initpage.html")
	if err != nil {
		log.Printf("<adminhandler.go> Erreur dans la génération du template adminUsers : %v", err)
		utils.InternalServError(w)
		return
	}

	err = pageToLoad.Execute(w, data)
	if err != nil {
		log.Print("<adminhandler.go> Erreur dans la lecture du template adminUsers : ", err)
		utils.InternalServError(w)
		return
	}
}

func adminHome(categories []models.Category, topics []models.Topic, stats models.Stats, users []models.User, w http.ResponseWriter, currentUser models.UserLoggedIn, postList []models.LastPost) {
	data := struct {
		PageName    string
		Categories  []models.Category
		Topics      []models.Topic
		Users       []models.User
		Stats       models.Stats
		PostList    []models.LastPost
		CurrentUser models.UserLoggedIn
	}{
		PageName:    "Panneau d'administration",
		Categories:  categories,
		Topics:      topics,
		Users:       users,
		Stats:       stats,
		PostList:    postList,
		CurrentUser: currentUser,
	}

	pageToLoad := template.Must(template.New("admin.html").Funcs(funcShort).ParseFiles("templates/admin/admin.html",
		"templates/admin/adminheader.html",
		"templates/initpage.html",
		"templates/login.html"))

	// pageToLoad, err := template.Funcs(funcShort).ParseFiles("templates/admin/admin.html", "templates/admin/adminheader.html", "templates/initpage.html", "templates/login.html")
	// if err != nil {
	// 	log.Printf("<adminhandler.go> Erreur dans la génération du template adminHome : %v", err)
	// 	utils.InternalServError(w)
	// 	return
	// }

	err := pageToLoad.Execute(w, data)
	if err != nil {
		log.Printf("<adminhandler.go> Erreur dans la lecture du template adminHome : %v", err)
		utils.InternalServError(w)
		return
	}
}

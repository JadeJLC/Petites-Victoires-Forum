package authhandlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/Mathis-Pain/Forum/sessions"
	"github.com/Mathis-Pain/Forum/utils"
	"github.com/Mathis-Pain/Forum/utils/getdata"
)

var funcMap = template.FuncMap{
	"preview": getdata.Preview,
}

var HomeHtml = template.Must(template.New("home.html").Funcs(funcMap).ParseFiles(
	"templates/home.html", "templates/login.html", "templates/header.html", "templates/initpage.html",
))

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	url := "/" + parts[1]

	if url == "/login" {
		url = "/"
	}

	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			utils.InternalServError(w)
			return
		}
		// Récupération des données
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Vérifie login + mot de passe (utils.Authentification s’occupe de la DB)
		db, err := sql.Open("sqlite3", "./data/forum.db")
		if err != nil {
			utils.InternalServError(w)
			return
		}
		defer db.Close()

		user, loginErr := utils.Authentification(db, username, password)
		if loginErr != nil {
			if strings.Contains(loginErr.Error(), "db") {
				// En cas d'erreur dans la base de données
				utils.InternalServError(w)
			} else {
				// En cas d'erreur qui ne vient pas de la base de données
				// Création d'une session temporaire et anonyme
				session, err := sessions.CreateSession(0)
				if err != nil {
					utils.InternalServError(w)
					return
				}
				// Stockage du message d'erreur dans la session
				session.Data["LoginErr"] = loginErr.Error()

				// Update de la session
				if err := sessions.SaveSessionToDB(session); err != nil {
					utils.InternalServError(w)
					return
				}
				// Pose du cookie
				http.SetCookie(w, &http.Cookie{
					Name:     "session_id",
					Value:    session.ID,
					Expires:  session.ExpiresAt,
					HttpOnly: true,
					Secure:   false, // false en local, true si HTTPS
					Path:     "/",
				})
				// Redirection vers la page d'origine
				referer := r.Header.Get("Referer")
				if referer == "" {
					referer = "/"
				}
				fmt.Println(referer)
				http.Redirect(w, r, referer, http.StatusSeeOther)
				return
			}
		}

		//Invalider toutes les sessions existantes
		if err := sessions.InvalidateUserSessions(user.ID); err != nil {
			utils.InternalServError(w)
			return
		}

		// Créer une nouvelle session
		session, err := sessions.CreateSession(user.ID)
		if err != nil {
			utils.InternalServError(w)
			return
		}

		// Ajoute des infos dans la session
		session.Data["user"] = user.Username
		if err := sessions.SaveSessionToDB(session); err != nil {
			utils.InternalServError(w)
			return
		}

		// Pose le cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    session.ID,
			Expires:  session.ExpiresAt,
			HttpOnly: true,
			Secure:   false, // false en local, true si HTTPS
			Path:     "/",
		})

		http.Redirect(w, r, url, http.StatusSeeOther)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

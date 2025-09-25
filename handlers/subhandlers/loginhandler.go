package subhandlers

import (
	"database/sql"
	"html/template"
	"log"
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
	// si l'utilisateur demande le formulaire
	// case http.MethodGet:
	// 	if err := .Execute(w, nil); err != nil {
	// 		utils.InternalServError(w)
	// 		return
	// 	}
	// Si l'utilisateur envoi le formulaire
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			utils.InternalServError(w)
			return
		}
		// Verification username et password non nul
		username := r.FormValue("username")
		password := r.FormValue("password")
		// if username == "" || password == "" {
		// 	http.Error(w, "Tous les champs sont requis", http.StatusBadRequest)
		// 	return
		// }

		// Vérifie login + mot de passe (utils.Authentification s’occupe de la DB)
		db, err := sql.Open("sqlite3", "./data/forum.db")
		if err != nil {
			utils.InternalServError(w)
			return
		}
		defer db.Close()

		user, err := utils.Authentification(db, username, password)
		if err != nil {
			if strings.Contains(err.Error(), "db") {
				// En cas d'erreur dans la base de données
				utils.InternalServError(w)
			} else {
				// En cas d'erreur qui ne vient pas de la base de données
				data := struct {
					LoginErr string
				}{
					LoginErr: err.Error(),
				}

				err = HomeHtml.Execute(w, data)
				if err != nil {
					log.Printf("<loginhandler.go> Could not execute template <home.html> : %v\n", err)
					utils.InternalServError(w)
					return
				}
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

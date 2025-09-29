package authhandlers

import (
	"fmt"
	"net/http"

	"github.com/Mathis-Pain/Forum/sessions"
	"github.com/Mathis-Pain/Forum/utils"
)

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		session, err := sessions.GetSessionFromRequest(r)
		if err != nil {
			fmt.Print("Erreur, supression session (deconnexion)")
			utils.InternalServError(w)
			return
		}
		sessions.DeleteSession(session.ID)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

package subhandlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Mathis-Pain/Forum/models"
	"github.com/Mathis-Pain/Forum/utils/getdata"
)

func UserEditHandler(r *http.Request, user *models.User) error {
	username := r.FormValue("username")
	status := r.FormValue("status")

	if username != "" {
		user.Username = username
	}

	if status != "" {
		user.Status = status
	}

	db, err := sql.Open("sqlite3", "./data/forum.db")
	if err != nil {
		log.Print("<profilhandler.go> Erreur à l'ouverture de la base de données :", err)
		return err
	}
	defer db.Close()

	sqlUpdate := `UPDATE user SET username = ?, role_id = ? WHERE id = ?`
	stmt, err := db.Prepare(sqlUpdate)
	if err != nil {
		log.Print(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Username, getdata.CodeUserStatus(user.Status), user.ID)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

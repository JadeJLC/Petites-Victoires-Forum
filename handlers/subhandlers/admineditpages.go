package subhandlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/Mathis-Pain/Forum/models"
	"github.com/Mathis-Pain/Forum/utils/getdata"
)

func UserEditHandler(r *http.Request, users []models.User) error {
	stringID := r.FormValue("userID")
	ID, err := strconv.Atoi(stringID)
	if err != nil {
		log.Print("<adminhandler.go adminUsers> Erreur dans la récupération de l'ID utilisateur : ", err)
		return err
	}

	var user models.User

	for _, current := range users {
		if current.ID == ID {
			user = current
			break
		}
	}

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
		log.Print("<admineditpages.go> Erreur à l'ouverture de la base de données :", err)
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

func BanUserHandler(stringID string) error {
	ID, err := strconv.Atoi(stringID)

	if err != nil {
		log.Print("<admineditpages.go> Erreur dans la récupération de l'utilisateur à bannir")
		return err
	}

	db, err := sql.Open("sqlite3", "./data/forum.db")
	if err != nil {
		log.Print("<admineditpages.go> Erreur à l'ouverture de la base de données :", err)
		return err
	}
	defer db.Close()

	sqlUpdate := `UPDATE user SET role_id = 4 WHERE id = ?`
	stmt, err := db.Prepare(sqlUpdate)
	if err != nil {
		log.Print(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(ID)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func UnbanUserHandler(stringID string) error {
	ID, err := strconv.Atoi(stringID)

	if err != nil {
		log.Print("<admineditpages.go> Erreur dans la récupération de l'utilisateur à débannir")
		return err
	}

	db, err := sql.Open("sqlite3", "./data/forum.db")
	if err != nil {
		log.Print("<admineditpages.go> Erreur à l'ouverture de la base de données :", err)
		return err
	}
	defer db.Close()

	sqlUpdate := `UPDATE user SET role_id = 3 WHERE id = ?`
	stmt, err := db.Prepare(sqlUpdate)
	if err != nil {
		log.Print(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(ID)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func DeleteUserHandler(stringID string) error {
	ID, err := strconv.Atoi(stringID)
	if err != nil {
		log.Print("<admineditpages.go> Erreur dans la récupération de l'utilisateur à bannir")
		return err
	}

	db, err := sql.Open("sqlite3", "./data/forum.db")
	if err != nil {
		log.Print("<admineditpages.go> Erreur à l'ouverture de la base de données :", err)
		return err
	}
	defer db.Close()

	if ID == 1 {
		log.Print("Tentative de suppression de Zoé")
		return nil
	}

	sqlUpdate := `DELETE FROM user WHERE id = ?`
	stmt, err := db.Prepare(sqlUpdate)
	if err != nil {
		log.Print(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(ID)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

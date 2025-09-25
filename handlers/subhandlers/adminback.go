package subhandlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/Mathis-Pain/Forum/models"
)

func AdminIsCatModified(r *http.Request, categories []models.Category) (models.Category, bool, error) {
	name := r.FormValue("name")
	description := r.FormValue("description")
	stringID := r.FormValue("catID")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		log.Print("<adminback.go adminUsers> Erreur dans la récupération de l'ID ce catégorie : ", err)
		return models.Category{}, false, err
	}

	var categ models.Category

	for _, current := range categories {
		if current.ID == ID {
			categ = current
			break
		}
	}

	if categ.Name == name && categ.Description == description {
		return categ, false, nil
	}

	return categ, true, nil

}

func AdminDeleteMessages(db *sql.DB, ID int) error {
	sqlUpdate := `DELETE FROM message WHERE topic_id = ?`

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

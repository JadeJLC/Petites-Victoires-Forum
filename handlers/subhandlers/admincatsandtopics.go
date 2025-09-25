package subhandlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/Mathis-Pain/Forum/models"
	"github.com/Mathis-Pain/Forum/utils/getdata"
)

func CatEditHandler(r *http.Request, categ models.Category) error {
	name := r.FormValue("name")
	description := r.FormValue("description")
	if name != "" {
		categ.Name = name
	}

	if description != "" {
		categ.Description = description
	}

	db, err := sql.Open("sqlite3", "./data/forum.db")
	if err != nil {
		log.Print("<admincatsandtopics.go> Erreur à l'ouverture de la base de données :", err)
		return err
	}
	defer db.Close()

	sqlUpdate := `UPDATE category SET name = ?, description = ? WHERE id = ?`
	stmt, err := db.Prepare(sqlUpdate)
	if err != nil {
		log.Print(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(categ.Name, categ.Description, categ.ID)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func DeleteCatHandler(stringID string) error {
	ID, err := strconv.Atoi(stringID)
	if err != nil {
		log.Print("<admincatsandtopics.go> Erreur dans la récupération de la catégorie à supprimer", err)
		return err
	}

	db, err := sql.Open("sqlite3", "./data/forum.db")
	if err != nil {
		log.Print("<admincatsandtopics.go> Erreur à l'ouverture de la base de données :", err)
		return err
	}
	defer db.Close()

	if ID == 1 {
		log.Print("Tentative de suppression de Plop")
		return nil
	}

	sqlUpdate := `DELETE FROM category WHERE id = ?`
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

	topicList, err := getdata.GetTopicList(db, ID)
	if err != nil {
		return err
	}

	for i := 0; i < len(topicList); i++ {
		err := AdminDeleteMessages(db, topicList[i].TopicID)
		if err != nil {
			log.Print("<admincatsandtopics.go Erreur dans la suppression des messages", err)
			return err
		}
	}

	sqlUpdate = `DELETE FROM topic WHERE category_id = ?`
	stmt, err = db.Prepare(sqlUpdate)
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

	log.Print("Catégorie et sujets liés supprimés avec succès.")

	return nil
}

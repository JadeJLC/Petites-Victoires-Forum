package subhandlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Mathis-Pain/Forum/models"
	"github.com/Mathis-Pain/Forum/utils/getdata"
)

func EditCatHandler(r *http.Request, categ models.Category) error {
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

func AddCatHandler(r *http.Request) error {
	name := r.FormValue("newcatname")
	description := r.FormValue("newcatdesc")

	db, err := sql.Open("sqlite3", "./data/forum.db")
	if err != nil {
		return err
	}
	defer db.Close()

	sqlUpdate := `INSERT INTO category (name, description) VALUES(?, ?)`
	_, err = db.Exec(sqlUpdate, name, description)
	if err != nil {
		return err
	}

	return nil
}

func EditTopicHandler(r *http.Request, topics []models.Topic) error {
	name := r.FormValue("topicname")
	topicID := r.FormValue("topicID")
	stringID := r.FormValue("catID")

	ID, err := strconv.Atoi(topicID)
	if err != nil {
		return nil
	}

	catID, err := strconv.Atoi(stringID)
	if err != nil {
		return nil
	}

	var topic models.Topic
	for _, current := range topics {
		if current.TopicID == ID {
			topic = current
			break
		}
	}

	if name != "" {
		topic.Name = name
	}

	db, err := sql.Open("sqlite3", "./data/forum.db")
	if err != nil {
		log.Print("<admincatsandtopics.go> Erreur à l'ouverture de la base de données :", err)
		return err
	}
	defer db.Close()

	sqlUpdate := `UPDATE topic SET name = ?, category_id = ? WHERE id = ?`
	stmt, err := db.Prepare(sqlUpdate)
	if err != nil {
		log.Print(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(topic.Name, catID, ID)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func DeleteTopicHandler(stringID string) error {
	ID, err := strconv.Atoi(stringID)
	if err != nil {
		log.Print("<admincatsandtopics.go> Erreur dans la récupération du sujet à supprimer", err)
		return err
	}

	fmt.Println(ID)

	db, err := sql.Open("sqlite3", "./data/forum.db")
	if err != nil {
		log.Print("<admincatsandtopics.go> Erreur à l'ouverture de la base de données :", err)
		return err
	}
	defer db.Close()

	sqlUpdate := `DELETE FROM topic WHERE id = ?`
	stmt, err := db.Prepare(sqlUpdate)
	if err != nil {
		log.Print("<admincatsandtopics.go> Erreur dans la suppression du sujet", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(ID)
	if err != nil {
		return err
	}

	err = AdminDeleteMessages(db, ID)
	if err != nil {
		log.Print("<admincatsandtopics.go> Erreur dans la suppression des messagesS", err)
		return err
	}

	log.Print("Sujets et messages supprimés avec succès.")

	return nil
}

package getdata

import (
	"database/sql"
	"log"

	"github.com/Mathis-Pain/Forum/models"
)

func LastMonthPost() ([]models.LastPost, int, error) {
	db, err := sql.Open("sqlite3", "./data/forum.db")
	if err != nil {
		log.Printf("<adminhandler.go> Erreur à l'ouverture de la base de données : %v\n", err)
		return nil, 0, err
	}
	defer db.Close()

	sqlQuery := `
        SELECT
            m.id,
            m.topic_id,
            m.content,
            m.created_at,
            m.user_id,
			u.username,
            t.name
        FROM message m
        JOIN topic t ON m.topic_id = t.id
		JOIN user u ON m.user_id = u.id
		WHERE m.created_at >= DATETIME('now', '-30 days')
        ORDER BY m.created_at DESC
    `
	rows, err := db.Query(sqlQuery)
	if err != nil {
		log.Print("<lastmonthpost.go> Erreur dans la récupération des derniers messages :", err)
		return nil, 0, err
	}
	defer rows.Close()

	var lastMontPosts []models.LastPost
	for rows.Next() {
		var currentPost models.LastPost

		err := rows.Scan(&currentPost.MessageID, &currentPost.TopicID, &currentPost.Content, &currentPost.Created, &currentPost.Author.ID, &currentPost.Author.Username, &currentPost.TopicName)
		if err != nil {
			log.Print("<lastmonthpost.go> Erreur dans le parcours de la base de données :", err)
			return nil, 0, err
		}
		lastMontPosts = append(lastMontPosts, currentPost)
	}

	return lastMontPosts, len(lastMontPosts), nil
}

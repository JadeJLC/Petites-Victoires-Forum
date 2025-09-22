package admin

import (
	"database/sql"
	"log"

	"github.com/Mathis-Pain/Forum/models"
	"github.com/Mathis-Pain/Forum/utils/getdata"
)

func GetAllTopics(categories []models.Category, db *sql.DB) ([]models.Topic, error) {
	var topics []models.Topic

	for i := 0; i < len(categories); i++ {
		topicList, err := getdata.GetTopicList(db, categories[i].ID)
		if err != nil {
			log.Print("<adminhandler.go> Erreur dans la récupération des sujets :", err)
			return nil, err
		}
		topics = append(topics, topicList...)
	}

	return topics, nil
}

func GetStats(topics []models.Topic) ([]models.LastPost, models.Stats, []models.User, error) {
	var stats models.Stats
	var users []models.User
	var err error

	users, stats.TotalUsers, err = GetAllUsers()
	if err != nil {
		log.Print("<displaydashboard.go, GetStats> Erreur dans la récupération des utilisateurs", err)
		return nil, models.Stats{}, nil, err
	}

	stats.LastUser = users[len(users)-1].Username
	stats.TotalTopics = len(topics)

	var lastMonthPosts []models.LastPost
	lastMonthPosts, stats.LastMonthPost, err = getdata.LastMonthPost()

	return lastMonthPosts, stats, users, nil
}

func GetAllUsers() ([]models.User, int, error) {
	db, err := sql.Open("sqlite3", "./data/forum.db")
	if err != nil {
		log.Print("<displaydashboard.go, GetAllUsers> Erreur à l'ouverture de la base de données : ", err)
		return nil, 0, err
	}
	defer db.Close()

	var users []models.User
	var totalUsers int

	sqlQuery := `SELECT MAX(id) FROM user`
	err = db.QueryRow(sqlQuery).Scan(&totalUsers)
	if err != nil && err != sql.ErrNoRows {
		log.Print("<displaydashboard.go, GetAllUsers> Erreur dans la récupération du dernier ID utilisateur : ", err)
		return nil, 0, err
	}

	for i := 1; i <= totalUsers; i++ {
		user, err := getdata.GetUserInfoFromID(db, i)
		if err != nil {
			log.Print("<displaydashboard.go, GetAllUsers> Erreur dans la récupération des données utilisateurs : ", err)
			return nil, 0, err
		}
		users = append(users, user)
	}

	return users, totalUsers, nil
}

package subhandlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Mathis-Pain/Forum/models"
	"github.com/Mathis-Pain/Forum/utils"
	"github.com/Mathis-Pain/Forum/utils/getdata"
)

func BuildHeader(r *http.Request, w http.ResponseWriter, db *sql.DB) ([]models.Category, models.UserLoggedIn, error) {
	categories, err := categoriesDropDownMenu()
	if err != nil {
		log.Print("<buildheader.go> Erreur dans la récupération de la liste des catégories :", err)
		utils.InternalServError(w)
		return nil, models.UserLoggedIn{}, err
	}

	var currentUser models.UserLoggedIn

	currentUser.LogStatus = CheckLogStatus(r)

	if currentUser.LogStatus {
		// Récupère le pseudo et l'ID de l'utilisateur si un utilisateur est en ligne
		currentUser.Username, currentUser.ID, err = utils.GetUserNameAndIDByCookie(r, db)
		if err != nil {
			log.Print("<buildheader.go> Erreur dans la récupération des données utilisateur :", err)
			utils.InternalServError(w)
			return categories, currentUser, err
		}
		return categories, currentUser, nil
	}

	return categories, currentUser, nil

}

// Vérifie si un utilisateur est connecté
func CheckLogStatus(r *http.Request) bool {
	userLoggedIn := false
	_, err := r.Cookie("session_id")
	if err == nil {
		userLoggedIn = true
	}

	return userLoggedIn

}

// Fabrique la liste des catégories pour le menu déroulant
func categoriesDropDownMenu() ([]models.Category, error) {
	categories, err := getdata.GetCatList()
	if err != nil {
		return []models.Category{}, err
	}

	return categories, nil
}

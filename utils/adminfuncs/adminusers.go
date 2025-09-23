package admin

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/Mathis-Pain/Forum/models"
)

func AdminEditUser(users []models.User, url *url.URL) error {
	parts := strings.Split(url.Path, "/")
	stringID := parts[len(parts)-1]

	userID, err := strconv.Atoi(stringID)
	if err != nil {
		return err
	}
	user := users[userID-1]

	fmt.Println(user)

	return nil
}

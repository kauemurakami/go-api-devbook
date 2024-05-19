package users_functions

import (
	"api-social-media/app/core/db"
	responses "api-social-media/app/core/helpers/response"
	"api-social-media/app/models"
	"context"
	"net/http"
	"strings"
)

func GetUsersByNickOrName(w http.ResponseWriter, r *http.Request) {
	conn := db.SetupDB()
	defer conn.Close(context.Background())

	query := "SELECT * FROM users lower(name) LIKE $1 OR lower(nick) LIKE $1"
	var users []models.User
	nameOrNick := "%" + strings.ToLower(r.URL.Query().Get("user")) + "%"
	rows, err := conn.Query(context.Background(),
		query,
		nameOrNick,
	)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Pass, &user.CreatedAt); err != nil {
			responses.Err(w, http.StatusBadRequest, err)

		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)

	}
	responses.JSON(w, http.StatusOK, users)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

}

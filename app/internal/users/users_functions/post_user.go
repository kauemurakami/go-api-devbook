package users_functions

import (
	"api-social-media/app/core/db"
	"api-social-media/app/models"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	conn := db.SetupDB()
	defer conn.Close(context.Background())
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("erro ao receber body")
	}
	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal("ero ao converter usuario em json")
	}
	// query := "insert into users(name,email,pass,nick) values($1,$2,$3,$4)"
	var insertedUser models.User
	err = conn.QueryRow(context.Background(),
		"INSERT INTO users (name, email, pass, nick) VALUES ($1, $2, $3, $4) RETURNING *",
		user.Name,
		user.Email,
		user.Pass,
		user.Nick,
	).Scan(
		&insertedUser.ID,
		&insertedUser.Name,
		&insertedUser.Email,
		&insertedUser.Pass,
		&insertedUser.Nick,
		&insertedUser.CreatedAt,
	)
	if err != nil {
		log.Fatal(err)
		log.Fatalf("erro ao inserir usuario no banco")
	}
	fmt.Printf("Registro inserido: %+v\n", insertedUser)
	u_JSON, err := json.Marshal(insertedUser)
	if err != nil {
		http.Error(w, "Erro ao serializar usuário para JSON", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(u_JSON)
	if err != nil {
		http.Error(w, "Erro ao escrever resposta JSON", http.StatusInternalServerError)
		return
	}
	// w.WriteHeader(http.StatusOK)
	// w.Write(u_json)

}

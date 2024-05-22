package posts_functions

import (
	responses "api-social-media/app/core/helpers/response"
	"api-social-media/app/data/db"
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func DeletePost(w http.ResponseWriter, r *http.Request) {
	// Obter o parâmetro da URL usando o Mux
	postID := mux.Vars(r)["id"]

	// Converter o postID para o tipo UUID
	id, err := uuid.Parse(postID)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, errors.New("ID do post inválido"))
		return
	}

	// Inicializar a conexão com o banco de dados
	conn := db.SetupDB()
	defer conn.Close(context.Background())

	// Consulta SQL para deletar o post pelo ID
	query := `
		DELETE FROM posts
		WHERE id = $1
	`

	// Executar a consulta SQL para deletar o post
	_, err = conn.Exec(context.Background(), query, id)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	// Responder com uma mensagem de sucesso
	responses.JSON(w, http.StatusOK, map[string]string{"message": "Post deletado com sucesso"})
}

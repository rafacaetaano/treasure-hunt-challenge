package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/models"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/service"
)

func CreateUserHandler(svc *service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User

		if err := json.NewDecoder(ctx.Request.Body).Decode(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
			return
		}

		if err := svc.CreateUser(ctx, &user); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
			return
		}

		ctx.JSON(http.StatusCreated, user)
	}
}

// // função Handler que retorna uma Handler func que é o que o Go pede para Handlers
// func CreateUserHandler(svc *service.UserService) http.HandlerFunc {
// 	//função handler que deve ser retornar pela função CreateUserHandler
// 	//vai sempre receber o w que é o canal de resposta para o "usuário" e o r request
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		//aqui criamos a struct User que vai ser preenchida pelo JSON
// 		var user models.User

// 		//fazemos o decode do JSON para struct
// 		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 			http.Error(w, "JSON inválido", http.StatusBadRequest)
// 			return
// 		}

// 		//chamamos a service e criamos o usuário no banco
// 		err := svc.CreateUser(r.Context(), &user)
// 		if err != nil {
// 			http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
// 			return
// 		}

// 		//aqui fixa o statusCode da resposta
// 		w.WriteHeader(http.StatusCreated)

// 		//fazemos o encode para retornar a resposta para o usuário, no caso o JSON da struct
// 		if err := json.NewEncoder(w).Encode(user); err != nil {
// 			http.Error(w, "Erro ao criar resposta", http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }

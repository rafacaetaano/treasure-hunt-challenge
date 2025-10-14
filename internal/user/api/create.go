package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/response"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/api/dto"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/models"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/service"
)

func CreateUserHandler(svc *service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User

		err := json.NewDecoder(ctx.Request.Body).Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewErrorResponse("JSON inválido"))
			return
		}

		err = svc.CreateUser(ctx.Request.Context(), &user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewErrorResponse(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, response.NewSuccessResponse("Usuário criado com sucesso", gin.H{"id": user.ID}))
	}
}

func Login(svc *service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var in dto.LoginRequest

		// 1) Ler o JSON do corpo
		if err := ctx.ShouldBindJSON(&in); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewErrorResponse("JSON inválido"))
			return
		}

		_, err := svc.Login(ctx, in.Username, in.Password)
		if err != nil {
			ctx.JSON(http.StatusNotFound, response.NewErrorResponse("Credenciais inválidas"))
			return
		}
		ctx.JSON(http.StatusOK, response.NewSuccessResponse("Login realizado com sucesso", nil))
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

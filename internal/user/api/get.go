package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/response"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/service"
)

func GetUserByIDHandler(svc *service.UserService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewErrorResponse("ID inválido"))
			return
		}

		user, err := svc.GetUserByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, response.NewErrorResponse("JSON inválido"))
			return
		}

		ctx.JSON(http.StatusOK, response.NewSuccessResponse("Usuário encontrado", user))
	}

}

func GetAllUsers(svc *service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := svc.GetAllUsers(ctx)
		if err != nil {
			ctx.JSON(http.StatusNotFound, response.NewErrorResponse("Erro ao buscar usuários"))
		}

		usersCount := len(users)
		ctx.JSON(http.StatusOK, response.NewSuccessResponse("Usuários encontrados", gin.H{"users": users, "count": usersCount}))
	}
}

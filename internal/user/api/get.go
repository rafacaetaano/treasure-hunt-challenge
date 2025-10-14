package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/response"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/api/dto"
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

		user, err := svc.GetUserByID(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, response.NewErrorResponse("JSON inválido"))
			return
		}
		userResponse := dto.ToUserResponse(user)
		ctx.JSON(http.StatusOK, response.NewSuccessResponse("Usuário encontrado", userResponse))
	}

}

func GetAllUsersHandler(svc *service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := svc.GetAllUsers(ctx.Request.Context())
		if err != nil {
			ctx.JSON(http.StatusNotFound, response.NewErrorResponse("Erro ao buscar usuários"))
			return
		}

		userResponseList := dto.ToUserResponseList(users)
		usersCount := len(userResponseList)
		ctx.JSON(http.StatusOK, response.NewSuccessResponse("Usuários encontrados", gin.H{"users": userResponseList, "count": usersCount}))
	}
}

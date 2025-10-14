package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/response"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/models"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/service"
)

func UpdateUserByIDHandler(svc *service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewErrorResponse("ID inválido"))
			return
		}

		var user models.User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewErrorResponse("JSON inválido"))
			return
		}

		err = svc.UpdateUserByID(ctx.Request.Context(), id, &user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewErrorResponse("Erro ao atualizar usuário"))
			return
		}
		ctx.JSON(http.StatusOK, response.NewSuccessResponse("Usuário atualizado", gin.H{"id": id}))
	}
}

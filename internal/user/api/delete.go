package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/response"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/service"
)

// TODO: ver como passar o erro que estamos usando na service na requisição para retornar quando bater no postman
func DeleteUserByIDHandler(svc *service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewErrorResponse("JSON inválido"))
			return
		}

		err = svc.DeleteUserByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewErrorResponse("Erro ao remover usuário"))
			return
		}
		ctx.JSON(http.StatusOK, response.NewSuccessResponse("Usuário removido com sucesso", nil))
	}
}

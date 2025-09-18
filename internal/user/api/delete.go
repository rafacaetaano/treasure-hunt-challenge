package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/service"
)

func DeleteUserByIDHandler(svc *service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		err = svc.DeleteUserByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar usuário"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Usuário removido com sucesso"})
	}
}

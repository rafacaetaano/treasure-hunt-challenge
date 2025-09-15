package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/service"
)

func GetUserByIDHandler(svc *service.UserService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		user, err := svc.GetUserByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "ID inválido"})
			return
		}

		ctx.JSON(http.StatusOK, user)
	}

}

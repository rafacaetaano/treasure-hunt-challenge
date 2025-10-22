package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/response"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/service"
)

func DeleteUserByIDHandler(svc *service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewErrorResponse("JSON inválido"))
			return
		}

		err = svc.DeleteUserByID(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewErrorResponse(err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, response.NewSuccessResponse("Usuário removido com sucesso", nil))
	}
}

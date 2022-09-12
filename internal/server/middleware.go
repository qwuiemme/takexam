package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func authRequired(ctx *gin.Context, successful func()) {
	if _, err := ctx.Cookie(authCookieName); err == nil {
		successful()
	} else {
		notAuth := action{
			ActionHeader:      "Доступ запрещён!",
			ActionDescription: "Вы не авторизованы в системе.",
		}

		ctx.HTML(http.StatusUnauthorized, "action-complete", notAuth)
	}
}

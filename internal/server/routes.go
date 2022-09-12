package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hnnngn/take-exam/internal/account"
	"github.com/hnnngn/take-exam/internal/task"
)

const authCookieName = "auth"

func DetermineRoutes() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("pages/*")
	router.Static("/public", "public/")

	tasksGroup := router.Group("/tasks")
	{
		tasksGroup.GET("/", handleGETTasks)

		tasksGroup.GET("/create", handleGETTasksCreate)
		tasksGroup.POST("/create", handlePOSTTasksCreate)
	}

	accountGroup := router.Group("/account")
	{
		accountGroup.GET("/create", handleGETAccountCreate)
		accountGroup.POST("/create", handlePOSTAccountCreate)

		accountGroup.GET("/auth", handleGETAccountAuth)
		accountGroup.POST("/auth", handlePOSTAccountAuth)
	}

	return router
}

func handleGETTasks(ctx *gin.Context) {
	authRequired(ctx, func() {
		login, _ := ctx.Cookie(authCookieName)
		tasks := task.GetFromDatabase(login)

		ctx.HTML(http.StatusAccepted, "tasks", tasks)
	})
}

func handleGETTasksCreate(ctx *gin.Context) {
	authRequired(ctx, func() {
		login, _ := ctx.Cookie(authCookieName)

		ctx.HTML(http.StatusAccepted, "create-task", login)
	})
}

func handlePOSTTasksCreate(ctx *gin.Context) {
	authRequired(ctx, func() {
		time, err := time.Parse(task.TimeFormat, ctx.PostForm("complete-before"))

		if err != nil {
			log.Fatal(err)
		}

		task := task.Task{
			BindedTo:       ctx.PostForm("login"),
			CompleteBefore: time,
			IsCompleted:    false,
			Name:           ctx.PostForm("name"),
			Description:    ctx.PostForm("description"),
		}

		task.InsertIntoDatabase()

		taskCreated := action{
			ActionHeader:      "Задача создана!",
			ActionDescription: "Задача была успешно создана.",
		}

		ctx.HTML(http.StatusOK, "action-complete", taskCreated)
	})
}

func handleGETAccountCreate(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "create-account", nil)
}

func handlePOSTAccountCreate(ctx *gin.Context) {
	acc := account.Account{
		Login:      ctx.PostForm("login"),
		Password:   account.GetSHA256Hash(ctx.PostForm("password")),
		AvatarLink: "",
	}

	if acquiredData := account.GetFromDatabase(acc.Login); len(acquiredData) == 0 {
		acc.InsertIntoDatabase()

		regComplete := action{
			ActionHeader:      "Регистрация завершена!",
			ActionDescription: "Регистрация успешно завершена, ваш логин для входа: " + acc.Login,
		}

		ctx.HTML(http.StatusOK, "action-complete", regComplete)
	} else {
		regFailed := action{
			ActionHeader:      "Регистрация не удалась!",
			ActionDescription: "В системе уже найден пользователь с таким логином.",
		}

		ctx.HTML(http.StatusOK, "action-complete", regFailed)
	}
}

func handleGETAccountAuth(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "authorize", nil)
}

func handlePOSTAccountAuth(ctx *gin.Context) {
	if acquiredData := account.GetFromDatabase(ctx.PostForm("login")); len(acquiredData) == 1 {
		if acquiredData[0].Password == account.GetSHA256Hash(ctx.PostForm("password")) {
			ctx.SetCookie(authCookieName, acquiredData[0].Login, int(24*time.Hour), "", "", true, false)

			authComplete := action{
				ActionHeader:      "Авторизация успешна!",
				ActionDescription: "Теперь вам доступны все возможности сервиса takexam.",
			}

			ctx.HTML(http.StatusOK, "action-complete", authComplete)
		} else {
			authFailed := action{
				ActionHeader:      "Авторизация не удалась!",
				ActionDescription: "Неправильно введен пароль.",
			}

			ctx.HTML(http.StatusOK, "action-complete", authFailed)
		}
	} else {
		authFailed := action{
			ActionHeader:      "Авторизация не удалась!",
			ActionDescription: "Неправильно введен логин.",
		}

		ctx.HTML(http.StatusOK, "action-complete", authFailed)
	}
}

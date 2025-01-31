package main

import (
	"news/internal/config"
	"news/internal/controllers"
	"news/internal/database/connection"
	"news/internal/database/repository"
	"news/internal/logger"
	"news/internal/middleware"
	"news/internal/services"

	_ "news/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/subosito/gotenv"
)

const idRule = "<int;min(1);max(2147483647)>"

func init() {
	gotenv.Load()
}

//	@title			News API
//	@version		1.0
//	@description	API for news.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		127.0.0.1:8000
//	@BasePath	/api/v1

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	env := config.NewEnv()
	app := fiber.New()
	conn := connection.Db()
	repo := repository.NewRepo(conn)
	groupsService := services.NewNewsGroupService(&repo)
	groupsController := controllers.NewNewsGroupController(&groupsService)
	articlesService := services.NewNewsArticleService(&repo)
	articlesController := controllers.NewNewsArticleController(&articlesService)
	filesService := services.NewFileUploadService(&repo)
	filesController := controllers.NewFileUploadController(&filesService)

	// статичні файли
	app.Static("/uploads", env.UploadPath)
	// документація
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(middleware.Protected()) // перевірка авторизації через jwt токен

	// групи новин
	groupsGroup := app.Group("/api/v1/groups")
	groupsGroup.Get("/", groupsController.GetNewsGroups)                       // список груп новин
	groupsGroup.Get(idRoute(), groupsController.GetNewsGroup)                  // інформація про групу новин
	groupsGroup.Post("/", groupsController.AddNewsGroup)                       // створення групи новин
	groupsGroup.Put(idRoute(), groupsController.UpdateNewsGroup)               // оновлення групи новин
	groupsGroup.Patch(idRoute()+"/trash", groupsController.TrashNewsGroup)     // м'яке видалення групи новин
	groupsGroup.Patch(idRoute()+"/recover", groupsController.RecoverNewsGroup) // відновлення групи новин
	groupsGroup.Delete(idRoute(), groupsController.DeleteNewsGroup)            // остаточне видалення групи новин

	// завантажені файли
	filesGroup := app.Group("/api/v1/files")
	filesGroup.Get("/", filesController.GetFileUploads)
	filesGroup.Get(idRoute(), filesController.GetFileUpload)
	filesGroup.Post("/", filesController.AddFileUpload)
	filesGroup.Delete(idRoute(), filesController.DeleteFileUpload)

	// статті
	articlesGroup := app.Group("/api/v1/news")
	articlesGroup.Get("/", articlesController.GetNewsArticles)      // список груп новин
	articlesGroup.Get(idRoute(), articlesController.GetNewsArticle) // інформація про групу новин
	articlesGroup.Post("/", articlesController.AddNewsArticle)      // створення групи новин

	// поточний користувач може здійснювати операції тільки над своїми статтями
	app.Use(middleware.CheckAuthor(middleware.Config{
		Service: &articlesService,
	}))
	articlesGroup.Put(idRoute(), articlesController.UpdateNewsArticle)               // оновлення групи новин
	articlesGroup.Patch(idRoute()+"/trash", articlesController.TrashNewsArticle)     // м'яке видалення групи новин
	articlesGroup.Patch(idRoute()+"/recover", articlesController.RecoverNewsArticle) // відновлення групи новин
	articlesGroup.Delete(idRoute(), articlesController.DeleteNewsArticle)            // остаточне видалення групи новин

	logger.Log().Fatal(app.Listen(":" + env.WebPort))
}

func idRoute() string {
	return "/:id" + idRule
}

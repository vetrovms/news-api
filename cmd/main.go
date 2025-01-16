package main

import (
	"news/internal/config"
	"news/internal/controllers"
	"news/internal/database/connection"
	"news/internal/database/repository"
	"news/internal/logger"
	"news/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/subosito/gotenv"
)

const idRule = "<int;min(1);max(2147483647)>"

func init() {
	gotenv.Load()
}

func main() {
	env := config.NewEnv()
	app := fiber.New()
	conn := connection.Db()
	repo := repository.NewRepo(conn)
	gs := services.NewNewsGroupService(&repo)
	groupsController := controllers.NewNewsGroupController(&gs)
	as := services.NewNewsArticleService(&repo)
	articlesController := controllers.NewNewsArticleController(&as)
	fs := services.NewFileUploadService(&repo)
	filesController := controllers.NewFileUploadController(&fs)

	// статті
	articlesGroup := app.Group("/api/v1/news")
	articlesGroup.Get("/", articlesController.GetNewsArticles)                       // список груп новин
	articlesGroup.Get(idRoute(), articlesController.GetNewsArticle)                  // інформація про групу новин
	articlesGroup.Post("/", articlesController.AddNewsArticle)                       // створення групи новин
	articlesGroup.Put(idRoute(), articlesController.UpdateNewsArticle)               // оновлення групи новин
	articlesGroup.Patch(idRoute()+"/trash", articlesController.TrashNewsArticle)     // м'яке видалення групи новин
	articlesGroup.Patch(idRoute()+"/recover", articlesController.RecoverNewsArticle) // відновлення групи новин
	articlesGroup.Delete(idRoute(), articlesController.DeleteNewsArticle)            // остаточне видалення групи новин

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

	// статичні файли
	app.Static("/uploads", env.UploadPath)

	logger.Log().Fatal(app.Listen(":" + env.WebPort))
}

func idRoute() string {
	return "/:id" + idRule
}

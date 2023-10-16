package handler

import (
	h "mongo_fiber/app/handler/v1"

	"github.com/eaciit/dbox"
	"github.com/gofiber/fiber/v2"
)

type HttpHandler struct {
	MongoFiber      *h.MongoFibertHandler
	CommentsHandler *h.MongoCommentsHandler
}

func GetHttpHander(db dbox.IConnection) (*HttpHandler, error) {
	mongoHandler, err := h.NewMongoFiberHandler(db)
	if err != nil {
		return nil, err
	}
	commentsHandler, err := h.NewMongoCommentsHandler(db)
	if err != nil {
		return nil, err
	}

	return &HttpHandler{
		MongoFiber:      mongoHandler,
		CommentsHandler: commentsHandler,
	}, nil
}

func GetRoutes(r *fiber.App, db dbox.IConnection) (*HttpHandler, error) {
	handler, err := GetHttpHander(db)
	if err != nil {
		return nil, err
	}

	cmt := r.Group("/api/v1/cmt")
	art := r.Group("/api/v1/art")

	art.Get("/", handler.MongoFiber.WelcomeHandler)
	art.Get("/articles", handler.MongoFiber.GetArticles)
	art.Post("/article", handler.MongoFiber.CreateArticle)
	art.Put("/article/:id", handler.MongoFiber.UpdateArticle)
	art.Delete("/article/:id", handler.MongoFiber.DeleteArticle)

	cmt.Get("/", handler.CommentsHandler.WelcomeHandler)
	cmt.Post("/comment", handler.CommentsHandler.AddComment)
	cmt.Get("/comments", handler.CommentsHandler.GetAllComments)
	cmt.Get("/comments/:id", handler.CommentsHandler.GetCommentsOnComments)
	cmt.Get("/article/comments/:id", handler.CommentsHandler.GetArticleComments)

	return nil, nil
}

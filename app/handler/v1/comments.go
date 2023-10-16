package v1

import (
	"encoding/json"
	"fmt"
	"mongo_fiber/app/models"
	"mongo_fiber/app/repository"
	v1 "mongo_fiber/app/repository/v1"
	"net/http"
	"strconv"

	"github.com/eaciit/dbox"
	"github.com/gofiber/fiber/v2"
)

type MongoCommentsHandler struct {
	repo repository.MongoCommentsRepo
	db   dbox.IConnection
}

func NewMongoCommentsHandler(db dbox.IConnection) (*MongoCommentsHandler, error) {
	return &MongoCommentsHandler{
		repo: v1.NewMongoCommentsRepo(),
		db:   db,
	}, nil
}

func (m *MongoCommentsHandler) WelcomeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello from comments Handler!")
}

func (m *MongoCommentsHandler) AddComment(c *fiber.Ctx) error {
	comment := &models.Comment{}
	err := json.Unmarshal(c.Request().Body(), &comment)
	if err != nil {
		c.JSON(map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"err":        err.Error(),
		})
	}
	fmt.Println("comments", comment)
	comments, err := m.repo.AddCommentRepo(m.db, comment)
	if err != nil {
		c.JSON(map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"err":        err.Error(),
		})
	}
	fmt.Println("comments", comments)
	return c.JSON(map[string]interface{}{
		"statusCode": http.StatusOK,
		"data":       comments,
	})
}

func (m *MongoCommentsHandler) GetAllComments(c *fiber.Ctx) error {
	comments, err := m.repo.GetAllCommentsRepo(m.db)
	if err != nil {
		c.JSON(map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"err":        err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"statusCode": http.StatusOK,
		"data":       comments,
	})
}

func (m *MongoCommentsHandler) GetCommentsOnComments(c *fiber.Ctx) error {
	id := c.Params("id")
	commentId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"err":        err.Error(),
		})
	}
	comments, err := m.repo.GetCommentsOnCommentsRepo(m.db, commentId)
	if err != nil {
		c.JSON(map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"err":        err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"statusCode": http.StatusOK,
		"data":       comments,
	})
}

func (m *MongoCommentsHandler) GetArticleComments(c *fiber.Ctx) error {
	id := c.Params("id")
	articleId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"err":        err.Error(),
		})
	}
	comments, err := m.repo.GetArticleCommentsRepo(m.db, articleId)
	if err != nil {
		c.JSON(map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"err":        err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"statusCode": http.StatusOK,
		"data":       comments,
	})
}

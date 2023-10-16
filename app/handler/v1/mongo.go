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

type MongoFibertHandler struct {
	repo repository.MongoFiberRepo
	db   dbox.IConnection
}

func NewMongoFiberHandler(db dbox.IConnection) (*MongoFibertHandler, error) {
	return &MongoFibertHandler{
		repo: v1.NewMongoFiberRepo(),
		db:   db,
	}, nil
}

func (m *MongoFibertHandler) WelcomeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func (m *MongoFibertHandler) GetArticles(c *fiber.Ctx) error {
	users, err := m.repo.GetAllArticlesRepo(m.db)
	if err != nil {
		c.JSON(map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"err":        err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"statusCode": http.StatusOK,
		"data":       users,
	})
}

func (m *MongoFibertHandler) CreateArticle(c *fiber.Ctx) error {
	article := &models.Article{}
	err := json.Unmarshal(c.Request().Body(), &article)
	if err != nil {
		c.JSON(map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"err":        err.Error(),
		})
	}
	fmt.Println("articles", article)
	users, err := m.repo.CreateArticleRepo(m.db, article)
	if err != nil {
		c.JSON(map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"err":        err.Error(),
		})
	}
	fmt.Println("users", users)
	return c.JSON(map[string]interface{}{
		"statusCode": http.StatusOK,
		"data":       users,
	})
}

func (m *MongoFibertHandler) UpdateArticle(c *fiber.Ctx) error {
	article := &models.Article{}
	err := json.Unmarshal(c.Request().Body(), &article)
	if err != nil {
		c.JSON(map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"err":        err.Error(),
		})
	}
	fmt.Println("articles", article)
	id := c.Params("id")
	articleId, err := strconv.Atoi(id)
	fmt.Println("article ids", id, articleId)
	if err != nil {
		c.JSON(map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"err":        err.Error(),
		})
	}
	users, err := m.repo.UpdateArticleRepo(m.db, article, articleId)
	if err != nil {
		c.JSON(map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"err":        err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"statusCode": http.StatusOK,
		"data":       users,
	})
}

func (m *MongoFibertHandler) DeleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	articleId, err := strconv.Atoi(id)
	fmt.Println("article ids", id, articleId)
	if err != nil {
		c.JSON(map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"err":        err.Error(),
		})
	}
	err = m.repo.DeleteArticleRepo(m.db, articleId)
	if err != nil {
		c.JSON(map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"err":        err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"statusCode": http.StatusOK,
		"message":    "Article Deleted Successfully",
	})
}

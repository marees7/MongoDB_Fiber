package repository

import (
	"mongo_fiber/app/models"

	"github.com/eaciit/dbox"
)

type MongoFiberRepo interface {
	GetAllArticlesRepo(db dbox.IConnection) ([]map[string]interface{}, error)
	CreateArticleRepo(db dbox.IConnection, req *models.Article) (map[string]interface{}, error)
	UpdateArticleRepo(db dbox.IConnection, req *models.Article, id int) (map[string]interface{}, error)
	DeleteArticleRepo(db dbox.IConnection, id int) error
}

type MongoCommentsRepo interface {
	AddCommentRepo(db dbox.IConnection, req *models.Comment) (map[string]interface{}, error)
	GetAllCommentsRepo(db dbox.IConnection) ([]map[string]interface{}, error)
	GetCommentsOnCommentsRepo(db dbox.IConnection, id int) ([]map[string]interface{}, error)
	GetArticleCommentsRepo(db dbox.IConnection, id int) ([]map[string]interface{}, error)
}

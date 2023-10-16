package repository_v1

import (
	"fmt"
	"mongo_fiber/app/constant"
	"mongo_fiber/app/models"
	"mongo_fiber/app/repository"
	"time"

	"github.com/eaciit/dbox"
	"github.com/eaciit/toolkit"
	log "github.com/sirupsen/logrus"
)

type MongoFiberRepo struct {
}

// Returns new task repository
func NewMongoFiberRepo() repository.MongoFiberRepo {
	return &MongoFiberRepo{}
}

func (MongoFiberRepo) GetAllArticlesRepo(db dbox.IConnection) ([]map[string]interface{}, error) {
	log.WithFields(log.Fields{"service": constant.ServiceName, "ended_at": time.Now()}).Info("GetAllArticlesRepo - started")
	usersData, err := db.NewQuery().From("Articles").Cursor(nil)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	var result []map[string]interface{}
	err = usersData.Fetch(&result, 0, false)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	return result, nil
}

func (MongoFiberRepo) CreateArticleRepo(db dbox.IConnection, req *models.Article) (map[string]interface{}, error) {
	log.WithFields(log.Fields{"service": constant.ServiceName, "ended_at": time.Now()}).Info("GetAllArticlesRepo - started")
	article := toolkit.M{
		"articleId":    req.ArticleID,
		"title":        req.Title,
		"content":      req.Content,
		"nickname":     req.Nickname,
		"creationDate": time.Now(),
	}
	err := db.NewQuery().From("Articles").Insert().Exec(toolkit.M{"data": article, "batch": 1})
	if err != nil {
		return map[string]interface{}{}, err
	}
	return article, nil
}

func (MongoFiberRepo) UpdateArticleRepo(db dbox.IConnection, req *models.Article, id int) (map[string]interface{}, error) {
	log.WithFields(log.Fields{"service": constant.ServiceName, "ended_at": time.Now()}).Info("GetAllArticlesRepo - started")
	article := toolkit.M{
		"articleId":    id,
		"title":        req.Title,
		"content":      req.Content,
		"nickname":     req.Nickname,
		"creationDate": time.Now(),
	}
	filter := dbox.Eq("articleId", id)
	updateData := toolkit.M{"data": article, "batch": 1}
	err := db.NewQuery().From("Articles").Where(filter).Update().Exec(updateData)
	if err != nil {
		fmt.Println("test err", err)
		return map[string]interface{}{}, err
	}
	return article, nil
}

func (MongoFiberRepo) DeleteArticleRepo(db dbox.IConnection, id int) error {
	log.WithFields(log.Fields{"service": constant.ServiceName, "ended_at": time.Now()}).Info("GetAllArticlesRepo - started")
	filter := dbox.Eq("articleId", id)
	updateData := toolkit.M{"data": nil, "batch": 1}
	err := db.NewQuery().From("Articles").Where(filter).Delete().Exec(updateData)
	if err != nil {
		fmt.Println("test err", err)
		return err
	}
	return nil
}

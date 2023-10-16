package repository_v1

import (
	"mongo_fiber/app/constant"
	"mongo_fiber/app/models"
	"mongo_fiber/app/repository"
	"time"

	"github.com/eaciit/dbox"
	"github.com/eaciit/toolkit"
	log "github.com/sirupsen/logrus"
)

type MongoCommentsRepo struct {
}

// Returns new task repository
func NewMongoCommentsRepo() repository.MongoCommentsRepo {
	return &MongoCommentsRepo{}
}

func (MongoCommentsRepo) AddCommentRepo(db dbox.IConnection, req *models.Comment) (map[string]interface{}, error) {
	log.WithFields(log.Fields{"service": constant.ServiceName, "ended_at": time.Now()}).Info("AddComment - started")
	comment := toolkit.M{
		"commentId":       req.CommentID,
		"articleId":       req.ArticleID,
		"content":         req.Content,
		"parentCommentId": req.ParentCommentID,
		"nickname":        req.Nickname,
		"creationDate":    time.Now(),
	}
	err := db.NewQuery().From("Comments").Insert().Exec(toolkit.M{"data": comment, "batch": 1})
	if err != nil {
		return map[string]interface{}{}, err
	}
	return comment, nil
}

func (MongoCommentsRepo) GetAllCommentsRepo(db dbox.IConnection) ([]map[string]interface{}, error) {
	log.WithFields(log.Fields{"service": constant.ServiceName, "ended_at": time.Now()}).Info("GetAllCommentsRepo - started")
	usersData, err := db.NewQuery().From("Comments").Cursor(nil)
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

func (MongoCommentsRepo) GetCommentsOnCommentsRepo(db dbox.IConnection, id int) ([]map[string]interface{}, error) {
	log.WithFields(log.Fields{"service": constant.ServiceName, "ended_at": time.Now()}).Info("GetAllArticlesRepo - started")
	filter := dbox.Eq("parentCommentId", id)
	comments, err := db.NewQuery().From("Comments").Where(filter).Cursor(nil)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	var result []map[string]interface{}
	err = comments.Fetch(&result, 0, false)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	return result, nil
}

func (MongoCommentsRepo) GetArticleCommentsRepo(db dbox.IConnection, id int) ([]map[string]interface{}, error) {
	log.WithFields(log.Fields{"service": constant.ServiceName, "ended_at": time.Now()}).Info("GetAllArticlesRepo - started")
	filter := dbox.Eq("articleId", id)
	comments, err := db.NewQuery().From("Comments").Where(filter).Cursor(nil)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	var result []map[string]interface{}
	err = comments.Fetch(&result, 0, false)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	return result, nil
}

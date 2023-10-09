package db

import (
	"encoding/json"
	"fmt"

	"github.com/glebpepega/new1/internal/models"
	"gorm.io/gorm/clause"
)

type result struct {
	Id         int
	Title      string
	Content    string
	Categories string
}

func (db *DB) UpdateNews(n *models.News, nc *models.NewsCategories) error {
	result := db.DB.Updates(&n)

	if result.RowsAffected == 0 {
		return fmt.Errorf("no such id")
	}

	for _, v := range n.Categories {
		nc.CategoryId = v
		db.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&nc)
	}

	db.DB.Delete(&models.NewsCategories{}, "news_id = ? AND category_id not in ?", nc.NewsId, n.Categories)

	return nil
}

func (db *DB) GetAllNews(pageNum int, pageVolume int) ([]models.News, error) {
	var queryResult []result

	if pageNum != 0 && pageVolume != 0 {
		db.DB.Model(&models.News{}).
			Select("news.id, news.title, news.content, json_agg(news_categories.category_id order by news_categories.category_id) as Categories").
			Limit(pageVolume).
			Offset((pageNum - 1) * pageVolume).
			Joins("left join news_categories on news_categories.news_id = news.id").
			Group("news.id").
			Scan(&queryResult)
	} else {
		db.DB.Model(&models.News{}).
			Select("news.id, news.title, news.content, json_agg(news_categories.category_id order by news_categories.category_id) as Categories").
			Joins("left join news_categories on news_categories.news_id = news.id").
			Group("news.id").
			Scan(&queryResult)
	}

	if len(queryResult) == 0 {
		return nil, fmt.Errorf("no rows found")
	}

	var news []models.News
	var categories []int

	for _, v := range queryResult {
		n := models.News{}
		n.Id = v.Id
		n.Title = v.Title
		n.Content = v.Content

		if err := json.Unmarshal([]byte(v.Categories), &categories); err != nil {
			return nil, err
		}

		n.Categories = categories

		news = append(news, n)
	}

	return news, nil
}

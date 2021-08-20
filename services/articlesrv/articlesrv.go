package articlesrv

import (
	"beego-hananoq-blog/component/errcode"
	"beego-hananoq-blog/component/utils"
	"beego-hananoq-blog/models"
	"sort"
	"strings"
)

type TimeLine struct {
	Year    int               `json:"year"`
	Article []*models.Article `json:"article"`
}

func SearchTag() (articleTags []string) {
	tags := models.SearchTag()
	for _, v := range tags {
		split := strings.Split(v, ",")
		for _, s := range split {
			if s != "" {
				articleTags = append(articleTags, s)
			}
		}
	}
	return utils.RemoveDuplicate(articleTags)
}

func InsertArticle(article *models.Article) (errCode uint) {
	models.InsertArticle(article)
	return errcode.SUCCESS
}

func GetArticleList(pageSize, pageIndex int) (total int, articles []*models.Article) {
	articles = models.List(pageSize, pageIndex)
	total = len(articles)
	return
}

func GetArticleListOrderByCreateTime() []TimeLine {
	articles := models.GetArticleListOrderByCreateTime()
	yearArtMap := make(map[int][]*models.Article)
	for _, article := range articles {
		year := article.CreateTime.Year()
		yearArtMap[year] = append(yearArtMap[year], article)
	}
	keys := make([]int, 0, len(yearArtMap))
	for key := range yearArtMap {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	res := make([]TimeLine, 0, len(keys))
	for _, v := range keys {
		res = append(res, TimeLine{v, yearArtMap[v]})
	}
	return res
}

func GetById(id int) *models.Article {
	return models.GetById(id)
}

func Search(keywords string, pageIndex, pageSize int) (total int, articles []*models.Article) {
	if strings.Trim(keywords, " ") == "" {
		return GetArticleList(pageSize, pageIndex)
	}
	articles = models.Search(keywords, pageIndex, pageSize)
	total = len(articles)
	return
}

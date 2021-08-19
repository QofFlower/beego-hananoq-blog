package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type Article struct {
	Id               int       `json:"id"`
	ArticleHeadPic   string    `json:"articleHeadPic"`
	ArticleName      string    `json:"articleName"`
	ArticleTag       string    `json:"articleTag"`
	ArticleRemark    string    `json:"articleRemark"`
	ArticleReadCount int       `json:"articleReadCount"`
	ArticleState     int       `json:"articleState"`
	ManagerId        int       `json:"managerId"`
	ManagerName      string    `json:"managerName"`
	ArticleContent   string    `json:"articleContent"`
	CreateTime       time.Time `json:"createTime"`
	ArticleType      int       `json:"articleType"`
	ArticleStarNum   int       `json:"articleStarNum"`
	ArticleConNum    int       `json:"articleConNum"`
	Enclosure        string    `json:"enclosure"`
}

func SearchTag() (articleTags []string) {
	sql := `select article_tag from hananoq_blog.blog_article where article_state = 0`
	o := orm.NewOrm()
	if _, err := o.Raw(sql).QueryRows(&articleTags); err != nil {
		logs.Error("Exec SearchTag SQL error:", err)
	}
	return
}

func InsertArticle(article *Article) {
	sql := `insert into hananoq_blog.blog_article(%s) values(%s)`
	article.CreateTime = time.Now()
	cols := []string{"article_head_pic", "article_name", "article_tag", "article_remark",
		"article_read_count", "article_state", "manager_id", "manager_name", "article_content",
		"create_time", "article_type", "article_star_num", "article_con_num", "enclosure"}
	val := make([]string, len(cols))
	for i := 0; i < len(val); i++ {
		val[i] = "?"
	}
	values := []interface{}{
		article.ArticleHeadPic, article.ArticleName, article.ArticleTag, article.ArticleRemark, article.ArticleReadCount,
		article.ArticleState, article.ManagerId, article.ManagerName, article.ArticleContent, article.CreateTime,
		article.ArticleType, article.ArticleStarNum, article.ArticleConNum, article.Enclosure}
	sql = fmt.Sprintf(sql, strings.Join(cols, ","), strings.Join(val, ","))
	fmt.Println(sql)
	o := orm.NewOrm()
	_, err := o.Raw(sql, values).Exec()
	if err != nil {
		logs.Error("Insert Article error: ", err)
		return
	}
}

func List(pageSize, pageIndex int) (articles []*Article) {
	sql := `select * from hananoq_blog.blog_article where article_state = 0 limit ? offset ?`
	o := orm.NewOrm()
	if _, err := o.Raw(sql, pageSize, (pageIndex-1)*pageSize).QueryRows(&articles); err != nil {
		return
	}
	return
}

func GetArticleListOrderByCreateTime() (articles []*Article) {
	sql := `select * from hananoq_blog.blog_article where article_state = 0 order by create_time desc`
	o := orm.NewOrm()
	if _, err := o.Raw(sql).QueryRows(&articles); err != nil {
		logs.Error(err)
		return
	}
	return
}

func GetById(id int) (article *Article) {
	sql := `select * from hananoq_blog.blog_article where id = ? and article_state = 0`
	o := orm.NewOrm()
	if err := o.Raw(sql, id).QueryRow(&article); err != nil {
		logs.Error(err)
		return
	}
	return
}

package model

import (
	"gorm.io/gorm"
	"time"
)

// ArticleLike 文章点赞模型
type ArticleLike struct {
	ID         int64     `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	ArticleID  int       `gorm:"not null;index;comment:文章ID" json:"article_id"`
	UserID     int64     `gorm:"not null;index;comment:用户ID" json:"user_id"`
	CreateTime time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:点赞时间" json:"create_time"`
	UpdateTime time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;on_update:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`
}

func (ArticleLike) TableName() string {
	return "tb_blog_article_like"
}

// BeforeCreate 在创建数据之前触发
func (like *ArticleLike) BeforeCreate(tx *gorm.DB) (err error) {
	like.CreateTime = time.Now()
	like.UpdateTime = time.Now()
	return nil
}

// BeforeUpdate 在更新数据之前触发
func (like *ArticleLike) BeforeUpdate(tx *gorm.DB) (err error) {
	like.UpdateTime = time.Now()
	return nil
}

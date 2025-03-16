package repository

import (
	"errors"
	"gorm.io/gorm"
	"website-gin/config"
	"website-gin/internal/models"
)

type TopicRepository struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) *TopicRepository {
	return &TopicRepository{
		db: config.DB,
	}
}

// CreateTopic 创建 Topic
func (r *TopicRepository) CreateTopic(topic *models.Topic) error {
	return r.db.Create(topic).Error
}

// QueryTopics 根据条件查询 Topic
func (r *TopicRepository) QueryTopics(conditions map[string]interface{}) ([]models.Topic, error) {
	var topics []models.Topic
	query := r.db
	for key, value := range conditions {
		query = query.Where(key+" = ?", value)
	}
	err := query.Find(&topics).Error
	if err != nil {
		return nil, err
	}
	return topics, nil
}

// QueryTopicByID 根据ID查询 Topic
func (r *TopicRepository) QueryTopicByID(id uint) (*models.Topic, error) {
	var topic models.Topic
	query := r.db
	query = query.Where("id = ?", id)
	err := query.First(&topic).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &topic, nil
}

// QueryTopicByTopicName 根据名称查询Topic
func (r *TopicRepository) QueryTopicByTopicName(topicName string) (*models.Topic, error) {
	var topic models.Topic
	query := r.db
	query = query.Where("topic_name = ?", topicName)
	err := query.First(&topic).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &topic, nil
}

// UpdateTopicByID 更新 Topic
func (r *TopicRepository) UpdateTopicByID(id uint, updates map[string]interface{}) error {
	record := r.db.Model(&models.Topic{}).Where("id = ?", id).Updates(updates)
	if record.Error != nil {
		return record.Error
	}
	if record.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// DeleteTopicByID 根据 ID 删除 Topic
func (r *TopicRepository) DeleteTopicByID(id uint) error {
	record := r.db.Where("id = ?", id).Delete(&models.Topic{})
	if record.Error != nil {
		return record.Error
	}
	if record.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

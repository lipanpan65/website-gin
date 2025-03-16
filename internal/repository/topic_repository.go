package repository

import (
	"errors"
	"gorm.io/gorm"
	"website-gin/internal/models"
)

type TopicRepository struct {
	db *gorm.DB
}

// 校验分页参数
func validatePaginationParams(page, pageSize int) (int, int) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	return page, pageSize
}

// 构建查询条件
func buildQuery(query *gorm.DB, conditions map[string]interface{}) *gorm.DB {
	for key, value := range conditions {
		query = query.Where(key+" = ?", value)
	}
	return query
}

func NewTopicRepository(db *gorm.DB) *TopicRepository {
	return &TopicRepository{
		db: db,
	}
}

// CreateTopic 创建 Topic
func (r *TopicRepository) CreateTopic(topic *models.Topic) error {
	return r.db.Create(topic).Error
}

// QueryTopics 根据条件查询 Topic
func (r *TopicRepository) QueryTopics(conditions map[string]interface{}, page, pageSize int) ([]models.Topic, int, error) {
	var topics []models.Topic
	var total int64

	// 校验分页参数
	page, pageSize = validatePaginationParams(page, pageSize)

	// 构建基础查询
	query := r.db

	// 添加查询条件
	query = buildQuery(query, conditions)

	// 查询总记录数
	err := query.Model(&models.Topic{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 应用分页
	err = query.Offset(offset).Limit(pageSize).Find(&topics).Error
	if err != nil {
		return nil, 0, err
	}

	return topics, int(total), nil
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

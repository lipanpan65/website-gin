package repositories

import (
	"website-gin/config"
	"website-gin/models"
)

func CreateTopic(topic models.Topic) (models.Topic, error) {
	// 插入数据库
	if err := config.DB.Create(&topic).Error; err != nil {
		return models.Topic{}, err
	}
	return topic, nil

}

func QueryTopics() ([]models.Topic, error) {
	return nil, nil
}

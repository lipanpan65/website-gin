package services

import (
	"fmt"
	"website-gin/internal/models"
)

func CreateTopic(topicVo models.TopicVo) (*models.TopicVo, error) {

	topic := models.TopicVo{}
	fmt.Println(topic)
	return &models.TopicVo{}, nil
}

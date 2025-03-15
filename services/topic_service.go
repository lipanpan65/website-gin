package services

import (
	"fmt"
	"website-gin/models"
)

func CreateTopic(topicVo models.TopicVo) (*models.TopicVo, error) {

	topic := models.TopicVo{}
	fmt.Println(topic)
	return &models.TopicVo{}, nil
}

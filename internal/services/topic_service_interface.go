package services

import (
	"website-gin/dto/request"
	"website-gin/dto"
)

// TopicServiceInterface 专题服务接口
type TopicServiceInterface interface {
	// CreateTopic 创建 Topic
	CreateTopic(topicDTO *request.TopicDTO) (*dto.TopicVo, error)

	// QueryTopicByID 根据 ID 查询 Topic
	QueryTopicByID(id uint) (*dto.TopicVo, error)

	// QueryTopics 查询多个 Topics，支持分页
	QueryTopics(conditions map[string]interface{}, page, pageSize int) ([]*dto.TopicVo, int64, error)
}

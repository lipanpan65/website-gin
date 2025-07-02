package services

import (
	"website-gin/dto/request"
	"website-gin/internal/model"
	"website-gin/internal/repository"
	"website-gin/internal/vo"
	"website-gin/utils/errors/application"
	"website-gin/utils/errors/infrastructure"
)

// TopicService 专题服务
//type TopicService struct {
//	topicRepo *repository.TopicRepository
//}

// NewTopicService 创建Topic服务实例
//func NewTopicService(topicRepo *repository.TopicRepository) *TopicService {
//	return &TopicService{
//		topicRepo: topicRepo,
//	}
//}

type TopicService struct {
	topicRepo *repository.TopicRepository
}

// NewTopicService 创建 Topic 服务实例，返回 TopicServiceInterface 接口类型
func NewTopicService(topicRepo *repository.TopicRepository) TopicServiceInterface {
	return &TopicService{
		topicRepo: topicRepo,
	}
}

// CreateTopic 创建Topic
func (s *TopicService) CreateTopic(topicDTO *request.TopicDTO) (*vo.TopicVo, error) {
	existingTopic, err := s.topicRepo.QueryTopicByTopicName(topicDTO.TopicName)
	if err != nil {
		return nil, err
	}
	if existingTopic != nil {
		//return nil, fmt.Errorf("topic already exists")
		return nil, application.DataExisted
	}
	topic := &model.Topic{
		TopicName: topicDTO.TopicName,
		Enable:    topicDTO.Enable,
		Remark:    topicDTO.Remark,
	}

	err = s.topicRepo.CreateTopic(topic)
	if err != nil {
		return nil, err
	}

	topicVo := &vo.TopicVo{
		Id:        topic.ID,
		TopicName: topic.TopicName,
		Enabled:   topic.Enable,
		Remark:    topic.Remark,
	}
	return topicVo, nil
}

// QueryTopicByID 根据ID查询 Topic
func (s *TopicService) QueryTopicByID(id uint) (*vo.TopicVo, error) {
	topic, err := s.topicRepo.QueryTopicByID(id)
	if err != nil {
		return nil, err
	}
	if topic == nil {
		return nil, infrastructure.DatabaseError
	}
	topicVo := &vo.TopicVo{
		Id: topic.ID,
	}
	return topicVo, nil
}

func (s *TopicService) QueryTopics(conditions map[string]interface{}, page, pageSize int) ([]*vo.TopicVo, int64, error) {
	return s.topicRepo.QueryTopics(conditions, page, pageSize)
}

package services

import (
	"website-gin/dto/request"
	"website-gin/internal/models"
	"website-gin/internal/repository"
	"website-gin/internal/vo"
	"website-gin/utils/errors/application"
	"website-gin/utils/errors/infrastructure"
)

// TopicService 专题服务
type TopicService struct {
	topicRepo *repository.TopicRepository
}

// NewTopicService 创建Topic服务实例
func NewTopicService(topicRepo *repository.TopicRepository) *TopicService {
	return &TopicService{
		topicRepo: topicRepo,
	}
}

// CreateTopic 创建Topic
func (s *TopicService) CreateTopic(topicDTO *request.TopicDTO) (*vo.TopicVo, error) {
	existingTopic, err := s.topicRepo.QueryTopicByTopicName(topicDTO.Name)
	if err != nil {
		return nil, err
	}
	if existingTopic != nil {
		//return nil, fmt.Errorf("topic already exists")
		return nil, application.ErrorCreateDict
	}
	topic := &models.Topic{
		TopicName: topicDTO.Name,
	}
	err = s.topicRepo.CreateTopic(topic)
	if err != nil {
		return nil, err
	}
	topicVo := &vo.TopicVo{
		Id:        topic.ID,
		TopicName: topic.TopicName,
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

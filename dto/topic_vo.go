package dto

type TopicVo struct {
	Id        int    `json:"id"`
	TopicName string `json:"topic_name"`
	Enabled   int    `json:"enabled"`
	Remark    string `json:"remark"`
}

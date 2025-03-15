package infrastructure

import "website-gin/utils/errors"

var (
	DatabaseError = errors.NewTechnicalError(500, "数据库异常")
)

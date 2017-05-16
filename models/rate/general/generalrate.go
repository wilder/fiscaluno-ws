package general

import (
	"time"
)

type GeneralRate struct {
	Id string
	Description string
	Rate float32
	RatedBy string
	Institution string
	//Course course.Course
	RatedAt time.Time
}

// Institution constructor
func New(id, desc, ratedBy string, rate float32, institution string) *GeneralRate {
	return &GeneralRate{
		Id:id,
		Description:desc,
		Rate:rate,
		RatedBy:ratedBy,
		Institution: institution,
		RatedAt: time.Now(),
	}
}
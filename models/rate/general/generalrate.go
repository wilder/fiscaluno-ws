package general

import (
	"fiscaluno-ws/models/institution"
	"time"
)

type GeneralRate struct {
	Id string
	Description string
	Rate float32
	RatedBy string
	Institution institution.Institution
	//Course course.Course
	RatedAt time.Time
}

// Institution constructor
func New(id, desc, ratedBy string, rate float32, institution institution.Institution) *GeneralRate {
	return &GeneralRate{
		Id:id,
		Description:desc,
		Rate:rate,
		RatedBy:ratedBy,
		Institution: institution,
		RatedAt: time.Now(),
	}
}
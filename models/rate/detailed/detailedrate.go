package specific

import (
	"time"
	"fiscaluno-ws/models/institution"
)

/*
Id
		Instituicao
		Avaliacao (numerico das estrelas)
		Data
		Categoria
		curso
		idAluno
 */

type DetailedRate struct {
	Id string
	Description string
	Category string
	Rate float32
	RatedBy string
	Institution institution.Institution
	//Course course.Course
	RatedAt time.Time
}

// Institution constructor
func New(id, desc, category, ratedBy string, rate float32, institution institution.Institution) *DetailedRate {
	return &DetailedRate{
		Id:id,
		Description:desc,
		Rate:rate,
		Category: category,
		RatedBy:ratedBy,
		Institution: institution,
		RatedAt: time.Now(),
	}
}
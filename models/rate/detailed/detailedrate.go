package specific

import (
	"time"
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
	Institution string
	Course string
	RatedAt time.Time
}

// Institution constructor
func New(id, desc, category, ratedBy string, rate float32, institution string, course string) *DetailedRate {
	return &DetailedRate{
		Id:id,
		Description:desc,
		Rate:rate,
		Category: category,
		RatedBy:ratedBy,
		Institution: institution,
		Course: course,
		RatedAt: time.Now(),
	}
}
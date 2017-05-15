package rate

import (
	"fiscaluno-ws/database/filter"
	"fiscaluno-ws/dao/detailedratedao"
)

func getRatingByUser(filter[] filter.Filter) (interface{}, error) {
	return detailedratedao.GetRatingsByUserId(filter)
}
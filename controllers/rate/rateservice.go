package rate

import (
	"fiscaluno-ws/database/filter"
	"fiscaluno-ws/dao/detailedratedao"
	"fiscaluno-ws/models/rate/general"
	"fiscaluno-ws/controllers/institution"
	"log"
	"fiscaluno-ws/dao/institutiondao"
)

func getRatingByUser(filter[] filter.Filter) (interface{}, error) {
	return detailedratedao.GetRatingsByUserId(filter)
}

func newGeneralRate(rate general.GeneralRate) (error) {
	err := detailedratedao.NewGeneralRate(rate)
	log.Print("rate's institution id: ", rate.Institution)

	if err != nil {
		//updates the rate in the institution node
		institution, _ := institution.GetInstitutionById(rate.Institution)
		//TODO: consider rating count
		institution.Rate += rate.Rate
		err = institutiondao.UpdateInstitution(institution)
	}

	return err
}
package rate

import (
	"github.com/emicklei/go-restful"
	"fiscaluno-ws/database/filter"
	"fiscaluno-ws-bkp/models/rate/general"
)

// Gets all ratings by user id
func RatedBy(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	var filterList = [] filter.Filter{*filter.New("RatedBy", id, "=")}
	ratings, err := getRatingByUser(filterList)
	if err != nil {
		response.WriteEntity(err)
	}else {
		response.WriteEntity(ratings)

	}
}

func GeneralRate(request *restful.Request, response *restful.Response) {
	rate := new(general.GeneralRate)
	err := request.ReadEntity(&rate)
	newGeneralRate(rate)
	if err != nil {
		response.WriteEntity(err)
	} else {
		response.WriteEntity(response.StatusCode())
	}
}

package rate

import (
	"github.com/emicklei/go-restful"
	"fiscaluno-ws/database/filter"
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
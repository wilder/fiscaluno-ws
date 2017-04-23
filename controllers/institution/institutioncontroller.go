package institution

import (
	"github.com/emicklei/go-restful"
)

func New() *restful.WebService {
	service := new(restful.WebService)
	service.
	Path("/institutions").
		Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	service.Route(service.GET("/{institution-id}").To(FindInstitution))
	service.Route(service.GET("/rate/{institution-id}").To(GetInstitutionRate))

	return service
}

// Find Institution by Id
func FindInstitution(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("institution-id")
	response.WriteEntity(findInstitution(id))
}

// Get institution rate by institution Id
func GetInstitutionRate(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("institution-id")
	institution := findInstitution(id)
	response.WriteEntity(institution.Rate)
}

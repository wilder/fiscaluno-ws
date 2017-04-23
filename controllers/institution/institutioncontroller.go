package institution

import (
	"github.com/emicklei/go-restful"
)

type InstitutionService struct {}

// Find Institution by Id
func FindInstitution(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("institution-id")
	response.WriteEntity(findInstitution(id))
}

// Get institution rate by institution Id
func GetInstitutionRate(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("institution-id")
	response.WriteEntity(getRate(id))
}

package institution

import (
    "fiscaluno-ws/dao/institutiondao"
    "fiscaluno-ws/routes/institutionrouter"
)

func NewInstitutionService() *restful.WebService {
    return institutionrouter.New()
}

// Find Institution by Id
func FindInstitution(request *restful.Request, response *restful.Response) {
    id := request.PathParameter("institution-id")
    institution := institutiondao.FindInstitutionById(id)
    response.WriteEntity(institution)
}

// Get institution rate by institution Id
func GetInstitutionRate(request *restful.Request, response *restful.Response) {
    id := request.PathParameter("institution-id")
    institution := institutiondao.FindInstitutionById(id)
    response.WriteEntity(institution.Rate)
}
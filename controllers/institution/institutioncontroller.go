package institution

import (
            "strconv"
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

// Creates new institution at institution node
func NewInstitution(request *restful.Request, response *restful.Response) {
      institution_id := request.PathParameter("institution-id")
      institution_rate, _ := strconv.ParseFloat( request.PathParameter("institution-rate"), 64)
      institution_name := request.PathParameter("institution-name")
      CreateNewInstitution( institution_id, institution_name, float32(institution_rate) )
      response.WriteEntity("Instituição inserida com sucesso")
}

func NewDetailedRateForInstitution(request *restful.Request, response *restful.Response) {
      institution_id := request.PathParameter("institution-id")
      stored_institution, _ := getInstitutionById(institution_id)
      SaveDetailedRateForInstitution(*stored_institution)
      response.WriteEntity("opa")
}
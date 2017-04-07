package institution

import (
    "github.com/emicklei/go-restful"
    "fiscaluno-ws/models/institution"
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
    usr := *institution.New(id, "faculdade impacta", 5.4)
    response.WriteEntity(usr)
}

func GetInstitutionRate(request *restful.Request, response *restful.Response) {
    id := request.PathParameter("institution-id")
    inst := institution.New(id, "faculdade impacta", 5.4)
    response.WriteEntity(inst.Rate)
}
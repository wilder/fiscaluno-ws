package institutionrouter

import (
    "github.com/emicklei/go-restful"
    "fiscaluno-ws/controllers/institution"
)

func New() *restful.WebService {
    service := new(restful.WebService)
    service.
        Path("/institutions").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)
        
    service.Route(service.GET("/{institution-id}").To(institution.FindInstitution))
    service.Route(service.GET("/rate/{institution-id}").To(institution.GetInstitutionRate))
        
    return service
}
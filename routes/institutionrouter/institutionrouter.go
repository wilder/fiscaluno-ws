package institutionrouter

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
package router

import (
    "github.com/emicklei/go-restful"
    "fiscaluno-ws/controllers/institution"
    "fiscaluno-ws/controllers/generic"
)

func New() *restful.WebService {
    service := new(restful.WebService)
    service.
        Path("/ws").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)
        
    service.Route(service.GET("/institution/{institution-id}").To(institution.FindInstitution))
    service.Route(service.GET("/rate/{institution-id}").To(institution.GetInstitutionRate))
    service.Route(service.GET("/institution/new/{institution-id}/{institution-name}/{institution-rate}").To(institution.NewInstitution))

    service.Route(service.GET("/all/{node}").To(generic.FindAll))
        
    return service
}
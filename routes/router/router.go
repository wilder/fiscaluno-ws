package institutionrouter

import (
    "github.com/emicklei/go-restful"
    //"fiscaluno-ws/controllers/institution"
    "fiscaluno-ws/controllers/generic"
)

func New() *restful.WebService {
    service := new(restful.WebService)
    service.
        Path("/all").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)
        
    //service.Route(service.GET("/{institution-id}").To(institution.FindInstitution))
    //service.Route(service.GET("/rate/{institution-id}").To(institution.GetInstitutionRate))

    service.Route(service.GET("/generic/{node}").To(generic.FindAll))
        
    return service
}
package institutionservice

import (
    "github.com/emicklei/go-restful"
)

type Institution struct {
    Id, Name string
}


func New() *restful.WebService {
    service := new(restful.WebService)
    service.
        Path("/institutions").
        Consumes(restful.MIME_XML, restful.MIME_JSON).
        Produces(restful.MIME_XML, restful.MIME_JSON)
        
    service.Route(service.GET("/{institution-id}").To(FindInstitution))
    //service.Route(service.POST("").To(UpdateInstitution))
    //service.Route(service.PUT("/{institution-id}").To(CreateInstitution))
    //service.Route(service.DELETE("/{institution-id}").To(RemoveInstitution))
        
    return service
}


func FindInstitution(request *restful.Request, response *restful.Response) {
    id := request.PathParameter("institution-id")
    // here you would fetch user from some persistence system
    usr := &Institution{Id: id, Name: "Faculdade Impacta"}
    response.WriteEntity(usr)
}
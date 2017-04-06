package institutionservice

import (
    "github.com/emicklei/go-restful"
    "fiscaluno-ws/models"
)

func New() *restful.WebService {
    service := new(restful.WebService)
    service.
        Path("/institutions").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)
        
    service.Route(service.GET("/{institution-id}").To(FindInstitution))
    service.Route(service.GET("/rate/{institution-id}").To(GetInstitutionRate))
    //service.Route(service.POST("").To(UpdateInstitution))
    //service.Route(service.PUT("/{institution-id}").To(CreateInstitution))
    //service.Route(service.DELETE("/{institution-id}").To(RemoveInstitution))
        
    return service
}


func FindInstitution(request *restful.Request, response *restful.Response) {
    id := request.PathParameter("institution-id")
    // here you would fetch user from some persistence system
    usr := models.FindInstitution(id)
    response.WriteEntity(usr)
}

func GetInstitutionRate(request *restful.Request, response *restful.Response) {
    id := request.PathParameter("institution-id")
    institution := models.FindInstitution(id)
    response.WriteEntity(institution.GetRate())
}
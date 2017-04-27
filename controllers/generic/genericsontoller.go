package generic

import "github.com/emicklei/go-restful"

//Get all the values for the node passed as parameter
func FindAll(request *restful.Request, response *restful.Response) {
	node := request.PathParameter("node")
	response.WriteEntity(getAll(node))
}

package controllers

import "github.com/emicklei/go-restful"

func FindAll(request *restful.Request, response *restful.Response) {
	node := request.PathParameter("node")
	response.WriteEntity(getAll(node))
}

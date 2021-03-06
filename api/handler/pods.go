package handler

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/fest-research/IoT-apiserver/api/proxy"
)

type PodService struct {
	proxy proxy.IServerProxy
}

func NewPodService(proxy proxy.IServerProxy) PodService {
	return PodService{proxy: proxy}
}

func (this PodService) Register(ws *restful.WebService) {
	// List pods
	ws.Route(
		ws.Method("GET").
				Path("/pods").
				To(this.listPods).
				Returns(http.StatusOK, "OK", nil).
				Writes(nil),
	)

	// Watch pods
	ws.Route(
		ws.Method("GET").
				Path("/pods").
				To(this.watchPods).
				Returns(http.StatusOK, "OK", nil).
				Writes(nil),
	)
}

func (this PodService) listPods(req *restful.Request, resp *restful.Response) {
	response, err := this.proxy.Get(req)
	if err != nil {
		handleInternalServerError(resp, err)
	}

	resp.AddHeader("Content-Type", "application/json")
	resp.Write(response)
}

func (this PodService) watchPods(req *restful.Request, resp *restful.Response) {
	response, err := this.proxy.Get(req)
	if err != nil {
		handleInternalServerError(resp, err)
	}

	resp.AddHeader("Content-Type", "application/json")
	resp.Write(response)
}

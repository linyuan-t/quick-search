package server

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/linyuan-t/quick-search/pkg/middlewares"
)

//NewPrometheus ..
func NewPrometheus() *middlewares.Prometheus {
	prome := middlewares.NewPrometheus("http")
	prome.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		url := c.Request.URL.Path
		for _, p := range c.Params {
			switch p.Key {
			case "id":
				url = strings.Replace(url, p.Value, ":id", 1)
			case "status":
				url = strings.Replace(url, p.Value, ":status", 1)
			case "tagid":
				url = strings.Replace(url, p.Value, ":tagid", 1)
			case "ext":
				url = strings.Replace(url, p.Value, ":ext", 1)
			}
		}
		return url
	}
	return prome
}

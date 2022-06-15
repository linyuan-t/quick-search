package server

import (
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/linyuan-t/quick-search/backend/services"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// initRouter is used to init router.
func initRouter(s *services.Services) *gin.Engine {
	router := gin.Default()

	// 添加prometheus 监控中间件
	prome := NewPrometheus()
	prome.Use(router)

	// router.Use(loggerMiddware(s.Logger))
	router.Use(cors())
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.GET("/healthcare", healthcare)
	return router
}

func healthcare(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

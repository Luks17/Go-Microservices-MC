package devutils

import (
	"net/http/httptest"

	"github.com/Luks17/Go-Microservices-MC/api"
	"github.com/Luks17/Go-Microservices-MC/db/repository"
	"github.com/gin-gonic/gin"
)

func NewMockServer(store repository.Store) (*gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	server := gin.New()
	server.Use(gin.Recovery())
	api.LoadRouter(server)

	recorder := httptest.NewRecorder()

	return server, recorder
}

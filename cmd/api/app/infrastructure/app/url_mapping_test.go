package app

import (
	"net/http"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/giojimen3z/Rate_Limited_Notification_Service/pkg/routerhandlers"
)

func TestStartApp(t *testing.T) {

	routeUrlPing := ""
	routeMethodPing := ""
	router := routerhandlers.DefaultRouter()

	MapUrls(router)

	var routes []gin.RouteInfo
	for _, r := range router.Routes() {
		if !strings.Contains(r.Path, "/debug") {
			routes = append(routes, r)
			if strings.Contains(r.Path, "/ping") {
				routeUrlPing = r.Path
				routeMethodPing = r.Method
			}

		}
	}

	assert.NotEmpty(t, routes)
	assert.EqualValues(t, http.MethodGet, routeMethodPing)
	assert.EqualValues(t, "/ping", routeUrlPing)
}

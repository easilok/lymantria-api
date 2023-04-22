package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	c "github.com/easilok/lymantria-api/controllers"
	"github.com/easilok/lymantria-api/routes"
	"github.com/easilok/lymantria-api/test"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetCatalogNothing(t *testing.T) {
	db := test.ConnectTestDatabase()

	w := httptest.NewRecorder()
	r := gin.Default()
	gin.SetMode(gin.TestMode)

	controllers := c.NewBaseHandler(db)

	routes.ApiRoutes(r, controllers)

	t.Run("get json data", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
	})
}

func TestRouter(t *testing.T) {
	db := test.ConnectTestDatabase()

	// w := httptest.NewRecorder()
	r := gin.Default()
	gin.SetMode(gin.TestMode)

	controllers := c.NewBaseHandler(db)

	routes.ApiRoutes(r, controllers)

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("get catalog without login", func(t *testing.T) {
		// Make a request to our server with the {base url}/ping
		resp, err := http.Get(fmt.Sprintf("%s/ping", ts.URL))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.StatusCode != 404 {
			t.Fatalf("Expected status code 404, got %v", resp.StatusCode)
		}
	})
}

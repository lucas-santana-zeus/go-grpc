package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-grpc/service/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestingRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func TestGetBlockById(t *testing.T) {
	r := setupTestingRoutes()
	r.GET("/blocks/:id", GetBlockById)

	t.Run("Getting existent block", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/blocks/F1", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)
		var gotBlock models.Block
		_ = json.Unmarshal(res.Body.Bytes(), &gotBlock)
		fmt.Println(&gotBlock)
		assert.Equal(t, models.GetBlockById("F1"), gotBlock)
	})

	t.Run("Getting non existent block", func(t *testing.T) {

	})
}

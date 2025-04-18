package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIssueTokensHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/api/token", IssueTokensHandler())

	t.Run("return 201 when request is valid", func(t *testing.T) {
		body := map[string]string{
			"guid": "6B29FC40-CA47-1067-B31D-00DD010162DA",
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest(http.MethodPost, "/api/token", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		req.RemoteAddr = "192.168.0.1:1234"

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("return 400 when JSON body is invalid", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/api/token", bytes.NewBuffer([]byte("{not json}")))
		req.Header.Set("Content-Type", "application/json")
		req.RemoteAddr = "192.168.0.1:1234"

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("return 400 when GUID is not valid", func(t *testing.T) {
		// невалидный guid
		body := map[string]string{
			"guid": "invalid-213",
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest(http.MethodPost, "/api/token", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		req.RemoteAddr = "192.168.0.1:1234"

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Должна быть ошибка guid невылидный
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("return 400 when GUID is not unique", func(t *testing.T) {
		// не уникальный guid
		body := map[string]string{
			"guid": "6B29FC40-CA47-1067-B31D-00DD010162DA",
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest(http.MethodPost, "/api/token", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		req.RemoteAddr = "192.168.0.1:1234"

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

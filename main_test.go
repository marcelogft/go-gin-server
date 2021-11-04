// Test Service
package main

import (
	"kstarter/service/controller"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	var controller controller.Controller = &controller.DefaultController{}
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	controller.Get(c)

	t.Run("get json data", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
	})
}

func TestGetById(t *testing.T) {
	var controller controller.Controller = &controller.DefaultController{}
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	controller.GetById(c)

	t.Run("get json data", func(t *testing.T) {
		assert.Equal(t, 501, w.Code)
	})
}

func TestCreate(t *testing.T) {
	var controller controller.Controller = &controller.DefaultController{}
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	controller.Create(c)

	t.Run("get json data", func(t *testing.T) {
		assert.Equal(t, 501, w.Code)
	})
}

func TestUpdate(t *testing.T) {
	var controller controller.Controller = &controller.DefaultController{}
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	controller.Update(c)

	t.Run("get json data", func(t *testing.T) {
		assert.Equal(t, 501, w.Code)
	})
}

func TestDelete(t *testing.T) {
	var controller controller.Controller = &controller.DefaultController{}
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	controller.Delete(c)

	t.Run("get json data", func(t *testing.T) {
		assert.Equal(t, 501, w.Code)
	})
}

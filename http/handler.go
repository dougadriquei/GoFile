package http

import (
	"GoFile/controller"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
}

// NewHandler configuração das rotas
func NewHandler() http.Handler {
	handler := &handler{}
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Global middlewares
	router.Use(gin.Logger())
	router.Use(handler.Recovery())
	router = SetRoutesApp(router, handler)
	return router
}

type result struct {
	QuantityInserted int     `json:"quantity_inserted,omitempty"`
	Error            []error `json:"error,omitempty"`
}

func (h *handler) Read(c *gin.Context) {
	fmt.Println("Passou 1")
	pathFile := "test/base_teste.txt"
	count, error := controller.ReadFileController(pathFile)
	data := result{
		QuantityInserted: count,
		Error:            error,
	}
	js, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Finalizou", count)
	fmt.Println("json", fmt.Sprintf("%v", js))
	c.JSON(http.StatusOK, js)
}

//SetRoutesApp seta configuração das rotas
func SetRoutesApp(router *gin.Engine, handler *handler) *gin.Engine {
	v1 := router.Group("/api/v1")
	v1.GET("/read", handler.Read)
	return router
}

func (h *handler) Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"ie-project-back/model"
	"log"
	"net/http"
)

type Handler struct {
	echo *echo.Echo
	db   *model.Database
}

func (handler *Handler) Init(db *model.Database) {
	handler.db = db
	handler.echo = echo.New()
	handler.echo.GET("/api/categories/all", handler.handleGetCategories)
	handler.echo.POST("/api/signup", handler.handleSignup)
	err := handler.echo.Start("127.0.0.1:7000")
	if err != nil {
		return
	}
}

func (handler *Handler) handleGetCategories(context echo.Context) error {
	log.Println(fmt.Sprintf("[Server]: requested for categories"))
	raw := handler.db.GetCategories()
	_json, err := json.Marshal(raw)
	if err != nil {
		log.Println(err)
		return context.String(http.StatusServiceUnavailable, "")
	} else {
		log.Println(fmt.Sprintf("[Server]: categories: %s", string(_json)))
		return context.String(http.StatusOK, string(_json))
	}
}

func (handler *Handler) handleSignup(context echo.Context) error {
	log.Println(fmt.Sprintf("[Server]: requested for signup"))
	var json map[string]interface{} = map[string]interface{}{}
	err := context.Bind(&json)
	if err != nil {
		log.Println(err)
		return context.String(http.StatusBadRequest, "")
	}
	fmt.Println(json)
	return context.String(http.StatusOK, "you have been registered")
}

package handler

import (
	"net/http"
	"github.com/jinzhu/gorm"
	echo "github.com/labstack/echo/v4"
)

type Server struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) http.Handler {
	e := echo.New()
	s := &Server{db: db}
	return e
}

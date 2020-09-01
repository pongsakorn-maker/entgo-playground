package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pongsakorn-maker/entgo-playground/ent"
)

type VideoController struct {
	client *ent.Client
	router gin.IRouter
}

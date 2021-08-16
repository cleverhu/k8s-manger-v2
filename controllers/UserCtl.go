package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type UserCtl struct {
}

func NewUserCtl() *UserCtl {
	return &UserCtl{}
}

func (*UserCtl) Name() string {
	return "UserCtl"
}

func (this *UserCtl) Index(c *gin.Context) goft.Json {
	return gin.H{"message": "this is index"}
}

func (this *UserCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/index", this.Index)
}

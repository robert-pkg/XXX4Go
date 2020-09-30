package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/api"
	http_svr "github.com/robert-pkg/micro-go/rpc/server/http"
)

type loginHandler struct {
}

func (s *loginHandler) SendVerifyCode(c *gin.Context) {

	req := &api.SendVerifyCodeReq{}
	if err := c.BindJSON(req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := http_svr.GetContext(c)
	defer cancel() // terminal routine tree

	resp, err := loginSvc.SendVerifyCode(ctx, req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s *loginHandler) Login(c *gin.Context) {

	req := &api.LoginReq{}
	if err := c.BindJSON(req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := http_svr.GetContext(c)
	defer cancel() // terminal routine tree

	resp, err := loginSvc.Login(ctx, req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s *loginHandler) VerifyToken(c *gin.Context) {

	req := &api.VerifyTokenReq{}
	if err := c.BindJSON(req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := http_svr.GetContext(c)
	defer cancel() // terminal routine tree

	resp, err := loginSvc.VerifyToken(ctx, req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

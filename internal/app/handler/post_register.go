package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *handler) postRegister(c *gin.Context) {
	var req RegisterReq
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"errors": "invalid input"})
		return
	}
	err = h.di.UseCases.UserUC.Register(c, req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"errors": "invalid input"})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "success"})
}

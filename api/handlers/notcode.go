package handlers

import (
	"github.com/Dostonlv/NotCode.git/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Compiler(c *gin.Context) {
	var req models.Req
	var test models.Test
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handlerResponse(c, "error in parse data", http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.storages.Compiler().Compile(req, test)
	if err != nil {
		h.handlerResponse(c, "error in compile response", http.StatusBadRequest, err.Error())
		return
	}
	h.handlerResponse(c, "create notcode", http.StatusOK, res)
}

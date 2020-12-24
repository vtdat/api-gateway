package controllers

import (
	"benchmark/api-gateway/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{
	}
}

func (h *Handler) MakeHandler (g *gin.RouterGroup) {
	group := g.Group("/statistics")
	group.POST("/quantile", h.caclQuantile)
}

// Statistics godoc
// @Summary Calculate Quantile point value
// @Description Calculate Quantile point value
// @Description Validate Information :If input type is not json, return 400 - Bad request;
// @Description Else If not contain any value in pool, return 400 - Bad request;
// @Description Else If not contain valid percentile, return 400 - Bad request;
// @Description Else return status ok with quantile value
// @Accept  json
// @Produce  json
// @tags statistics
// @Param quantile body models.Quantile true "calculate quantile"
// @Success 200 {object} controllers.AppResponse
// @Failure 400 {object} controllers.AppResponse
// @Router /statistics [post]
func (h *Handler) caclQuantile(c *gin.Context)  {
	var quantile models.Quantile

	if err := c.ShouldBindJSON(&quantile); err != nil {
		Abort(c, http.StatusBadRequest, nil, err.Error())
		return
	}

	if len(quantile.Pool) == 0 {
		Abort(c, http.StatusBadRequest, nil, "No value in the pool")
		return
	}

	if len(quantile.Pool) >= 1000 {
		Abort(c, http.StatusBadRequest, nil, "Simulate DDOS attack! Request over 1000 values is not acceptable")
		return
	}

	if quantile.Percentile < 0 || quantile.Percentile > 100 {
		Abort(c, http.StatusBadRequest, nil, "Percentile must be float type in range (0-100]")
		return
	}

	quantileValue := quantile.CalcQuantile()

	desc := fmt.Sprintf("The point which %v%% of values in the pool are less than or equal is %v", quantile.Percentile, quantileValue)
	c.JSON(http.StatusOK, Response(http.StatusOK, quantileValue, desc))
}
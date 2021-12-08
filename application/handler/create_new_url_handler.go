package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/misaghrm/urlshortner/application/model"
	"github.com/misaghrm/urlshortner/pkg/util"
	"net/http"
)

func NewURL(c *fiber.Ctx) error {
	request := model.CreateURLRequest{}
	//response := model.ResponseModel{}
	err := c.BodyParser(&request)
	if err != nil {
		return c.JSON(model.ResponseModel{
			Code:    http.StatusBadRequest,
			Message: "request body is wrong",
		})
	}
	if !util.IsURLValid(request.URL) {
		return c.JSON(model.ResponseModel{
			Code:    http.StatusBadRequest,
			Message: "url is wrong",
		})
	}
	return c.JSON(model.ResponseModel{
		Code: http.StatusOK,
	})
}

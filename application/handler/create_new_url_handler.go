package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/misaghrm/urlshortener/application/model"
	"github.com/misaghrm/urlshortener/pkg/util"
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
	validUrl, err := util.IsURLValid(request.URL)
	if err != nil {
		return c.JSON(model.ResponseModel{
			Code:    http.StatusBadRequest,
			Message: "url is wrong",
		})
	}
	marshal, err := json.Marshal(validUrl)
	if err != nil {
		fmt.Println("err:")
		fmt.Println(err)
		return err
	}
	fmt.Println(string(marshal))
	return c.JSON(model.ResponseModel{
		Code: http.StatusOK,
	})
}

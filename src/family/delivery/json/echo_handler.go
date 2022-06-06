package json

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yanuar-nc/go-boiler-plate/src/family/domain"
	"github.com/yanuar-nc/go-boiler-plate/src/family/usecase"
	"github.com/yanuar-nc/golang/helper"
)

// EchoHandler structure
type EchoHandler struct {
	familyUsecase usecase.FamilyUsecase
}

// NewEchoHandler function
// Returns *EchoHandler
func NewEchoHandler(familyUsecase usecase.FamilyUsecase) *EchoHandler {
	return &EchoHandler{familyUsecase: familyUsecase}
}

// Mount function
// Params : *echo.Group
func (h *EchoHandler) Mount(group *echo.Group) {
	group.POST("", h.Save)
}

// Save handler
func (h *EchoHandler) Save(c echo.Context) error {

	response := new(helper.JSONSchemaTemplate)

	param := domain.Family{}

	err := c.Bind(&param)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		response.Code = http.StatusBadRequest
		response.SetData(helper.Empty{})
		return response.ShowHTTPResponse(c)
	}

	err = h.familyUsecase.Save(c.Request().Context(), param)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		response.Code = http.StatusBadRequest
		response.SetData(helper.Empty{})
		return response.ShowHTTPResponse(c)
	}

	response.Success = true
	response.Message = "Post Family Response"
	response.Code = http.StatusOK

	return response.ShowHTTPResponse(c)
}

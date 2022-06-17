package json

import (
	"github.com/labstack/echo/v4"
	"github.com/yanuar-nc/lineage/src/usecase"
)

// EchoHandler structure
type EchoHandler struct {
	usecase usecase.Usecase
}

// NewEchoHandler function
// Returns *EchoHandler
func NewEchoHandler(usecase usecase.Usecase) *EchoHandler {
	return &EchoHandler{usecase: usecase}
}

// Mount function
// Params : *echo.Group
func (h *EchoHandler) Mount(group *echo.Group) {
	// group.POST("", h.Save)
}

/*
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
*/

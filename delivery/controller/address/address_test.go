package address

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"together/be8/entities"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TEST METHODE CREATE_ADDRESS
func TestCreateAddress(t *testing.T) {
	t.Run("Create Success", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"recipient":    "Galih",
			"hp":           "21343555",
			"street":       "Jl Buntu",
			"subDistrict":  "Bangun Rejo",
			"UrbanVillage": "Pagar Alam Utara",
			"City":         "Pagar Alam",
			"zip":          "23413",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address")
		AddressC := NewControlAddress(&mockAddress{}, validator.New())
		AddressC.CreateAddress()(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 201, result.Code)
		assert.Equal(t, "Success Create Address", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"recipient":    "Galih",
			"hp":           "21343555",
			"street":       "Jl Buntu",
			"subDistrict":  "Bangun Rejo",
			"UrbanVillage": "Pagar Alam Utara",
			"City":         "Pagar Alam",
			"zip":          "23413",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address")
		AddressC := NewControlAddress(&errMockAddress{}, validator.New())
		AddressC.CreateAddress()(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Bind", func(t *testing.T) {
		e := echo.New()
		requestBody := "Jalan Gunung"
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address")
		AddressC := NewControlAddress(&errMockAddress{}, validator.New())
		AddressC.CreateAddress()(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 415, result.Code)
		assert.Equal(t, "Cannot Bind Data", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Validate", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"recipient": "Galih",
			"hp":        "21343555",
			"street":    "Jl Buntu",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address")
		AddressC := NewControlAddress(&errMockAddress{}, validator.New())
		AddressC.CreateAddress()(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Validate Error", result.Message)
		assert.False(t, result.Status)
	})
}

// TEST GET ALL ADDRESS
func TestGetAllAddress(t *testing.T) {
	t.Run("Success Get All Address", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address")
		AddressC := NewControlAddress(&mockAddress{}, validator.New())
		AddressC.GetAllAddress()(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get All data", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address")
		AddressC := NewControlAddress(&errMockAddress{}, validator.New())
		AddressC.GetAllAddress()(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
	})
}

// TEST GET ADDRESS BY ID
func TestGetAddressID(t *testing.T) {
	t.Run("Success Get Address By ID", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		GetAddress := NewControlAddress(&mockAddress{}, validator.New())
		GetAddress.GetAddressID()(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response

		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get Data ID", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Not Found", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		GetAddress := NewControlAddress(&errMockAddress{}, validator.New())
		GetAddress.GetAddressID()(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 404, result.Code)
		assert.Equal(t, "Data Not Found", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Convert ID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		GetAddress := NewControlAddress(&errMockAddress{}, validator.New())
		GetAddress.GetAddressID()(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Cannot Convert ID", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Bind", func(t *testing.T) {
		e := echo.New()
		requestBody := "Jalan Mentari"
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address")
		AddressC := NewControlAddress(&errMockAddress{}, validator.New())
		AddressC.UpdateAddress()(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 415, result.Code)
		assert.Equal(t, "Cannot Bind Data", result.Message)
		assert.False(t, result.Status)
	})
}

// TEST UPDATE ADDRESS BY ID
func TestUpdateAddress(t *testing.T) {
	t.Run("Update Success", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"recipient":    "Galih",
			"hp":           "21343555",
			"street":       "Jl Buntu",
			"subDistrict":  "Bangun Rejo",
			"UrbanVillage": "Pagar Alam Utara",
			"City":         "Pagar Alam",
			"zip":          "23413",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		AddressC := NewControlAddress(&mockAddress{}, validator.New())
		AddressC.UpdateAddress()(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Updated", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Not Found", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		GetAddress := NewControlAddress(&errMockAddress{}, validator.New())
		GetAddress.UpdateAddress()(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 404, result.Code)
		assert.Equal(t, "Data Not Found", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Convert ID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		GetAddress := NewControlAddress(&errMockAddress{}, validator.New())
		GetAddress.UpdateAddress()(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Cannot Convert ID", result.Message)
		assert.False(t, result.Status)
	})
}

// TEST DELETE ADDRESS BY ID
func TestDeleteAddress(t *testing.T) {
	t.Run("Success Delete Address", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		GetAddress := NewControlAddress(&mockAddress{}, validator.New())
		GetAddress.DeleteAddress()(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Deleted", result.Message)
		assert.True(t, result.Status)
	})
	t.Run("Error Delete Address", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		GetAddress := NewControlAddress(&errMockAddress{}, validator.New())
		GetAddress.DeleteAddress()(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Convert ID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		GetAddress := NewControlAddress(&errMockAddress{}, validator.New())
		GetAddress.DeleteAddress()(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Cannot Convert ID", result.Message)
		assert.False(t, result.Status)
	})
}

// MOCK SUCCESS
type mockAddress struct {
}

//METHOD MOCK SUCCESS
func (m *mockAddress) CreateAddress(newAdd entities.Address) (entities.Address, error) {
	return entities.Address{Recipient: "Galih", HP: "123456"}, nil
}
func (m *mockAddress) GetAllAddress() ([]entities.Address, error) {
	return []entities.Address{{Recipient: "Galih", HP: "123456"}, {Recipient: "Nando", HP: "433112"}}, nil
}
func (m *mockAddress) GetAddressID(x uint) (entities.Address, error) {
	return entities.Address{Recipient: "Galih", HP: "123456"}, nil
}
func (m *mockAddress) UpdateAddress(id uint, updatedAddress entities.Address) (entities.Address, error) {
	return entities.Address{Recipient: "Galih", HP: "123456"}, nil
}

func (m *mockAddress) DeleteAddress(id uint) error {
	return nil
}

// MOCK ERROR
type errMockAddress struct {
}

// METHOD MOCK ERROR
func (e *errMockAddress) CreateAddress(newAdd entities.Address) (entities.Address, error) {
	return entities.Address{}, errors.New("Access Database Error")
}

func (e *errMockAddress) GetAllAddress() ([]entities.Address, error) {
	return nil, errors.New("Access Database Error")
}

func (e *errMockAddress) GetAddressID(x uint) (entities.Address, error) {
	return entities.Address{}, errors.New("Access Database Error")
}

func (e *errMockAddress) UpdateAddress(id uint, updatedAddress entities.Address) (entities.Address, error) {
	return entities.Address{}, errors.New("Access Database Error")
}

func (e *errMockAddress) DeleteAddress(id uint) error {
	return errors.New("Access Database Error")
}
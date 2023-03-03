package exception

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	responsemodel "github.com/tensuqiuwulu/bupda-reset-otp/model/http_model/response_model"
)

func ErrorHandler(err error, e echo.Context) {
	errS := ErrorStruct{}
	json.Unmarshal([]byte(err.Error()), &errS)
	if errS.Code != 0 {
		response := responsemodel.Response{Code: errS.Code, Mssg: errS.Mssg, Data: []string{}, Error: errS.Error}
		e.JSON(errS.Code, response)
	} else {
		response := responsemodel.Response{Data: []string{}, Error: []string{"Internal Server Error"}}
		e.JSON(http.StatusInternalServerError, response)
	}
}

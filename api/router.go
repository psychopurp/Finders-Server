package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"unsafe"

	"example.com/example/dao"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

var (
	DB *gorm.DB
)

type PagedResults struct {
	Page         int64       `json:"page"`
	PageSize     int64       `json:"page_size"`
	Data         interface{} `json:"data"`
	TotalRecords int         `json:"total_records"`
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// example for init the database:
//
//  DB, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/employees?charset=utf8&parseTime=true")
//  if err != nil {
//  	panic("failed to connect database: " + err.Error())
//  }
//  defer db.Close()

func ConfigRouter() http.Handler {
	router := httprouter.New()
	configUserInfosRouter(router)
	configUsersRouter(router)

	return router
}

func ConfigGinRouter(router gin.IRoutes) {
	configGinUserInfosRouter(router)
	configGinUsersRouter(router)

	return
}

func ConverHttprouterToGin(f httprouter.Handle) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params httprouter.Params
		_len := len(c.Params)
		if _len == 0 {
			params = nil
		} else {
			params = ((*[1 << 10]httprouter.Param)(unsafe.Pointer(&c.Params[0])))[:_len]
		}

		f(c.Writer, c.Request, params)
	}
}

func readInt(r *http.Request, param string, v int64) (int64, error) {
	p := r.FormValue(param)
	if p == "" {
		return v, nil
	}

	return strconv.ParseInt(p, 10, 64)
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func writeRowsAffected(w http.ResponseWriter, rowsAffected int64) {
	data, _ := json.Marshal(rowsAffected)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func readJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, v)
}

func returnError(w http.ResponseWriter, r *http.Request, err error) {
	status := 0
	switch err {
	case dao.NotFound:
		status = http.StatusBadRequest
	case dao.UnableToMarshalJson:
		status = http.StatusBadRequest
	case dao.UpdateFailedError:
		status = http.StatusBadRequest
	case dao.InsertFailedError:
		status = http.StatusBadRequest
	case dao.DeleteFailedError:
		status = http.StatusBadRequest
	case dao.BadParamsError:
		status = http.StatusBadRequest
	default:
		status = http.StatusBadRequest
	}
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}

	SendJson(w, r, er.Code, er)
}

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

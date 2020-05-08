package api

import (
	"net/http"

	"example.com/example/dao"
	"example.com/example/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
)

func configUserInfosRouter(router *httprouter.Router) {
	router.GET("/userinfos", GetAllUserInfos)
	router.POST("/userinfos", AddUserInfo)
	router.GET("/userinfos/:id", GetUserInfo)
	router.PUT("/userinfos/:id", UpdateUserInfo)
	router.DELETE("/userinfos/:id", DeleteUserInfo)
}

func configGinUserInfosRouter(router gin.IRoutes) {
	router.GET("/userinfos", ConverHttprouterToGin(GetAllUserInfos))
	router.POST("/userinfos", ConverHttprouterToGin(AddUserInfo))
	router.GET("/userinfos/:id", ConverHttprouterToGin(GetUserInfo))
	router.PUT("/userinfos/:id", ConverHttprouterToGin(UpdateUserInfo))
	router.DELETE("/userinfos/:id", ConverHttprouterToGin(DeleteUserInfo))
}

// GetAllUserInfo is a function to get a slice of record(s) from user_infos table in the employees database
// @Summary Get list of UserInfo
// @Tags UserInfo
// @Description GetAllUserInfo is a handler to get a slice of record(s) from user_infos table in the employees database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.UserInfo}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userinfos [get]
func GetAllUserInfos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		returnError(w, r, dao.BadParamsError)
		return
	}

	pagesize, err := readInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		returnError(w, r, dao.BadParamsError)
		return
	}

	order := r.FormValue("order")

	records, totalRows, err := dao.GetAllUserInfos(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetUserInfo is a function to get a single record to user_infos table in the employees database
// @Summary Get record from table UserInfo by id
// @Tags UserInfo
// @ID record id
// @Description GetUserInfo is a function to get a single record to user_infos table in the employees database
// @Accept  json
// @Produce  json
// @Param  id path int true "record id"
// @Success 200 {object} model.UserInfo
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "NotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /userinfos/{id} [get]
func GetUserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	record, err := dao.GetUserInfo(r.Context(), id)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddUserInfo add to add a single record to user_infos table in the employees database
// @Summary Add an record to user_infos table
// @Description add to add a single record to user_infos table in the employees database
// @Tags UserInfo
// @Accept  json
// @Produce  json
// @Param UserInfo body model.UserInfo true "Add UserInfo"
// @Success 200 {object} model.UserInfo
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userinfos [post]
func AddUserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userinfo := &model.UserInfo{}

	if err := readJSON(r, userinfo); err != nil {
		returnError(w, r, dao.BadParamsError)
		return
	}

	if err := userinfo.BeforeSave(); err != nil {
		returnError(w, r, dao.BadParamsError)
	}

	userinfo.Prepare()

	if err := userinfo.Validate(model.Create); err != nil {
		returnError(w, r, dao.BadParamsError)
		return
	}

	var err error
	userinfo, _, err = dao.AddUserInfo(r.Context(), userinfo)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, userinfo)
}

// UpdateUserInfo Update a single record from user_infos table in the employees database
// @Summary Update an record in table user_infos
// @Description Update a single record from user_infos table in the employees database
// @Tags UserInfo
// @Accept  json
// @Produce  json
// @Param  id path int true "Account ID"
// @Param  UserInfo body model.UserInfo true "Update UserInfo record"
// @Success 200 {object} model.UserInfo
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userinfos/{id} [patch]
func UpdateUserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	userinfo := &model.UserInfo{}
	if err := readJSON(r, userinfo); err != nil {
		returnError(w, r, dao.BadParamsError)
		return
	}

	if err := userinfo.BeforeSave(); err != nil {
		returnError(w, r, dao.BadParamsError)
	}

	userinfo.Prepare()

	if err := userinfo.Validate(model.Update); err != nil {
		returnError(w, r, dao.BadParamsError)
		return
	}

	userinfo, _, err := dao.UpdateUserInfo(r.Context(), id, userinfo)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, userinfo)
}

// DeleteUserInfo Delete a single record from user_infos table in the employees database
// @Summary Delete a record from user_infos
// @Description Delete a single record from user_infos table in the employees database
// @Tags UserInfo
// @Accept  json
// @Produce  json
// @Param  id path int true "ID" Format(int64)
// @Success 204 {object} model.UserInfo
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /userinfos/{id} [delete]
func DeleteUserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	rowsAffected, err := dao.DeleteUserInfo(r.Context(), id)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}

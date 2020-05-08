package api

import (
	"net/http"

	"example.com/example/dao"
	"example.com/example/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
)

func configUsersRouter(router *httprouter.Router) {
	router.GET("/users", GetAllUsers)
	router.POST("/users", AddUser)
	router.GET("/users/:id", GetUser)
	router.PUT("/users/:id", UpdateUser)
	router.DELETE("/users/:id", DeleteUser)
}

func configGinUsersRouter(router gin.IRoutes) {
	router.GET("/users", ConverHttprouterToGin(GetAllUsers))
	router.POST("/users", ConverHttprouterToGin(AddUser))
	router.GET("/users/:id", ConverHttprouterToGin(GetUser))
	router.PUT("/users/:id", ConverHttprouterToGin(UpdateUser))
	router.DELETE("/users/:id", ConverHttprouterToGin(DeleteUser))
}

// GetAllUser is a function to get a slice of record(s) from users table in the employees database
// @Summary Get list of User
// @Tags User
// @Description GetAllUser is a handler to get a slice of record(s) from users table in the employees database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.User}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /users [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllUsers(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetUser is a function to get a single record to users table in the employees database
// @Summary Get record from table User by id
// @Tags User
// @ID record id
// @Description GetUser is a function to get a single record to users table in the employees database
// @Accept  json
// @Produce  json
// @Param  id path int true "record id"
// @Success 200 {object} model.User
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "NotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	record, err := dao.GetUser(r.Context(), id)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddUser add to add a single record to users table in the employees database
// @Summary Add an record to users table
// @Description add to add a single record to users table in the employees database
// @Tags User
// @Accept  json
// @Produce  json
// @Param User body model.User true "Add User"
// @Success 200 {object} model.User
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /users [post]
func AddUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := &model.User{}

	if err := readJSON(r, user); err != nil {
		returnError(w, r, dao.BadParamsError)
		return
	}

	if err := user.BeforeSave(); err != nil {
		returnError(w, r, dao.BadParamsError)
	}

	user.Prepare()

	if err := user.Validate(model.Create); err != nil {
		returnError(w, r, dao.BadParamsError)
		return
	}

	var err error
	user, _, err = dao.AddUser(r.Context(), user)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, user)
}

// UpdateUser Update a single record from users table in the employees database
// @Summary Update an record in table users
// @Description Update a single record from users table in the employees database
// @Tags User
// @Accept  json
// @Produce  json
// @Param  id path int true "Account ID"
// @Param  User body model.User true "Update User record"
// @Success 200 {object} model.User
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /users/{id} [patch]
func UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	user := &model.User{}
	if err := readJSON(r, user); err != nil {
		returnError(w, r, dao.BadParamsError)
		return
	}

	if err := user.BeforeSave(); err != nil {
		returnError(w, r, dao.BadParamsError)
	}

	user.Prepare()

	if err := user.Validate(model.Update); err != nil {
		returnError(w, r, dao.BadParamsError)
		return
	}

	user, _, err := dao.UpdateUser(r.Context(), id, user)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, user)
}

// DeleteUser Delete a single record from users table in the employees database
// @Summary Delete a record from users
// @Description Delete a single record from users table in the employees database
// @Tags User
// @Accept  json
// @Produce  json
// @Param  id path int true "ID" Format(int64)
// @Success 204 {object} model.User
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	rowsAffected, err := dao.DeleteUser(r.Context(), id)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}

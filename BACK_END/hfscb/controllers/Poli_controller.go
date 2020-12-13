package controllers


import (
  "github.com/labstack/echo"
  "hsfcb/repository"
  "hsfcb/helper"
  "net/http"
//   "fmt"
//   Model "hsfcb/model"
)



func Get_list_poli(c echo.Context) error {
	  	var each = helper.Response{}
		u := repository.Get_list_poli();
		each.Status = "true"
		each.Message = "Succes"
		each.Data = u
		return c.JSON(http.StatusOK, each)
}


package controllers


import (
  "github.com/labstack/echo"
  "hsfcb/repository"
  "hsfcb/helper"
  "net/http"
//   "fmt"
//   Model "hsfcb/model"
)



func Get_antrian_ticketing(c echo.Context) error {
		var each = helper.Response{}
		id_rs := c.Param("id_rs")
		id_poli := c.Param("id_poli")
		// fmt.print(id_rs)
		// fmt.print(id_poli)
		
		u := repository.Get_ticketing(id_rs,id_poli);
		each.Status = "true"
		each.Message = "Succes"
		each.Data = u
		return c.JSON(http.StatusOK, each)
}


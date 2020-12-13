package controllers


import (
  "github.com/labstack/echo"
  "hsfcb/repository"
  "hsfcb/helper"
  "net/http"
//   "fmt"
//   Model "hsfcb/model"
)


func Get_order_ambulance(c echo.Context) error {
		var each = helper.Response{}
		id := c.Param("id")
  		u := repository.Get_order_ambulance(id);
		each.Status = "true"
		each.Message = "Succes"
		each.Data = u
		return c.JSON(http.StatusOK, each)
}


func Get_order_riwayat_ambulance(c echo.Context) error {
	var each = helper.Response{}
	id := c.Param("id")
	  u := repository.Get_order_riwayat_ambulance(id);
	each.Status = "true"
	each.Message = "Succes"
	each.Data = u
	return c.JSON(http.StatusOK, each)
}
func Update_fasilitas_ambulance_status(c echo.Context) error {
	var each = helper.Response{}
	id := c.Param("id")
	u := repository.Update_fasilitas_ambulance_status(id);
	each.Status = "true"
	each.Message = "Succes"
	each.Data = u
	return c.JSON(http.StatusOK, each)
}



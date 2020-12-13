package controllers


import (
  "github.com/labstack/echo"
  "hsfcb/repository"

  "hsfcb/helper"
  "net/http"
  // Model "hsfcb/model"
)

//create user
// func CreateUser(c echo.Context) error {
//  user := new(models.User)
//  var err error
//  return c.JSON(http.StatusCreated, user)
// }

func Get_detail_user(c echo.Context) error {
  var each = helper.Response{}
       
	// var result []model_data.Mahasiswa
  // result=dao.Get_all_kelas_mahasiswa()
  
  bpjs := c.Param("bpjs")
  
  u := repository.Get_detail_user(bpjs);
  each.Status = "true"
  each.Message = "Succes"
  each.Data = u

	return c.JSON(http.StatusOK, each)
}

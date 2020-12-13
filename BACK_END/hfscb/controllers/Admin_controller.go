package controllers


import (
  "github.com/labstack/echo"
  "hsfcb/repository"
  "hsfcb/helper"
  "net/http"
//   "fmt"
  Model "hsfcb/model"
)


func Login_auth(c echo.Context) error {
  var each = helper.Response{}
  each.Status = "true"
  each.Message = "Succes"
  // each.Data = u

	u := new(Model.Admin)
    if err := c.Bind(u); err != nil {
        return err
  }

//   if  err:=repository.Login_auth(u);err != nil{
//     // fmt.print(err)
//     fmt.Println(err.Error())
//     each.Status = "false"
//     each.Message = "User Not Found"
   
//     return c.JSON(http.StatusOK, each)    
//   }
 err:=repository.Login_auth(u)

  each.Data = err

  return c.JSON(http.StatusOK, each)
}

func Get_list_rs(c echo.Context) error {
	  	var each = helper.Response{}
		u := repository.Get_all_rs();
		each.Status = "true"
		each.Message = "Succes"
		each.Data = u
		return c.JSON(http.StatusOK, each)
}

func Get_count_admin(c echo.Context) error {
	var each = helper.Response{}
  u := repository.Get_count_admin();
  each.Status = "true"
  each.Message = "Succes"
  each.Data = u
  return c.JSON(http.StatusOK, each)
}



// func Get_detail_logistik(c echo.Context) error {
//   id := c.Param("id")
//   var each = helper.Response{}
//   u := repository.Get_detail_logistik(id);
//   each.Status = "true"
//   each.Message = "Succes"
//   each.Data = u
//   return c.JSON(http.StatusOK, each)

// }



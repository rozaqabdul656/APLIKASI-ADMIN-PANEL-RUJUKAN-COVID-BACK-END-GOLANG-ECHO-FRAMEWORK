package controllers


import (
  "github.com/labstack/echo"
  "hsfcb/repository"
  "hsfcb/helper"
  "net/http"
  "fmt"
  Model "hsfcb/model"
)


func Add_fasilitas(c echo.Context) error {
  var each = helper.Response{}
  each.Status = "true"
  each.Message = "Succes"
  // each.Data = u

	u := new(Model.Fasilitas)
    if err := c.Bind(u); err != nil {
        return err
  }

  if  err:=repository.Add_fasilitas(u);err != nil{
    // fmt.print(err)
    fmt.Println(err.Error())
    each.Status = "false"
    each.Message = "Error"
   
    return c.JSON(http.StatusOK, each)    
  }
  return c.JSON(http.StatusOK, each)
}

func Get_all_fasilitas(c echo.Context) error {
	  	id := c.Param("id")
		var each = helper.Response{}
		u := repository.Get_all_fasilitas(id);
		each.Status = "true"
		each.Message = "Succes"
		each.Data = u
		return c.JSON(http.StatusOK, each)
}

func Get_master_fasilitas(c echo.Context) error {
	var each = helper.Response{}
	u := repository.Get_master_fasilitas();
	each.Status = "true"
	each.Message = "Succes"
	each.Data = u
	return c.JSON(http.StatusOK, each)
}

func Detail_fasilitas(c echo.Context) error {
  // User ID from path `users/:id`
  id := c.Param("id")
  
  
  var each = helper.Response{}
  
  u := repository.Detail_fasilitas(id);
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


func Update_fasilitas(c echo.Context) error {
  var each = helper.Response{}
  each.Status = "true"
  each.Message = "Succes"
  // each.Data = u

	u := new(Model.Fasilitas)
    if err := c.Bind(u); err != nil {
        return err
  }

  if  err:=repository.Update_fasilitas(u);err != nil{
    // fmt.print(err)
    fmt.Println(err.Error())
    each.Status = "false"
    each.Message = "Error"
   
    return c.JSON(http.StatusOK, each)    
  }
  return c.JSON(http.StatusOK, each)

}



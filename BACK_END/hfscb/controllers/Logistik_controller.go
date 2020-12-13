package controllers


import (
  "github.com/labstack/echo"
  "hsfcb/repository"
  "hsfcb/helper"
  "net/http"
  "fmt"
  Model "hsfcb/model"
)


func Add_logistik(c echo.Context) error {
  var each = helper.Response{}
  each.Status = "true"
  each.Message = "Succes"
  // each.Data = u

	u := new(Model.Logistik)
    if err := c.Bind(u); err != nil {
        return err
  }

  if  err:=repository.Add_logistik(u);err != nil{
    // fmt.print(err)
    fmt.Println(err.Error())
    each.Status = "false"
    each.Message = "Error"
   
    return c.JSON(http.StatusOK, each)    
  }
  return c.JSON(http.StatusOK, each)
}

func Get_all_logistik(c echo.Context) error {
	  	id := c.Param("id")
		var each = helper.Response{}
		u := repository.Get_all_logistik(id);
		each.Status = "true"
		each.Message = "Succes"
		each.Data = u
		return c.JSON(http.StatusOK, each)
}
func Get_all_master_logistik(c echo.Context) error {
	var each = helper.Response{}
	u := repository.Get_all_master_logistik();
	each.Status = "true"
	each.Message = "Succes"
	each.Data = u
	return c.JSON(http.StatusOK, each)
}
func Get_all_order_logistik(c echo.Context) error {
  var each = helper.Response{}
    id := c.Param("id")

	u := repository.Get_all_order_logistik(id);
	each.Status = "true"
	each.Message = "Succes"
	each.Data = u
	return c.JSON(http.StatusOK, each)
}
func Get_all_riwayat_order_logistik(c echo.Context) error {
  var each = helper.Response{}
    id := c.Param("id")

	u := repository.Get_all_riwayat_order_logistik(id);
	each.Status = "true"
	each.Message = "Succes"
	each.Data = u
	return c.JSON(http.StatusOK, each)
}

func Get_detail_order_logistik(c echo.Context) error {
  var each = helper.Response{}
    id := c.Param("id")

	u := repository.Get_detail_order_logistik(id);
	each.Status = "true"
	each.Message = "Succes"
	each.Data = u
	return c.JSON(http.StatusOK, each)
}


// func Get_detail_berita(c echo.Context) error {
//   // User ID from path `users/:id`
//   id := c.Param("id")
  
//   var each = helper.Response{}
  
//   u := repository.Get_detail_berita(id);
//   each.Status = "true"
//   each.Message = "Succes"
//   each.Data = u

//   return c.JSON(http.StatusOK, each)

// }

// func Update_berita_status(c echo.Context) error {
//   // User ID from path `users/:id`
//   id := c.Param("id_berita")
//   status := c.Param("status")
//   var each = helper.Response{}
//   // fmt.Println(id)
//   // fmt.Println(status)
  
//   u := repository.Update_status(id,status);
//   each.Status = "true"
//   each.Message = "Succes"
//   each.Data = u

//   return c.JSON(http.StatusOK, each)

// }

func Get_detail_logistik(c echo.Context) error {
  id := c.Param("id")
  var each = helper.Response{}
  u := repository.Get_detail_logistik(id);
  each.Status = "true"
  each.Message = "Succes"
  each.Data = u
  return c.JSON(http.StatusOK, each)

}

func Update_logistik(c echo.Context) error {
  var each = helper.Response{}
  each.Status = "true"
  each.Message = "Succes"
  // each.Data = u

	u := new(Model.Logistik)
    if err := c.Bind(u); err != nil {
        return err
  }

  if  err:=repository.Update_logistik(u);err != nil{
    // fmt.print(err)
    fmt.Println(err.Error())
    each.Status = "false"
    each.Message = "Error"
   
    return c.JSON(http.StatusOK, each)    
  }
  return c.JSON(http.StatusOK, each)

}


func Edit_order_logistik(c echo.Context) error {
  var each = helper.Response{}
  each.Status = "true"
  each.Message = "Succes"
	u := new(Model.Order_logistik)
    if err := c.Bind(u); err != nil {
        return err
  }

  if  err:=repository.Edit_order_logistik(u);err != nil{
    // fmt.print(err)
    fmt.Println(err.Error())
    each.Status = "false"
    each.Message = "Error"
   
    return c.JSON(http.StatusOK, each)    
  }
  return c.JSON(http.StatusOK, each)
}



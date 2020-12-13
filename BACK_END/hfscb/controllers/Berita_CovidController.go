package controllers


import (
  "github.com/labstack/echo"
  "hsfcb/repository"

  "hsfcb/helper"
  "net/http"
  "fmt"
  Model "hsfcb/model"
)


func Add_berita(c echo.Context) error {
  var each = helper.Response{}
  each.Status = "true"
  each.Message = "Succes"
  // each.Data = u

	u := new(Model.Berita)
    if err := c.Bind(u); err != nil {
        return err
  }

  if  err:=repository.Add_berita(u);err != nil{
    // fmt.print(err)
            fmt.Println(err.Error())
      each.Status = "false"
      each.Message = "Error"
   
    return c.JSON(http.StatusOK, each)    
  }
  return c.JSON(http.StatusOK, each)
}

func Update_berita(c echo.Context) error {
  var each = helper.Response{}
  each.Status = "true"
  each.Message = "Succes"
  // each.Data = u

	u := new(Model.Berita)
    if err := c.Bind(u); err != nil {
        return err
  }

  if  err:=repository.Update_berita(u);err != nil{
      fmt.Println(err.Error())
      each.Status = "false"
      each.Message = "Error"
   
    return c.JSON(http.StatusOK, each)    
  }
  return c.JSON(http.StatusOK, each)
}

func Get_all_berita(c echo.Context) error {
  var each = helper.Response{}
       
	// var result []model_data.Mahasiswa
  // result=dao.Get_all_kelas_mahasiswa()

  u := repository.Get_all_berita();
  each.Status = "true"
  each.Message = "Succes"
  each.Data = u


  // if err = c.Bind(u); err != nil {
  //     return
  // }
	return c.JSON(http.StatusOK, each)
}
func Get_detail_berita(c echo.Context) error {
  // User ID from path `users/:id`
  id := c.Param("id")
  
  var each = helper.Response{}
  
  u := repository.Get_detail_berita(id);
  each.Status = "true"
  each.Message = "Succes"
  each.Data = u

  return c.JSON(http.StatusOK, each)

}

func Update_berita_status(c echo.Context) error {
  // User ID from path `users/:id`
  id := c.Param("id_berita")
  status := c.Param("status")
  var each = helper.Response{}
  // fmt.Println(id)
  // fmt.Println(status)
  
  u := repository.Update_status(id,status);
  each.Status = "true"
  each.Message = "Succes"
  each.Data = u

  return c.JSON(http.StatusOK, each)

}

func Delete_berita(c echo.Context) error {
  id := c.Param("id_berita")
  var each = helper.Response{}
  u := repository.Delete_berita(id);
  each.Status = "true"
  each.Message = "Succes"
  each.Data = u
  return c.JSON(http.StatusOK, each)

}


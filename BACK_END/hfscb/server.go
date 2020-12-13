package main

import (
	"net/http"
	
	"github.com/labstack/echo"
	"hsfcb/controllers"
)


type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)
type M map[string]interface{}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// e.GET("/users", controllers.GetAllUsers)
	e.GET("/Berita/list_berita_admin", controllers.Get_all_berita)
	e.GET("/Berita/show_edit/id_berita/:id", controllers.Get_detail_berita)
	e.POST("/Berita/add_berita", controllers.Add_berita)
	e.POST("/Berita/update_berita", controllers.Update_berita)
	e.GET("/Berita/edit_berita_status/id_berita/:id_berita/status/:status", controllers.Update_berita_status)
	e.GET("/Berita/delete_berita/id_berita/:id_berita", controllers.Delete_berita)
	
	
	e.GET("/logistik/jenis_logistik", controllers.Get_all_master_logistik)
	e.POST("/Logistik/add_logistik", controllers.Add_logistik)
	e.GET("/logistik/list_logistik/id_user/:id", controllers.Get_all_logistik)
	e.GET("Logistik/show_edit/id_logistik/:id", controllers.Get_detail_logistik)
	e.POST("/Logistik/update_logistik", controllers.Update_logistik)
	
	// 
	e.GET("/fasilitas/list_fasilitas/id_user/:id", controllers.Get_all_fasilitas)
	e.GET("/fasilitas/get_jenis_fasilitas", controllers.Get_master_fasilitas)
	e.POST("/fasilitas/add_fasilitas", controllers.Add_fasilitas)
	
	e.GET("/fasilitas/show_edit/id_fasilitas/:id", controllers.Detail_fasilitas)
	e.POST("/fasilitas/update_fasilitas", controllers.Update_fasilitas)
	
	e.GET("/admin/list_rs", controllers.Get_list_rs)
	
	e.GET("/poli/list_poli", controllers.Get_list_poli)
	e.GET("/ticketing_irj/list_adminantrian_irj/:id_rs/:id_poli", controllers.Get_antrian_ticketing)
	
	e.GET("/user/detail_user/:bpjs", controllers.Get_detail_user)
	
	e.GET("/logistik/list_order_covid/id_user/:id", controllers.Get_all_order_logistik)
	e.GET("/logistik/list_riwayat_order_covid/id_user/:id", controllers.Get_all_riwayat_order_logistik)
	
	

	e.GET("/logistik/detail_logistik/id_logistik/:id", controllers.Get_detail_order_logistik)
	e.POST("/logistik/status_change", controllers.Edit_order_logistik)

	e.GET("/fasilitas/list_order_fasilitas_ambulance/id_rs/:id", controllers.Get_order_ambulance)
	
	e.GET("/fasilitas/riwayat_ambulance_rs/id_rs/:id", controllers.Get_order_riwayat_ambulance)
	
	
	e.GET("/fasilitas/update_fasilitas_ambulance/id_order_ambulance/:id", controllers.Update_fasilitas_ambulance_status)
	e.POST("/admin/admin_login", controllers.Login_auth)
	e.GET("/logistik/get_count_covid_admin/", controllers.Get_count_admin)

	// e.GET("/logistik/get_count_covid_user/", controllers.Get_count_covid_user)

	
	
	

	
	
	

	

	e.GET("/json", func(ctx echo.Context) error {
		data := M{"Message": "Hello", "Counter": 2}
		return ctx.JSON(http.StatusOK, data)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

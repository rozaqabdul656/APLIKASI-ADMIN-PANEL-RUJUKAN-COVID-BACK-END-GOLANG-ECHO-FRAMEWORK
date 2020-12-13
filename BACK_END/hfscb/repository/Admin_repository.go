package repository

import koneksi "hsfcb/helper"
import  Model "hsfcb/model"
import  "fmt"
// import  "context"

func Get_all_rs()([]Model.Admin){
	db, err := koneksi.Connect()
    result := make([]Model.Admin, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("SELECT * FROM tb_admin where level_akses ='1'")
	
    if err != nil {
        fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()


    for rows.Next() {
		var each = Model.Admin{}
		 var err = rows.Scan(&each.Id_admin,&each.Nama,&each.Username,&each.Password,&each.Level_akses,&each.No_telp,&each.Alamat,&each.Email,&each.Url,&each.Foto,&each.Created_at)
		if err != nil {
            // fmt.Println(err.Error())
            return result
     }

        result = append(result, each)
	}

	return result

}

func Login_auth(obj *Model.Admin)([]Model.Admin){
	db, err := koneksi.Connect()
    result := make([]Model.Admin, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("SELECT * FROM tb_admin where username=? and password=?",obj.Username,obj.Password)
	
    if err != nil {
        fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()


    for rows.Next() {
		var each = Model.Admin{}
		 var err = rows.Scan(&each.Id_admin,&each.Nama,&each.Username,&each.Password,&each.Level_akses,&each.No_telp,&each.Alamat,&each.Email,&each.Url,&each.Foto,&each.Created_at)
		if err != nil {
            // fmt.Println(err.Error())
            return result
     }

        result = append(result, each)
	}

	return result

}
// ([]map[string]interface{})
func Get_count_admin()([]interface{}){
	db, err := koneksi.Connect()
    // result := make([]Model.Admin, 0)
    if err != nil {
        // fmt.Println(err.Error())
        // return result
    }
    defer db.Close()
	rows, err := db.Query("SELECT id_admin,nama FROM tb_admin WHERE LEVEL_AKSES = '1'")
	
    if err != nil {
        fmt.Println(err.Error())
        // return result
    }
	
	defer rows.Close()
    
    // result :=[]interface
    result := make([]interface{},0)
    
    for rows.Next() {
		var each = Model.Admin{}
		 var err = rows.Scan(&each.Id_admin,&each.Nama)
		if err != nil {
            fmt.Println(err.Error())
            // return result
     }
            rows_count := db.QueryRow("SELECT count(*) as total_per_rs  FROM tbl_order_logistik   where id_rs=? and hasil='Positif'",each.Id_admin)
            var (

                total_per_rs   string
            )
            
            defer rows.Close()
            if err := rows_count.Scan(&total_per_rs); err != nil {

                fmt.Println(err.Error())
    
            }
        
        //  names:=[]string
        //  names:=append("rs",total_per_rs)
         names:=map[string]string{each.Nama:total_per_rs}
         

         result = append(result,names)

        // result = append(result, each)
    }
    

    // for rows.Next() {
    //     id_rumahsakit_select:=rows.Id_admin;
	// 	fmt.Println(id_rumahsakit)			
	// 	//  var err = rows.Scan(&each.Id_admin,&each.Nama,&each.Username,&each.Password,&each.Level_akses,&each.No_telp,&each.Alamat,&each.Email,&each.Url,&each.Foto,&each.Created_at)
        
    //     //  result = append(rows.nama, )


	// }

	return result

}

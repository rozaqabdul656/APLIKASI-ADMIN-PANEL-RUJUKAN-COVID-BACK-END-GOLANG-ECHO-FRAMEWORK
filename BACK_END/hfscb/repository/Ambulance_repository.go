package repository

import koneksi "hsfcb/helper"
import  Model "hsfcb/model"
import  "fmt"
// import  "context"

func Get_order_ambulance(id_rs string)([]Model.Ambulance){
	db, err := koneksi.Connect()
    result := make([]Model.Ambulance, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("SELECT a.id_order_ambulance,a.no_telp,a.alamat,a.keluhan,b.nama FROM order_ambulance a, tb_user b where a.bpjs = b.bpjs and a.status = '0' and a.id_rs =?",id_rs)
	
    if err != nil {
        fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()


    for rows.Next() {
		var each = Model.Ambulance{}
		 var err = rows.Scan(&each.Id_order_ambulance,&each.No_telp,&each.Alamat,&each.Keluhan,&each.Nama)
		if err != nil {
            // fmt.Println(err.Error())
            return result
     }

        result = append(result, each)
	}

	return result

}

func Get_order_riwayat_ambulance(id_rs string)([]Model.Ambulance){
	db, err := koneksi.Connect()
    result := make([]Model.Ambulance, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("SELECT a.id_order_ambulance,a.no_telp,a.alamat,a.keluhan,b.nama  FROM order_ambulance a, tb_user b where a.bpjs = b.bpjs and a.status = '1' and a.id_rs =?",id_rs)
	
    if err != nil {
        fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()


    for rows.Next() {
		var each = Model.Ambulance{}
		 var err = rows.Scan(&each.Id_order_ambulance,&each.No_telp,&each.Alamat,&each.Keluhan,&each.Nama)
		if err != nil {
            fmt.Println(err.Error())
            return result
     }

        result = append(result, each)
	}

	return result

}





func Update_fasilitas_ambulance_status( id string) error {
    db, err := koneksi.Connect()

    if err != nil {
        fmt.Println(err.Error())
        // return result
	}
    defer db.Close()

        _, err = db.Exec("update  order_ambulance set status='1' where id_order_ambulance=?",id)
   
    if err != nil {
        fmt.Println("Error Query" ,err.Error())
        return err
    }
    // fmt.Println("insert success!")
    return nil
}
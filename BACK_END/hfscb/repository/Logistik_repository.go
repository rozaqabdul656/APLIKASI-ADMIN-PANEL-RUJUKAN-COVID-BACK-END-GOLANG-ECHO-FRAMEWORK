package repository

import koneksi "hsfcb/helper"
import  Model "hsfcb/model"
import  "fmt"
// import  "context"


func Get_all_logistik(id string)([]Model.Logistik){
	db, err := koneksi.Connect()
    result := make([]Model.Logistik, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("select a.id_logistik,a.stok_persediaan,b.nama_logistik from  tb_logistik a ,master_logistik b where a.nama_logistik=b.id_master_logistik and id_user=?",id)
	
    if err != nil {
        fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()

	// var result []Mahasiswa

    for rows.Next() {
		var each = Model.Logistik{}
		 var err = rows.Scan(&each.Id_logistik,&each.Stok_persediaan,&each.Nama_logistik)
		if err != nil {
            // fmt.Println(err.Error())
            return result
     }

        result = append(result, each)
	}

	return result

}

func Get_all_master_logistik()([]Model.Master_Logistik){
	db, err := koneksi.Connect()
    result := make([]Model.Master_Logistik, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("select * from  master_logistik")
	
    if err != nil {
        fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()

	// var result []Mahasiswa

    for rows.Next() {
        var each = Model.Master_Logistik{}
 		var err = rows.Scan(&each.Id_master_logistik,&each.Nama_logistik_master)
		if err != nil {
            // fmt.Println(err.Error())
            return result
     }

        result = append(result, each)
	}

	return result

}



func Get_detail_logistik(id string)([]Model.Logistik){
	db, err := koneksi.Connect()
    result := make([]Model.Logistik, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("select * from  tb_logistik where id_logistik=?",id)
	
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()

	// var result []Mahasiswa

    for rows.Next() {
		var each = Model.Logistik{}
		 var err = rows.Scan(&each.Id_logistik,&each.Nama_logistik,&each.Stok_persediaan,&each.Id_user)
		if err != nil {
            // fmt.Println(err.Error())
            return result
     }

        result = append(result, each)
	}

	return result

}



func Update_logistik( modb * Model.Logistik) error {
    db, err := koneksi.Connect()

    if err != nil {
        // fmt.Println(err.Error())
        // return result
	}
    defer db.Close()

        _, err = db.Exec("update  tb_logistik set stok_persediaan=?  where id_logistik=?",modb.Stok_persediaan,modb.Id_logistik)
   
    if err != nil {
        fmt.Println("Error Query" ,err.Error())
        return err
    }
    // fmt.Println("insert success!")
    return nil
}

func Add_logistik( modb * Model.Logistik) error {
    db, err := koneksi.Connect()

    if err != nil {
        // fmt.Println(err.Error())
        // return result
	}
    defer db.Close()
    _, err = db.Exec("insert into tb_logistik values (?,?,?,?)","",modb.Nama_logistik,modb.Stok_persediaan,modb.Id_user)
    if err != nil {
        // fmt.Println("Error Query" ,err.Error())
        return err
    }
    // fmt.Println("insert success!")
    return nil
}


func Get_all_order_logistik(id_rs string)([]Model.Order_logistik){
	db, err := koneksi.Connect()
    result := make([]Model.Order_logistik, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("SELECT a.id_order_logsitik,b.nama,b.nik,b.bpjs,a.waktu_order FROM tbl_order_logistik a,tb_user b where  a.bpjs=b.bpjs and a.id_rs=? and a.hasil is  null",id_rs)
	// fmt.Println(rows)
    if err != nil {
        fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()

	// var result []Mahasiswa

    for rows.Next() {
        var each = Model.Order_logistik{}
 		var err = rows.Scan(&each.Id_order_logistik,&each.Nama,&each.Nik,&each.Bpjs,&each.Waktu_order)
		if err != nil {
            fmt.Println(err.Error())
            return result
     }

        result = append(result, each)
	}

	return result

}

func Get_all_riwayat_order_logistik(id_rs string)([]Model.Order_logistik){
	db, err := koneksi.Connect()
    result := make([]Model.Order_logistik, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("SELECT a.id_order_logsitik,b.nama,b.nik,b.bpjs,a.waktu_order FROM tbl_order_logistik a,tb_user b where  a.bpjs=b.bpjs and a.id_rs=? and a.hasil is not null",id_rs)
	// fmt.Println(rows)
    if err != nil {
        fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()

	// var result []Mahasiswa

    for rows.Next() {
        var each = Model.Order_logistik{}
 		var err = rows.Scan(&each.Id_order_logistik,&each.Nama,&each.Nik,&each.Bpjs,&each.Waktu_order)
		if err != nil {
            fmt.Println(err.Error())
            return result
     }

        result = append(result, each)
	}

	return result

}

func Get_detail_order_logistik(id_logistik string)([]Model.Order_logistik){
	db, err := koneksi.Connect()
    result := make([]Model.Order_logistik, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("SELECT a.id_order_logsitik,b.nama,b.nik,b.bpjs,a.waktu_order FROM tbl_order_logistik a,tb_user b where  a.bpjs=b.bpjs and a.id_order_logsitik=? and a.hasil is  null",id_logistik)
	// fmt.Println(rows)
    if err != nil {
        fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()

	// var result []Mahasiswa

    for rows.Next() {
        var each = Model.Order_logistik{}
 		var err = rows.Scan(&each.Id_order_logistik,&each.Nama,&each.Nik,&each.Bpjs,&each.Waktu_order)
		if err != nil {
            fmt.Println(err.Error())
            return result
     }

        result = append(result, each)
	}

	return result

}

func Edit_order_logistik(Orl *Model.Order_logistik)error{
    db, err := koneksi.Connect()

    if err != nil {
        // fmt.Println(err.Error())
        // return result
	}
    defer db.Close()

        _, err = db.Exec("update  tbl_order_logistik set hasil=? , tindakan=? , report=?  where id_order_logsitik=?",Orl.Hasil,Orl.Tindakan,Orl.Report,Orl.Id_order_logistik)
   
    if err != nil {
        fmt.Println("Error Query" ,err.Error())
        return err
    }
    // fmt.Println("insert success!")
    return nil
}



package repository

import koneksi "hsfcb/helper"
import  Model "hsfcb/model"
import  "fmt"
// import  "context"


func Get_all_fasilitas(id string)([]Model.Fasilitas){
	db, err := koneksi.Connect()
    result := make([]Model.Fasilitas, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("select a.*,b.* from  tb_fasilitas a ,master_fasilitas b where a.jenis_fasilitas=b.id_master_fasilitas and a.id_user=?",id)
	
    if err != nil {
        fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()


    for rows.Next() {
		var each = Model.Fasilitas{}
		 var err = rows.Scan(&each.Id_fasilitas,&each.Jenis_fasilitas,&each.Persediaan,&each.Id_user,&each.Id_master_fasilitas,&each.Nama_fasilitas)
		if err != nil {
            // fmt.Println(err.Error())
            return result
     }

        result = append(result, each)
	}

	return result

}

func Get_master_fasilitas()([]Model.Master_fasilitas){
	db, err := koneksi.Connect()
    result := make([]Model.Master_fasilitas, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("select * from  master_fasilitas")
	
    if err != nil {
        fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()

	// var result []Mahasiswa

    for rows.Next() {
        var each = Model.Master_fasilitas{}
 		var err = rows.Scan(&each.Id_master_fasilitas,&each.Nama_fasilitas)
		if err != nil {
            // fmt.Println(err.Error())
            return result
     }

        result = append(result, each)
	}

	return result

}



func Detail_fasilitas(id string)([]Model.Fasilitas){
	db, err := koneksi.Connect()
    result := make([]Model.Fasilitas, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("select * from  tb_fasilitas where id_fasilitas=?",id)
	
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()

	// var result []Mahasiswa

    for rows.Next() {
		var each = Model.Fasilitas{}
		 var err = rows.Scan(&each.Id_fasilitas,&each.Jenis_fasilitas,&each.Persediaan,&each.Id_user)
		if err != nil {
            // fmt.Println(err.Error())
            return result
     }

        result = append(result, each)
	}

	return result

}



func Update_fasilitas( modb * Model.Fasilitas) error {
    db, err := koneksi.Connect()

    if err != nil {
        // fmt.Println(err.Error())
        // return result
	}
    defer db.Close()

        _, err = db.Exec("update  tb_fasilitas set persediaan=?  where id_fasilitas=?",modb.Persediaan,modb.Id_fasilitas)
   
    if err != nil {
        fmt.Println("Error Query" ,err.Error())
        return err
    }
    // fmt.Println("insert success!")
    return nil
}

func Add_fasilitas( modb * Model.Fasilitas) error {
    db, err := koneksi.Connect()

    if err != nil {
        // fmt.Println(err.Error())
        // return result
	}
    defer db.Close()
    _, err = db.Exec("insert into tb_fasilitas values (?,?,?,?)","",modb.Jenis_fasilitas,modb.Persediaan,modb.Id_user)
    if err != nil {
        // fmt.Println("Error Query" ,err.Error())
        return err
    }
    // fmt.Println("insert success!")
    return nil
}


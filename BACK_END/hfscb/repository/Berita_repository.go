package repository

import koneksi "hsfcb/helper"
import  Model "hsfcb/model"
import  "fmt"
// import  "context"


func Get_all_berita()([]Model.Berita){
	db, err := koneksi.Connect()
    result := make([]Model.Berita, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("select * from  berita_covid")
	
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()

	// var result []Mahasiswa

    for rows.Next() {
		var each = Model.Berita{}
		 var err = rows.Scan(&each.Id_berita,&each.Judul,&each.Cover,&each.Content,&each.Tanggal,&each.Status)
		if err != nil {
            // fmt.Println(err.Error())
            return result
     }

        result = append(result, each)
	}
	// fmt.Println("rza")
	// return "aasdadas"
	return result
	// return 
	//   u.SetId(userName)
	//   return u
}


func Get_detail_berita(id string)([]Model.Berita){
	db, err := koneksi.Connect()
    result := make([]Model.Berita, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("select * from  berita_covid where id_berita=?",id)
	
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()

	// var result []Mahasiswa

    for rows.Next() {
		var each = Model.Berita{}
		 var err = rows.Scan(&each.Id_berita,&each.Judul,&each.Cover,&each.Content,&each.Tanggal,&each.Status)
		if err != nil {
            // fmt.Println(err.Error())
            return result
     }

        result = append(result, each)
	}
	// fmt.Println("rza")
	// return "aasdadas"
	return result
	// return 
	//   u.SetId(userName)
	//   return u
}



func Update_berita( modb * Model.Berita) error {
    db, err := koneksi.Connect()

    if err != nil {
        // fmt.Println(err.Error())
        // return result
	}
    defer db.Close()
    if modb.Cover == ""{
        _, err = db.Exec("update  berita_covid set judul=?,content=? where id_berita=?",modb.Judul,modb.Content,modb.Id_berita)
   
    }else{
        _, err = db.Exec("update  berita_covid set judul=? , cover=?,content=? where id_berita=?",modb.Judul,modb.Cover,modb.Content,modb.Id_berita)
   
    }
    if err != nil {
        fmt.Println("Error Query" ,err.Error())
        return err
    }
    // fmt.Println("insert success!")
    return nil
}

func Add_berita( modb * Model.Berita) error {
    db, err := koneksi.Connect()

    if err != nil {
        // fmt.Println(err.Error())
        // return result
	}
    defer db.Close()
    _, err = db.Exec("insert into berita_covid values (?,?,?,?,?,?)","",modb.Judul,modb.Cover,modb.Content,modb.Tanggal,modb.Status)
    if err != nil {
        // fmt.Println("Error Query" ,err.Error())
        return err
    }
    // fmt.Println("insert success!")
    return nil
}



func Update_status( id string,status string) error {
    db, err := koneksi.Connect()

    if err != nil {
        fmt.Println(err.Error())
        // return result
	}
    defer db.Close()

        _, err = db.Exec("update  berita_covid set status=? where id_berita=?",status,id)
   
    if err != nil {
        fmt.Println("Error Query" ,err.Error())
        return err
    }
    // fmt.Println("insert success!")
    return nil
}
func Delete_berita( id string) error {
    db, err := koneksi.Connect()

    if err != nil {
        fmt.Println(err.Error())
        // return result
	}
    defer db.Close()

        _, err = db.Exec("delete from berita_covid where id_berita=?",id)
   
    if err != nil {
        fmt.Println("Error Query" ,err.Error())
        return err
    }
    // fmt.Println("insert success!")
    return nil
}

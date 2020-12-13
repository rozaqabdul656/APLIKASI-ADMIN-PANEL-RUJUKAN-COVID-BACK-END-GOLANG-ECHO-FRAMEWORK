package repository

import koneksi "hsfcb/helper"
import  Model "hsfcb/model"
import  "fmt"
// import  "context"


func Get_detail_user(bpjs string)([]Model.User){
	db, err := koneksi.Connect()
    result := make([]Model.User, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("select * from tb_user where bpjs=?",bpjs)
    if err != nil {
        fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()

	// var result []Mahasiswa

    for rows.Next() {
        var each = Model.User{}
        var err = rows.Scan(&each.Id_user,&each.Bpjs,&each.Nama,&each.Username,&each.Password,&each.Id_rs,&each.Nik,&each.Alamat,&each.Tanggal_lahir,&each.No_hp)
	    if err != nil {
            // fmt.Println(err.Error())
            return result
     }

        result = append(result, each)
	}
	return result

}

// // func Get_all_kelas_mahasiswa()([]Model.Mahasiswa){
// // 	db, err := koneksi.Connect()

// //     if err != nil {
// //         // fmt.Println(err.Error())
// //         // return result
// // 	}
	
// //     defer db.Close()
// // 	rows, err := db.Query("select a.*,b.* from tb_mahasiswa a,tb_kelas b where a.id=b.id_mahasiswa")
// // 	// rows,err :=db.QueryRow("select a.*,b.* from tb_mahasiswa a,tb_kelas b where a.id=b.id_mahasiswa")
// // 	result := make([]Model.Mahasiswa,0)
// //     if err != nil {
// //         // fmt.Println(err.Error())
// //         return result
// //     }
	
// // 	defer rows.Close()

	
// // 	// var result []Mahasiswa

// //     for rows.Next() {
// // 		var each = Model.Mahasiswa{}

// // 		var err = rows.Scan(&each.Id,&each.Nama,&each.Umur,&each.Id_kelas,&each.Id_mahasiswa,&each.Kelas_nama)
// // 		// var err = rows.StructScan(&each)
// // 		// var err = StructScan(rows, &each)

		
// // 		// .Id,&each.Nama,&each.Umur,&each.Id_kelas,&each.Id_mahasiswa,&each.Kelas_nama
// // 		// fmt.Println(&each.Nama)
// //         if err != nil {
// //             // fmt.Println(err.Error())
// //             return result
// //      }

// //         result = append(result, each)
// // 	}
// // 	// fmt.Println("rza")
// // 	// return "aasdadas"
// // 	return result
// // 	// return 
// // 	//   u.SetId(userName)
// // 	//   return u
// // }

// // // func get_all_stuctscan()([]Model.Mahasiswa){
// // // 	db, err := koneksi.Connect()

// // //     if err != nil {
// // //         // fmt.Println(err.Error())
// // //         // return result
// // // 	}
	
// // //     defer db.Close()
	
// // // 	rows, _ := db.Query("select a.*,b.* from tb_mahasiswa a,tb_kelas b where a.id=b.id_mahasiswa")
// // // 	defer rows.Close()
// // // 	s := Model.Mahasiswa{}
// // // 	err = structScan(rows, &s)
// // // 	var result []Model.Mahasiswa
// // // 	result = append(result, s)
	
// // // 	return result
// // // }

// // func Add_mhs( mhs Model.Mahasiswa) error {
// //     db, err := koneksi.Connect()

// //     if err != nil {
// //         // fmt.Println(err.Error())
// //         // return result
// // 	}
// //     defer db.Close()
// //     _, err = db.Exec("insert into tb_mahasiswa values (?,?, ?)","", mhs.Nama, mhs.Umur)
// //     if err != nil {
// //         fmt.Println("Error Query" ,err.Error())
// //         return err
// //     }
// //     fmt.Println("insert success!")
// //     return nil
// // }

// // func Delete_mhs( id string) error {
// //     db, err := koneksi.Connect()

// //     if err != nil {
// //         // fmt.Println(err.Error())
// //         return err
// // 	}
// //     defer db.Close()
// // 	 _, err = db.Exec("DELETE from tb_mahasiswa where id=?",id)
// //     if err != nil {
// //         fmt.Println("Error Query" ,err.Error())
// //         return err
// //     }
// //     fmt.Println("Delete success!")
// //     return nil
// // }



// // func Update_mhs( mhs Model.Mahasiswa) error {
// //     db, err := koneksi.Connect()

// //     if err != nil {
// //         fmt.Println("Koneksi" ,err.Error())
// //         return err
// // 	}
// //     defer db.Close()
    
// //     _, err = db.Exec("update tb_mahasiswa set nama = ?,umur=? where id = ?", mhs.Nama,mhs.Umur,mhs.Id)
// //     if err != nil {
// //         fmt.Println("Error Query" ,err.Error())
// //         return err
// //     }
// //     fmt.Println("Update  success!")
// //     return nil
// // }
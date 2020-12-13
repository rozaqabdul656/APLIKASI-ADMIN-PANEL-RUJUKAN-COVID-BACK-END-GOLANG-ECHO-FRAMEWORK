package repository

import koneksi "hsfcb/helper"
import  Model "hsfcb/model"
// import  "fmt"
// import  "context"


func Get_list_poli()([]Model.Poli){
	db, err := koneksi.Connect()
    result := make([]Model.Poli, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("select * from  tb_poli")
	
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()

	// var result []Mahasiswa

    for rows.Next() {
		var each = Model.Poli{}
		 var err = rows.Scan(&each.Id_poli,&each.Nama_poli)
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




package repository

import koneksi "hsfcb/helper"
import  Model "hsfcb/model"
// import  "fmt"
// import  "context"


func Get_ticketing(id_rs string,id_poli string )([]Model.Ticketing_irj){
	db, err := koneksi.Connect()
    result := make([]Model.Ticketing_irj, 0)
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
    defer db.Close()
	rows, err := db.Query("SELECT a.id_ticketing_irj,a.antrian,c.nama,c.bpjs,a.time FROM tbl_ticketing_irj a,tb_poli b,tb_user c  WHERE a.id_poli=b.id_poli and a.bpjs=c.bpjs and a.id_rs=?  and b.id_poli=? AND a.status is NULL group by a.antrian",id_rs,id_poli)
	
    if err != nil {
        // fmt.Println(err.Error())
        return result
    }
	
	defer rows.Close()

	// var result []Mahasiswa

    for rows.Next() {
		var each = Model.Ticketing_irj{}
		 var err = rows.Scan(&each.Id_ticketing_irj,&each.Antrian,&each.Nama,&each.Bpjs,&each.Time)
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




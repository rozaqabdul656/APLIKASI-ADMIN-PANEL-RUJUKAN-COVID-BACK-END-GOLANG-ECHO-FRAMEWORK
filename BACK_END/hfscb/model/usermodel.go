package Model

type User struct {
	Id_user   int `json:"id_user"`
	Bpjs int `json:"bpjs"`
	Nama string `json:"nama"`
	Username string `json:"username"`
	Password string `json:"password"`
	Id_rs string `json:"id_rs"`
	Nik string `json:"nik"`
	Alamat string `json:"alamat"`
	Tanggal_lahir string `json:"tanggal_lahir"`
	No_hp string `json:"no_hp"`

}

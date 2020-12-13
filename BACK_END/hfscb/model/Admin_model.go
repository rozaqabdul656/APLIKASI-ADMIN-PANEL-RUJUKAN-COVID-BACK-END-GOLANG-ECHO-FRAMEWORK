package Model

type Admin struct {
	Id_admin   int `json:"id_admin"`
	Nama string `json:"nama"`
    Username string `json:"username"`
	Password string `json:"password"`
	Level_akses string `json:"level_akses"`
	No_telp string  `json:"no_telp"`
	Alamat string  `json:"alamat"`
	Email string  `json:"email"`
	Url string  `json:"url"`
	Foto string  `json:"foto"`
	Created_at string  `json:"created_at"`
}
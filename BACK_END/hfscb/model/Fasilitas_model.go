package Model


type Master_fasilitas struct {
	Id_master_fasilitas   int `json:"id_master_fasilitas"`
	Nama_fasilitas string `json:"nama_fasilitas"`
}

type Fasilitas struct {
	Id_fasilitas   int `json:"id_fasilitas"`
	Jenis_fasilitas string `json:"jenis_fasilitas"`
    Persediaan string `json:"persediaan"`
	Id_user string `json:"id_user"`
	Master_fasilitas

}
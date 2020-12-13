package Model


type Master_Logistik struct {
	Id_master_logistik   int `json:"id_master_logistik"`
	Nama_logistik_master string `json:"nama_logistik_master"`
}

type Logistik struct {
	Id_logistik   int `json:"id_logistik"`
	Nama_logistik string `json:"nama_logistik"`
    Stok_persediaan string `json:"stok_persediaan"`
	Id_user string `json:"id_user"`
	Master_Logistik

}

type Order_logistik struct {
	Id_order_logistik   int `json:"id_order_logsitik"`
	bpjs string `json:"bpjs"`
    Tanggal_order string `json:"tanggal_order"`
	Id_rs string `json:"id_rs"`
	Jenis_order string `json:"jenis_order"`
	Hasil string `json:"hasil"`
	Waktu_order string `json:"waktu_order"`
	Tindakan string `json:"tindakan"`
	Report string `json:"report"`
	User	

}
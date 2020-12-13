package Model

type Ambulance struct {
	Id_order_ambulance   int `json:"id_order_ambulance"`
	Bpjs string `json:"bpjs"`
    No_telp string `json:"no_telp"`
	Alamat string `json:"alamat"`
	Keluhan string `json:"keluhan"`
	Status string  `json:"status"`
	Id_rs string  `json:"id_rs"`
	User
}
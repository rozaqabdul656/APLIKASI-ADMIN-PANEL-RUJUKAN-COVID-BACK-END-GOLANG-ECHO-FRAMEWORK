package Model


type Ticketing_irj struct {
	Id_ticketing_irj   int `json:"id_ticketing_irj"`
	Bpjs string `json:"bpjs"`
	Id_poli string `json:"id_poli"`
	Time string `json:"time"`
	Antrian string `json:"antrian"`
	Id_rs string `json:"id_rs"`
	Status string `json:"status"`
	Id_dokter string `json:"id_dokter"`
	Poli
	User
}

package Model

type Berita struct {
	Id_berita   int `json:"id_berita"`
	Judul string `json:"judul"`
    Cover string `json:"cover"`
	Content string `json:"content"`
	Tanggal string `json:"tanggal"`
	Status string  `json:"status"`
}
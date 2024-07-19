package web

type AddressCreateRequest struct {
	NamaPenerima string `json:"nama_penerima"`
	Kabupaten    string `json:"kabupaten"`
	Kecamatan    string `json:"kecamatan"`
	Kelurahan    string `json:"kelurahan"`
	Alamat       string `json:"alamat"`
	NoHP         string `json:"nohp"`
	Note         string `json:"note"`
}

type AddressUpdateRequest struct {
	Id           int    `json:"id"`
	NamaPenerima string `json:"nama_penerima"`
	Kabupaten    string `json:"kabupaten"`
	Kecamatan    string `json:"kecamatan"`
	Kelurahan    string `json:"kelurahan"`
	Alamat       string `json:"alamat"`
	NoHP         string `json:"nohp"`
	Note         string `json:"note"`
}

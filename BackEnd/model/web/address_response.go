package web

type AddressResponse struct {
	Id           int            `json:"id"`
	UserId       int            `json:"user_id"`
	NamaPenerima string         `json:"nama_penerima"`
	Kabupaten    string         `json:"kabupaten"`
	Kecamatan    string         `json:"kecamatan"`
	Kelurahan    string         `json:"kelurahan"`
	Alamat       string         `json:"alamat"`
	NoHP         string         `json:"nohp"`
	Note         string         `json:"note"`
	User         []UserResponse `json:"user"`
}

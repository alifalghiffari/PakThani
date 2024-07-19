package domain

type Address struct {
	Id           int
	UserId       int
	NamaPenerima string
	Kabupaten    string
	Kecamatan    string
	Kelurahan    string
	Alamat       string
	NoHP         string
	Note         string
	User         []User
}

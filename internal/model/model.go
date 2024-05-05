package model

type User struct {
	IdUser   int    `gorm:"primaryKey;column:id_user" json:"id_user"`
	Nama     string `gorm:"column:nama" json:"nama"`
	Npm      int    `gorm:"column:npm" json:"npm"`
	Kelas    string `gorm:"column:kelas" json:"kelas"`
	Asalkota string `gorm:"column:asal_kota" json:"asal_kota"`
}

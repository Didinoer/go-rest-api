package models

// type product struct{
// 	id int64 'gorm:"primaryKey" json:"id"'
// 	nama string 'gorm:"type:varchar(300)" json:"nama"'
// 	deskripsi string 'gorm:"type:varchar(300)" json:"deskripsi"'
// }

type Product struct {
	ID        int64  `gorm:"primaryKey" json:"id`
	Nama      string `gorm:"type:varchar(300)" json:"nama"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
}

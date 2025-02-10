package model

// import (
// 	"gorm.io/gorm"
// )

type List3Origin struct {
	ID    int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Judul string `gorm:"type:varchar(256)" json:"judul"`
	// Halaman *int    `json:"halaman"`
	// Chapter *string `gorm:"type:varchar(10)" json:"chapter"`
	// Gendre  *string `gorm:"type:varchar(128)" json:"gendre"`
	// Artis   *string `gorm:"type:varchar(128)" json:"artis"`
	// Grup    *string `gorm:"type:varchar(128)" json:"grup"`
}

func (List3Origin) TableName() string {
	return "list3_origin"
}

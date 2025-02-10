package repositories

import (
	"backend/config"

	"gorm.io/gorm"
)

// func GetAllDataFromDatabase() ([]model.List3Origin, error) {
func GetAllDataFromDatabase() ([]string, error) {
	// var listDataList3 []model.List3Origin
	var listDataList3 []string
	var tmpResult *gorm.DB

	db := config.DB

	// tmpResult = db.Raw("SELECT * FROM list3_origin").Find(&listDataList3)
	// tmpResult = db.Raw("SELECT id, judul FROM list3_origin ORDER BY id DESC LIMIT 10").Find(&listDataList3)
	// tmpResult = db.Raw("SELECT judul FROM list3_origin ORDER BY id DESC LIMIT 1000").Scan(&listDataList3)
	tmpResult = db.Raw("SELECT judul FROM list3").Scan(&listDataList3)

	return listDataList3, tmpResult.Error
}

// func SelectedData() ([]model.List3Origin, error) {
func SelectedData() ([]string, error) {
	root := "../new"
	// var result model.BaseResponseModel

	// parsing data ke array
	dataFromDatabase, err := GetAllDataFromDatabase()
	dataFromLocale, err := GetAllDataFromLocal(root)

	// ==========================================

	// Buat map untuk menyimpan data dari database
	dbMap := make(map[string]bool)
	for _, item := range dataFromDatabase {
		dbMap[item] = true
	}

	// Cari elemen yang ada di dataFromLocale tapi tidak di dataFromDatabase
	var newData []string
	for _, item := range dataFromLocale {
		if dbMap[item] { // Jika item tidak ada di database
			// newData = append(newData, item)
		} else {
			newData = append(newData, item)
		}
	}

	// ========================================
	// 				DEBUGING ONLY
	// ========================================
	// fmt.Printf(" newData : %d\n", len(newData))
	// fmt.Printf(" dataFromLocale : %d\n", len(dataFromLocale))
	// fmt.Printf(" dataFromDatabase : %d\n", len(dataFromDatabase))

	// return dataFromDatabase, err
	return newData, err
}

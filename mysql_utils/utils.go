package mysql_utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"golang_practice/log_for_app/log"
)

var db *gorm.DB

func init() {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		"sit-saemon-postgresql-tw-01.silkrode.in",
		"3732",
		"sitsaemon",
		"silkrode_saemon",
		"Mfm7fsP$zwmV^*Js",
	)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Print("gorm open err: ", err.Error())
		panic(err)
	}
	fmt.Println(db)
}

func GetTopUpRecordsByOrderIDs(orderIDs []string) {
	var records []TopUpRecord
	db.Where("order_id IN ?", orderIDs).Find(&records)
	log.Println("GetTopUpRecordsByOrderIDs: ", records)
}

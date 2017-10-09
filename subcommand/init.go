package subcommand

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// TODO 機能実現スピード最優先での実装なので要リファクタ
func ExecInit() {
	os.Create("./data.db")

	db, err := gorm.Open("sqlite3", "./data.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.CreateTable(&TargetUrl{})
}

type TargetUrl struct {
	gorm.Model
	URL    string `gorm:"size:255"`
	Source string `gorm:"size:123456789"`
}

// set User's table name to be `profiles`
func (TargetUrl) TableName() string {
	return "target_url"
}

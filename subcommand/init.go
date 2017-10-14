package subcommand

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"os/exec"
	"path/filepath"
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

	prevDir, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}
	defer func() {
		err = os.Chdir(prevDir)
		if err != nil {
			panic(err)
		}
	}()

	err = os.Chdir("./store")
	if err != nil {
		panic(err)
	}

	gi := exec.Command("git", "init")
	err = gi.Run()
	if err != nil {
		panic(err)
	}

}

type TargetUrl struct {
	gorm.Model
	URL    string `gorm:"size:255"`
	Source string `gorm:"size:123456789"`
}

func (TargetUrl) TableName() string {
	return "target_url"
}

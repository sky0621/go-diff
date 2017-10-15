package subcommand

import (
	"os"

	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"

	"os/exec"
	"path/filepath"

	"fmt"

	"github.com/sky0621/go-diff/static"
)

// TODO 機能実現スピード最優先での実装なので要リファクタ
func ExecInit() {
	//os.Create(static.Storage)
	//
	//db, err := gorm.Open("sqlite3", static.Storage)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer db.Close()
	//
	//db.CreateTable(&TargetUrl{})

	prevDir, err := filepath.Abs(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err = os.Chdir(prevDir)
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	err = os.Chdir(static.StorePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	gi := exec.Command("git", "init")
	err = gi.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

}

//type TargetUrl struct {
//	gorm.Model
//	URL    string `gorm:"size:255"`
//	Source string `gorm:"size:123456789"`
//}
//
//func (TargetUrl) TableName() string {
//	return "target_url"
//}

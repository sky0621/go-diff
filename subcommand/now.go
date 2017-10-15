package subcommand

import _ "github.com/jinzhu/gorm/dialects/sqlite"

// TODO 機能実現スピード最優先での実装なので要リファクタ
func ExecNow() {
	//db, err := gorm.Open("sqlite3", "./data.db")
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//defer db.Close()
	//
	//urls := viper.GetStringSlice("urls")
	//for _, url := range urls {
	//	out := &TargetUrl{}
	//	db.Where("url = ?", url).Find(out)
	//	fmt.Println("================================================")
	//	fmt.Println(out)
	//}

}

package subcommand

import (
	"io/ioutil"
	"net/http"
	"os"

	"fmt"

	"github.com/sky0621/go-diff/static"
	"github.com/spf13/viper"

	"os/exec"

	"path/filepath"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// TODO 機能実現スピード最優先での実装なので要リファクタ
func ExecCheck() {
	url := viper.GetString("url")

	source := crawl(url)

	//prevs := prevSearch(url)
	//if len(prevs) == 0 {
	//	fmt.Println("==========================================")
	//	fmt.Println("No Previous Record, so Just Save.")
	//	fmt.Println("==========================================")
	//
	//	save(source)
	//	return
	//}

	diffStr := save2(url, source)
	fmt.Println(diffStr)
}

func crawl(url string) string {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer func() {
		if res != nil {
			res.Body.Close()
		}
	}()

	baBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(baBody)
}

//func prevSearch(url string) []*TargetUrl {
//	db, err := gorm.Open("sqlite3", static.Storage)
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//	defer db.Close()
//
//	outs := []*TargetUrl{}
//	db.Where("url = ?", url).Find(outs)
//	return outs
//}
//
//func save(source string) {
//	err := ioutil.WriteFile(static.StorePath+viper.GetString("title"), []byte(source), os.ModePerm)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	prevDir, err := filepath.Abs(".")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer func() {
//		err = os.Chdir(prevDir)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//	}()
//
//	err = os.Chdir(static.StorePath)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	gitaddCmd := exec.Command("git", "add", viper.GetString("title"))
//	err = gitaddCmd.Run()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	gitcommitCmd := exec.Command("git", "commit", "-m", "update")
//	err = gitcommitCmd.Run()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//}

func save2(url, source string) string {
	err := ioutil.WriteFile(static.StorePath+viper.GetString("title"), []byte(source), os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	prevDir, err := filepath.Abs(".")
	if err != nil {
		fmt.Println(err)
		return ""
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
		return ""
	}

	gitaddCmd := exec.Command("git", "add", viper.GetString("title"))
	err = gitaddCmd.Run()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	gitcommitCmd := exec.Command("git", "commit", "-m", "update")
	err = gitcommitCmd.Run()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	diffRes, err := exec.Command("git", "diff", "HEAD^").Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(diffRes))

	//db, err := gorm.Open("sqlite3", static.Storage)
	//if err != nil {
	//	fmt.Println(err)
	//	return ""
	//}
	//defer db.Close()
	//
	//db.Create(&TargetUrl{URL: url, Source: string(diffRes)})

	return string(diffRes)
}

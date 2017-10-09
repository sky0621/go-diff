package subcommand

import (
	"io/ioutil"
	"net/http"

	"fmt"

	"github.com/spf13/viper"

	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// TODO 機能実現スピード最優先での実装なので要リファクタ
func ExecCheck() {
	urls := viper.GetStringSlice("urls")
	for _, url := range urls {
		source := crawl(url)
		prev := prevSearch(url)
		if prev.ID < 1 {
			fmt.Println("=============================")
			fmt.Println("No prev. Just Save.")
			fmt.Println("=============================")

			save(url, source)
		} else {
			prvSrcs := strings.Split(prev.Source, "\n")
			aprvSrcs := []string{}
			for _, prvSrc := range prvSrcs {
				tprvSrc := strings.TrimSpace(prvSrc)
				if tprvSrc != "" {
					aprvSrcs = append(aprvSrcs, tprvSrc)
				}
			}

			nowSrcs := strings.Split(source, "\n")
			anowSrcs := []string{}
			for _, nowSrc := range nowSrcs {
				tnowSrc := strings.TrimSpace(nowSrc)
				if tnowSrc != "" {
					anowSrcs = append(anowSrcs, tnowSrc)
				}
			}

			if len(aprvSrcs) > len(anowSrcs) {
				for idx, aprvSrc := range aprvSrcs {
					fmt.Println(lineDiff(aprvSrc, anowSrcs[idx]))
				}
			} else {
				for idx, anowSrc := range anowSrcs {
					fmt.Println(lineDiff(aprvSrcs[idx], anowSrc))
				}
			}
		}
	}
}

func crawl(url string) string {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() {
		if res != nil {
			res.Body.Close()
		}
	}()

	baBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(baBody)
}

func prevSearch(url string) *TargetUrl {
	db, err := gorm.Open("sqlite3", "./data.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	out := &TargetUrl{}
	db.Where("url = ?", url).Find(out)
	return out
}

func save(url, source string) {
	db, err := gorm.Open("sqlite3", "./data.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.Create(&TargetUrl{URL: url, Source: source})
}

func lineDiff(src1, src2 string) []diffmatchpatch.Diff {
	dmp := diffmatchpatch.New()
	a, b, c := dmp.DiffLinesToChars(src1, src2)
	diffs := dmp.DiffMain(a, b, false)
	result := dmp.DiffCharsToLines(diffs, c)
	return result
}

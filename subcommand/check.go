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
)

// TODO 機能実現スピード最優先での実装なので要リファクタ
func ExecCheck() {
	source := crawl(viper.GetString(static.TargetURL))

	err := ioutil.WriteFile(static.StorePath+viper.GetString(static.SaveFile), []byte(source), os.ModePerm)
	if err != nil {
		fmt.Printf("[check][01]%#v", err.Error())
		return
	}

	prevDir, err := filepath.Abs(".")
	if err != nil {
		fmt.Printf("[check][02]%#v", err.Error())
		return
	}
	defer func() {
		err = os.Chdir(prevDir)
		if err != nil {
			fmt.Printf("[check][03]%#v", err.Error())
			return
		}
	}()

	err = os.Chdir(static.StorePath)
	if err != nil {
		fmt.Printf("[check][04]%#v", err.Error())
		return
	}

	gitaddCmd := exec.Command("git", "add", viper.GetString(static.SaveFile))
	err = gitaddCmd.Run()
	if err != nil {
		fmt.Printf("[check][05]%#v", err.Error())
		return
	}
	gitcommitCmd := exec.Command("git", "commit", "-m", "update")
	err = gitcommitCmd.Run()
	if err != nil {
		fmt.Printf("[check][06]%#v", err.Error())
		return
	}

	diffRes, err := exec.Command("git", "diff", "HEAD^").Output()
	if err != nil {
		fmt.Printf("[check][07]%#v", err.Error())
		return
	}

	fmt.Println(string(diffRes))
}

func crawl(url string) string {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("[crawl][01]%#v", err)
		return ""
	}
	defer func() {
		if res != nil {
			res.Body.Close()
		}
	}()

	baBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("[crawl][02]%#v", err)
		return ""
	}
	return string(baBody)
}

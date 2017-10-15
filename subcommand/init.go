package subcommand

import (
	"os"

	"os/exec"
	"path/filepath"

	"fmt"

	"github.com/sky0621/go-diff/static"
)

// TODO 機能実現スピード最優先での実装なので要リファクタ
func ExecInit() {
	prevDir, err := filepath.Abs(".")
	if err != nil {
		fmt.Printf("[init][01]%#v", err.Error())
		return
	}
	defer func() {
		err = os.Chdir(prevDir)
		if err != nil {
			fmt.Printf("[init][02]%#v", err.Error())
			return
		}
	}()

	err = os.Chdir(static.StorePath)
	if err != nil {
		fmt.Printf("[init][03]%#v", err.Error())
		return
	}

	gi := exec.Command("git", "init")
	err = gi.Run()
	if err != nil {
		fmt.Printf("[init][04]%#v", err.Error())
		return
	}
}

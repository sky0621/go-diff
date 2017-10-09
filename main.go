// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"

	"github.com/sky0621/go-diff/cmd"
	"github.com/spf13/viper"
)

// TODO 機能実現スピード最優先での実装なので要リファクタ
func main() {
	cmd.Execute()
}

func init() {
	cfg := flag.String("f", "config.toml", "Config File")
	flag.Parse()

	viper.SetConfigFile(*cfg)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

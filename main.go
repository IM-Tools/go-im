/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"flag"
	"go_im/config"
	"go_im/im"
	"go_im/im/tcp"
	"go_im/pkg/wordsfilter"
)

func init()  {
	config.Initialize()
	wordsfilter.SetTexts()
}

func main() {
	var serve string
	flag.StringVar(&serve, "serve", "", "选择运行的服务🚀")
	flag.Parse()
	switch serve {
	case "http":
		im.StartHttp()
	case "tcp-serve":
		tcp.StartTcpServe()
	case "tcp-client":
		tcp.StartTcpClient()
	default:
		im.StartHttp()
	}
}

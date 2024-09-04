package main

import (
	"fmt"
	"log"
	"slogv2/src/main/controller"
	"slogv2/src/main/entity"
)

const (
	DEV  string = "dev"
	PROD string = "prod"
)

func main() {
	log.Println("START")
	Env := PROD

	switch Env {
	case DEV:
		fmt.Println("dev")
		//fmt.Println(test.GenTestArticle())
		//fmt.Println(test.RegexTest(utils.TestText))
		break

	case PROD:
		fmt.Println("prod")
		//TODO 实现环境变量配置
		entity.DbInit()
		controller.InitRouter()
		break
	}

	log.Println("END")
}

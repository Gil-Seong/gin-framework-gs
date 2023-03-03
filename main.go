package main

import (
	"fmt"
	"os"
	"strconv"

	"gin-framework-gs/config"
	"gin-framework-gs/lib/logging"
	"gin-framework-gs/routers"

	"github.com/spf13/viper"
)

// main 패키지에 init() 메서드를 만들어놓으면 main()보다 먼저 실행됨
func init() {
	profile := initProfile()
	setRuntimeConfig(profile)
	logging.LoggerConfig()

}

// PROFILE을 기반으로 config파일을 읽고 전역변수에 언마샬링
func setRuntimeConfig(profile string) {
	viper.AddConfigPath(".")
	// 환경변수에서 읽어온 profile이름의 yaml파일을 configPath로 설정
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// viper는 읽어온 설정파일의 정보를 가지고있으니, 전역변수에 언마샬링
	// 애플리케이션의 원하는곳에서 사용
	err = viper.Unmarshal(&config.RuntimeConf)
	if err != nil {
		panic(err)
	}
}

// 환경변수는 PROFILE을 확인하기 위해 하나만 설정
func initProfile() string {
	var profile string
	profile = os.Getenv("GO_PROFILE")
	if len(profile) <= 0 {
		profile = "local"
	}
	fmt.Println("GOLANG_PROFILE: " + profile)
	return profile
}

func main() {
	strPort := ":" + strconv.Itoa(config.RuntimeConf.Server.Port)

	// routers.TestServer().Run(":8080")
	// routers.NewServer().Run(":8080")
	routers.NewServer().Run(strPort)

}

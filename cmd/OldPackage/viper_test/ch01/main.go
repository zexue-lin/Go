package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	ServiceName string `mapstructure:"name"`
	Port        int    `mapstructure:"port"`
}

func main() {
	v := viper.New()

	// 文件路径如何设置，这里使用 go run 运行
	// 编辑配置 -> 运行/调试-> 工作目录是 F:\Depository\Golang-Learn\Go\cmd\OldPackage
	// 才能正确找到目标文件
	v.SetConfigFile("viper_test/ch01/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)
	fmt.Printf("%v", v.Get("name"))

}

/*
获取配置文件内容信息
先 new 一个 。再 设置文件路径。再直接Get 就可以拿到了

还有一个功能就是可以直接把yaml里面的文件 映射成 go中的  struct
*/

package main

import (
	"fmt"
	"time"
  "github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
type ServerConfig struct {
	ServiceName string      `mapstructure:"name"`
	MysqlInfo   MysqlConfig `mapstructure:"mysql"`
}

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func main() {
	// 这边设置了系统环境变量之后输出的是false ，只要重启编辑器，重新运行即可
	debug := GetEnvInfo("GOSHOP_DEBUG")

	configFilePrefix := "config" // 自定义文件前缀
	configFileName := fmt.Sprintf("./%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("./%s-debug.yaml", configFilePrefix)
	}
	v := viper.New()

	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)

	// viper的功能 - 动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&serverConfig)

    fmt.Println(serverConfig)
	})

  time.Sleep(time.Second*300)

}

// 如何将测试和线上的配置文件分开，不用改任何代码。
// 本地电脑配置一个环境变量

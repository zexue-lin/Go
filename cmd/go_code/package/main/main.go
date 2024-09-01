package main

// 如果想要在这里访问到user目录下定义的东西。要使用import关键字
// 引入的时候看go.mod文件，看里面的module是什么路径就从哪里开始
import (
	_ "Go/cmd/go_code/package/user" // 包的路径
	u "Go/cmd/go_code/package/user" // 包的别名（应对多个目录下包名一样的情况）
	// . "Go/cmd/go_code/package/user"
	// 这种方法可以把需要导入包里面的全部东西都引入过来，下面用的时候就不用每次都加包的前缀.变量
	// 但是这种方法音量少用，代码可读性差，而且容易出错
	"fmt"
)

func main() {
	c := u.User{
		Name: "Tom",
	}
	fmt.Println(u.GetUser(c))
}

/*
go list -m all  查看项目所依赖的东西

go list -m -version github.com/gin-gonic/gin  查看可用版本

go get github.com/gin-gonic/gin@v1.8.0  切换使用的版本为1.8.0  这里是举例

// 1.编译器一键下载

go get github.com/redis/go-redis/v9  2.go get 也可以下载依赖 (有版本的就在后面@版本号) 会修改go.mod文件

go mod tidy 3.tidy也可以下载依赖 先把代码辅复制过来才行

go install也行好像

go get -u  升级某个包 到最新的次要版本或修订版本
go get -u=patch  升级到最新的修订版本


如果A项目仓库是project-A，但是仓库中的go.mod中设置的是github.com/tom/A
B项目由于依赖项目A，import的github.com/tom/A，go get命令的时候，由于package和代码仓库的名称不一样会报错
使用replace

go mod edit -replace github.com/tom/A=github.com/tom/project-A v1.0.0

执行这个之后，go.mod文件中出现了一条替换的语句，此时下载依赖的时候就会从替换后的地址里面下载(后面的地址)
*/

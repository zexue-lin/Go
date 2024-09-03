package main

// 豆包写的
import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type FounderData struct {
	Message string
	Count   int
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT founder,month,day,channel_code,newFriends FROM companymtl WHERE day = 30")
	if err != nil {
		fmt.Println("查询失败", err)
		return
	}
	defer rows.Close()

	founders := make(map[string]*FounderData)
	for rows.Next() {
		var founder string
		var month, day int
		var channelCode string
		var newFriends int
		err := rows.Scan(&founder, &month, &day, &channelCode, &newFriends)
		if err != nil {
			fmt.Println("扫描数据失败", err)
			return
		}
		if _, ok := founders[founder]; !ok {
			founders[founder] = &FounderData{
				Message: fmt.Sprintf("%s 老师中午好，您有 %d 月 %d 日创建的渠道活码【%s】新增了 %d 个好友，请在运营中台 MLTL 及时新增推广。<br>", founder, month, day, channelCode, newFriends),
				Count:   1,
			}
		} else {
			founders[founder].Message += fmt.Sprintf("以及渠道活码【%s】新增了 %d 个好友。<br>", channelCode, newFriends)
			founders[founder].Count++
		}
	}

	if err := rows.Err(); err != nil {
		fmt.Println("遍历数据错误", err)
		return
	}

	// 获取当前文件的绝对路径
	absPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("获取当前路径失败", err)
		return
	}

	// 生成文件名
	outputFileName := filepath.Join(absPath, "output.html")

	// 将结果写入文件
	file, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("创建文件失败", err)
		return
	}
	defer file.Close()

	var allMessages strings.Builder
	for _, data := range founders {
		allMessages.WriteString(data.Message)
	}

	_, err = file.WriteString(allMessages.String())
	if err != nil {
		fmt.Println("写入文件失败", err)
		return
	}

	fmt.Println("结果已写入", outputFileName)
}

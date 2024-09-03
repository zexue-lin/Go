package main

// Chat-GPT写的
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html"
	"os"
)

func main() {
	// 数据库连接信息
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	defer db.Close()

	// 检查数据库连接是否成功
	if err := db.Ping(); err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}

	// 查询语句
	query := "SELECT founder, month, day, channel_code, newFriends FROM companymtl WHERE day = 1"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}
	defer rows.Close()

	// 保存查询结果
	founders := make(map[string]struct {
		message string
		count   int
	})

	// 处理查询结果
	for rows.Next() {
		var founder, channelCode string
		var month, day, newFriends int

		err := rows.Scan(&founder, &month, &day, &channelCode, &newFriends)
		if err != nil {
			fmt.Println("读取行数据失败:", err)
			return
		}

		if data, ok := founders[founder]; !ok {
			message := fmt.Sprintf(
				"%s老师下午好，您有%d月%d日创建的渠道活码【%s】新增了%d个好友，请在运营中台 MLTL 及时新增推广。若已上表请忽略<br>",
				html.EscapeString(founder), month, day, html.EscapeString(channelCode), newFriends)
			founders[founder] = struct {
				message string
				count   int
			}{
				message: message,
				count:   1,
			}
		} else {
			data.message += fmt.Sprintf("以及渠道活码【%s】新增了%d个好友。若已上表请忽略<br>", html.EscapeString(channelCode), newFriends)
			data.count++
			founders[founder] = data
		}
	}

	// 生成输出文件
	output := "output.html"
	file, err := os.Create(output)
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}
	defer file.Close()

	// 写入数据到文件
	for _, data := range founders {
		file.WriteString(data.message)
	}

	// 提示生成文件成功
	fmt.Println("数据已写入到", output)
}

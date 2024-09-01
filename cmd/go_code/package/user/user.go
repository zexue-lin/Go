package user

// package 用来组织源码，是多个go源码的集合，代码的复用，自带的包有：fmt、os、io
// 每个源码文件开始都要必须申明所属的package
// python不需要申明package，
// php c# 通过 namespace 定义

// go中 同一个目录下的源文件不可以出现两个不同名字的包名

type User struct {
	Name string
}

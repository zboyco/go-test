package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	log.Println("测试redigo")

	conn, err := redis.Dial("tcp", "192.168.2.99:6379")
	checkErr(err)
	defer conn.Close()

	// 插入数据,返回OK
	addResult, err := conn.Do("SET", "mykey", "content")
	checkErr(err)
	log.Println(addResult)

	// 插入数据并设置过期时间,返回OK
	addExResult, err := conn.Do("SET", "mykeyex", "content", "EX", 10)
	checkErr(err)
	log.Println(addExResult)

	// 对已有key设置5s过期时间,返回影响数
	exRows, err := conn.Do("EXPIRE", "mykeyex", 5)
	checkErr(err)
	log.Println(exRows)

	// 判断key是否存在,返回bool
	has, err := redis.Bool(conn.Do("EXISTS", "mykey"))
	checkErr(err)
	log.Println(has)

	// 获取数据bytes,返回bytes
	bytes, err := conn.Do("GET", "mykey")
	checkErr(err)
	bytesContent := string(bytes.([]byte))
	log.Println(bytesContent)

	// 获取数据并直接转成string,返回string
	content, err := redis.String(conn.Do("GET", "mykey"))
	checkErr(err)
	log.Println(content)

	// 删除key,返回影响数
	delRows, err := conn.Do("DEL", "mykey")
	checkErr(err)
	log.Println(delRows)

	// 插入列表,返回list中的元素个数
	for i := 0; i < 3; i++ {
		listRows, err := conn.Do("RPUSH", "list", fmt.Sprint(i))
		checkErr(err)
		log.Println(listRows)
	}

	// 读取列表for
	values, err := redis.Values(conn.Do("LRANGE", "list", 0, 100))
	checkErr(err)
	for _, v := range values {
		log.Println(string(v.([]byte)))
	}
	conn.Do("DEL", "list")
}

// checkErr 错误检查
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var c, err = redis.Dial("tcp", "192.168.50.253:6379")

func init() {

	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	fmt.Println("Connect to redis sucess!")

}
func expire() {

	_, err = c.Do("SET", "myname", "Rosco", "EX", "5") //EX -> expire過期
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "myname"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	time.Sleep(8 * time.Second)

	username, err = redis.String(c.Do("GET", "myname"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get myname: %v \n", username)
	}
}
func exists() {
	_, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	is_key_exit, err := redis.Bool(c.Do("EXISTS", "mykey1"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("exists or not: %v \n", is_key_exit)
	}
}
func jsontest() {
	key := "LinChi"
	imap := map[string]string{"username": "Rosco", "phonenumber": "0985944005"}
	value, _ := json.Marshal(imap)

	n, err := c.Do("SET", key, value)

	if err != nil {
		fmt.Println(err)
	}
	if n == int64(1) {
		fmt.Println("success", int64(1))
	}

	var imapGet map[string]string

	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}
	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}
	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["phonenumber"])
}
func delkey() {
	_, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	_, err = c.Do("DEL", "mykey")
	if err != nil {
		fmt.Println("redis delelte failed:", err)
	}

	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}
func main() {

	// c, err := redis.Dial("tcp", "192.168.50.253:6379")
	// if err != nil {
	// 	fmt.Println("Connect to redis error", err)
	// 	return
	// }
	// fmt.Println("sucess!")
	// defer c.Close()

	// ////測試過期
	expire()
	// ////測試key是否存在
	// exists()
	// /////測試JSON
	// jsontest()
	// //////測試DEL
	// delkey()

}

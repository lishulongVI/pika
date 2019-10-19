package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

func main() {

	//创建map
	s1 := make(map[string]string)
	fmt.Println(s1)

	s1["hello"] = "world"

	s := s1["hello"]

	fmt.Println(s1)
	fmt.Println(s)

	i, ok := s1["world"]

	if !ok {
		fmt.Println("查询不到对应的key")
	} else {
		fmt.Println(i)
	}

	// 添加key
	s1["world"] = "good"
	//删除key
	delete(s1, "hello")
	fmt.Println(s1)
	//更新key
	s1["world"] = "wonderful"
	fmt.Println(s1)

	for i, v := range s1 {
		fmt.Println(i, v)
	}

	//初始化随机
	rand.Seed(time.Now().UnixNano())

	map_ := make(map[string]int)
	for i := 0; i < 100; i++ {
		a := fmt.Sprintf("stu_%d", i)
		map_[a] = i
	}

	keys := make([]string, 0, 200)
	for i, v := range map_ {
		fmt.Println(i, v)
		keys = append(keys, i)
	}

	fmt.Println(map_)
	fmt.Println(keys)
	//获取keys
	mapKeys := reflect.ValueOf(map_).MapKeys()
	fmt.Println(mapKeys)

	//键值 转换
	m1 := make(map[int]string)
	for i, v := range map_ {
		m1[v] = i
	}
	fmt.Println(m1)
	// 键排序
	type p struct {
		n int
		v string
	}

	var pp []p

	for k, v := range m1 {
		pp = append(pp, p{k, v})
	}
	fmt.Println(pp)

	sort.Slice(pp, func(i, j int) bool {
		//return pp[i].n > pp[j].n
		return pp[i].n < pp[j].n
	})

	fmt.Println(pp)

}

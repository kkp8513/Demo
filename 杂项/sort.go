package main

import (
	"fmt"
	"sort"
)

//动物分类
type Type int

//分类顺序
const (
	Bird Type = iota
	Beast
	Fish
)

//结构
type Animal struct {
	Name string
	Type
}

//定义结构类型
type Animals []*Animal

//长度
func (a Animals) Len() int {
	return len(a)
}

//比较
func (a Animals) Less(i,j int) bool {
	if a[i].Type != a[j].Type {
		return a[i].Type > a[j].Type 
	}
	return a[i].Name > a[j].Name
}

//执行
func (a Animals) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	animals := Animals {
		&Animal {
			"鸟",
			Bird,
		},
		&Animal {
			"鱼",
			Fish,
		},
		&Animal {
			"兽2",
			Beast,
		},
		&Animal {
			"兽",
			Beast,
		},
	}

	sort.Sort(animals)

	for _, v := range animals {
		fmt.Println(v)
	}

}
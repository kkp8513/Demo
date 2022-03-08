package main

import "fmt"

type node struct {
	Value int
	Next *node
}

//初始化
func InitNode(b []int) *node{
	head := &node{ 
		0,
		nil,
	}
	temp := head
	for _, v := range b {
		node := &node {
			v,
			nil,
		}
		temp.Next = node
		temp = node
	}
	return head
}

//遍历
func Display(n *node) {
	for {
		if n.Next == nil {
			break
		}
		fmt.Printf("p:%p, value:%d, next:%p\n", n, n.Value, n.Next)
		n = n.Next
	}
	fmt.Println("end")
}

//插入
func Add(n *node, site, v int) {
	temp := n
	i := 0
	for {
		i++
		n = n.Next
		if i == site {
			temp.Next = &node {
				v,
				n,
			}
			break 
		}
		if i > site {
			fmt.Println("输入位置异常")
			break
		}
		temp = n
	}
}

//删除
func Del(n *node, v int) *node{
	temp := n
	for {
		n = n.Next
		if v == n.Value {
			temp.Next = n.Next
			return n
		}
		if n.Next == nil {
			fmt.Println("未找到该节点")
			break
		}
		temp = n
	}
	return nil
}

//查找
func Find(n *node, v int) int {
	i := 0
	for {
		i++
		n = n.Next
		if n.Value == v {
			return i
		}
		if n.Next == nil {
			return -1
		}		
	}
}

//更新
func Update(n *node, site, v int) {
	i := 0
	for {
		i++
		n = n.Next
		if i == site {
			n.Value = v
			break
		}
		if n.Next == nil {
			fmt.Println("位置不存在")
			break
		}
	}
}

func main() {
	head := InitNode([]int{1, 3, 5, 6, 8})
	Display(head)

	i := Find(head, 3)
	fmt.Printf("i:%d\n", i)

	Add(head, i, 33)
	Display(head)

	Update(head, i, 44)
	Display(head)

	node :=Del(head, 5)
	fmt.Printf("delete:%+v\n", node)
	Display(head)
	
}
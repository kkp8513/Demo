package main

import "fmt"

type table struct {
	Arr []int
	Len int
	Size int
}

//创建
func InitTable(len,size int) *table{
	if len > size {
		fmt.Println("初始化长度不能大于容量")
	}
	arr := make([]int, len, size)
	for i := 0; i < len; i++ {
		arr[i] = i+10
	}
	return &table{
		arr,
		len,
		size,
	}
}

//遍历
func (t *table) Display() {
	for i := 0; i < t.Len; i++ {
		fmt.Printf("%d:%d ", i, t.Arr[i])
	}
	fmt.Printf("Len:%v\n", t.Len)
}

//添加
func (t *table) Add(i, v int) {
	if i > t.Len+1 || i < 0{
		fmt.Println("位置不合理")
	}
	if t.Len == t.Size {
		fmt.Println("无空间")
	}
	t.Arr = append(t.Arr[:i], append([]int{v}, t.Arr[i:]...)...)
	t.Len++
}

//查找
func (t *table) Find(v int) int {
	for i := 0; i < t.Len; i++ {
		if v == t.Arr[i] {
			return i
		}
	}
	return -1
}

//编辑
func (t *table) Update(i, v int) {
	if i > t.Len || i < 0 {
		fmt.Println("位置不合理")
	}
	t.Arr[i] = v
}

//删除
func (t *table) Del(i int) (v int){
	if i > t.Len || i < 0 {
		fmt.Println("位置不合理")
	}
	v = t.Arr[i]
	t.Arr = append(t.Arr[:i], t.Arr[i+1:]...)
	t.Len--
	return v
}

func main() {
	table := InitTable(6,10)
	table.Display()

	table.Add(2, 22)
	table.Display()

	i := table.Find(22)
	table.Update(i, 44)
	table.Display()

	table.Del(0)
	table.Display()
}
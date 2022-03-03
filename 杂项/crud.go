package main

import "fmt"

type Operation struct {}

func (o *Operation) Creater() {
	fmt.Println("method: creater")
}

func (o *Operation) Select(field []string, where string) {
	fmt.Println("method: select")
}

func (o *Operation) Update(where string, datas map[string]interface{}) {
	fmt.Println("method: update")
}

func (o *Operation) Delete(where string) {
	fmt.Println("method: delete")
}

type Library struct {
	Operation
}

func main() {
	library := new(Library)
	library.Creater()
	library.Select([]string{"name", "age"}, "where....")
	library.Update("where...", map[string]interface{}{"name": "大哥", "age": 30})
	library.Delete("where...")
}
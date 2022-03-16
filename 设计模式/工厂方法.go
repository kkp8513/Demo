package main

import (
	"fmt"
)

type Monster interface {
	Create() infoFactory
}

type infoFactory interface {
	setName(name string)
	setAge(age int)
	result()
}

type info struct {
	name string
	age int
}

func (i *info) setName(name string) {
	i.name = name
}

func (i *info) setAge(age int) {
	i.age = age
}

type godzillaFactory struct {}

func (godzillaFactory) Create() infoFactory{
	return &godzilla{
		&info{},
	}
}

type godzilla struct {
	*info
}

func (g *godzilla) result() {
	fmt.Println("godzilla:", g.info.name, g.info.age)
}

type orangutansFactory struct {}

func (orangutansFactory) Create() infoFactory{
	return &orangutans{
		&info{},
	}
}

type orangutans struct {
	*info
}

func (o *orangutans) result() {
	fmt.Println("orangutans:", o.info.name, o.info.age)
}

func born(monster Monster, name string, age int) {
	fa := monster.Create()
	fa.setName(name)
	fa.setAge(age)
	fa.result()
}

func main() {
	var monster Monster
	monster = godzillaFactory{}
	born(monster, "哥斯拉", 11)
	monster = orangutansFactory{}
	born(monster, "猩猩", 12)
}
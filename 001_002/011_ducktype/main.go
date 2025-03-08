package main

import "fmt"

type Duck interface {
	Gaga()
	Walk()
	Swimming()
}
type pskDuck struct {
	name string
}

func (p *pskDuck) Gaga() {
	fmt.Println(p.name, "gagagagaga~")
}
func (p *pskDuck) Walk() {
	fmt.Println(p.name, "walk~ 121 121")
}
func (p *pskDuck) Swimming() {
	fmt.Println(p.name, "swimming~")
}

func main() {
	//鸭子类型
	//go语言中到处都是interface，到处都是duck typing
	//当看到一只鸟走起来像鸭子，游泳起来像鸭子，叫起来也像鸭子，那么这只鸟就是鸭子
	//强调的是外部暴露出来的方法，而不是结构

	var d Duck = &pskDuck{name: "gg1"}
	d.Walk()
	d.Swimming()
	d.Gaga()

}

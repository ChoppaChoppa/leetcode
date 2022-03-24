package main

import "fmt"

func main() {

	//водитель сперва ездит на машине
	driver := &Driver{}
	auto := &Auto{}
	driver.Travel(auto)

	// так как у структуры camelToTransport есть метод drive,
	// который является реализацие интерфейса iTransport
	// и в этом методe выполняется метод move унаследованный от camel
	// мы можем оставить машину и ходить по пескам на верблюде
	c := &Camel{}
	camelToTransport := CamelToTransportAdapter{
		camel: *c,
	}

	driver.Travel(camelToTransport)
}

type ITransport interface {
	Drive()
}

type Auto struct {
}

func (a *Auto) Drive() {
	fmt.Println("машина едет по дороге")
}

type IAnimal interface {
	Move()
}
type Camel struct {
}

func (c *Camel) Move() {
	fmt.Println("верблюд идет по пескам")
}

type CamelToTransportAdapter struct {
	camel Camel
}

func (camel CamelToTransportAdapter) Drive() {
	camel.camel.Move()
}

type Driver struct {
}

func (d *Driver) Travel(transport ITransport) {
	transport.Drive()
}

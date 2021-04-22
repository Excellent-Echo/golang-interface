package main

import "fmt"

type BangunDatar interface {
	luas() float64
}

type Persegi struct {
	sisi float64
}

func (p Persegi) luas() float64 {
	return p.sisi * p.sisi
}

type PersegiPanjang struct {
	panjang float64
	lebar   float64
}

func (pp PersegiPanjang) luas() float64 {
	return pp.panjang * pp.lebar
}

type Segitiga struct {
	panjang float64
	tinggi  float64
}

func (s Segitiga) luas() float64 {
	return (1 / 2) * s.panjang * s.tinggi
}

func luasBangunDatar(bd BangunDatar) float64 {
	return bd.luas()
}

func main() {
	// interface go

	// var p = Persegi{sisi: 5}
	// var luasP = luasBangunDatar(p)

	// var pp = PersegiPanjang{panjang: 4, lebar: 6}
	// var luasPP = luasBangunDatar(pp)

	// var sgt = Segitiga{panjang: 10, tinggi: 12}
	// var luasSgt = luasBangunDatar(sgt)
	// fmt.Println(luasP, luasPP, luasSgt)

	// casting

	var a interface{} // disarankan tipe data primitif (string, int, float, bool)
	var b = 21
	var c = "impact "

	a = 29
	sum := a.(int) + b

	a = "byte"

	gabung := c + a.(string)

	fmt.Println(sum, gabung)

	// interface kosong

	// var random map[string]interface{}

	// random := map[string]interface{}{
	// 	"name":     "impact byte",
	// 	"age":      21,
	// 	"language": []string{"js", "golang", "ruby"},
	// }

	// random["name"] = "impact byte"
	// random["age"] = 21
	// random["language"] = []string{"js", "golang", "ruby"}

	// random = "impact byte"
	// fmt.Println(random)
	// random = 21
	// fmt.Println(random)
	// random = []string{"haha", "hihi", "huhu"}
	// fmt.Println(random)
}

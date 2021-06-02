package main

import (
	"fmt"
	"math"
)

type coordinate struct{
	d, m, s float64
	h rune
}

type location struct{
	lat, long float64
}

type world struct{
	name string
	radius float64
}
//Перевод координат в десятичный вид
func (c coordinate) decimal() float64{
	sign := 1.0
	switch c.h{
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func newWorld(name string, radius float64) world{
	return world{name, radius}
}
//Перевод десятичных координат в радианы
func rad(deg float64) float64{
	return deg * math.Pi / 180
}
//Расчёт дистанции от точки p1 до точки p2
func (w world) distance(p1, p2 location) float64{
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}
func main(){
	earth := newWorld("Earth", 6378) //Другие планеты - https://nssdc.gsfc.nasa.gov/planetary/factsheet/
												//Но программа работает максимально точно именно для Земли
	latHome := coordinate{59,52,38.1,'N'}
	longHome := coordinate{29,53,26.6,'E'}
	//ITMO University
	//59°55'38.6"N 30°20'17.7"E
	latDst := coordinate{59,55,38.6,'N'}
	longDst := coordinate{30,20,17.7,'E'}

	home := location{latHome.decimal(), longHome.decimal()}
	dst := location{latDst.decimal(), longDst.decimal()}

	dist := earth.distance(home, dst)
	fmt.Printf("%.2f km\n", dist)



}

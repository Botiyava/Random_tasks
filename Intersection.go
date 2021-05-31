package main

import "fmt"
// На вход подаются два неупорядоченных массива любой длины.
// Необходимо написать функцию, которая возвращает пересечение массивов

func intersection(a []int, b []int) (result []int){
	list := make(map[int]int)
	//анализируем первый срез
	for _, elem := range a{
		if _,ok := list[elem]; !ok{
			list[elem] = 1
		}else{
			list[elem]++
		}
	}
	//ищем одинаковые числа в обоих срезах и заносим в итоговый срез
	for _,elem := range b{
		if _,ok := list[elem]; ok && list[elem] >=1{
			list[elem]--
			result = append(result, elem)
		}
	}
	return
}
func main(){
	a := []int{4,6,7,4}
	b := []int{4,1,7,3,7}
	fmt.Printf("даны два среза: %v, %v\n", a, b)
	fmt.Printf("результат операции пересечения: %v",intersection(a,b))
}

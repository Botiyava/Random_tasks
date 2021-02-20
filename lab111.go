package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main(){
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//ВВЕДИТЕ СЮДА ЖЕЛАЕМОЕ КОЛИЧЕСТВО ИНТЕРВАЛОВ И ЭКСПЕРИМЕНТАЛЬНУЮ ВЫБОРКУ t_k
	listOfIntervals := [3]int{}
	for rr:= 0; rr < 3; rr++{
		fmt.Println("Введите n = ", rr+1)
		fmt.Fscan(os.Stdin, &listOfIntervals[rr])
	}
	fmt.Println("dozspis:", listOfIntervals)
	//listOfIntervals := [3]int{11,15,17}
	rawData := []float64{68.63,823.26,353.06,217.66,216.60,821.00,160.98,863.55,509.23,
		378.66,360.73,158.75,972.65,880.97,678.66,665.12,502.98,697.36,369.11,86.56,896.68,
		880.63,198.57,392.96,999.92,693.17,592.69,208.61,159.68,198.16,803.65,166.05,900.65,
		236.58,703.73,193.31,171.29,725.58,51.16,736.68,762.61,129.03,681.15,666.78,332.53,561.02,
		506.38,370.17,953.08,727.69,791.82,669.15,1007.71,876.75,657.56,238.68,108.67,900.38,291.09,
		1010.00,909.72,885.67,159.00,331.66,759.68,676.63,858.18,1007.07,673.35,828.99,990.06,262.09,
		563.13,516.11,68.65,856.18,855.25,875.26,560.16,999.99,500.60,169.28,261.31,210.52,239.98,680.73,
		238.07,108.20,660.56,1018.06,766.97,822.96,332.20,268.73,59.71,361.09,131.51, 138.15 ,697.11, 650.05}

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////
	sort.Float64s(rawData) // Сортировка данных по возрастанию
	name := "/home/botiyava/go/src/theory_lab_1/resultToGraph.csv"
	rows := make([][]string,listOfIntervals[0] + listOfIntervals[1] + listOfIntervals[2] + 3)

	rows = append(rows, []string{strconv.Itoa(listOfIntervals[0]),strconv.Itoa(listOfIntervals[1]), strconv.Itoa(listOfIntervals[2])}) //csv
	/*for i:= 0; i < N; i++{

}*/
	rows[0] = append(rows[0], "Интервалы", "Кол-во отказов", "Инт-ть отказов", "Част отказов",
		"Ф-ия надежности", "λ = p/α", "n") //csv
for iterator:= 0; iterator < 3; iterator++ {
	numberOfIntervals := listOfIntervals[iterator]//количество интервалов
	var currentI int
	switch iterator{
	case 0: currentI = 0
	case 1: currentI = listOfIntervals[0]
	case 2: currentI = listOfIntervals[0] + listOfIntervals[1]
	}
	fmt.Println("При n = ", iterator)
	N := len(rawData)
	var sum float64
	interval := 0 //Текущий интервал времени
	fmt.Println("lol1: ", (rawData[len(rawData)-1]/float64(numberOfIntervals)))
	fmt.Println("lol12: ",math.Round((rawData[len(rawData)-1]/float64(numberOfIntervals))))
	fmt.Println("lol12: ", int(math.Round((rawData[len(rawData)-1]/float64(numberOfIntervals)))))
		gap := math.Round((rawData[len(rawData)-1]/float64(numberOfIntervals))))
	overallR := 0 // Общее количество вышедших из строя объектов

	list := make(map[int][]float64) //списки объектов по интервалам
	gapLoop := gap
itr := 0
	// распределение часов на работки по интервалам
	for i := 0; i < N; i++ {
		sum += rawData[i]
	this:
		if rawData[i] < float64(gap) {
			list[interval] = append(list[interval], rawData[i])
		} else {
			itr++
			overallR += len(list[interval])

			// Δr - количество отказов в данных интервал времени, r - общее количество отказов за интервал времени
			// от начала работы до текущего времени включительно.
			/*fmt.Print("Δr(" + strconv.Itoa(interval) + ") = ")
			fmt.Print(len(list[interval]))
			fmt.Print("\tr(" + strconv.Itoa(interval) + ") = ")
			fmt.Println(overallR)*/

			interval += gapLoop
			gap += gapLoop
			goto this
		}

	}
	// Δr - количество отказов в данных интервал времени, r - общее количество отказов за интервал времени
	// от начала работы до текущего времени включительно. ( Для последнего интервала)
	overallR += len(list[interval])
	/*fmt.Print("Δr(" + strconv.Itoa(interval) + ") = ")
	fmt.Print(len(list[interval]))
	fmt.Print("\tr(" + strconv.Itoa(interval) + ") = ")
	fmt.Println(overallR)*/
	interval = 0
	overallR = 0
	gap = int(math.Round((rawData[len(rawData)-1]/float64(numberOfIntervals))))

	/*rows[0] = append(rows[0], "Интервалы", "Кол-во отказов", "Инт-ть отказов", "Част отказов",
		"Ф-ия надежности", "λ = p/α") //csv*/
	//--------------         Рассчёт наработки на отказ

	T := sum / float64(len(rawData))
	formatT := fmt.Sprintf("%.4f", T)

	fmt.Println("\n---------------------------------------------------------------------------------------------")
	fmt.Print("|  Интервалы\t|Кол-во отказов\t|Инт-ть отказов  | Част отказов |Ф-ия надежности| λ = p/α   |")
	for i := 0; i < itr + 1; i++ { //решение тут
		//2 : 19,20,21,22
		//3 : 23-25
		//4 : 28,29
		//5 : 26, 27, 30
		//6 : 31

		fmt.Println("\n---------------------------------------------------------------------------------------------")
		if len(list[interval]) == 0 && i == itr+1 {
			break
		}

		// --------------   Вероятность безотказной работы (Пятый столбец вывода)

		P := float64(1.0 - (float64(overallR) / float64(N)))
		formatP := fmt.Sprintf("%.2f", P) //Форматирование вывода (Сам вывод в конце)

		//---------------   Текущий интервал времени (Первый столбец в таблице)
		fmt.Print("| (", interval, ",", interval+gap, ")\t")
		rows[i+currentI+1] = append(rows[i+currentI+1], strconv.Itoa(interval+gap)) //csv

		// --------------   Количество отказов  (Второй столбец в таблице)
		fmt.Print("|\t ", len(list[interval]), "\t")
		rows[i+currentI+1] = append(rows[i+currentI+1], strconv.Itoa(len(list[interval]))) //csv

		// --------------   Интенсивность отказов (Третий столбец в таблице)
		CurrentN := N - overallR
		overallR += len(list[interval])
		lambda := float64(len(list[interval])) / (((float64(CurrentN) + float64(N-overallR)) / 2) * float64(gap))
		formatLambda := fmt.Sprintf("%.5f", lambda) //Форматирование вывода
		fmt.Print("|\t ", formatLambda, " |\t")
		rows[i+currentI+1] = append(rows[i+currentI+1], formatLambda) //csv

		// --------------   Частота отказа (Пятый столбец в таблице)
		frequency := float64(len(list[interval])) / (float64(N) * float64(gap))
		formatFrequency := fmt.Sprintf("%.3f", frequency) //Форматирование вывода
		fmt.Print(formatFrequency, "\t| ")
		rows[i+currentI+1] = append(rows[i+currentI+1], formatFrequency) //csv

		//// --------------  Контроль (Шестой столбец в таблице)
		fmt.Print(formatP, "\t\t| ")           //Печать вероятности безотказной работы
		rows[i+currentI+1] = append(rows[i+currentI+1], formatP) //csv
		control := frequency / P
		formatControl := fmt.Sprintf("%.5f", control)
		fmt.Print(formatControl, "   | ")
		rows[i+currentI+1] = append(rows[i+currentI+1], formatControl) //csv

			//rows[i+currentI+1] = append(rows[i+currentI+1], strconv.Itoa(iterator)) //csv

		interval += gapLoop //Переходим к рассчету количественных характеристик надежности для следующего интервала

	}
	fmt.Println("\n---------------------------------------------------------------------------------------------")
	fmt.Println("\nНаработка на отказ = ", formatT, " часов")

	fmt.Println(list)

}
	writeOrders(name, rows)
}
func writeOrders(name string, rows [][]string){
	f, err := os.Create(name)
	if err != nil{
		fmt.Println(err)
	}
	defer func(){
		e := f.Close()
		if e != nil{
			fmt.Println(e)
		}
	}()
	w := csv.NewWriter(f)
	err = w.WriteAll(rows)
	if err != nil{
		fmt.Println(err)
	}
}

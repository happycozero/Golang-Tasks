/*
1. Проверить истинность высказывания: "Данные целые числа X и Y являются координатами точки, лежащей во второй координатной четверти".
2. В вещественном массиве известны данные о количестве осадков, выпавших за каждый день месяца N (N - любой месяц в году).
Найти общее число осадков, выпавших по четным числам месяца.
3. Дан целочисленный массив, состоящий из N элементов (N > 0). Найти сумму и произведение всех чисел из данного массива.
4. Написать функцию double DegToRad(D) вещественного типа, находящую величину угла в радианах, если дана его величина D в градусах
(D — вещественное число, 0 ≤ D ≤ 360). Воспользоваться следующим соотношением: 180° = pi радианов. В качестве значения PI
использовать предопределенную константу из библиотеки языка программирования.
5. Вводится строка, состоящая из слов, разделенных подчеркиваниями (одним или несколькими). Длина строки может быть разной.
Определить и вывести количество слов, которые содержат ровно одну букву 'h'.
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Задание 1")

	file, err := os.Open("input1.txt")
	if err != nil {
		fmt.Println("При открытии файла input1.txt возникла ошибка!")
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	var XY [2]float64
	for i := 0; i < 2 && scanner.Scan(); i++ {
		temp, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("При чтении файла input1.txt возникла ошибка!")
			return
		}
		XY[i] = temp
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("При чтении файла input1.txt возникла ошибка!")
		return
	}

	file, err = os.OpenFile("output1.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("При открытии файла output1.txt возникла ошибка!")
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	if XY[0] < 0 && XY[1] > 0 {
		if _, err = fmt.Fprint(file, "True"); err != nil {
			fmt.Println("При записи в файл output1.txt возникла ошибка!")
		}
	} else {
		if _, err = fmt.Fprint(file, "False"); err != nil {
			fmt.Println("При записи в файл output1.txt возникла ошибка!")
		}
	}

	fmt.Println("Файл output1.txt создан и записан.")

	fmt.Println("Задание 2")

	file2, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println("При открытии файла input2.txt возникла ошибка!")
	}

	var N = 0
	var arr [50]float64

	scanner = bufio.NewScanner(file2)

	for scanner.Scan() {
		temp, _ := strconv.ParseFloat(scanner.Text(), 64)
		arr[N] = temp
		N++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("При чтении файла input2.txt возникла ошибка!")
	}

	err = file2.Close()
	if err != nil {
		return
	}

	var sum float64
	for i := 0; i < N; i++ {
		if i%2 != 0 {
			sum += arr[i]
		}
	}

	file2, err = os.OpenFile("output2.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("При открытии файла output2.txt возникла ошибка!")
	}

	strtemp := "Число осадков по нечетным дням: " + fmt.Sprintf("%f", sum)
	_, err = io.WriteString(file2, strtemp)
	if err != nil {
		fmt.Println("При записи в файл output2.txt возникла ошибка!")
	}

	err = file2.Close()
	if err != nil {
		return
	}
	fmt.Println("Файл output2.txt создан и записан.")

	fmt.Println("Задание 3")

	file3, err := os.Open("input3.txt")
	if err != nil {
		log.Fatal("При открытии файла input3.txt возникла ошибка:", err)
	}
	defer func(file3 *os.File) {
		err := file3.Close()
		if err != nil {

		}
	}(file3)

	var arr1 []int
	scanner = bufio.NewScanner(file3)
	for scanner.Scan() {
		temp, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Ошибка при чтении числа!", err)
		}
		arr1 = append(arr1, temp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("При чтении файла input3.txt возникла ошибка!", err)
	}

	var summ, mull float64
	for _, val := range arr1 {
		summ += float64(val)
		mull *= float64(val)
	}

	strtemp = "Сумма чисел равна: " + fmt.Sprintf("%f", summ)
	strtemp1 := "\nПроизведение чисел равно: " + fmt.Sprintf("%f", mull)

	file3, err = os.OpenFile("output3.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatal("При открытии файла output3.txt возникла ошибка!", err)
	}
	defer func(file3 *os.File) {
		err := file3.Close()
		if err != nil {

		}
	}(file3)

	if _, err := io.WriteString(file3, strtemp); err != nil {
		log.Fatal("Ошибка при записи в файл!", err)
	}
	if _, err := io.WriteString(file3, strtemp1); err != nil {
		log.Fatal("Ошибка при записи в файл!", err)
	}

	fmt.Println("Файл output3.txt создан и записан.")

	fmt.Println("Задание 4")

	file4, err := os.Open("input4.txt")
	if err != nil {
		fmt.Println("При открытии файла input4.txt возникла ошибка!")
	}

	d, err := ioutil.ReadAll(file4)
	if err != nil {
		fmt.Println("При открытии файла input4.txt возникла ошибка!")
	}

	err = file4.Close()
	if err != nil {
		return
	}

	D, err := strconv.ParseFloat(string(d), 64)

	result := DegToRad(D)

	strtemp = "Угол в радианах: " + fmt.Sprintf("%f", result)

	file4, err = os.OpenFile("output4.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("При открытии файла output4.txt возникла ошибка!")
	}

	io.WriteString(file4, strtemp)

	fmt.Println("Файл output4.txt создан и записан.")

	fmt.Println("Задание 5")

	file5, err := os.Open("input5.txt")
	if err != nil {
		log.Fatal("При открытии файла input5.txt возникла ошибка!", err)
	}
	defer func(file5 *os.File) {
		err := file5.Close()
		if err != nil {

		}
	}(file5)

	str, err := ioutil.ReadAll(file5)
	if err != nil {
		log.Fatal("При чтении файла input5.txt возникла ошибка!", err)
	}

	strnew := strings.Split(string(str), "_")

	var count int

	for _, s := range strnew {
		if strings.Count(s, "h") == 1 {
			count++
		}
	}

	file5, err = os.OpenFile("output5.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatal("При открытии файла output5.txt возникла ошибка!", err)
	}
	defer func(file5 *os.File) {
		err := file5.Close()
		if err != nil {

		}
	}(file5)

	_, err = fmt.Fprintf(file5, "Количество слов содержащих ровно одну букву 'h': %d", count)
	if err != nil {
		log.Fatal("При записи в файл output5.txt возникла ошибка!", err)
	}

	fmt.Println("Файл output5.txt создан и записан.")

}

func DegToRad(deg float64) float64 {
	return deg * math.Pi / 180
}

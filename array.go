package main

import "fmt"

//"math"
func main() {

	

	a := []int{4, 5, 6}
	b := make([]int, 3, 3)
	n := copy(a, b)
	fmt.Printf("a = %v\n", a)                  // a = [1 2 3]
	fmt.Printf("b = %v\n", b)                  // b = [1 2 3]
	fmt.Printf("Скопировано %d элемента\n", n) // Скопировано 3 элемента
	//=====================
	// arr := []int{1,2,3,4,5,6,7,8,9,10}
	// arr = append(arr[0:2], arr[3:]...)
	// fmt.Println(arr)
	//----------------------------
	// arr := []int{1,2,3,4,5,6,7,8}
	// arr = append(arr,9,10)
	// fmt.Println(arr)
	// user1 := arr[2:6]
	// user2 := arr[:4]
	// user3 := arr[3:]
	// fmt.Println(user1)
	// fmt.Println(user2)
	// fmt.Println(user3)
	//======================
	// a := [3]int {1,2,3}
	// b := [3]int {1,2,3}
	// c := [3]int {3,2,1}
	// fmt.Println(a==b)
	// fmt.Println(a==c)
	//=======================
	// baseArray := [10]int{0,1,2,3,4,5,6,7,8,9}
	// fmt.Printf("Базовый массив: %v\n", baseArray)
	// baseSlice := baseArray[5:8]
	// fmt.Printf(
	// 	"Срез, основанный на базовом массиве длиной %d и емкостью %d: %v\n",
	// len(baseSlice),
	// cap(baseSlice),
	// baseSlice)
	// baseSlice = append(baseSlice,8,9,10)
	// fmt.Println(baseSlice)
	//=====================
	// var num [5] int =  [5] int{1,2,3,4,5}
	// fmt.Println(num[0])
	// fmt.Println(num[4])
	// num[0] = 21
	// fmt.Println(num)

	//==========================
	// a := [5]int {1,2,3,4,5}
	// fmt.Println(a)
	//  for i := 0; i<len(a); i++{
	// 	fmt.Println(a[i])
	//  }

	//==================
	// a := [5]int {1,2,3,4,5}
	// fmt.Println(a)
	// //При итерации по массиву мы можем использовать ключевое слово range, тогда цикл будет иметь следующий вид:
	// for idx, elem := range a{
	// 	fmt.Printf("Element s indeksom %d: %d\n", idx, elem)
	// }

	//=============
	// a := [5]int {1,2,3,4,5}
	// for idx := range a{
	// 	fmt.Println(a[idx])
	// }
	// for idx, _ := range a{
	// 	fmt.Println(a[idx])
	// }
	// for _, elem := range a{
	// 	fmt.Println(a[elem])
	// }
	//================================
	// var n  int
	// s := 0
	// fmt.Scan(&n)
	// for i := n; i>0;i/=10  {
	//     s = s*10+i%10

	// }
	// fmt.Println(s)
	//====================================
	// var k, h, m int
	// fmt.Scan(&k)
	// if k <= 0 || k > 86399 {
	// 	fmt.Println("In khato ast,boyad 0<k<86399 boshad:")
	// 	return
	// }
	// h = k / 3600
	// m = k % 3600 / 60
	// fmt.Println("It is", h, "hours", m, "minute")

	//===========================
// 	var n int
//     s := 0
//     for  {
//         fmt.Scan(&n)
//         if n==0 {
//         s++
//     } else {
// 		break

// 	}
	
//   } 
//   fmt.Scanln()
//   fmt.Println(s)
  
	//====================================
	//Array1. Дано целое число N (> 0).
	//Сформировать и вывести целочисленный массив размера N, содержащий N первых положительных нечетных чисел: 1, 3, 5, … .
	// var N int
	// fmt.Println("Введите размер массива: ")
	// fmt.Scan(&N)
	// arr :=make ([]int , N)
	// num := 1
	// for i := 0; i<N; i++ {
	// 	arr[i] = num
	// 	num += 2
	// }
	// fmt.Println(arr)

	//====================
	//Array2. Дано целое число N (> 0). Сформировать и вывести целочисленный массив размера N,
	// содержащий степени двойки от первой до N-й: 2, 4, 8, 16, … .
	// var n int
	// fmt.Println("Введите размер массива: ")
	// fmt.Scan(&n)
	// arr := make ([]int, n)
	// num := 2
	// s := 1
	// for i := 0; i<n; i++ {
	// 	s = s * num
	// 	arr[i] = s
	// }
	// fmt.Println(arr)

	//===============================
	//Array3. Дано целое число N (> 1), а также первый член A и разность D арифметической прогрессии.
	// Сформировать и вывести массив размера N, содержащий N первых членов данной прогрессии: A,    A + D,    A + 2·D,    A + 3·D,    … .
	// var n int
	// var a,d float64
	// fmt.Println("Введите размер массива: ")
	// fmt.Scan(&n)
	// fmt.Println("Введите a: ")
	// fmt.Scan(&a)
	// fmt.Println("Введите d: ")
	// fmt.Scan(&d)
	// if n<=1 {
	// 	fmt.Println("In khato ast, boyad n > 1 boshad: ")
	// 	return
	// }
	// arr := make ([]float64, n)
	// s := 0.0
	// for i:=0; i<n; i++ {
	// 	s = a + (float64(i) * d)
	// 	arr[i] = s
	// }
	// fmt.Println(arr)

	//Array4°. Дано целое число N (> 1), а также первый член A и знаменатель D геометрической прогрессии.
	//Сформировать и вывести массив размера N, содержащий N первых членов данной прогрессии: A,    A·D,    A·D^2,    A·D^3,    … .
	// var n int
	// var a,d float64
	// fmt.Println("Введите размер массива: ")
	// fmt.Scan(&n)
	// fmt.Println("Введите a: ")
	// fmt.Scan(&a)
	// fmt.Println("Введите d: ")
	// fmt.Scan(&d)
	// if n<=1 {
	// 	fmt.Println("In khato ast, boyad n > 1 boshad: ")
	// 	return
	// }
	// arr := make([]float64, n)
	// s := 0.0
	// for i := 0; i<n; i++ {
	// 	s = a * math.Pow(d,float64(i))
	// 	arr[i] = s
	// }
	// fmt.Println(arr)

	//=================================  ???? [0 1 1 2 3 5 8 13 21] bovad chunin baroyad: [ 1 1 2 3 5 8 13 21 34]
	// Array5. Дано целое число N (> 2). Сформировать и вывести целочисленный массив размера N,
	// содержащий N первых элементов последовательности чисел Фибоначчи FK: F1 = 1,  F2 = 1,  FK = FK−2 + FK−1, K = 3, 4, … .
	// var n int
	// fmt.Println("Vvedite razmera massiva : ")
	// fmt.Scan(&n)
	// if n<=2 {
	// 	fmt.Println("In khato ast, boyad n > 0 boshad")
	// 	return
	// }
	// F1 := 1
	// F2 := 1
	// Fk := 0
	// arr := make ([]int, n)
	// arr[0] = F1
	// arr[1] = F2
	// for i := 2; i<n; i++ {
	// 		Fk = F1+F2
	// 		arr[i] = Fk
	// 		F1 = F2
	// 		F2 = Fk
	// }
	// fmt.Println(arr)

	//=============================
	//Array6. Даны целые числа N (> 2), A и B.
	//Сформировать и вывести целочисленный массив размера N, первый элемент которого равен A, второй равен B,
	//а каждый последующий элемент равен сумме всех предыдущих.
	// 	var n, a, b int
	// fmt.Println("Vvedite razmera massiva : ")
	// fmt.Scan(&n)
	// if n <= 2 {
	//     fmt.Println("In khato ast, boyad n > 0 boshad")
	//     return
	// }
	// fmt.Println("Vvedite a:")
	// fmt.Scan(&a)
	// fmt.Println("Vvedite b:")
	// fmt.Scan(&b)
	// arr := make([]int, n)
	// arr[0] = a
	// arr[1] = b
	// for i := 2; i < n; i++ {
	//     arr[i] = 0
	//     for j := 0; j < i; j++ {
	//         arr[i] += arr[j]
	//     }
	// }
	// fmt.Println(arr)

	//Одномерные массивы: вывод элементов
	//Array7°. Дан массив размера N. Вывести его элементы в обратном порядке.
	// var n int
	// fmt.Println("Vvedite razmera massiva : ")
	// fmt.Scan(&n)
	// arr := [5]int{1,2,3,4,5}
	// for i := len(arr)-1; i>=0; i-- {
	// 	fmt.Println(arr[i])
	// }

	//========================

	//Array8.Дан целочисленный массив размера N.
	// Вывести все содержащиеся в данном массиве нечетные числа в порядке возрастания их индексов, а также их количество K.
	// var n, k int
	// fmt.Println("Vvedite razmera massiva : ")
	// fmt.Scan(&n)
	// for i:=1; i<n; i++ {
	// 	if i%2==1 {
	// 		fmt.Println(i)
	// 		k ++
	// 	}

	// }
	// fmt.Println("количество k:", k)

	//=========================
	//Array9. Дан целочисленный массив размера N. Вывести все содержащиеся в данном массиве четные числа в порядке убывания их индексов,
	// а также их количество K.
	// var n, k int
	// fmt.Println("Vvedite razmera massiva : ")
	// fmt.Scan(&n)
	// for i := 0; i<n; i++ {
	// 	if i%2==0 {
	// 		fmt.Println(i)
	// 		k++
	// 	}
	// }
	// fmt.Println("количество k :",k)

	//---------------------------------------
	//Array10. Дан целочисленный массив размера N.
	//Вывести вначале все содержащиеся в данном массиве четные числа в порядке возрастания их индексов,
	//а затем — все нечетные числа в порядке убывания их индексов.

	//что такое map и как ей ползуется в Golang

	
}

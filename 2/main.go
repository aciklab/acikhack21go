package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Ali"
	fmt.Println(name)
	name = "Veli"
	fmt.Println(name)
	// name = 3
	fmt.Println(reflect.TypeOf(name))

	var isim string
	isim = "Test"
	fmt.Println(isim)

	///////
	var numbers []int                     // slice
	numbers = append(numbers, 2)          // slice boş tanımlanırsa append çalışır
	numbers = append(numbers, 3, 4, 5, 6) // birden fazla eleman append leyebiliriz
	fmt.Println(numbers)

	names := make([]string, 2)
	names = append(names, "Ali")
	names = append(names, "Veli")
	fmt.Println(names)
	fmt.Println(names[0])
	// boş olmasının sebebi 1 büyüklüğünde slice tanımlamış olmamız
	// append tanımlanan kısım sonrasına atıyor

	for i, val := range names {
		fmt.Println(i, val)
	}

	//var employeeSalary map[string]int
	//employeeSalary := map[string]int{}
	employeeSalary := map[string]int{
		"Ali":  1000,
		"Veli": 2000,
	}
	employeeSalary["Mehmet"] = 5000
	fmt.Println(employeeSalary)

	employeeSalary2 := map[string]int{}
	employeeSalary2["Emin"] = 2000
	fmt.Println(employeeSalary2)
}

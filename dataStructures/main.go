package main

import "fmt"

/*
   - array
   - slice     -> a list
   - map       -> key / value storage
   - struct    -> a data structure, individual object
*/
func main() {

	// array
	var a [58]string
	fmt.Println(a)      // [   ]
	fmt.Println(len(a)) // 58
	fmt.Println(a[42])  //
	for i := 65; i < 122; i++ {
		a[i-65] = string(i)
	}
	fmt.Println(a)      // [A B C D E F ...]
	fmt.Println(len(a)) // 58
	fmt.Println(a[42])  // k

	// slice
	mySlice := []int{1, 3, 5, 7, 9, 11, 42}
	fmt.Printf("%T \n", mySlice)
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])
	fmt.Println(mySlice[5])

	mySlice2 := make([]int, 0, 5)
	fmt.Println("capacity ", cap(mySlice2))
	fmt.Println("length ", len(mySlice2))
	// if slice at the end of cap it will double 5 -> 10 -> 20 -> ...
	for i := 0; i < 80; i++ {
		mySlice2 = append(mySlice2, i)
		fmt.Println("Len:", len(mySlice2), "Capacity:", cap(mySlice2), "Value:", mySlice2[i])
	}

	// 2d slice
	records := make([][]string, 0)
	// student1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"
	// store
	records = append(records, student1)
	// student2
	student2 := make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.00"
	student2[3] = "24.00"
	// store
	records = append(records, student2)
	// Print
	fmt.Println(records)

	// maps
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 42
	m["test"] = 66

	delete(m, "test")
	fmt.Println("map:", m)
	fmt.Println(m["k1"])

	// maps
	mm := map[string]int{"foo": 1, "bar": 2, "bim": 3}
	fmt.Println(mm)
	_, ok := mm["foo"]
	fmt.Println("ok:", ok, "len", len(mm))

	// maps
	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)

	// maps example 2
	greeting := map[int]string{
		0: "Good morning!",
		1: "Bonour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}
	fmt.Println(greeting)

	if value, exists := greeting[2]; exists {
		delete(greeting, 2)
		fmt.Println("val:", value)
		fmt.Println("exists:", exists)
	} else {
		fmt.Println("That value doesnt exist")
		fmt.Println("val:", value)
		fmt.Println("exists:", exists)
	}

	// for loop through map
	for key, val := range greeting {
		fmt.Println(key, " - ", val)
	}
}

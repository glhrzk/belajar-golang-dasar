package main

import "fmt"

func main() {
	/*	slice ada potongan data dari array

		membuat slice array dari index low sampai index sebelum high
		array[low:high]
		array[:high]

		membuat slice array dari index low sampai index di akhir array
		array[low:]
		array[:]
	*/

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	fmt.Println("Array months: ", months)

	var slice1 = months[4:7]
	fmt.Println("Slice1: ", slice1)
	fmt.Println("Len Slice1: ", len(slice1))
	fmt.Println("Cap Slice1: ", cap(slice1))

	//months[4] = "Mei-lagi"
	//fmt.Println("Slice1: ", slice1)

	slice1[0] = "Mei-slice"
	fmt.Println("Array months: ", months)
	fmt.Println("Slice1: ", slice1)

	var slice2 = months[10:]
	fmt.Println("Slice2: ", slice2)

	// ketika meng append dan capacity array tersebut sudah full/penuh maka akan terputus dari referensi array sumber, dan akan membuat array baru.
	var slice3 = append(slice2, "galih rizki")
	fmt.Println("Slice3: ", slice3)

	slice3[0] = "November-slice3"
	fmt.Println("Slice3 : ", slice3)
	fmt.Println(months)

	// membuat slice baru tanpa referensi ke array yang sudah ada
	newSlice := make([]string, 2, 5)
	newSlice[0] = "galih"
	newSlice[1] = "rizki"

	fmt.Println("newSlice: ", newSlice)

	// copy slice, harus identik size
	copySlice := make([]string, len(newSlice), cap(newSlice))
	// sourceCopy to destination
	copy(copySlice, newSlice)
	fmt.Println("copySlice: ", copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println("ini adalah array", iniArray)
	fmt.Println("ini adalah slice", iniSlice)
}

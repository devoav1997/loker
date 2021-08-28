package main

import "fmt"

// type VSlice []int
// func (s VSlice) String() string {
//     var str string
//     for _, i := range s {
//         str += fmt.Sprintf("%s\n", i)
//     }
//     return str
// }

func main() {

	var size int
	fmt.Scanln(&size)
	fmt.Printf("%d", size)
	fmt.Println("")



	var arr = make([]int, size)

	var tipeidentitas string
	var arrTipeIdentitas []string


	switch loker := size; loker {
	case loker:

		fmt.Printf("Berhasil membuat loker dengan jumlah %v\n", len(arr))

		// var arrTipeIdentitas = make(map[string]string)

		for i := 0; i < size; i++ {
			fmt.Print("input ")
			fmt.Scanf("%s %d", &tipeidentitas, &arr[i])
			// fmt.Scanln(&tipeidentitas , &nomoridentitas)
		
			// for _, item := range arrTipeIdentitas {

			// 	// Prints index of all the
			// 	// characters in the string
			// 	fmt.Printf("%s \n", item)
			// }
			fmt.Printf("Kartu identitas tersimpan di loker nomor %d \n", i + 1)

		}

		fmt.Println("Your array is: ", arr)


		var statuss string
		fmt.Scanln(&statuss)

		switch status := statuss; status {
		case status:
			arrTipeIdentitas = []string{tipeidentitas}
			for _, item2 := range arr {
			for _, item := range arrTipeIdentitas {
					fmt.Printf("input %s %d \n", item, item2)
				}	

			}
		default:
			fmt.Println("Invalid")
		}
	default:
		fmt.Println("Invalid")
	}

}

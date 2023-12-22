package main

func main() {
	myList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	search := 10

	test := simplySearch(myList, search)
	test2 := binarySearch(myList, search)

	println(test)
	println(test2)
}

func simplySearch(list []int, item int) int {
	for i, v := range list {
		println("entrou no simples ", i)
		if v == item {
			return i
		}
	}

	return -1
}

func binarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		println("entrou no binário")
		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			return mid
		}

		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}

	return -1
}

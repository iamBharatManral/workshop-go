package main

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				t := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = t
			}
		}
	}
}

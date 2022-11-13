package main

func main() {
	slice := []int{1, 3, 4, 6, 2, 5, 4, 1, 9}

	visited := make(map[int]bool)
	
	for i := 0; i < len(slice); i++ {
		if visited[slice[i]] == true {
			println("duplicate value is", slice[i])
		} else {
			visited[slice[i]] = true
		}
	}

}

package main

// For an Unsorted data, we have to loop through the array to find it.
func linearSearch (data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++{
		if data[i] == value{
			return true
		}
		
	}
	return false
}
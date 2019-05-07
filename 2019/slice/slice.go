package slice

func SplitSlice(slice []int, n int) [][]int {
	var sTemp [][]int
	s := len(slice) / n
	for i := 0; i < s; i++ {
		ll := slice[i*n : (i+1)*n]
		sTemp = append(sTemp, ll)
	}
	sTemp = append(sTemp, slice[s*n:])
	return sTemp
}
func SplitSliceLen(slice []int, length int) [][]int {
	var sTemp [][]int
	n := len(slice) / length
	for i := 0; i < n; i++ {
		ll := slice[i*length : (i+1)*length]
		sTemp = append(sTemp, ll)
	}
	sTemp = append(sTemp, slice[length*n:])
	return sTemp
}

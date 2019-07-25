package _struct

func AppendSlice(list []int, i int) bool {
	if len(list) > 10 {
		return true
	} else {
		list = append(list, i)
		return AppendSlice(list, i+1)
	}
}

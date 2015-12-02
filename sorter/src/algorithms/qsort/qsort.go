package qsort

func qulicksort(values []int){
	if len(values) <= 1{
		return
	}

	key, p := values[0], 1
	head, tail := 0, len(values)-1

	for head < tail{
		if values[p] < key{
			values[p], values[tail] = values[tail], values[p]
			tail--
		} else {
			values[p], values[head] = values[head], values[p]
			p++
		}
	}
	values[head] = key
	qulicksort(values[:head])
	qulicksort(values[head+1:])	
}
func QuickSort(values []int){
	println("Thanks use QuickSort")
	qulicksort(values)
}
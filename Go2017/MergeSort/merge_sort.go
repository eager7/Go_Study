package MergeSort


func MergeSort(list []int)[]int{
	length := len(list)
	if length < 2{
		return list
	}else if length == 2{
		listRet := make([]int, 2)
		if list[0] > list[1]{
			listRet[0], listRet[1] = list[1], list[0]
		}else{
			listRet[1], listRet[0] = list[1], list[0]
		}
		return listRet
	}
	list1, list2 := list[:length/2], list[length/2:]
	temp1 := MergeSort(list1)
	temp2 := MergeSort(list2)

	listEnd := Merge(temp1, temp2)
	return listEnd
}

func Merge(list1, list2 []int)[]int{
	l := make([]int, 0)
	for i, length := 0, len(list1)+len(list2); i < length; i ++ {
		if list1[0] < list2[0]{
			l = append(l , list1[0])
			if len(list1) == 1{
				l = append(l, list2...)
				break
			}
			list1 = list1[1:]
		}else {
			l = append(l, list2[0])
			if len(list2) == 1{
				l = append(l, list1...)
				break
			}
			list2 = list2[1:]
		}
	}
	return l
}

func MergeSortGo(list []int, chRet chan []int){
	length := len(list)
	if length == 1{
		chRet <- list
		return
	}else if length == 2{
		listRet := make([]int, 2)
		if list[0] > list[1]{
			listRet[0], listRet[1] = list[1], list[0]
		}else{
			listRet[1], listRet[0] = list[1], list[0]
		}
		chRet <- listRet
		return
	}
	list1, list2 := list[:length/2], list[length/2:]
	chRet1, chRet2 := make(chan []int), make(chan []int)
	go MergeSortGo(list1, chRet1)
	go MergeSortGo(list2, chRet2)

	temp1, temp2 := <-chRet1, <-chRet2

	go MergeGo(temp1, temp2, chRet)
	return
}

func MergeGo(list1, list2 []int, chRet chan []int){
	l := make([]int, 0)
	for i, length := 0, len(list1)+len(list2); i < length; i ++ {
		if list1[0] < list2[0]{
			l = append(l , list1[0])
			if len(list1) == 1{
				l = append(l, list2...)
				break
			}
			list1 = list1[1:]
		}else {
			l = append(l, list2[0])
			if len(list2) == 1{
				l = append(l, list1...)
				break
			}
			list2 = list2[1:]
		}
	}
	chRet <- l
}

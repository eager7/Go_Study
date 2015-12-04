package bubblesort

func BubbleSort(values []int){
	println("Thanks use BubbleSort, the len is", len(values))
	
	for i:=0; i < len(values) - 1; i++{
		for j:=0; j < len(values) - i - 1;j++{
			if values[j] > values[j+1]{
				values[j], values[j+1] = values[j+1],values[j]
			}
		}
	}
}
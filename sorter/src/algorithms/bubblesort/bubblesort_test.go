package bubblesort
import "testing"

func TestBubbleSort1(t *testing.T){
	values := []int {5,4,3,2,1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5{
		t.Error("BubbleSort Test Error. Got", values, "Expected 1 2 3 4 5")
		return
	}else{
		println("Test1 Pass")
	}
}
func TestBubbleSort2(t *testing.T){
	values := []int {6,6,3,2,1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 6 || values[4] != 6{
		t.Error("BubbleSort Test Error. Got", values, "Expected 1 2 3 6 6")
		return 
	}else{
		println("Test2 Pass")
	}
}
func TestBubbleSort3(t *testing.T){
	values := []int {5}
	BubbleSort(values)
	if values[0] != 5{
		t.Error("BubbleSort Test Error. Got", values, "Expected 5")
		return
	}else{
		println("Test3 Pass")
	}
}
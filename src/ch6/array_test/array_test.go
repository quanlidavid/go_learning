package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 3, 5, 7, 9}
	arr1[1] = 99

	t.Log(arr, arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr2 := [...]int{1, 3, 5, 7, 9}

	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i])
	}

	for idx /*索引*/, e /*元素*/ := range arr2 {
		t.Log(idx, e)
	}

	for _ /*占位符*/, e /*元素*/ := range arr2 {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr2 := [...]int{1, 3, 5, 7, 9}
	//arr2Sec := arr2[:3]
	arr2Sec := arr2[:]
	t.Log(arr2Sec)
}

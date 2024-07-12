package MaxCounters

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxCounters_t1(t *testing.T) {
	t1 := MaxCounters(5, []int{3, 4, 4, 6, 1, 4, 4})
	pippo := reflect.DeepEqual(t1, []int{3, 2, 2, 4, 2})
	if pippo == false {
		t.Fatalf(`"Expected ",%v, ", your result however is:", %v`, []int{3, 2, 2, 4, 2}, t1)
	}
	fmt.Println("Correct result : ", []int{3, 2, 2, 4, 2})
}

func TestMaxCounters_Empty(t *testing.T) {
	t1 := MaxCounters(5, []int{})
	pippo := reflect.DeepEqual(t1, []int{0, 0, 0, 0, 0})
	if pippo == false {
		t.Fatalf(`"Expected ",%v, ", your result however is:", %v`, []int{0, 0, 0, 0, 0}, t1)
	}
}

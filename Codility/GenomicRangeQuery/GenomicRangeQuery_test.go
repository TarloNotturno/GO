package GenomicRangeQuery

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
)

func TestMaxCounters_t1(t *testing.T) {
	val, error := GenomicRangeQuery("CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6})
	if error == nil {
		fmt.Println("The weight related to 'CAGCCTA' for input", []int{2, 5, 0}, []int{4, 5, 6}, "is equal to:\n", val)
	} else {
		t.Fatalf(`"Unexpected error"`)
		log.Fatal(error)
	}

}

func TestMaxCounters_tEmptyInput(t *testing.T) {
	_, error := GenomicRangeQuery("CAGCCTA", []int{}, []int{4, 5, 6})
	if error == nil {
		t.Fatalf(`"P and Q has different dimension, expected an error dealing but no error returned"`)
		log.Fatal(error)
	}

}

func TestMaxCounters_tRandomInput(t *testing.T) {
	P := []int{4, 5, 6}
	Q := []int{22, 12, 33}
	for i := 0; i < 40; i++ {
		adding_val := rand.Intn(int((len("CAGTCCTACCAGGGGTCCTACAGCAGTCCTACCAGGGGTCCTACAGCAGTCCTACCAGGGGTCCTACAGCAGTCCTACCAGGGGTCCTACAG") - 2) / 2))
		P = append(P, adding_val)

		Q = append(Q, adding_val*2)
	}
	fmt.Println(P)
	val, error := GenomicRangeQuery("CAGTCCTACCAGGGGTCCTACAGCAGTCCTACCAGGGGTCCTACAGCAGTCCTACCAGGGGTCCTACAGCAGTCCTACCAGGGGTCCTACAG", P, Q)
	if error != nil {
		t.Fatalf(`"No error should be present"`)
		log.Fatal(error)
	} else {
		fmt.Println(val)
	}

}

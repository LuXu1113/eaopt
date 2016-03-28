package gago

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestTournament(t *testing.T) {
	var source = rand.NewSource(time.Now().UnixNano())
	var generator = rand.New(source)
	var size = 3
	var indis = make(Individuals, size)
	for i := 0; i < size; i++ {
		indis[i] = Individual{make([]interface{}, 2), float64(i)}
	}
	var original = make([]Individual, len(indis))
	copy(original, indis)
	var selector = Tournament{3}
	var sample = selector.apply(size, indis, generator)
	// Check the size of the sample
	if len(sample) != size {
		t.Error("Wrong sample size")
	}
	// Check the original population hasn't changed
	for i := range indis {
		if reflect.DeepEqual(indis[i], original[i]) == false {
			t.Error("Population has been modified")
		}
	}
	// Check the individual is from the initial population
	if reflect.DeepEqual(indis[0], sample[0]) == false {
		t.Error("Problem with tournament selection")
	}
}
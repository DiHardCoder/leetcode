package leetcode815

import (
	"fmt"
	"testing"
)

func TestTrainCount(t *testing.T) {
	A := [][]int{{1, 2, 7}, {5, 3, 6, 7}}
	cnt := GetTrainsCount(A, 1, 5)
	if 2 != cnt {
		t.Error(fmt.Sprintf("getting wrong result :%d", cnt))
		return
	}
}

func TestTrainCountNegativeScenario(t *testing.T) {
	A := [][]int{{1, 2, 7}, {5, 3, 6, 7}}

	if -1 != GetTrainsCount(A, 1, 8) {
		t.Error("getting wrong result")
		return
	}
}

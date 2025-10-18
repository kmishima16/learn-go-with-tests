package array

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	t.Run("複数のスライスの合計を計算する", func(t *testing.T) {

		got := SumAll([]int{1, 2}, []int{3, 4})
		want := []int{3, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}

func TestSum(t *testing.T) {
	t.Run("配列の合計を計算する", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers[:])
		want := 15

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("スライスの合計を計算する", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

package array

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {

	checkSum := func(t *testing.T, got, want []int) {
		t.Helper()
		for !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("複数の配列の末尾以外の合計を計算する", func(t *testing.T) {
		numbers1 := [4]int{1, 2, 3, 4}
		numbers2 := [3]int{0, 9, 8}
		got := SumAllTails(numbers1[:], numbers2[:])
		want := []int{9, 17}

		checkSum(t, got, want)
	})

	t.Run("空の配列を含む場合", func(t *testing.T) {
		numbers1 := [4]int{1, 2, 3, 4}
		numbers2 := [0]int{}
		got := SumAllTails(numbers1[:], numbers2[:])
		want := []int{9, 0}

		checkSum(t, got, want)
	})
}

func TestSumAll(t *testing.T) {

	checkSum := func(t *testing.T, got, want []int) {
		t.Helper()
		for !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("複数の配列の合計を計算する", func(t *testing.T) {
		numbers1 := [3]int{1, 2, 3}
		numbers2 := [2]int{4, 5}
		got := SumAll(numbers1[:], numbers2[:])
		want := []int{6, 9}

		checkSum(t, got, want)
	})

	t.Run("空の配列を含む場合", func(t *testing.T) {
		numbers1 := [3]int{1, 2, 3}
		numbers2 := [0]int{}
		got := SumAll(numbers1[:], numbers2[:])
		want := []int{6, 0}

		checkSum(t, got, want)
	})
}

func TestSum(t *testing.T) {
	t.Run("配列の合計を計算する", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := Sum(numbers[:])
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

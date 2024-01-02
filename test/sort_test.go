package test_demo

import (
	c "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"os"
	"testing"
)

func TestSort(t *testing.T) {
	// 只需要在顶层的Convey调用时传入t
	c.Convey("测试两个用例", t, func() {
		tests := []struct {
			name  string
			input []int
			want  []int
		}{
			{"equal1", []int{34, 54, 5645, 2, 423}, []int{2, 34, 54, 423, 5645}},
			{"equal2", []int{3, 5, 1, 2, 6, 7}, []int{1, 2, 3, 5, 6, 7}},
		}
		for _, test := range tests {
			c.Convey(test.name, func() {
				c.So(Sort(test.input), c.ShouldEqual, test.want)
			})
		}
	})
}

func BenchmarkSort(b *testing.B) {
	benchmarkSort(b, 100000000)
}
func BenchmarkSortParallel(b *testing.B) {
	benchmarkSortParallel(b, 100)
}

func benchmarkSort(b *testing.B, n int) {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(10000)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sort(arr)
	}
}

func benchmarkSortParallel(b *testing.B, n int) {
	b.SetParallelism(4) // 设置使用的CPU数
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(10000)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Sort(arr)
		}
	})
}
func BenchmarkSort10(b *testing.B) {
	benchmarkSort(b, 10)
}

func BenchmarkSort100(b *testing.B) {
	benchmarkSort(b, 100)
}

func BenchmarkSort1000(b *testing.B) {
	benchmarkSort(b, 1000)
}
func setup() {
	//fmt.Println("111")
}
func teardown() {
	//fmt.Println("222")
}
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

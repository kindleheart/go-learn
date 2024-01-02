package func_demo

import (
	"github.com/agiledragon/gomonkey"
	"reflect"
	"strings"
	"testing"
)

// func_test.go

func TestFunc(t *testing.T) {
	// 对 GetInfoByUID 进行打桩
	patches1 := gomonkey.ApplyFunc(GetInfoByID, func(id int64) (*UserInfo, error) {
		return &UserInfo{
			Name: "jason",
			Age:  18,
		}, nil
	})
	defer patches1.Reset()
	user := &UserInfo{}

	// 对UserInfo.SelectCourse打桩
	patches2 := gomonkey.ApplyMethod(reflect.TypeOf(user), "SelectCourse", func(*UserInfo, int64) string {
		return "math"
	})
	defer patches2.Reset()

	ret := MyFunc(123, 222)
	if !strings.Contains(ret, "jason") || !strings.Contains(ret, "math") {
		t.Fatal()
	}
}

func TestFuncSeq(t *testing.T) {
	// 对函数打一个序列桩
	outputs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{&UserInfo{
			Name: "jason",
			Age:  18,
		}, nil}},
		{Values: gomonkey.Params{&UserInfo{
			Name: "jay",
			Age:  25,
		}, nil}},
	}
	patches := gomonkey.ApplyFuncSeq(GetInfoByID, outputs)
	defer patches.Reset()
	ret := MyFunc(111, 1111)
	if !strings.Contains(ret, "jason") {
		t.Fatal()
	}
	ret = MyFunc(222, 2222)
	if !strings.Contains(ret, "jay") {
		t.Fatal()
	}
}

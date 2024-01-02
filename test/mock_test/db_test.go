package interface_demo

import (
	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"goLearn/test/mock_test/mocks"
	"reflect"
	"testing"
)

// db_test.go

func TestGetFromDB(t *testing.T) {
	// 创建gomock控制器，用来记录后续的操作信息
	ctrl := gomock.NewController(t)
	// 断言期望的方法都被执行
	// Go1.14+的单测中不再需要手动调用该方法
	defer ctrl.Finish()
	// 调用mockgen生成代码中的NewMockDB方法
	// 这里mocks是我们生成代码时指定的package名称
	m := mocks.NewMockDB(ctrl)
	// 打桩（stub）
	// 当传入Get函数的参数为liwenzhou.com时返回1和nil
	m.
		EXPECT().
		Get("age").      // 参数
		Return(18, nil). // 返回值
		Times(1)         // 调用次数
	m.
		EXPECT().
		Add("age", 18).
		Return(nil).
		Times(1)
	// 调用GetFromDB函数时传入上面的mock对象m
	assert.Equal(t, nil, AddToDB(m, "age", 18))
	assert.Equal(t, 18, GetFromDB(m, "age"))
}

func TestInterface(t *testing.T) {
	myDB := &MyDB{}
	gomonkey.ApplyFunc(NewDB, func(string) DB {
		return myDB
	})
	db := NewDB("my")
	gomonkey.ApplyMethod(reflect.TypeOf(myDB), "Get", func(*MyDB, string) (int, error) {
		return 18, nil
	})
	gomonkey.ApplyMethod(reflect.TypeOf(myDB), "Add", func(*MyDB, string, int) error {
		return nil
	})
	assert.Equal(t, nil, AddToDB(db, "age", 18))
	assert.Equal(t, 18, GetFromDB(db, "age"))
}

package func_demo

import (
	"fmt"
)

type UserInfo struct {
	Name string
	Age  int
}

func (u *UserInfo) SelectCourse(id int64) string {
	return ""
}

func GetInfoByID(id int64) (*UserInfo, error) {
	// 还没写完
	return &UserInfo{}, nil
}

func MyFunc(uid, courseID int64) string {
	u, err := GetInfoByID(uid)
	if err != nil {
		return "welcome"
	}
	courseName := u.SelectCourse(courseID)
	// 这里是一些逻辑代码...

	return fmt.Sprintf("hello %s,  you select course %s\n", u.Name, courseName)
}

package mygomock

import (
	gen "test/mygomock/gen" //包名以test/开头,就可以找到当前项目的目录
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

//生成命令 cd src/gomock && mocken . Foo
//mockgen -source=foo.go -destination=gen.go Foo
// vscode 运行要以  "program": "${workspaceFolder}" 方式
//代码放在 mygomock 目录中,目录叫gomock不行的
func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := gen.NewMockFoo(ctrl)

	// Does not make any assertions. Executes the anonymous functions and returns
	// its result when Bar is invoked with 99.
	m.
		EXPECT().
		Bar(gomock.Eq(99)).
		DoAndReturn(func(_ int) int {
			time.Sleep(1 * time.Second)
			return 101
		}).
		AnyTimes()

	// Does not make any assertions. Returns 103 when Bar is invoked with 101.
	m.
		EXPECT().
		Bar(gomock.Eq(101)).
		Return(103).
		AnyTimes()

	SUT(m)
}

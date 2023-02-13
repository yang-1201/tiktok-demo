package testpackage

import (
	"testing"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	//"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	defer logs.Flush()
	//测试增删改查
	TestLikesInsert(t)
	TestLikesQuery(t)
	TestVideoUpdate(t)
	TestVideoDelete(t)
	logs.Infof("test go package")
}

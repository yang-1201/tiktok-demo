package testpackage

import (
	"testing"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
<<<<<<< HEAD
	//"github.com/stretchr/testify/assert"
=======
	"github.com/stretchr/testify/assert"
>>>>>>> f3bcb08 (publish完成)
)

func TestAll(t *testing.T) {
	defer logs.Flush()
<<<<<<< HEAD
	//测试增删改查
	TestLikesInsert(t)
	TestLikesQuery(t)
	TestVideoUpdate(t)
	TestVideoDelete(t)
=======
	assert.Equal(t, 1, 2)
>>>>>>> f3bcb08 (publish完成)
	logs.Infof("test go package")
}

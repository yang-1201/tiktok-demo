package mysql

import (
	//"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDatabaseOp(t *testing.T) {
	InitDB()
	err := CreateLike(123, 321, 111)
	assert.NoError(t, err)
	id, err := FindIDinLike(321, 111)
	assert.NoError(t, err)
	del := DelLike(id)
	assert.Nil(t, del)
}

func TestData(t *testing.T) {
	//defer logs.Flush()
	InitDB()
	// err := CreateLike(123, 321, 111)
	// assert.NoError(t, err)
	// id, err := FindIDinLike(321, 111)
	// assert.NoError(t, err)
	// print(id)
}

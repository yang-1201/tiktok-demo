package testpackage

import (
	"fmt"
	"go_tiktok_project/common/dal/mysql"
	"testing"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/stretchr/testify/assert"
)

func TestLikesInsert(t *testing.T) {
	defer logs.Flush()
	mysql.InitDB()
	//插入数据key_id=123 owner_id=321 video_id=111
	err := mysql.CreateLike(124, 321, 111)
	assert.NoError(t, err)
	logs.Infof("test Likes Insert")
}

func TestLikesQuery(t *testing.T) {
	defer logs.Flush()
	mysql.InitDB()
	//查询owner_id=321 video_id=111 的key_id
	id, err := mysql.FindIDinLike(321, 111)
	assert.NoError(t, err)
	assert.Nil(t, err)
	fmt.Println("查询like表主键为 ", id)
	logs.Infof("test Likes Query")

}

func TestVideoUpdate(t *testing.T) {
	defer logs.Flush()
	mysql.InitDB()
	//更新video_id=的点赞数
	err := mysql.UpdateLikeCount(111)
	assert.NoError(t, err)
	logs.Infof("test Video Update")
}

func TestVideoDelete(t *testing.T) {
	defer logs.Flush()
	mysql.InitDB()
	//删除Likes表key_id=123的数据
	err := mysql.DelLike(1234)
	//assert.NotNil(t, err)
	assert.Nil(t, err)
	assert.NoError(t, err)
	logs.Infof("test Likes Delete")
}

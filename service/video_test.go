package service

import (
<<<<<<< HEAD
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/stretchr/testify/assert"
	"testing"
=======
	"testing"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/stretchr/testify/assert"
>>>>>>> f3bcb08 (publish完成)
)

func TestGetUserVideo(t *testing.T) {

<<<<<<< HEAD
	list, err := GetUserVideo(123,123)
	assert.NoError(t, err)
	logs.Info("list: ", list)
	logs.Info("list: %T", list)
}

func TestPostUserVideo(t *testing.T) {

	err := PostUserVideo(123,"11","\\11","")
	assert.NoError(t, err)
	
}

func TestGetSnapshot(t *testing.T) {

	path,err := GetSnapshot("..//video_data//video//111_2_2549243961211321131.mp4 ","22",2)
	assert.NoError(t, err)
	logs.Info("path:",path)
	
=======
	list, err := GetUserVideo(2,4)
	assert.NoError(t, err)
	logs.Info("list: ", list)
}

func TestPostUserVideo(t *testing.T) {
	err := PostUserVideo(2, "11", "..//video_data//111//bear//bear.mp4", "")
	assert.NoError(t, err)

}

func TestGetSnapshot(t *testing.T) {
	path, err := GetSnapshot("..//video_data//111//bear//bear.mp4", "..//video_data//111//bear//33", 2)
	assert.NoError(t, err)
	logs.Info("path:", path)

>>>>>>> f3bcb08 (publish完成)
}

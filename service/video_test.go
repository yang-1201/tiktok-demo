package service

import (
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserVideo(t *testing.T) {

	list, err := GetUserVideo(123,123)
	assert.NoError(t, err)
	logs.Info("list: ", list)
	logs.Info("list: %T", list)
}

func TestPostUserVideo(t *testing.T) {

	err := PostUserVideo(123,"11","..//video_data//111//2_2549243961211321131//2_2549243961211321131.mp4","")
	assert.NoError(t, err)
	
}

func TestGetSnapshot(t *testing.T) {

	path,err := GetSnapshot("..//video_data//video//111_2_2549243961211321131.mp4 ","22",2)
	assert.NoError(t, err)
	logs.Info("path:",path)
	
}

package mysql

import (
	"fmt"
)

// FindIDinLike 失败时主键返回0和错误信息
func FindIDinLike(userID, videoID int64) (int64, error) {
	var like Like
	err := db.Where("owner_id = ? AND video_id = ?", userID, videoID).First(&like)
	if err.Error != nil {
		fmt.Println("查询like表主键出错, error: " + err.Error.Error())
		return 0, err.Error
	}
	return like.KeyID, nil
}

// FindVideoListinVideo 失败时主键返回0和错误信息
func FindVideoListinVideo(userID int64) ([]*Video, error) {
	var Video_list []*Video
	err := db.Where("auther_id = ?", userID).Find(&Video_list)

	if err.Error != nil {
		fmt.Println("查询Videl表主键出错, error: " + err.Error.Error())
		return nil, err.Error
	}

	return Video_list, nil
}

// FindUserInfoinUser 失败时主键返回0和错误信息
func FindUserInfoinUser(userID int64) (User, error) {
	var user User
	err := db.Where("user_id = ?", userID).Find(&user)

	if err.Error != nil {
		fmt.Println("查询User表主键出错, error: " + err.Error.Error())
		return user, err.Error
	}

	return user, nil
}

// FindUserInfoinUser 失败时主键返回0和错误信息
func FindCountinLikes(userID int64, videoID int64) (int64, error) {
	var total int64 = 0

	err := db.Where("owner_id = ? AND video_id = ?", userID, videoID).Model(Like{}).Count(&total)
	if err.Error != nil {
		fmt.Println("查询like表主键出错, error: " + err.Error.Error())
		return -1, err.Error
	}

	return total, nil
}

// FindUserInfoinUser 失败时主键返回0和错误信息
func FindMaxIdinVideos() (int64, error) {
	var video Video

	err := db.Order("video_id desc").First(&video)
	if err.Error != nil {
		fmt.Println("查询Video表最大id出错, error: " + err.Error.Error())
		return -1, err.Error
	}
    
	return video.VideoID, nil
}

// FindUserInfoinUser 失败时主键返回0和错误信息
func FindCountinFollows(watcher_id int64, watched_id int64) (int64, error) {
	var total int64 = 0

	err := db.Where("watcher_id = ? AND watched_id = ?", watcher_id, watched_id).Model(Follow{}).Count(&total)
	if err.Error != nil {
		fmt.Println("查询follows表主键出错, error: " + err.Error.Error())
		return -1, err.Error
	}

	return total, nil
}
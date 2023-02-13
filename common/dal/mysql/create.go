package mysql

import (
	"fmt"
	"time"
)

// CreateLike 根据雪花主键插入视频的点赞表
func CreateLike(keyID, userID, videoID int64) error {
	like := Like{KeyID: keyID, OwnerID: userID, VideoID: videoID, LikeTime: time.Now()}
	// 通过数据的指针来创建
	result := db.Create(&like)
	if result.Error != nil {
		fmt.Println("Like表格创建数据失败: " + result.Error.Error())
		return result.Error
	}
	return nil
}

// CreateLike 根据雪花主键插入视频的点赞表
func CreateVideo(videoID, autherID int64, play_url, cover_url string, likecount, commentcount int64, title, abstract string) error {
	video := Video{VideoID: videoID, AutherID: autherID, PlayUrl: play_url, CoverUrl: cover_url,
		LikeCount: likecount, CommentCount: commentcount, Title: title, Abstract: abstract}
	// 通过数据的指针来创建
	result := db.Create(&video)
	if result.Error != nil {
		fmt.Println("Video表格创建数据失败: " + result.Error.Error())
		return result.Error
	}
	return nil
}

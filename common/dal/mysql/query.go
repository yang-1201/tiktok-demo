package mysql

import (
<<<<<<< HEAD
	"fmt"
)

// FindIDinLike 失败时主键返回0和错误信息
func FindIDinLike(userID, videoID int64) (int64, error) {
=======
	"errors"
	"fmt"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/redis/go-redis/v9"
)

// FindIDinLike 失败时主键返回0和错误信息
func FindIDinLike(userID, videoID uint64) (int64, error) {
>>>>>>> f3bcb08 (publish完成)
	var like Like
	err := db.Where("owner_id = ? AND video_id = ?", userID, videoID).First(&like)
	if err.Error != nil {
		fmt.Println("查询like表主键出错, error: " + err.Error.Error())
		return 0, err.Error
	}
	return like.KeyID, nil
}

<<<<<<< HEAD
// FindVideoListinVideo 失败时主键返回0和错误信息
=======
// FindUserByName check whether username exist
// false, nil --> user isn't found
// false, err --> database error
// true, nil --> user is found
func FindUserByName(username string) (bool, error) {
	var user User
	err := db.Where("username = ?", username).First(&user)
	if err.Error == redis.Nil {
		// user isn't found
		return false, nil
	}
	// other error
	if err != nil {
		return false, err.Error
	}
	// user is found
	return true, nil
}

func FindUserByNameAndPass(username, password string) (User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user)
	if err.Error == redis.Nil {
		return user, errors.New("user doesn't exist")
	}
	if err != nil {
		logs.Errorf("mysql error during selecting: ", err.Error.Error())
		return user, err.Error
	}
	if user.Password != password {
		return user, errors.New("wrong password")
	}
	return user, nil
}

func FindUserById(userid uint64) (User, error) {
	var user User
	res := db.Where("user_id = ?", userid).First(&user)
	if res.Error == redis.Nil {
		return user, errors.New("user doesn't exist")
	}
	if res.Error != nil {
		logs.Errorf("mysql error during selecting: ", res.Error.Error())
		return user, res.Error
	}
	return user, nil
}

func FindComment(videoID int64) ([]Comment, error) {
	var comments []Comment
	// select * from comments where video_id = ? order by comment_time desc
	res := db.Where("video_id = ? ", videoID).Order("comment_time desc").Find(&comments)
	if res.Error != nil {
		fmt.Println("查询comment表主键出错, error: " + res.Error.Error())
		return nil, res.Error
	}
	var commentss []Comment
	for i := 0; i < int(res.RowsAffected); i++ {
		commentss = append(commentss, comments[i])
	}
	return commentss, nil
}

// 查询登录用户喜欢的视频列表
func FindLikeList(userID int64) (*[]int64, error) {
	var likes []Like
	res := db.Where("owner_id=?", userID).Find(&likes)
	if res.Error != nil {
		fmt.Println("无法获取用户喜爱列表,error: " + res.Error.Error())
		return nil, res.Error
	}
	var videoIDs []int64
	for i := 0; i < int(res.RowsAffected); i++ {
		videoIDs = append(videoIDs, likes[i].VideoID)
	}
	return &videoIDs, nil
}

// 查询视频点赞数
func FindLikeOfVideo(videoID int64) (int64, error) {
	var video Video
	err := db.Where("video_id=?", videoID).First(&video).Error
	if err != nil {
		fmt.Println("查询点赞出错, error: " + err.Error())
		return -1, err
	}
	return video.LikeCount, nil
}

// 根据查询userID发布的视频列表
>>>>>>> f3bcb08 (publish完成)
func FindVideoListinVideo(userID int64) ([]*Video, error) {
	var Video_list []*Video
	err := db.Where("auther_id = ?", userID).Find(&Video_list)

	if err.Error != nil {
		fmt.Println("查询Videl表主键出错, error: " + err.Error.Error())
		return nil, err.Error
	}

	return Video_list, nil
}

<<<<<<< HEAD
// FindUserInfoinUser 失败时主键返回0和错误信息
=======
// 查询userID对应的user信息
>>>>>>> f3bcb08 (publish完成)
func FindUserInfoinUser(userID int64) (User, error) {
	var user User
	err := db.Where("user_id = ?", userID).Find(&user)

	if err.Error != nil {
		fmt.Println("查询User表主键出错, error: " + err.Error.Error())
		return user, err.Error
	}

	return user, nil
}

<<<<<<< HEAD
// FindUserInfoinUser 失败时主键返回0和错误信息
=======
// 查询user给视频的点赞次数
>>>>>>> f3bcb08 (publish完成)
func FindCountinLikes(userID int64, videoID int64) (int64, error) {
	var total int64 = 0

	err := db.Where("owner_id = ? AND video_id = ?", userID, videoID).Model(Like{}).Count(&total)
	if err.Error != nil {
		fmt.Println("查询like表主键出错, error: " + err.Error.Error())
		return -1, err.Error
	}

	return total, nil
}

<<<<<<< HEAD
// FindUserInfoinUser 失败时主键返回0和错误信息
=======
// 查询视频表中最大的vedio_id
>>>>>>> f3bcb08 (publish完成)
func FindMaxIdinVideos() (int64, error) {
	var video Video

	err := db.Order("video_id desc").First(&video)
	if err.Error != nil {
		fmt.Println("查询Video表最大id出错, error: " + err.Error.Error())
		return -1, err.Error
	}
    
	return video.VideoID, nil
}

<<<<<<< HEAD
// FindUserInfoinUser 失败时主键返回0和错误信息
=======
// 查询wachter是否关注watched
>>>>>>> f3bcb08 (publish完成)
func FindCountinFollows(watcher_id int64, watched_id int64) (int64, error) {
	var total int64 = 0

	err := db.Where("watcher_id = ? AND watched_id = ?", watcher_id, watched_id).Model(Follow{}).Count(&total)
	if err.Error != nil {
		fmt.Println("查询follows表主键出错, error: " + err.Error.Error())
		return -1, err.Error
	}

	return total, nil
}
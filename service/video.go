package service

import (
	"bytes"
	"fmt"
	"go_tiktok_project/common/dal/mysql"

	//"log"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"

	//pb "go_tiktok_project/idl/pb"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	//"github.com/stretchr/testify/assert"
)

type User struct {
	Id            int64  `protobuf:"varint,1,req,name=id" json:"id"`                                            // 用户id
	Name          string `protobuf:"bytes,2,req,name=name" json:"name"`                                         // 用户名称
	FollowCount   int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount" json:"follow_count"`       // 关注总数
	FollowerCount int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount" json:"follower_count"` // 粉丝总数
	IsFollow      bool   `protobuf:"varint,5,req,name=is_follow,json=isFollow" json:"is_follow"`                // true-已关注，false-未关注
}

type Video struct {
	Id             int64  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`                                   // 视频唯一标识
	Author         User   `protobuf:"varint,2,req,name=author" json:"author"`                                     // 视频作者信息
	Play_url       string `protobuf:"varint,3,req,name=play_url" json:"play_url"`                                 // 视频播放地址
	Cover_url      string `protobuf:"varint,4,opt,name=cover_url_count,json=cover_url" json:"cover_url"`          // 视频封面地址
	Favorite_count int64  `protobuf:"varint,5,opt,name=favorite_count,json=favorite_count" json:"favorite_count"` // 视频的点赞总数
	Comment_count  int64  `protobuf:"varint,6,opt,name=comment_count,json=comment_count" json:"comment_count"`    // 视频的评论总数
	Is_favorite    bool   `protobuf:"varint,7,opt,name=is_favorite ,json=is_favorite" json:"is_favorite "`        // true-已点赞，false-未点赞
	Title          string `protobuf:"varint,8,opt,name=title,json=title" json:"title"`                            // 视频标题
	Abstract       string `protobuf:"varint,8,opt,name=absruct,json=abstract" json:"abstract"`                    // 视频简介
}

type douyin_publish_list_response struct {
	Status_code int32   `protobuf:"varint,1,req,name=status_code" json:"status_code"`               // 状态码，0-成功，其他值-失败
	Status_msg  string  `protobuf:"varint,2,opt,name=status_msg,json=status_msg" json:"status_msg"` // 返回状态描述
	Video_list  []Video `protobuf:"varint,3,req,name=video_list" json:"video_list"`                 // 用户发布的视频列表
}

func GetUserVideo(video_user_id, token_user_id int64) (*douyin_publish_list_response, error) {
	mysql.InitDB()
	userInfo, err := mysql.FindUserInfoinUser(video_user_id)
	if err != nil {
		logs.Errorf("查询User表出错, error: " + err.Error())
		return nil, err
	}

	var isFollow bool = false
	FollowCount, err := mysql.FindCountinFollows(token_user_id, video_user_id)
	if err != nil {
		logs.Errorf("查询Follow表出错, error: " + err.Error())
		//return nil, err
	}

	if FollowCount > 0 {
		isFollow = true
	}

	user := &User{
		Id:            userInfo.UserID,
		Name:          userInfo.Username,
		FollowCount:   userInfo.Follow_cnt,
		FollowerCount: userInfo.Follower_cnt,
		IsFollow:      isFollow,
	}

	fmt.Println("user_info:", user)
	logs.Info("user_info %T: ", user)
	fmt.Println("user_info:", user.Id)
	// //2. 创建LoginMsg结构体对象并序列化
	// user_info, err := json.Marshal(&user)
	// if err != nil {
	// 	fmt.Println("Client login data marshal failed, err = ", err)
	// 	return nil, err
	// }
	// // fmt.Println("user_info:",user_info)
	// fmt.Println("user_info:", string(user_info))
	videos, err := mysql.FindVideoListinVideo(video_user_id)
	if err != nil {
		logs.Errorf("查询Video表出错, error: " + err.Error())
		return nil, err
	}

	logs.Info("videos:", videos)

	var video_list []Video

	for _, v := range videos {
		var isLike bool = false
		LikesCount, err := mysql.FindCountinLikes(video_user_id, v.VideoID)
		if err != nil {
			logs.Errorf("查询Like表Count出错, error: " + err.Error())
			return nil, err
		}

		if LikesCount > 0 {
			isLike = true
		}

		video := &Video{
			Id:             v.VideoID,      // 视频唯一标识
			Author:         *user,          // 视频作者信息
			Play_url:       v.PlayUrl,      // 视频播放地址
			Cover_url:      v.CoverUrl,     // 视频封面地址
			Favorite_count: v.LikeCount,    // 视频的点赞总数
			Comment_count:  v.CommentCount, // 视频的评论总数
			Is_favorite:    isLike,         // true-已点赞，false-未点赞
			Title:          v.Title,        // 视频标题
			Abstract:       v.Abstract,     //视频简介
		}
		video_list = append(video_list, *video)
		fmt.Println("video_info:", video)
		logs.Info("video_info %T: ", video)

	}
	logs.Info("video_list: ", video_list)
	//var resp = new(douyin_publish_list_response)
	var resp = new(douyin_publish_list_response)
	logs.Info("video_list:%T ", video_list)
	//logs.Info("VideoList:%T ", resp)
	resp.Status_code = 0
	resp.Status_msg = "获取用户发布列表成功"
	// for i, v := range video_list {
	// 	resp.Video_list[i] = *v
	// }

	resp.Video_list = video_list
	return resp, nil
}

func PostUserVideo(user_id int64, title string, filepath string, filedata string) error {
	mysql.InitDB()
	max_video_id, err := mysql.FindMaxIdinVideos()
	if err != nil {
		logs.Errorf("查询Video表主键出错, error: " + err.Error())
		return err
	}
	var video_id int64 = max_video_id + 1

	logs.Info("video_id: ", video_id)

	strArray := strings.Split(filepath, ".")

	ImageName := strArray[0]
	//logs.Info("imagePath: ", imagePath)
	imagePath, err := GetSnapshot(filepath, ImageName, 1)
	if err != nil {
		logs.Errorf("截取视频第一帧错误, error: " + err.Error())
		return err
	}
	logs.Info("imagePath: ", imagePath)

	err_createvideo := mysql.CreateVideo(video_id, user_id, filepath, imagePath, 0, 0, title, filedata)
	if err_createvideo != nil {
		fmt.Println("创建Video数据出错, error: " + err.Error())
		return err
	}
	return nil
}

func GetSnapshot(videoPath, imageName string, frameNum int) (ImagePath string, err error) {
	//var snapshotPath string = "..//video_data//image//" + imageName
	//var snapshotPath string = imageName
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		logs.Errorf("截图失败, error: ", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		logs.Info("生成缩略图失败：", err)
		return "", err
	}

	err = imaging.Save(img, imageName+".png")
	if err != nil {
		logs.Info("生成缩略图失败：", err)
		return "", err
	}

	imgPath := imageName + ".png"

	return imgPath, nil
}

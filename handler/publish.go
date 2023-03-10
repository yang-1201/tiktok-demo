// Code generated by hertz generator.

package handler

import (
	"context"
<<<<<<< HEAD

	//"go_tiktok_project/common/dal/mysql"
=======
>>>>>>> f3bcb08 (publish完成)
	"fmt"
	"os"
	"strconv"
	"strings"

<<<<<<< HEAD
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	//pb "go_tiktok_project/idl/pb"
	"go_tiktok_project/service"

=======
	"go_tiktok_project/common"
	"go_tiktok_project/service"
>>>>>>> f3bcb08 (publish完成)
	"path/filepath"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
<<<<<<< HEAD
)

// // Ping .
func GetUserVideo(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: %s", path)
	// logs.Info("ctx: %s", ctx)
	// logs.Info("c: %s", c)

	// req := new(pb.DouyinUserRequest)
	// if err := c.BindAndValidate(&req); err != nil {
	// 	logs.Info("错误")
	// 	c.String(400, err.Error())
	// 	return
	// }
	// logs.Info("req: %s", req)

	user := c.Query("user_id")
	logs.Info("user_id: %s", user)
	token := c.Query("token")
	logs.Info("token: %s", token)

	video_user_id, err := strconv.ParseInt(user, 10, 64)
	if err != nil {
		logs.Errorf("转换user_id为int失败, error: " + err.Error())
		c.JSON(400,  utils.H{
			"status_code": 1,
			"status_msg":  "获得用户发布视频失败",
			"video_list":  nil,
		})
		return
	}
	logs.Info("video_user_id: %T", video_user_id)

	token_user_id, err := strconv.ParseInt(token, 10, 64)
	if err != nil {
		logs.Errorf("转换token为int失败, error: " + err.Error())
		c.JSON(400,  utils.H{
			"status_code": 1,
			"status_msg":  "获得用户发布视频失败",
			"video_list":  nil,
		})
		return
	}
	
	// // //鉴权token
	// mysql.InitDB()
	// video_list, err := mysql.FindIDinVideo(b)
	video_list, err := service.GetUserVideo(video_user_id,token_user_id)
	if err != nil {
		logs.Errorf("service err, error: " + err.Error())
		c.JSON(400,  utils.H{
				"status_code": 1,
				"status_msg":  "获得用户发布视频失败",
				"video_list":  nil,
			})
		return 
	}
	// video_list1,_:=json.Marshal(&video_list)
	// logs.Info("resp:", string(video_list1))

	//user_id := c.user_id

	//logs.Info("user_id:%s", user_id)
	// req := new(pb.DouyinUserRequest)
	// if err := c.BindAndValidate(&req); err != nil {
	// 	c.String(400, err.Error())
	// 	return
	// }
	// c.JSON(consts.StatusOK, utils.H{
	// 	"status_code": 0,
	// 	"status_msg":  "获得用户发布列表",
	// 	"video_list":  string(video_list1),
	// })
	c.JSON(consts.StatusOK, video_list)
}

=======
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// 生成pb会显示有冲突，就在这里构造了结构体
type douyin_publish_list_response struct {
	Status_code int32           `protobuf:"varint,1,req,name=status_code" json:"status_code"`              // 状态码，0-成功，其他值-失败
	Status_msg  string          `protobuf:"bytes,2,opt,name=status_msg,json=status_msg" json:"status_msg"` // 返回状态描述
	Video_list  []service.Video `protobuf:"varint,3,req,name=video_list" json:"video_list"`                // 用户发布的视频列表
}

// 获得发布列表
func GetUserVideo(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: %s", path)

	//鉴权token
	token := c.Query("token")
	token_id, _ := strconv.ParseInt(token, 10, 64)
	token1, _ := service.GenerateToken(uint64(token_id), "222")

	token_userID, err := common.Token2UserID(token1)
	token_user_id := int64(token_userID)
	if err != nil {
		logs.Errorf("鉴权token错误, error: " + err.Error())
		c.JSON(400, douyin_publish_list_response{
			Status_code: common.TokenFailed,
			Status_msg:  common.TokenFailedMsg,
			Video_list:  nil,
		})
		return
	}

	user := c.Query("user_id")
	video_user_id, err := strconv.ParseInt(user, 10, 64)
	if err != nil {
		logs.Errorf("转换user_id为int失败, error: " + err.Error())
		c.JSON(400, douyin_publish_list_response{
			Status_code: common.GetUserVideoFailed,
			Status_msg:  common.GetUserVideoFailedMsg,
			Video_list:  nil,
		})
		return
	}

	logs.Info("token_user_id: %s", token_user_id)

	//获取用户发布列表
	video_list, err := service.GetUserVideo(video_user_id, token_user_id)
	if err != nil {
		logs.Errorf("service err, error: " + err.Error())
		c.JSON(400, douyin_publish_list_response{
			Status_code: common.GetUserVideoFailed,
			Status_msg:  common.GetUserVideoFailedMsg,
			Video_list:  nil,
		})
		return
	}

	//返回respose
	resp := new(douyin_publish_list_response)
	resp.Status_code = common.GetUserVideoSuccess
	resp.Status_msg = common.GetUserVideoSuccessMsg
	resp.Video_list = video_list
	c.JSON(consts.StatusOK, resp)
}

// 用户视频投稿
>>>>>>> f3bcb08 (publish完成)
func PostUserVideo(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: %s", path)

<<<<<<< HEAD
	// data1, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Printf("ctx.Request.body: %v", string(data1))

	token, _ := c.GetPostForm("token") //认为是user_id
	// if err != nil {
	// 	logs.Error("获得token失败, error: " + err)
	// 	c.JSON(400, utils.H{
	// 		"status_code": 1,
	// 		"status_msg":  "发布视频失败",
	// 	})
	// 	return
	// }
	
	user_id, err := strconv.ParseInt(token, 10, 64)
	if err != nil {
		//logs.Errorf("转化user_id为int失败, error: "+ err)
		c.JSON(400, utils.H{
			"status_code": 1,
			"status_msg":  "发布视频失败",
		})
		return
	}
	logs.Info("user_id: %T", user_id)

	title, _ := c.GetPostForm("title")
	logs.Info("title: %T", title)
	// if err != nil {
	// 	//logs.Errorf("获得title失败, error: " + err)
	// 	c.JSON(400, utils.H{
	// 		"status_code": 1,
	// 		"status_msg":  "发布视频失败",
	// 	})
	// 	return
	// }
	// data := c.Query("data")
	// logs.Info("data: %s", data)
	// logs.Info("data: %T", data)

	data, err := c.FormFile("data")
	if err != nil {
		panic(err)
		c.JSON(400, utils.H{
			"status_code": 1,
			"status_msg":  "发布视频失败",
		})
		return
	}
	
	logs.Info("file: %s", data.Filename)
	filename := filepath.Base(data.Filename)
	fileInfo := strings.Split(filename, ".")
	filedir := fileInfo[0]
	logs.Info("file: %s", filedir)
	filedata := fmt.Sprintf("uesr:%d  video:%s", user_id, filename)
	filedir = fmt.Sprintf("./video_data/%s/%s", token, filedir)
	logs.Info("file: %s", filedir)

	_, erByStat := os.Stat(filedir)
	if erByStat != nil {
		logs.Errorf("os stat %s error......%s", filedir, erByStat.Error())
		c.JSON(400, utils.H{
			"status_code": 1,
			"status_msg":  "发布视频失败",
		})
		return
=======
	//resp := new(pb_video.DouyinPublishActionResponse)

	//鉴权token
	token, _ := c.GetPostForm("token") //认为是user_id
	token_id, _ := strconv.ParseInt(token, 10, 64)
	token1, _ := service.GenerateToken(uint64(token_id), "222")

	token_userID, err := common.Token2UserID(token1)
	user_id := int64(token_userID)
	if err != nil {
		logs.Errorf("鉴权token错误, error: " + err.Error())
		c.JSON(400, utils.H{
			"status_code": common.TokenFailed,
			"status_msg":  common.TokenFailedMsg,
		})
		return
	}
	logs.Info("user_id: %s", user_id)

	title, _ := c.GetPostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		panic(err)
	}
    
	//利用文件名获得文件信息和文件路径
	filename := filepath.Base(data.Filename)
	fileInfo := strings.Split(filename, ".")
	filedir := fileInfo[0]
	filedata := fmt.Sprintf("uesr:%d  video:%s", user_id, filename)
	filedir = fmt.Sprintf("./video_data/%s/%s", token, filedir)
	logs.Info("file: %s", data.Filename)
	logs.Info("filedir: %s", filedir)

	//创建存储文件夹
	_, erByStat := os.Stat(filedir)
	if erByStat != nil {
		logs.Errorf("os stat %s error......%s", filedir, erByStat.Error())
>>>>>>> f3bcb08 (publish完成)
	}
	//该判断主要是部分文件权限问题导致os.Stat()出错,具体看业务启用
	//使用os.IsNotExist()判断为true,说明文件夹不存在
	if os.IsNotExist(erByStat) {
		logs.Info("%s is not exist", erByStat.Error())
		err := os.MkdirAll(filedir, 0777)
		if err != nil {
<<<<<<< HEAD
			logs.Error("创建文件夹错误 , error:"+ err.Error())
			c.JSON(400, utils.H{
				"status_code": 1,
				"status_msg":  "发布视频失败",
=======
			logs.Error("创建文件夹错误 , error:" + err.Error())
			c.JSON(400, utils.H{
				"status_code": common.PublishFailed,
				"status_msg":  common.PublishFailedMsg,
>>>>>>> f3bcb08 (publish完成)
			})
			return
		} else {
			logs.Info("Create dir %s success!", filedir)
		}
	}
<<<<<<< HEAD
	//finalName := fmt.Sprintf("%s.%s", filedir, filetype)
	saveFile := filepath.Join(filedir, filename)

	logs.Info("filepath: %s", saveFile)
	logs.Info("filepath: %T", saveFile)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		logs.Error("保存视频失败,error:"+ err.Error())
		c.JSON(400, utils.H{
			"status_code": 1,
			"status_msg":  "发布视频失败",
=======

	//保存视频文件
	saveFile := filepath.Join(filedir, filename)
	logs.Info("filepath: %s", saveFile)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		logs.Error("保存视频失败,error:" + err.Error())
		c.JSON(400, utils.H{
			"status_code": common.PublishFailed,
			"status_msg":  common.PublishFailedMsg,
>>>>>>> f3bcb08 (publish完成)
		})
		return
	}

<<<<<<< HEAD
=======
	//service 保存视频数据到数据库
>>>>>>> f3bcb08 (publish完成)
	err_service := service.PostUserVideo(user_id, title, saveFile, filedata)
	if err_service != nil {
		logs.Error("server error,error :", err_service.Error())
		c.JSON(400, utils.H{
<<<<<<< HEAD
			"status_code": 1,
			"status_msg":  "发布视频失败",
=======
			"status_code": common.PublishFailed,
			"status_msg":  common.PublishFailedMsg,
>>>>>>> f3bcb08 (publish完成)
		})
		return
	}

<<<<<<< HEAD
	c.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "发布视频成功",
	})
}

// func PersonBind(ctx context.Context, c *app.RequestContext) {
// 	type user struct {
// 		user_id int `query:"user_id" json:"user_id"` // 从路径中获取参数
// 		token   int `query:"token" json:"token"` // 从query中获取参数
// 	}
// 	var p user
// 	if err := c.BindAndValidate(&p); err != nil {
// 		panic(err)
// 	}
// 	c.JSON(200, utils.H{
// 		"user": p,
// 	})
// }
=======
	//返回参数
	c.JSON(consts.StatusOK, utils.H{
		"status_code": common.PublishSuccess,
		"status_msg":  common.PublishSuccessMsg,
	})
}
>>>>>>> f3bcb08 (publish完成)

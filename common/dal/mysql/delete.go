package mysql

import "fmt"

func DelLike(keyID int64) error {
	like := Like{KeyID: keyID}
	// 修改结构体后是硬删除，修改前包含gorm.Model后生成了gorm.deletedat,自动变为软删除
<<<<<<< HEAD
	result := db.Delete(&like) 
=======
	result := db.Delete(&like)
>>>>>>> f3bcb08 (publish完成)
	if result.Error != nil {
		fmt.Println("删除like表数据失败,error: " + result.Error.Error())
		return result.Error
	}
	return nil
}
<<<<<<< HEAD
=======

func DelComment(keyID int64) error {
	comment := Comment{CommentID: keyID}
	result := db.Delete(&comment)
	if result.Error != nil {
		fmt.Println("删除comment表数据失败,error: " + result.Error.Error())
		return result.Error
	}
	return nil
}
>>>>>>> f3bcb08 (publish完成)

package mysql

import (
	"credit_gin/model"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type UserDetail struct {
	Id 			int64 		`gorm:"column:id"`
	UserId		int64 		`gorm:"column:user_id"`
	Gender		int64		`gorm:"column:gender"`     //0-男 1-女  2223333333333333333333333333333
	Birthday	string		`gorm:"column:birthday"`
	UserPic		string		`gorm:"column:user_pic"`
	UserDesc	string		`gorm:"column:user_desc"`
	CreateTime	int64		`gorm:"column:create_time"`
	UpdateTime	int64		`gorm:"column:update_time"`
}

func (*UserDetail) TableName() string {
	return "user_detail"
}

//AddUserDetail: 新增用户消息
func AddUserDetail(req model.UserDetailReq, db *gorm.DB) error{
	userDetail := UserDetail{
		UserId:     req.UserId,
		Gender:     req.Gender,
		Birthday:   req.Birthday,
		UserPic:    req.UserPic,
		UserDesc:   req.UserDesc,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	if err := db.Create(&userDetail).Error; err != nil {
		log.Println("[db]: add user detail err: ", err, ". user detail: ", userDetail)
		return err
	}
	log.Println("user detail ", userDetail)
	log.Println("[db]: add user detail success")
	return nil
}

//UpdateUserDetail: 更新用户详细信息
func UpdateUserDetail(req model.UserDetailReq, db *gorm.DB) error {
	updateModel := UserDetail{
		UserId:     req.UserId,
		Gender:     req.Gender,
		Birthday:   req.Birthday,
		UserPic:    req.UserPic,
		UserDesc:   req.UserDesc,
		UpdateTime: time.Now().Unix(),
	}
	var detail UserDetail
	if err := db.Where("user_id = ?", req.UserId).First(&detail).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			updateModel.CreateTime = time.Now().Unix()
			db.Create(&updateModel)
			log.Println("[db] create user detail success")
		}else {
			log.Println("[db] insert detail err: ", err)
			return err
		}
	}else {
		db.Model(&UserDetail{}).Where("user_id = ?", req.UserId).Update(&updateModel)
		log.Println("[db] update detail success")
	}
	return nil
}

//DeleteUserDetail: 删除用户详细信息
func DeleteUserDetail(userId int64, db *gorm.DB) error {
	userDetail := UserDetail{
		UserId: userId,
	}
	if err := db.Delete(&userDetail).Error; err != nil {
		log.Println("delete user detail err:", err,". user is: ", userId)
		return err
	}
	log.Println("delete user detail success, user id is: ", userId)
	return nil
}

//SelectUserDetailById：根据id查询用户详细信息
func SelectUserDetailById(userId int64, db *gorm.DB) (*UserDetail, error) {
	user := &UserDetail{}
	if err := db.Model(&UserDetail{}).Where("user_id = ?", userId).Find(user).Error; err != nil {
		log.Println("[db]: select user detail err: ", err)
		return nil, err
	}
	log.Println("[db]: select user detail success, user is ", user)
	return user, nil
}

//SelectUserDetailByPage：分页查询用户详细信息 limit=-1 && offset = 0表示查询全部
func SelectUserDetailByPage(limit int, offset int, db *gorm.DB) ([]*UserDetail, error) {
	users := make([]*UserDetail, 0 ,0)
	if err := db.Model(&UserDetail{}).Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		log.Println("select users detail err: ", err)
		return nil, err
	}
	log.Println("select user detail success, user num is ", len(users))
	return users, nil
}
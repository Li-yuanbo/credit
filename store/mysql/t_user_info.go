package mysql

import (
	"credit_gin/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type UserInfo struct {
	Id			int64			`gorm:"column:id"`
	Name		string			`gorm:"column:name"`
	Password	string			`gorm:"column:password"`
	Phone		string			`gorm:"column:phone"`
	IdCard		string			`gorm:"column:id_card"`
	Level		int64			`gorm:"column:level"`
	Gender		int64			`gorm:"column：gender"`
	Birthday	string			`gorm:"column:birthday"`
	Desc		string			`gorm:"column:desc"`
	CreateTime	int64			`gorm:"column:create_time"`
	UpdateTime	int64			`gorm:"column:update_time"`
}

func (*UserInfo) TableName() string {
	return "user_info"
}

//InsertUserInfo: 新增用户
func InsertUserInfo(req model.RegisterUserReq, db *gorm.DB) (int64, error){
	UserInfoModel := &UserInfo{
		Name:       req.UserName,
		Password:   req.Password,
		Phone:      req.Phone,
		IdCard:		req.IdCard,
		Gender:		req.Gender,
		Birthday:	req.Birthday,
		Desc:		req.Desc,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	err := db.Create(UserInfoModel).Error
	if err != nil {
		log.Println("InsertUser: insert user err: ", err)
		return -1, err
	}
	log.Println("Insert User Success, user:", UserInfoModel)
	return UserInfoModel.Id, nil
}

//UpdateUserInfo: 更新用户信息
func UpdateUserInfo(userReq model.UserInfoModel, db *gorm.DB) error {
	UpdateUserModel := map[string]interface{}{
		"name":        userReq.UserName,
		"password":    userReq.Password,
		"phone":       userReq.Phone,
		"update_time": time.Now().Unix(),
	}
	err := db.Model(&UserInfo{}).Where("id = ?", userReq.Id).Update(UpdateUserModel).Error
	if err != nil {
		log.Println("UpdateUser: update user info err: ", err)
		return err
	}
	log.Println("Update User Info Success, user: ", UpdateUserModel)
	return nil
}

//DeleteUserInfoById: 根据User Id删除用户
func DeleteUserInfoById(req model.DeleteUserReq, db *gorm.DB) error {
	user := UserInfo{
		Id: req.Id,
	}
	err := db.Delete(&user).Error
	if err != nil {
		log.Println("DeleteUserInfoById: delete user info err: ", err)
		return err
	}
	log.Println("Delete User Info Success, user id: ", req.Id)
	return nil
}

//SelectUserInfoById: 根据user id查询指定用户
func SelectUserInfoById(userId int64, db *gorm.DB) (UserInfo, error) {
	user := UserInfo{}
	err := db.Where("id = ?", userId).Find(&user).Error
	if err != nil {
		log.Println("[db]: SelectUserInfoById: select user info by id err: ", err)
		return user, err
	}
	log.Println("[db]: Select User By Id Success, user: ", user)
	return user, nil
}

//SelectUsers: 分页查询用户. offset=0 && limit=-1查询全部
func SelectUserInfos(limit int64, offset int64, db *gorm.DB)([]*UserInfo, error) {
	users := make([]*UserInfo, 0, 0)
	err := db.Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		log.Println("[db] Select Users: select users err: ", err)
		return users, err
	}
	log.Println("[db] Select Users Success, user num : ", len(users))
	return users, nil
}

//根据username查询用户
func SelectUserByUserName(username string, db *gorm.DB) (*UserInfo, error) {
	user := &UserInfo{}
	err := db.Where("name = ?", username).Find(&user).Error
	if err != nil {
		log.Println("[db]: SelectUserInfoById: select user info by id err: ", err)
		return user, err
	}
	log.Println("[db]: Select User By Id Success, user: ", user)
	return user, nil
}

//获取用户总数
func GetUserTotalCount(db *gorm.DB) (int64, error) {
	var total int64
	if err := db.Model(&UserInfo{}).Count(&total).Error; err != nil {
		log.Println("[db]: get user total count err: ", err)
		return -1, err
	}
	log.Println("[db]: get user total count success, count: ", total)
	return total, nil
}

//根据用户id查询用户信用卡信息
func SelectUserCardInfo(userid int64, db *gorm.DB) ([]*model.UserCard, error) {
	cards := make([]*model.UserCard, 0, 0)
	if err := db.Table("user_card").Select("bank.bank_name, user_card.bank_id,user_card.credit_card_num, user_card.balance").
		Joins("left join bank on user_card.bank_id = bank.id").Where("user_id = ?", userid).Scan(&cards).
		Error; err != nil {
			log.Println("[db]: select user card err: ", err)
			return nil, err
	}
	res := make([]*model.UserCard, 0, 0)
	for _, card := range cards {
		c := &model.UserCard{
			BankId:        card.BankId,
			BankName:      Uint8ToStr(card.BankName.([]uint8)),
			CreditCardNum: Uint8ToStr(card.CreditCardNum.([]uint8)),
			Balance:       card.Balance,
		}
		res = append(res, c)
		fmt.Println(c)
	}
	return res, nil
}

func Uint8ToStr(data []uint8) string {
	ba := []byte{}
	for _, b := range data {
		ba = append(ba, b)
	}
	return string(ba)
}

//获取用户收入/支出详情
func GetUserOperate(req model.GetUserOperateReq, db *gorm.DB) ([]*CreditCardFlow, error){
	ops := make([]*CreditCardFlow, 0, 0)
	if err := db.Where("user_id = ? and operate_type = ?", req.UserId, req.OperateType).Find(&ops).Error; err != nil {
		log.Println("[db]: select user ops err: ", err)
		return nil, err
	}
	log.Println("[db]: select user ops success")
	return ops, nil
}
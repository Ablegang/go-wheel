package show

import "time"

type Admin struct {
	Id          int64
	Name        string `xorm:"index"`
	Email       string `xorm:"unique"`
	Phone       string `xorm:"unique"`
	Salt        string
	Pwd         string `xorm:"varchar(200)"`
	LastLoginAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (admin *Admin) TableName() string {
	return "admins"
}
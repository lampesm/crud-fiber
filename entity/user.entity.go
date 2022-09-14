package entity

type User struct {
	ID       int64  `gorm:"primary_key:auto_increment" json:"-"`
	Username string `gorm:"unique;not null;type:varchar(100)" json:"-"`
	Password string `gorm:"not null;type:varchar(100)" json:"-"`
	Email    string `gorm:"unique;not null;type:varchar(100)" json:"-"`
}

package user

import (
	"time"

	"database/sql/driver"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	Admin    UserRole = "ADMIN"
	Courier  UserRole = "COURIER"
	Customer UserRole = "CUSTOMER"
)

func (as *UserRole) Scan(value interface{}) error {
	*as = UserRole(value.(string))
	return nil
}

func (as UserRole) Value() (driver.Value, error) {
	return string(as), nil
}

func (as *UserRole) String() string {
	return string(*as)
}

type UsersInter interface {
	Create(aPE *UserEntity) error
	Delete(id int64) error
	Update(id int64, title string, status string) (UserEntity, error)
	Get(id int64) (UserEntity, error)
	FindByLogin(login string) (UserEntity, error)
}

type UserEntity struct {
	ID       uuid.UUID `gorm:"column:uuid;type:uuid;primaryKey"`
	Login    string    `gorm:"column:login;uniqueIndex"`
	Password string    `gorm:"column:password"`
	Salt     string    `gorm:"column:salt"`
	Created  int64     `gorm:"column:created"`
	Role     UserRole  `gorm:"column:status;type:enum_user_role;default:'active'"`
}

func (UserEntity) TableName() string {
	return "tusers"
}

func (a *UserEntity) BeforeCreate(_ *gorm.DB) (err error) {
	a.ID = uuid.New()
	a.Created = time.Now().Unix()
	return
}

type users struct {
	db *gorm.DB
}

func New(dbr *gorm.DB) UsersInter {

	return &users{db: dbr}
}

func (aP users) Create(aPE *UserEntity) error {
	if err := aP.db.Create(aPE).Error; err != nil {
		return err
	}
	return nil
}

func (aP users) Get(id int64) (UserEntity, error) {
	var acPlan UserEntity

	if err := aP.db.Find(&acPlan, id).Error; err != nil {
		return UserEntity{}, err
	}
	return acPlan, nil
}

func (aP users) FindByLogin(login string) (UserEntity, error) {
	var acPlan UserEntity

	if err := aP.db.Where("login", login).Find(&acPlan).Error; err != nil {
		return UserEntity{}, err
	}
	return acPlan, nil
}

func (aP users) Delete(id int64) error {
	if err := aP.db.Delete(UserEntity{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (aP users) Update(id int64, title string, status string) (UserEntity, error) {

	return UserEntity{}, nil
}

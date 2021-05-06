package repository

import (
	"log"

	"github.com/roier/POS/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user models.User) models.User
	UpdateUser(user models.User) models.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicatedEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) models.User
	ProfileUser(UserID string) models.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user models.User) models.User {
	user.Password = encryptPassword([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user models.User) models.User {
  if user.Password != "" {
    user.Password = encryptPassword([]byte(user.Password))
  } else {
    var tempUser models.User
    db.connection.Find(&tempUser, user.ID)
    user.Password = tempUser.Password
  }
	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user models.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	log.Println("Error verifying credentials")
	return nil
}

func (db *userConnection) IsDuplicatedEmail(email string) (tx *gorm.DB) {
	var user models.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) models.User {
	var user models.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) models.User {
	var user models.User
	db.connection.Find(&user, userID)
	return user
}

func encryptPassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash the password")
	}
	return string(hash)
}

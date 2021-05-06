package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/roier/POS/dto"
	"github.com/roier/POS/models"
	"github.com/roier/POS/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface{
  VerifyCredential(email string, password string) interface{}
  CreateUser(user dto.Register) models.User
  FindByEmail(email string) models.User
  IsDuplicatedEmail(email string) bool
}

type authService struct {
  userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
  return &authService{
    userRepository: userRep,
  }
}

func (s *authService) VerifyCredential(email string, password string) interface{} {
  res := s.userRepository.VerifyCredential(email, password)
  if v, ok := res.(models.User); ok {
    comparedPassword := comparePassword(v.Password, []byte(password))
    if v.Email == email && comparedPassword {
      return res
    }
    return false
  }
  return false
}

func (s *authService) CreateUser(user dto.Register) models.User {
  newUser := models.User{}
  err := smapping.FillStruct(&newUser, smapping.MapFields(&user))
  if err != nil {
    log.Fatalf("Failed to map %v", err)
  }
  res := s.userRepository.InsertUser(newUser)
  return res
}

func (s *authService) FindByEmail(email string) models.User {
  return s.userRepository.FindByEmail(email)
}

func (s *authService) IsDuplicatedEmail(email string) bool {
  res := s.userRepository.IsDuplicatedEmail(email)
  return !(res.Error == nil)
}

func comparePassword(encryptedPassword string, plainPassword []byte) bool {
  byteHash := []byte(encryptedPassword)
  err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
  if err != nil {
    log.Println(err)
    return false
  }
  return true
}

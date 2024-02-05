package domain

import (
   "errors"

   "github.com/google/uuid"

   "net/mail"
   "unicode"
)

type User struct{
   UserId   string   `json:"userId"`
   Name     string   `json:"name"`
   Email    string   `json:"email"`
   Password string   `json:"password"`
}

var Users = []User{
   {UserId: "1", Name: "Patrick Carpio", Email: "patricik18483@gmail.com", Password: "BatmanIsBruceWayne"},
   {UserId: "1", Name: "Bruno Carpio", Email: "bruno@gmail.com", Password: "BatmanIsBruceWayne"},
}

type UserRepository interface {
	Fetch(*User, error)
   GetById(id string) (*User, error)
   Delete(id string)
   Update(User)
   Create(User)
}

func IsValid(user *User) (bool, error) {
   var errs []error

   if err := isValidUserId(user.ID); err != nil {
      errs = append(errs, err)
   }

   if err := isValidEmail(user.Email); err != nil {
      errs = append(errs, err)
   }

   if err := isValidPassword(user.Password); err != nil {
      errs = append(errs, err)
   }

   if len(errs) > 0 {
      return errors.New("Invalid user: " + errs[0].Error()) // Return a single error with context
   }

    return nil
}

func isValidUserId(id uuid.UUID) error {
   if _, err := uuid.Parse(id.String()); err != nil {
       return errors.New("Invalid user ID format")
   }
   return nil
}

func isValidEmail(email string) error {
   _, err := mail.ParseAddress(email)
   return err
}

func isValidPassword(password string) error {
   hasNumber := false
   hasUpper := false
   hasSpecial := false
   minLength := 7

   for _, c := range password {
       switch {
       case unicode.IsNumber(c):
           hasNumber = true
       case unicode.IsUpper(c):
           hasUpper = true
       case unicode.IsPunct(c) || unicode.IsSymbol(c):
           hasSpecial = true
       }
   }

   if len(password) < minLength || !hasNumber || !hasUpper || !hasSpecial {
       return errors.New("Password must be at least 7 characters long and contain a number, uppercase letter, and special character")
   }

   return nil
}
package go_micro_srv_user

import (
  "github.com/jinzhu/gorm"
  "log"
  "github.com/satori/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
  uuid, error := uuid.NewV4()

  if error != nil {
    log.Fatal(error)
  }

  return scope.SetColumn("Id", uuid.String())
}
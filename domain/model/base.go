package model

import (
	"github.com/asaskevich/govalidator"
	"time"
)

func init(){
	govalidator.SetFieldsRequiredByDefault( value: true)
}
type Base struct {
	id 			string `json:"id" valid:"uuid"`
	CreateAt 	time.Time `json:"create_at" valid:"-"`
	UpdatedAt 	time.Time `json:"updated_at" valid:"-"`
}
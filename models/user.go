package libs

import "gorm.io/gorm"


//masukan struc di model dan masukan nama biar rapi
type User struct {
	gorm.Model
	Name     string
	Password string
}

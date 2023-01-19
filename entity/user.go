package entity


type APIUser struct {
	ID                int         `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id" model:"id"`
	Name              string      `gorm:"column:name;type:string;" json:"name" model:"name"`
	Email             string      `gorm:"column:email;type:string;" json:"email" model:"email"`
}

package entity

type User struct {
	UID int64 `json:"uid"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	EMail string `json:"e_mail"`
	College string `json:"college"`
	Password string `json:"password"`
}
// 4.1
package user

type UserOrderKey int

const (
	Name UserOrderKey = iota + 1
	UserID
)

type UserRepository interface {
	FindByStatus(status UserStatus, userOrderKey UserOrderKey) []User
}

type UserStatus int

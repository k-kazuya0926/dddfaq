// 4.2
package repository

import "dddfaq/domain/user"

type UserRepository struct {
}

func (ur *UserRepository) findByStatus(status user.UserStatus, userOrderKey user.UserOrderKey) []user.User {
	return nil
}

// 3.1
package user

type User struct {
	userID        string
	name          string
	mailAddresses []MailAddress
}

type MailAddress struct {
	value string
}

// 8.6
package usersession

type UserSession interface {
}

type SomeRequireSessionUseCase struct {
}

func (u *SomeRequireSessionUseCase) Execute(session UserSession) {
	// 何らかの処理
}

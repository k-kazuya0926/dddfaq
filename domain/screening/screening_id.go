// 2.4
package screening

import "github.com/google/uuid"

type ScreeningID struct {
	value string
}

func NewScreeningID() ScreeningID {
	return ScreeningID{value: uuid.New().String()}
}

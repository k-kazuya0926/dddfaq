// 2.5
package position

import "github.com/google/uuid"

type PositionID struct {
	value string
}

func NewPositionID() PositionID {
	return PositionID{value: uuid.New().String()}
}

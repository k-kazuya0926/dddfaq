// 2.3
package screening

type ScreeningStatus struct {
	canAddInterview bool
}

var (
	InProgress ScreeningStatus = ScreeningStatus{true}
	Adopted    ScreeningStatus = ScreeningStatus{false}
	Rejected   ScreeningStatus = ScreeningStatus{false}
)

// 2.2
package screening

import (
	"dddfaq/domain/position"
	"errors"
	"time"
)

type Screening struct {
	screeningID   ScreeningID
	positionID    position.PositionID
	applicant     Applicant
	interviews    Interviews
	status        ScreeningStatus
	applyDateTime time.Time
}

func CreateScreening(
	positionID position.PositionID,
	applicant Applicant,
) *Screening {
	return &Screening{
		screeningID:   NewScreeningID(),
		positionID:    positionID,
		applicant:     applicant,
		interviews:    emptyInterviews(),
		status:        InProgress,
		applyDateTime: time.Now(),
	}
}

func ReconstructScreening(
	screeningID ScreeningID,
	positionID position.PositionID,
	applicant Applicant,
	interviews Interviews,
	status ScreeningStatus,
	applyDateTime time.Time,
) *Screening {
	return &Screening{
		screeningID:   screeningID,
		positionID:    positionID,
		applicant:     applicant,
		interviews:    interviews,
		status:        status,
		applyDateTime: applyDateTime,
	}
}

func (s *Screening) AddInterview(interviewDateTime time.Time) error {
	if !s.status.canAddInterview {
		return errors.New("面接が選考中ではありません")
	}
	interviews := &(s.interviews)
	s.interviews = *(interviews.addInterview(interviewDateTime))

	return nil
}

func (s *Screening) Adopt() {
	s.status = Adopted
}

func (s *Screening) Reject() {
	s.status = Rejected
}

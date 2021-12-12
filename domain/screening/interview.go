// 2.6
package screening

import "time"

type Interview struct {
	phase    int
	dateTime time.Time
}

type Interviews struct {
	value []Interview
}

func (i *Interviews) addInterview(interviewDateTime time.Time) *Interviews {
	newInterview := Interview{
		phase:    len(i.value) + 1,
		dateTime: interviewDateTime,
	}
	i.value = append(i.value, newInterview)
	return i
}

func emptyInterviews() Interviews {
	return Interviews{value: make([]Interview, 0)}
}

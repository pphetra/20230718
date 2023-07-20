package member_domain_event

import (
	"encoding/json"
	member_domain "taejai/internal/member/domain"
)

const IndividualMemberRegisteredEventName = "individual_member_registered"

type IndividualMemberRegisteredEvent struct {
	MemberId member_domain.MemberId `json:"member_id"`
}

func NewIndividualMemberRegisteredEvent(memberId member_domain.MemberId) IndividualMemberRegisteredEvent {
	return IndividualMemberRegisteredEvent{
		MemberId: memberId,
	}
}

func (e IndividualMemberRegisteredEvent) GetName() string {
	return IndividualMemberRegisteredEventName
}

func (e IndividualMemberRegisteredEvent) GetPayload() []byte {
	// serialize IndividualMemberRegisteredEvent struct to json
	// and return []byte
	jsonBytes, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	return jsonBytes
}

// use when parsing from event bus
func NewIndividualMemberRegisteredEventFromJsonBytes(jsonBytes []byte) (IndividualMemberRegisteredEvent, error) {
	var event IndividualMemberRegisteredEvent
	err := json.Unmarshal(jsonBytes, &event)
	if err != nil {
		return IndividualMemberRegisteredEvent{}, err
	}

	return event, nil
}

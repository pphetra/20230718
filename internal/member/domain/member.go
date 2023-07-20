package member_domain

import (
	value_object "taejai/internal/shared/value_object"
)

type MemberType int

const (
	Individual MemberType = iota
	Organization
)

type MemberId int64

type Member struct {
	Id      MemberId
	Name1   string
	Name2   string
	Type    MemberType
	Address value_object.Address
}

func NewIndividualMember(
	firstName string,
	lastName string,
	address value_object.Address,
) (Member, error) {
	return Member{
		Name1:   firstName,
		Name2:   lastName,
		Type:    Individual,
		Address: address,
	}, nil
}

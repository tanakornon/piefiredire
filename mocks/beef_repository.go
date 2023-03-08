package mocks

import "github.com/stretchr/testify/mock"

type BeefRepositoryMock struct {
	mock.Mock
}

func (m BeefRepositoryMock) GetText() (string, error) {
	args := m.Called()

	var r0 string
	v0 := args.Get(0)
	if v0 != nil {
		r0 = v0.(string)
	}

	var r1 error
	v1 := args.Get(1)
	if v1 != nil {
		r1 = v1.(error)
	}

	return r0, r1
}

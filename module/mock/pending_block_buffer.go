// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import flow "github.com/dapperlabs/flow-go/model/flow"
import messages "github.com/dapperlabs/flow-go/model/messages"
import mock "github.com/stretchr/testify/mock"

// PendingBlockBuffer is an autogenerated mock type for the PendingBlockBuffer type
type PendingBlockBuffer struct {
	mock.Mock
}

// Add provides a mock function with given fields: originID, proposal
func (_m *PendingBlockBuffer) Add(originID flow.Identifier, proposal *messages.BlockProposal) bool {
	ret := _m.Called(originID, proposal)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier, *messages.BlockProposal) bool); ok {
		r0 = rf(originID, proposal)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ByID provides a mock function with given fields: blockID
func (_m *PendingBlockBuffer) ByID(blockID flow.Identifier) (*flow.PendingBlock, bool) {
	ret := _m.Called(blockID)

	var r0 *flow.PendingBlock
	if rf, ok := ret.Get(0).(func(flow.Identifier) *flow.PendingBlock); ok {
		r0 = rf(blockID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.PendingBlock)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(flow.Identifier) bool); ok {
		r1 = rf(blockID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// ByParentID provides a mock function with given fields: parentID
func (_m *PendingBlockBuffer) ByParentID(parentID flow.Identifier) ([]*flow.PendingBlock, bool) {
	ret := _m.Called(parentID)

	var r0 []*flow.PendingBlock
	if rf, ok := ret.Get(0).(func(flow.Identifier) []*flow.PendingBlock); ok {
		r0 = rf(parentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*flow.PendingBlock)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(flow.Identifier) bool); ok {
		r1 = rf(parentID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// DropForParent provides a mock function with given fields: parentID
func (_m *PendingBlockBuffer) DropForParent(parentID flow.Identifier) {
	_m.Called(parentID)
}

// PruneByHeight provides a mock function with given fields: height
func (_m *PendingBlockBuffer) PruneByHeight(height uint64) {
	_m.Called(height)
}

// Size provides a mock function with given fields:
func (_m *PendingBlockBuffer) Size() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

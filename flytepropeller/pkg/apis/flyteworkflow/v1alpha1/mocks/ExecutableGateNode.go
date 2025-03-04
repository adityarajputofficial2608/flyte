// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/flyteorg/flyte/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// ExecutableGateNode is an autogenerated mock type for the ExecutableGateNode type
type ExecutableGateNode struct {
	mock.Mock
}

type ExecutableGateNode_GetApprove struct {
	*mock.Call
}

func (_m ExecutableGateNode_GetApprove) Return(_a0 *core.ApproveCondition) *ExecutableGateNode_GetApprove {
	return &ExecutableGateNode_GetApprove{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableGateNode) OnGetApprove() *ExecutableGateNode_GetApprove {
	c_call := _m.On("GetApprove")
	return &ExecutableGateNode_GetApprove{Call: c_call}
}

func (_m *ExecutableGateNode) OnGetApproveMatch(matchers ...interface{}) *ExecutableGateNode_GetApprove {
	c_call := _m.On("GetApprove", matchers...)
	return &ExecutableGateNode_GetApprove{Call: c_call}
}

// GetApprove provides a mock function with given fields:
func (_m *ExecutableGateNode) GetApprove() *core.ApproveCondition {
	ret := _m.Called()

	var r0 *core.ApproveCondition
	if rf, ok := ret.Get(0).(func() *core.ApproveCondition); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.ApproveCondition)
		}
	}

	return r0
}

type ExecutableGateNode_GetKind struct {
	*mock.Call
}

func (_m ExecutableGateNode_GetKind) Return(_a0 v1alpha1.ConditionKind) *ExecutableGateNode_GetKind {
	return &ExecutableGateNode_GetKind{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableGateNode) OnGetKind() *ExecutableGateNode_GetKind {
	c_call := _m.On("GetKind")
	return &ExecutableGateNode_GetKind{Call: c_call}
}

func (_m *ExecutableGateNode) OnGetKindMatch(matchers ...interface{}) *ExecutableGateNode_GetKind {
	c_call := _m.On("GetKind", matchers...)
	return &ExecutableGateNode_GetKind{Call: c_call}
}

// GetKind provides a mock function with given fields:
func (_m *ExecutableGateNode) GetKind() v1alpha1.ConditionKind {
	ret := _m.Called()

	var r0 v1alpha1.ConditionKind
	if rf, ok := ret.Get(0).(func() v1alpha1.ConditionKind); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.ConditionKind)
	}

	return r0
}

type ExecutableGateNode_GetSignal struct {
	*mock.Call
}

func (_m ExecutableGateNode_GetSignal) Return(_a0 *core.SignalCondition) *ExecutableGateNode_GetSignal {
	return &ExecutableGateNode_GetSignal{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableGateNode) OnGetSignal() *ExecutableGateNode_GetSignal {
	c_call := _m.On("GetSignal")
	return &ExecutableGateNode_GetSignal{Call: c_call}
}

func (_m *ExecutableGateNode) OnGetSignalMatch(matchers ...interface{}) *ExecutableGateNode_GetSignal {
	c_call := _m.On("GetSignal", matchers...)
	return &ExecutableGateNode_GetSignal{Call: c_call}
}

// GetSignal provides a mock function with given fields:
func (_m *ExecutableGateNode) GetSignal() *core.SignalCondition {
	ret := _m.Called()

	var r0 *core.SignalCondition
	if rf, ok := ret.Get(0).(func() *core.SignalCondition); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.SignalCondition)
		}
	}

	return r0
}

type ExecutableGateNode_GetSleep struct {
	*mock.Call
}

func (_m ExecutableGateNode_GetSleep) Return(_a0 *core.SleepCondition) *ExecutableGateNode_GetSleep {
	return &ExecutableGateNode_GetSleep{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableGateNode) OnGetSleep() *ExecutableGateNode_GetSleep {
	c_call := _m.On("GetSleep")
	return &ExecutableGateNode_GetSleep{Call: c_call}
}

func (_m *ExecutableGateNode) OnGetSleepMatch(matchers ...interface{}) *ExecutableGateNode_GetSleep {
	c_call := _m.On("GetSleep", matchers...)
	return &ExecutableGateNode_GetSleep{Call: c_call}
}

// GetSleep provides a mock function with given fields:
func (_m *ExecutableGateNode) GetSleep() *core.SleepCondition {
	ret := _m.Called()

	var r0 *core.SleepCondition
	if rf, ok := ret.Get(0).(func() *core.SleepCondition); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.SleepCondition)
		}
	}

	return r0
}

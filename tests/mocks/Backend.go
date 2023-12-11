// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	api "github.com/microsoft/durabletask-go/api"
	backend "github.com/microsoft/durabletask-go/backend"

	context "context"

	mock "github.com/stretchr/testify/mock"

	protos "github.com/microsoft/durabletask-go/internal/protos"
)

// Backend is an autogenerated mock type for the Backend type
type Backend struct {
	mock.Mock
}

type Backend_Expecter struct {
	mock *mock.Mock
}

func (_m *Backend) EXPECT() *Backend_Expecter {
	return &Backend_Expecter{mock: &_m.Mock}
}

// AbandonActivityWorkItem provides a mock function with given fields: _a0, _a1
func (_m *Backend) AbandonActivityWorkItem(_a0 context.Context, _a1 *backend.ActivityWorkItem) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AbandonActivityWorkItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *backend.ActivityWorkItem) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backend_AbandonActivityWorkItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AbandonActivityWorkItem'
type Backend_AbandonActivityWorkItem_Call struct {
	*mock.Call
}

// AbandonActivityWorkItem is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *backend.ActivityWorkItem
func (_e *Backend_Expecter) AbandonActivityWorkItem(_a0 interface{}, _a1 interface{}) *Backend_AbandonActivityWorkItem_Call {
	return &Backend_AbandonActivityWorkItem_Call{Call: _e.mock.On("AbandonActivityWorkItem", _a0, _a1)}
}

func (_c *Backend_AbandonActivityWorkItem_Call) Run(run func(_a0 context.Context, _a1 *backend.ActivityWorkItem)) *Backend_AbandonActivityWorkItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*backend.ActivityWorkItem))
	})
	return _c
}

func (_c *Backend_AbandonActivityWorkItem_Call) Return(_a0 error) *Backend_AbandonActivityWorkItem_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Backend_AbandonActivityWorkItem_Call) RunAndReturn(run func(context.Context, *backend.ActivityWorkItem) error) *Backend_AbandonActivityWorkItem_Call {
	_c.Call.Return(run)
	return _c
}

// AbandonOrchestrationWorkItem provides a mock function with given fields: _a0, _a1
func (_m *Backend) AbandonOrchestrationWorkItem(_a0 context.Context, _a1 *backend.OrchestrationWorkItem) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AbandonOrchestrationWorkItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *backend.OrchestrationWorkItem) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backend_AbandonOrchestrationWorkItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AbandonOrchestrationWorkItem'
type Backend_AbandonOrchestrationWorkItem_Call struct {
	*mock.Call
}

// AbandonOrchestrationWorkItem is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *backend.OrchestrationWorkItem
func (_e *Backend_Expecter) AbandonOrchestrationWorkItem(_a0 interface{}, _a1 interface{}) *Backend_AbandonOrchestrationWorkItem_Call {
	return &Backend_AbandonOrchestrationWorkItem_Call{Call: _e.mock.On("AbandonOrchestrationWorkItem", _a0, _a1)}
}

func (_c *Backend_AbandonOrchestrationWorkItem_Call) Run(run func(_a0 context.Context, _a1 *backend.OrchestrationWorkItem)) *Backend_AbandonOrchestrationWorkItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*backend.OrchestrationWorkItem))
	})
	return _c
}

func (_c *Backend_AbandonOrchestrationWorkItem_Call) Return(_a0 error) *Backend_AbandonOrchestrationWorkItem_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Backend_AbandonOrchestrationWorkItem_Call) RunAndReturn(run func(context.Context, *backend.OrchestrationWorkItem) error) *Backend_AbandonOrchestrationWorkItem_Call {
	_c.Call.Return(run)
	return _c
}

// AddNewOrchestrationEvent provides a mock function with given fields: _a0, _a1, _a2
func (_m *Backend) AddNewOrchestrationEvent(_a0 context.Context, _a1 api.InstanceID, _a2 *protos.HistoryEvent) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for AddNewOrchestrationEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, api.InstanceID, *protos.HistoryEvent) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backend_AddNewOrchestrationEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddNewOrchestrationEvent'
type Backend_AddNewOrchestrationEvent_Call struct {
	*mock.Call
}

// AddNewOrchestrationEvent is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 api.InstanceID
//   - _a2 *protos.HistoryEvent
func (_e *Backend_Expecter) AddNewOrchestrationEvent(_a0 interface{}, _a1 interface{}, _a2 interface{}) *Backend_AddNewOrchestrationEvent_Call {
	return &Backend_AddNewOrchestrationEvent_Call{Call: _e.mock.On("AddNewOrchestrationEvent", _a0, _a1, _a2)}
}

func (_c *Backend_AddNewOrchestrationEvent_Call) Run(run func(_a0 context.Context, _a1 api.InstanceID, _a2 *protos.HistoryEvent)) *Backend_AddNewOrchestrationEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(api.InstanceID), args[2].(*protos.HistoryEvent))
	})
	return _c
}

func (_c *Backend_AddNewOrchestrationEvent_Call) Return(_a0 error) *Backend_AddNewOrchestrationEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Backend_AddNewOrchestrationEvent_Call) RunAndReturn(run func(context.Context, api.InstanceID, *protos.HistoryEvent) error) *Backend_AddNewOrchestrationEvent_Call {
	_c.Call.Return(run)
	return _c
}

// CompleteActivityWorkItem provides a mock function with given fields: _a0, _a1
func (_m *Backend) CompleteActivityWorkItem(_a0 context.Context, _a1 *backend.ActivityWorkItem) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CompleteActivityWorkItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *backend.ActivityWorkItem) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backend_CompleteActivityWorkItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompleteActivityWorkItem'
type Backend_CompleteActivityWorkItem_Call struct {
	*mock.Call
}

// CompleteActivityWorkItem is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *backend.ActivityWorkItem
func (_e *Backend_Expecter) CompleteActivityWorkItem(_a0 interface{}, _a1 interface{}) *Backend_CompleteActivityWorkItem_Call {
	return &Backend_CompleteActivityWorkItem_Call{Call: _e.mock.On("CompleteActivityWorkItem", _a0, _a1)}
}

func (_c *Backend_CompleteActivityWorkItem_Call) Run(run func(_a0 context.Context, _a1 *backend.ActivityWorkItem)) *Backend_CompleteActivityWorkItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*backend.ActivityWorkItem))
	})
	return _c
}

func (_c *Backend_CompleteActivityWorkItem_Call) Return(_a0 error) *Backend_CompleteActivityWorkItem_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Backend_CompleteActivityWorkItem_Call) RunAndReturn(run func(context.Context, *backend.ActivityWorkItem) error) *Backend_CompleteActivityWorkItem_Call {
	_c.Call.Return(run)
	return _c
}

// CompleteOrchestrationWorkItem provides a mock function with given fields: _a0, _a1
func (_m *Backend) CompleteOrchestrationWorkItem(_a0 context.Context, _a1 *backend.OrchestrationWorkItem) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CompleteOrchestrationWorkItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *backend.OrchestrationWorkItem) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backend_CompleteOrchestrationWorkItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompleteOrchestrationWorkItem'
type Backend_CompleteOrchestrationWorkItem_Call struct {
	*mock.Call
}

// CompleteOrchestrationWorkItem is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *backend.OrchestrationWorkItem
func (_e *Backend_Expecter) CompleteOrchestrationWorkItem(_a0 interface{}, _a1 interface{}) *Backend_CompleteOrchestrationWorkItem_Call {
	return &Backend_CompleteOrchestrationWorkItem_Call{Call: _e.mock.On("CompleteOrchestrationWorkItem", _a0, _a1)}
}

func (_c *Backend_CompleteOrchestrationWorkItem_Call) Run(run func(_a0 context.Context, _a1 *backend.OrchestrationWorkItem)) *Backend_CompleteOrchestrationWorkItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*backend.OrchestrationWorkItem))
	})
	return _c
}

func (_c *Backend_CompleteOrchestrationWorkItem_Call) Return(_a0 error) *Backend_CompleteOrchestrationWorkItem_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Backend_CompleteOrchestrationWorkItem_Call) RunAndReturn(run func(context.Context, *backend.OrchestrationWorkItem) error) *Backend_CompleteOrchestrationWorkItem_Call {
	_c.Call.Return(run)
	return _c
}

// CreateOrchestrationInstance provides a mock function with given fields: _a0, _a1, _a2
func (_m *Backend) CreateOrchestrationInstance(_a0 context.Context, _a1 *protos.HistoryEvent, _a2 ...backend.OrchestrationIdReusePolicyOptions) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrchestrationInstance")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *protos.HistoryEvent, ...backend.OrchestrationIdReusePolicyOptions) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backend_CreateOrchestrationInstance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrchestrationInstance'
type Backend_CreateOrchestrationInstance_Call struct {
	*mock.Call
}

// CreateOrchestrationInstance is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *protos.HistoryEvent
//   - _a2 ...backend.OrchestrationIdReusePolicyOptions
func (_e *Backend_Expecter) CreateOrchestrationInstance(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *Backend_CreateOrchestrationInstance_Call {
	return &Backend_CreateOrchestrationInstance_Call{Call: _e.mock.On("CreateOrchestrationInstance",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *Backend_CreateOrchestrationInstance_Call) Run(run func(_a0 context.Context, _a1 *protos.HistoryEvent, _a2 ...backend.OrchestrationIdReusePolicyOptions)) *Backend_CreateOrchestrationInstance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]backend.OrchestrationIdReusePolicyOptions, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(backend.OrchestrationIdReusePolicyOptions)
			}
		}
		run(args[0].(context.Context), args[1].(*protos.HistoryEvent), variadicArgs...)
	})
	return _c
}

func (_c *Backend_CreateOrchestrationInstance_Call) Return(_a0 error) *Backend_CreateOrchestrationInstance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Backend_CreateOrchestrationInstance_Call) RunAndReturn(run func(context.Context, *protos.HistoryEvent, ...backend.OrchestrationIdReusePolicyOptions) error) *Backend_CreateOrchestrationInstance_Call {
	_c.Call.Return(run)
	return _c
}

// CreateTaskHub provides a mock function with given fields: _a0
func (_m *Backend) CreateTaskHub(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateTaskHub")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backend_CreateTaskHub_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTaskHub'
type Backend_CreateTaskHub_Call struct {
	*mock.Call
}

// CreateTaskHub is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Backend_Expecter) CreateTaskHub(_a0 interface{}) *Backend_CreateTaskHub_Call {
	return &Backend_CreateTaskHub_Call{Call: _e.mock.On("CreateTaskHub", _a0)}
}

func (_c *Backend_CreateTaskHub_Call) Run(run func(_a0 context.Context)) *Backend_CreateTaskHub_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Backend_CreateTaskHub_Call) Return(_a0 error) *Backend_CreateTaskHub_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Backend_CreateTaskHub_Call) RunAndReturn(run func(context.Context) error) *Backend_CreateTaskHub_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteTaskHub provides a mock function with given fields: _a0
func (_m *Backend) DeleteTaskHub(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTaskHub")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backend_DeleteTaskHub_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteTaskHub'
type Backend_DeleteTaskHub_Call struct {
	*mock.Call
}

// DeleteTaskHub is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Backend_Expecter) DeleteTaskHub(_a0 interface{}) *Backend_DeleteTaskHub_Call {
	return &Backend_DeleteTaskHub_Call{Call: _e.mock.On("DeleteTaskHub", _a0)}
}

func (_c *Backend_DeleteTaskHub_Call) Run(run func(_a0 context.Context)) *Backend_DeleteTaskHub_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Backend_DeleteTaskHub_Call) Return(_a0 error) *Backend_DeleteTaskHub_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Backend_DeleteTaskHub_Call) RunAndReturn(run func(context.Context) error) *Backend_DeleteTaskHub_Call {
	_c.Call.Return(run)
	return _c
}

// GetActivityWorkItem provides a mock function with given fields: _a0
func (_m *Backend) GetActivityWorkItem(_a0 context.Context) (*backend.ActivityWorkItem, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetActivityWorkItem")
	}

	var r0 *backend.ActivityWorkItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*backend.ActivityWorkItem, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *backend.ActivityWorkItem); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backend.ActivityWorkItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Backend_GetActivityWorkItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetActivityWorkItem'
type Backend_GetActivityWorkItem_Call struct {
	*mock.Call
}

// GetActivityWorkItem is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Backend_Expecter) GetActivityWorkItem(_a0 interface{}) *Backend_GetActivityWorkItem_Call {
	return &Backend_GetActivityWorkItem_Call{Call: _e.mock.On("GetActivityWorkItem", _a0)}
}

func (_c *Backend_GetActivityWorkItem_Call) Run(run func(_a0 context.Context)) *Backend_GetActivityWorkItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Backend_GetActivityWorkItem_Call) Return(_a0 *backend.ActivityWorkItem, _a1 error) *Backend_GetActivityWorkItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Backend_GetActivityWorkItem_Call) RunAndReturn(run func(context.Context) (*backend.ActivityWorkItem, error)) *Backend_GetActivityWorkItem_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrchestrationMetadata provides a mock function with given fields: _a0, _a1
func (_m *Backend) GetOrchestrationMetadata(_a0 context.Context, _a1 api.InstanceID) (*api.OrchestrationMetadata, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetOrchestrationMetadata")
	}

	var r0 *api.OrchestrationMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, api.InstanceID) (*api.OrchestrationMetadata, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, api.InstanceID) *api.OrchestrationMetadata); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.OrchestrationMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, api.InstanceID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Backend_GetOrchestrationMetadata_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrchestrationMetadata'
type Backend_GetOrchestrationMetadata_Call struct {
	*mock.Call
}

// GetOrchestrationMetadata is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 api.InstanceID
func (_e *Backend_Expecter) GetOrchestrationMetadata(_a0 interface{}, _a1 interface{}) *Backend_GetOrchestrationMetadata_Call {
	return &Backend_GetOrchestrationMetadata_Call{Call: _e.mock.On("GetOrchestrationMetadata", _a0, _a1)}
}

func (_c *Backend_GetOrchestrationMetadata_Call) Run(run func(_a0 context.Context, _a1 api.InstanceID)) *Backend_GetOrchestrationMetadata_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(api.InstanceID))
	})
	return _c
}

func (_c *Backend_GetOrchestrationMetadata_Call) Return(_a0 *api.OrchestrationMetadata, _a1 error) *Backend_GetOrchestrationMetadata_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Backend_GetOrchestrationMetadata_Call) RunAndReturn(run func(context.Context, api.InstanceID) (*api.OrchestrationMetadata, error)) *Backend_GetOrchestrationMetadata_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrchestrationRuntimeState provides a mock function with given fields: _a0, _a1
func (_m *Backend) GetOrchestrationRuntimeState(_a0 context.Context, _a1 *backend.OrchestrationWorkItem) (*backend.OrchestrationRuntimeState, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetOrchestrationRuntimeState")
	}

	var r0 *backend.OrchestrationRuntimeState
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backend.OrchestrationWorkItem) (*backend.OrchestrationRuntimeState, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backend.OrchestrationWorkItem) *backend.OrchestrationRuntimeState); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backend.OrchestrationRuntimeState)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backend.OrchestrationWorkItem) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Backend_GetOrchestrationRuntimeState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrchestrationRuntimeState'
type Backend_GetOrchestrationRuntimeState_Call struct {
	*mock.Call
}

// GetOrchestrationRuntimeState is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *backend.OrchestrationWorkItem
func (_e *Backend_Expecter) GetOrchestrationRuntimeState(_a0 interface{}, _a1 interface{}) *Backend_GetOrchestrationRuntimeState_Call {
	return &Backend_GetOrchestrationRuntimeState_Call{Call: _e.mock.On("GetOrchestrationRuntimeState", _a0, _a1)}
}

func (_c *Backend_GetOrchestrationRuntimeState_Call) Run(run func(_a0 context.Context, _a1 *backend.OrchestrationWorkItem)) *Backend_GetOrchestrationRuntimeState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*backend.OrchestrationWorkItem))
	})
	return _c
}

func (_c *Backend_GetOrchestrationRuntimeState_Call) Return(_a0 *backend.OrchestrationRuntimeState, _a1 error) *Backend_GetOrchestrationRuntimeState_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Backend_GetOrchestrationRuntimeState_Call) RunAndReturn(run func(context.Context, *backend.OrchestrationWorkItem) (*backend.OrchestrationRuntimeState, error)) *Backend_GetOrchestrationRuntimeState_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrchestrationWorkItem provides a mock function with given fields: _a0
func (_m *Backend) GetOrchestrationWorkItem(_a0 context.Context) (*backend.OrchestrationWorkItem, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetOrchestrationWorkItem")
	}

	var r0 *backend.OrchestrationWorkItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*backend.OrchestrationWorkItem, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *backend.OrchestrationWorkItem); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backend.OrchestrationWorkItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Backend_GetOrchestrationWorkItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrchestrationWorkItem'
type Backend_GetOrchestrationWorkItem_Call struct {
	*mock.Call
}

// GetOrchestrationWorkItem is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Backend_Expecter) GetOrchestrationWorkItem(_a0 interface{}) *Backend_GetOrchestrationWorkItem_Call {
	return &Backend_GetOrchestrationWorkItem_Call{Call: _e.mock.On("GetOrchestrationWorkItem", _a0)}
}

func (_c *Backend_GetOrchestrationWorkItem_Call) Run(run func(_a0 context.Context)) *Backend_GetOrchestrationWorkItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Backend_GetOrchestrationWorkItem_Call) Return(_a0 *backend.OrchestrationWorkItem, _a1 error) *Backend_GetOrchestrationWorkItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Backend_GetOrchestrationWorkItem_Call) RunAndReturn(run func(context.Context) (*backend.OrchestrationWorkItem, error)) *Backend_GetOrchestrationWorkItem_Call {
	_c.Call.Return(run)
	return _c
}

// PurgeOrchestrationState provides a mock function with given fields: _a0, _a1
func (_m *Backend) PurgeOrchestrationState(_a0 context.Context, _a1 api.InstanceID) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for PurgeOrchestrationState")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, api.InstanceID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backend_PurgeOrchestrationState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PurgeOrchestrationState'
type Backend_PurgeOrchestrationState_Call struct {
	*mock.Call
}

// PurgeOrchestrationState is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 api.InstanceID
func (_e *Backend_Expecter) PurgeOrchestrationState(_a0 interface{}, _a1 interface{}) *Backend_PurgeOrchestrationState_Call {
	return &Backend_PurgeOrchestrationState_Call{Call: _e.mock.On("PurgeOrchestrationState", _a0, _a1)}
}

func (_c *Backend_PurgeOrchestrationState_Call) Run(run func(_a0 context.Context, _a1 api.InstanceID)) *Backend_PurgeOrchestrationState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(api.InstanceID))
	})
	return _c
}

func (_c *Backend_PurgeOrchestrationState_Call) Return(_a0 error) *Backend_PurgeOrchestrationState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Backend_PurgeOrchestrationState_Call) RunAndReturn(run func(context.Context, api.InstanceID) error) *Backend_PurgeOrchestrationState_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: _a0
func (_m *Backend) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backend_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Backend_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Backend_Expecter) Start(_a0 interface{}) *Backend_Start_Call {
	return &Backend_Start_Call{Call: _e.mock.On("Start", _a0)}
}

func (_c *Backend_Start_Call) Run(run func(_a0 context.Context)) *Backend_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Backend_Start_Call) Return(_a0 error) *Backend_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Backend_Start_Call) RunAndReturn(run func(context.Context) error) *Backend_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields: _a0
func (_m *Backend) Stop(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backend_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type Backend_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Backend_Expecter) Stop(_a0 interface{}) *Backend_Stop_Call {
	return &Backend_Stop_Call{Call: _e.mock.On("Stop", _a0)}
}

func (_c *Backend_Stop_Call) Run(run func(_a0 context.Context)) *Backend_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Backend_Stop_Call) Return(_a0 error) *Backend_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Backend_Stop_Call) RunAndReturn(run func(context.Context) error) *Backend_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewBackend creates a new instance of Backend. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBackend(t interface {
	mock.TestingT
	Cleanup(func())
}) *Backend {
	mock := &Backend{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

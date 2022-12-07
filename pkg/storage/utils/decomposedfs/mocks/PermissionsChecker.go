// Copyright 2018-2022 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	node "github.com/cs3org/reva/v2/pkg/storage/utils/decomposedfs/node"

	providerv1beta1 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
)

// PermissionsChecker is an autogenerated mock type for the PermissionsChecker type
type PermissionsChecker struct {
	mock.Mock
}

// AssemblePermissions provides a mock function with given fields: ctx, n
func (_m *PermissionsChecker) AssemblePermissions(ctx context.Context, n *node.Node) (providerv1beta1.ResourcePermissions, error) {
	ret := _m.Called(ctx, n)

	var r0 providerv1beta1.ResourcePermissions
	if rf, ok := ret.Get(0).(func(context.Context, *node.Node) providerv1beta1.ResourcePermissions); ok {
		r0 = rf(ctx, n)
	} else {
		r0 = ret.Get(0).(providerv1beta1.ResourcePermissions)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *node.Node) error); ok {
		r1 = rf(ctx, n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPermissionsChecker interface {
	mock.TestingT
	Cleanup(func())
}

// NewPermissionsChecker creates a new instance of PermissionsChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPermissionsChecker(t mockConstructorTestingTNewPermissionsChecker) *PermissionsChecker {
	mock := &PermissionsChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

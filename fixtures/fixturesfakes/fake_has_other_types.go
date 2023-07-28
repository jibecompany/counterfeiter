// Code generated by counterfeiter. DO NOT EDIT.
package fixturesfakes

import (
	"sync"

	"github.com/jibecompany/counterfeiter/v6/fixtures"
)

type FakeHasOtherTypes struct {
	GetThingStub        func(fixtures.SomeString) fixtures.SomeFunc
	getThingMutex       sync.RWMutex
	getThingArgsForCall []struct {
		arg1 fixtures.SomeString
	}
	getThingReturns struct {
		result1 fixtures.SomeFunc
	}
	getThingReturnsOnCall map[int]struct {
		result1 fixtures.SomeFunc
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHasOtherTypes) GetThing(arg1 fixtures.SomeString) fixtures.SomeFunc {
	fake.getThingMutex.Lock()
	ret, specificReturn := fake.getThingReturnsOnCall[len(fake.getThingArgsForCall)]
	fake.getThingArgsForCall = append(fake.getThingArgsForCall, struct {
		arg1 fixtures.SomeString
	}{arg1})
	stub := fake.GetThingStub
	fakeReturns := fake.getThingReturns
	fake.recordInvocation("GetThing", []interface{}{arg1})
	fake.getThingMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHasOtherTypes) GetThingCallCount() int {
	fake.getThingMutex.RLock()
	defer fake.getThingMutex.RUnlock()
	return len(fake.getThingArgsForCall)
}

func (fake *FakeHasOtherTypes) GetThingCalls(stub func(fixtures.SomeString) fixtures.SomeFunc) {
	fake.getThingMutex.Lock()
	defer fake.getThingMutex.Unlock()
	fake.GetThingStub = stub
}

func (fake *FakeHasOtherTypes) GetThingArgsForCall(i int) fixtures.SomeString {
	fake.getThingMutex.RLock()
	defer fake.getThingMutex.RUnlock()
	argsForCall := fake.getThingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHasOtherTypes) GetThingReturns(result1 fixtures.SomeFunc) {
	fake.getThingMutex.Lock()
	defer fake.getThingMutex.Unlock()
	fake.GetThingStub = nil
	fake.getThingReturns = struct {
		result1 fixtures.SomeFunc
	}{result1}
}

func (fake *FakeHasOtherTypes) GetThingReturnsOnCall(i int, result1 fixtures.SomeFunc) {
	fake.getThingMutex.Lock()
	defer fake.getThingMutex.Unlock()
	fake.GetThingStub = nil
	if fake.getThingReturnsOnCall == nil {
		fake.getThingReturnsOnCall = make(map[int]struct {
			result1 fixtures.SomeFunc
		})
	}
	fake.getThingReturnsOnCall[i] = struct {
		result1 fixtures.SomeFunc
	}{result1}
}

func (fake *FakeHasOtherTypes) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getThingMutex.RLock()
	defer fake.getThingMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHasOtherTypes) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ fixtures.HasOtherTypes = new(FakeHasOtherTypes)

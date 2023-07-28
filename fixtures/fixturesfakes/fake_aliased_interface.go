// Code generated by counterfeiter. DO NOT EDIT.
package fixturesfakes

import (
	"sync"

	"github.com/jibecompany/counterfeiter/v6/fixtures"
	"github.com/jibecompany/counterfeiter/v6/fixtures/another_package"
)

type FakeAliasedInterface struct {
	AnotherMethodStub        func([]another_package.SomeType, map[another_package.SomeType]another_package.SomeType, *another_package.SomeType, another_package.SomeType, chan another_package.SomeType)
	anotherMethodMutex       sync.RWMutex
	anotherMethodArgsForCall []struct {
		arg1 []another_package.SomeType
		arg2 map[another_package.SomeType]another_package.SomeType
		arg3 *another_package.SomeType
		arg4 another_package.SomeType
		arg5 chan another_package.SomeType
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAliasedInterface) AnotherMethod(arg1 []another_package.SomeType, arg2 map[another_package.SomeType]another_package.SomeType, arg3 *another_package.SomeType, arg4 another_package.SomeType, arg5 chan another_package.SomeType) {
	var arg1Copy []another_package.SomeType
	if arg1 != nil {
		arg1Copy = make([]another_package.SomeType, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.anotherMethodMutex.Lock()
	fake.anotherMethodArgsForCall = append(fake.anotherMethodArgsForCall, struct {
		arg1 []another_package.SomeType
		arg2 map[another_package.SomeType]another_package.SomeType
		arg3 *another_package.SomeType
		arg4 another_package.SomeType
		arg5 chan another_package.SomeType
	}{arg1Copy, arg2, arg3, arg4, arg5})
	stub := fake.AnotherMethodStub
	fake.recordInvocation("AnotherMethod", []interface{}{arg1Copy, arg2, arg3, arg4, arg5})
	fake.anotherMethodMutex.Unlock()
	if stub != nil {
		fake.AnotherMethodStub(arg1, arg2, arg3, arg4, arg5)
	}
}

func (fake *FakeAliasedInterface) AnotherMethodCallCount() int {
	fake.anotherMethodMutex.RLock()
	defer fake.anotherMethodMutex.RUnlock()
	return len(fake.anotherMethodArgsForCall)
}

func (fake *FakeAliasedInterface) AnotherMethodCalls(stub func([]another_package.SomeType, map[another_package.SomeType]another_package.SomeType, *another_package.SomeType, another_package.SomeType, chan another_package.SomeType)) {
	fake.anotherMethodMutex.Lock()
	defer fake.anotherMethodMutex.Unlock()
	fake.AnotherMethodStub = stub
}

func (fake *FakeAliasedInterface) AnotherMethodArgsForCall(i int) ([]another_package.SomeType, map[another_package.SomeType]another_package.SomeType, *another_package.SomeType, another_package.SomeType, chan another_package.SomeType) {
	fake.anotherMethodMutex.RLock()
	defer fake.anotherMethodMutex.RUnlock()
	argsForCall := fake.anotherMethodArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeAliasedInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.anotherMethodMutex.RLock()
	defer fake.anotherMethodMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAliasedInterface) recordInvocation(key string, args []interface{}) {
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

var _ fixtures.AliasedInterface = new(FakeAliasedInterface)

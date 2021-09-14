// Code generated by counterfeiter. DO NOT EDIT.
package repositoryfakes

import (
	"sync"
	"time"

	"github.com/vmware-tanzu/cartographer/pkg/repository"
)

type FakeExpiringCache struct {
	GetStub        func(interface{}) (interface{}, bool)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 interface{}
	}
	getReturns struct {
		result1 interface{}
		result2 bool
	}
	getReturnsOnCall map[int]struct {
		result1 interface{}
		result2 bool
	}
	SetStub        func(interface{}, interface{}, time.Duration)
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		arg1 interface{}
		arg2 interface{}
		arg3 time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeExpiringCache) Get(arg1 interface{}) (interface{}, bool) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{arg1})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeExpiringCache) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeExpiringCache) GetCalls(stub func(interface{}) (interface{}, bool)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeExpiringCache) GetArgsForCall(i int) interface{} {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeExpiringCache) GetReturns(result1 interface{}, result2 bool) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 interface{}
		result2 bool
	}{result1, result2}
}

func (fake *FakeExpiringCache) GetReturnsOnCall(i int, result1 interface{}, result2 bool) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 bool
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 interface{}
		result2 bool
	}{result1, result2}
}

func (fake *FakeExpiringCache) Set(arg1 interface{}, arg2 interface{}, arg3 time.Duration) {
	fake.setMutex.Lock()
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		arg1 interface{}
		arg2 interface{}
		arg3 time.Duration
	}{arg1, arg2, arg3})
	stub := fake.SetStub
	fake.recordInvocation("Set", []interface{}{arg1, arg2, arg3})
	fake.setMutex.Unlock()
	if stub != nil {
		fake.SetStub(arg1, arg2, arg3)
	}
}

func (fake *FakeExpiringCache) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *FakeExpiringCache) SetCalls(stub func(interface{}, interface{}, time.Duration)) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = stub
}

func (fake *FakeExpiringCache) SetArgsForCall(i int) (interface{}, interface{}, time.Duration) {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	argsForCall := fake.setArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeExpiringCache) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeExpiringCache) recordInvocation(key string, args []interface{}) {
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

var _ repository.ExpiringCache = new(FakeExpiringCache)
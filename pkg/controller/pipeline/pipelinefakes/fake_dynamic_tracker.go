// Code generated by counterfeiter. DO NOT EDIT.
package pipelinefakes

import (
	"sync"

	"github.com/go-logr/logr"
	"github.com/vmware-tanzu/cartographer/pkg/controller/pipeline"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/handler"
)

type FakeDynamicTracker struct {
	WatchStub        func(logr.Logger, runtime.Object, handler.EventHandler) error
	watchMutex       sync.RWMutex
	watchArgsForCall []struct {
		arg1 logr.Logger
		arg2 runtime.Object
		arg3 handler.EventHandler
	}
	watchReturns struct {
		result1 error
	}
	watchReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDynamicTracker) Watch(arg1 logr.Logger, arg2 runtime.Object, arg3 handler.EventHandler) error {
	fake.watchMutex.Lock()
	ret, specificReturn := fake.watchReturnsOnCall[len(fake.watchArgsForCall)]
	fake.watchArgsForCall = append(fake.watchArgsForCall, struct {
		arg1 logr.Logger
		arg2 runtime.Object
		arg3 handler.EventHandler
	}{arg1, arg2, arg3})
	stub := fake.WatchStub
	fakeReturns := fake.watchReturns
	fake.recordInvocation("Watch", []interface{}{arg1, arg2, arg3})
	fake.watchMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDynamicTracker) WatchCallCount() int {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	return len(fake.watchArgsForCall)
}

func (fake *FakeDynamicTracker) WatchCalls(stub func(logr.Logger, runtime.Object, handler.EventHandler) error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = stub
}

func (fake *FakeDynamicTracker) WatchArgsForCall(i int) (logr.Logger, runtime.Object, handler.EventHandler) {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	argsForCall := fake.watchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDynamicTracker) WatchReturns(result1 error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	fake.watchReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDynamicTracker) WatchReturnsOnCall(i int, result1 error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	if fake.watchReturnsOnCall == nil {
		fake.watchReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.watchReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDynamicTracker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDynamicTracker) recordInvocation(key string, args []interface{}) {
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

var _ pipeline.DynamicTracker = new(FakeDynamicTracker)
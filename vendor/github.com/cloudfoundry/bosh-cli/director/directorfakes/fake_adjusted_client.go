// This file was generated by counterfeiter
package directorfakes

import (
	"net/http"
	"sync"

	"github.com/cloudfoundry/bosh-cli/director"
)

type FakeAdjustedClient struct {
	DoStub        func(*http.Request) (*http.Response, error)
	doMutex       sync.RWMutex
	doArgsForCall []struct {
		arg1 *http.Request
	}
	doReturns struct {
		result1 *http.Response
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAdjustedClient) Do(arg1 *http.Request) (*http.Response, error) {
	fake.doMutex.Lock()
	fake.doArgsForCall = append(fake.doArgsForCall, struct {
		arg1 *http.Request
	}{arg1})
	fake.recordInvocation("Do", []interface{}{arg1})
	fake.doMutex.Unlock()
	if fake.DoStub != nil {
		return fake.DoStub(arg1)
	} else {
		return fake.doReturns.result1, fake.doReturns.result2
	}
}

func (fake *FakeAdjustedClient) DoCallCount() int {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return len(fake.doArgsForCall)
}

func (fake *FakeAdjustedClient) DoArgsForCall(i int) *http.Request {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return fake.doArgsForCall[i].arg1
}

func (fake *FakeAdjustedClient) DoReturns(result1 *http.Response, result2 error) {
	fake.DoStub = nil
	fake.doReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeAdjustedClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeAdjustedClient) recordInvocation(key string, args []interface{}) {
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

var _ director.AdjustedClient = new(FakeAdjustedClient)

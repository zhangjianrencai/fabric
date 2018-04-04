// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric/common/chaincode"
	cc "github.com/hyperledger/fabric/core/cclifecycle"
)

type Listener struct {
	LifeCycleChangeListenerStub        func(channel string, chaincodes chaincode.MetadataSet)
	lifeCycleChangeListenerMutex       sync.RWMutex
	lifeCycleChangeListenerArgsForCall []struct {
		channel    string
		chaincodes chaincode.MetadataSet
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Listener) LifeCycleChangeListener(channel string, chaincodes chaincode.MetadataSet) {
	fake.lifeCycleChangeListenerMutex.Lock()
	fake.lifeCycleChangeListenerArgsForCall = append(fake.lifeCycleChangeListenerArgsForCall, struct {
		channel    string
		chaincodes chaincode.MetadataSet
	}{channel, chaincodes})
	fake.recordInvocation("LifeCycleChangeListener", []interface{}{channel, chaincodes})
	fake.lifeCycleChangeListenerMutex.Unlock()
	if fake.LifeCycleChangeListenerStub != nil {
		fake.LifeCycleChangeListenerStub(channel, chaincodes)
	}
}

func (fake *Listener) LifeCycleChangeListenerCallCount() int {
	fake.lifeCycleChangeListenerMutex.RLock()
	defer fake.lifeCycleChangeListenerMutex.RUnlock()
	return len(fake.lifeCycleChangeListenerArgsForCall)
}

func (fake *Listener) LifeCycleChangeListenerArgsForCall(i int) (string, chaincode.MetadataSet) {
	fake.lifeCycleChangeListenerMutex.RLock()
	defer fake.lifeCycleChangeListenerMutex.RUnlock()
	return fake.lifeCycleChangeListenerArgsForCall[i].channel, fake.lifeCycleChangeListenerArgsForCall[i].chaincodes
}

func (fake *Listener) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lifeCycleChangeListenerMutex.RLock()
	defer fake.lifeCycleChangeListenerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Listener) recordInvocation(key string, args []interface{}) {
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

var _ cc.LifeCycleChangeListener = new(Listener)

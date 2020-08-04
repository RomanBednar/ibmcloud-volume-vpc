// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"net/http"
	"sync"

	"github.com/IBM/ibmcloud-volume-vpc/common/vpcclient/instances"
	"github.com/IBM/ibmcloud-volume-vpc/common/vpcclient/models"
	"go.uber.org/zap"
)

type VolumeAttachService struct {
	AttachVolumeStub        func(*models.VolumeAttachment, *zap.Logger) (*models.VolumeAttachment, error)
	attachVolumeMutex       sync.RWMutex
	attachVolumeArgsForCall []struct {
		arg1 *models.VolumeAttachment
		arg2 *zap.Logger
	}
	attachVolumeReturns struct {
		result1 *models.VolumeAttachment
		result2 error
	}
	attachVolumeReturnsOnCall map[int]struct {
		result1 *models.VolumeAttachment
		result2 error
	}
	DetachVolumeStub        func(*models.VolumeAttachment, *zap.Logger) (*http.Response, error)
	detachVolumeMutex       sync.RWMutex
	detachVolumeArgsForCall []struct {
		arg1 *models.VolumeAttachment
		arg2 *zap.Logger
	}
	detachVolumeReturns struct {
		result1 *http.Response
		result2 error
	}
	detachVolumeReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	GetVolumeAttachmentStub        func(*models.VolumeAttachment, *zap.Logger) (*models.VolumeAttachment, error)
	getVolumeAttachmentMutex       sync.RWMutex
	getVolumeAttachmentArgsForCall []struct {
		arg1 *models.VolumeAttachment
		arg2 *zap.Logger
	}
	getVolumeAttachmentReturns struct {
		result1 *models.VolumeAttachment
		result2 error
	}
	getVolumeAttachmentReturnsOnCall map[int]struct {
		result1 *models.VolumeAttachment
		result2 error
	}
	ListVolumeAttachmentsStub        func(*models.VolumeAttachment, *zap.Logger) (*models.VolumeAttachmentList, error)
	listVolumeAttachmentsMutex       sync.RWMutex
	listVolumeAttachmentsArgsForCall []struct {
		arg1 *models.VolumeAttachment
		arg2 *zap.Logger
	}
	listVolumeAttachmentsReturns struct {
		result1 *models.VolumeAttachmentList
		result2 error
	}
	listVolumeAttachmentsReturnsOnCall map[int]struct {
		result1 *models.VolumeAttachmentList
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *VolumeAttachService) AttachVolume(arg1 *models.VolumeAttachment, arg2 *zap.Logger) (*models.VolumeAttachment, error) {
	fake.attachVolumeMutex.Lock()
	ret, specificReturn := fake.attachVolumeReturnsOnCall[len(fake.attachVolumeArgsForCall)]
	fake.attachVolumeArgsForCall = append(fake.attachVolumeArgsForCall, struct {
		arg1 *models.VolumeAttachment
		arg2 *zap.Logger
	}{arg1, arg2})
	fake.recordInvocation("AttachVolume", []interface{}{arg1, arg2})
	fake.attachVolumeMutex.Unlock()
	if fake.AttachVolumeStub != nil {
		return fake.AttachVolumeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.attachVolumeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *VolumeAttachService) AttachVolumeCallCount() int {
	fake.attachVolumeMutex.RLock()
	defer fake.attachVolumeMutex.RUnlock()
	return len(fake.attachVolumeArgsForCall)
}

func (fake *VolumeAttachService) AttachVolumeCalls(stub func(*models.VolumeAttachment, *zap.Logger) (*models.VolumeAttachment, error)) {
	fake.attachVolumeMutex.Lock()
	defer fake.attachVolumeMutex.Unlock()
	fake.AttachVolumeStub = stub
}

func (fake *VolumeAttachService) AttachVolumeArgsForCall(i int) (*models.VolumeAttachment, *zap.Logger) {
	fake.attachVolumeMutex.RLock()
	defer fake.attachVolumeMutex.RUnlock()
	argsForCall := fake.attachVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *VolumeAttachService) AttachVolumeReturns(result1 *models.VolumeAttachment, result2 error) {
	fake.attachVolumeMutex.Lock()
	defer fake.attachVolumeMutex.Unlock()
	fake.AttachVolumeStub = nil
	fake.attachVolumeReturns = struct {
		result1 *models.VolumeAttachment
		result2 error
	}{result1, result2}
}

func (fake *VolumeAttachService) AttachVolumeReturnsOnCall(i int, result1 *models.VolumeAttachment, result2 error) {
	fake.attachVolumeMutex.Lock()
	defer fake.attachVolumeMutex.Unlock()
	fake.AttachVolumeStub = nil
	if fake.attachVolumeReturnsOnCall == nil {
		fake.attachVolumeReturnsOnCall = make(map[int]struct {
			result1 *models.VolumeAttachment
			result2 error
		})
	}
	fake.attachVolumeReturnsOnCall[i] = struct {
		result1 *models.VolumeAttachment
		result2 error
	}{result1, result2}
}

func (fake *VolumeAttachService) DetachVolume(arg1 *models.VolumeAttachment, arg2 *zap.Logger) (*http.Response, error) {
	fake.detachVolumeMutex.Lock()
	ret, specificReturn := fake.detachVolumeReturnsOnCall[len(fake.detachVolumeArgsForCall)]
	fake.detachVolumeArgsForCall = append(fake.detachVolumeArgsForCall, struct {
		arg1 *models.VolumeAttachment
		arg2 *zap.Logger
	}{arg1, arg2})
	fake.recordInvocation("DetachVolume", []interface{}{arg1, arg2})
	fake.detachVolumeMutex.Unlock()
	if fake.DetachVolumeStub != nil {
		return fake.DetachVolumeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.detachVolumeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *VolumeAttachService) DetachVolumeCallCount() int {
	fake.detachVolumeMutex.RLock()
	defer fake.detachVolumeMutex.RUnlock()
	return len(fake.detachVolumeArgsForCall)
}

func (fake *VolumeAttachService) DetachVolumeCalls(stub func(*models.VolumeAttachment, *zap.Logger) (*http.Response, error)) {
	fake.detachVolumeMutex.Lock()
	defer fake.detachVolumeMutex.Unlock()
	fake.DetachVolumeStub = stub
}

func (fake *VolumeAttachService) DetachVolumeArgsForCall(i int) (*models.VolumeAttachment, *zap.Logger) {
	fake.detachVolumeMutex.RLock()
	defer fake.detachVolumeMutex.RUnlock()
	argsForCall := fake.detachVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *VolumeAttachService) DetachVolumeReturns(result1 *http.Response, result2 error) {
	fake.detachVolumeMutex.Lock()
	defer fake.detachVolumeMutex.Unlock()
	fake.DetachVolumeStub = nil
	fake.detachVolumeReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *VolumeAttachService) DetachVolumeReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.detachVolumeMutex.Lock()
	defer fake.detachVolumeMutex.Unlock()
	fake.DetachVolumeStub = nil
	if fake.detachVolumeReturnsOnCall == nil {
		fake.detachVolumeReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.detachVolumeReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *VolumeAttachService) GetVolumeAttachment(arg1 *models.VolumeAttachment, arg2 *zap.Logger) (*models.VolumeAttachment, error) {
	fake.getVolumeAttachmentMutex.Lock()
	ret, specificReturn := fake.getVolumeAttachmentReturnsOnCall[len(fake.getVolumeAttachmentArgsForCall)]
	fake.getVolumeAttachmentArgsForCall = append(fake.getVolumeAttachmentArgsForCall, struct {
		arg1 *models.VolumeAttachment
		arg2 *zap.Logger
	}{arg1, arg2})
	fake.recordInvocation("GetVolumeAttachment", []interface{}{arg1, arg2})
	fake.getVolumeAttachmentMutex.Unlock()
	if fake.GetVolumeAttachmentStub != nil {
		return fake.GetVolumeAttachmentStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getVolumeAttachmentReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *VolumeAttachService) GetVolumeAttachmentCallCount() int {
	fake.getVolumeAttachmentMutex.RLock()
	defer fake.getVolumeAttachmentMutex.RUnlock()
	return len(fake.getVolumeAttachmentArgsForCall)
}

func (fake *VolumeAttachService) GetVolumeAttachmentCalls(stub func(*models.VolumeAttachment, *zap.Logger) (*models.VolumeAttachment, error)) {
	fake.getVolumeAttachmentMutex.Lock()
	defer fake.getVolumeAttachmentMutex.Unlock()
	fake.GetVolumeAttachmentStub = stub
}

func (fake *VolumeAttachService) GetVolumeAttachmentArgsForCall(i int) (*models.VolumeAttachment, *zap.Logger) {
	fake.getVolumeAttachmentMutex.RLock()
	defer fake.getVolumeAttachmentMutex.RUnlock()
	argsForCall := fake.getVolumeAttachmentArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *VolumeAttachService) GetVolumeAttachmentReturns(result1 *models.VolumeAttachment, result2 error) {
	fake.getVolumeAttachmentMutex.Lock()
	defer fake.getVolumeAttachmentMutex.Unlock()
	fake.GetVolumeAttachmentStub = nil
	fake.getVolumeAttachmentReturns = struct {
		result1 *models.VolumeAttachment
		result2 error
	}{result1, result2}
}

func (fake *VolumeAttachService) GetVolumeAttachmentReturnsOnCall(i int, result1 *models.VolumeAttachment, result2 error) {
	fake.getVolumeAttachmentMutex.Lock()
	defer fake.getVolumeAttachmentMutex.Unlock()
	fake.GetVolumeAttachmentStub = nil
	if fake.getVolumeAttachmentReturnsOnCall == nil {
		fake.getVolumeAttachmentReturnsOnCall = make(map[int]struct {
			result1 *models.VolumeAttachment
			result2 error
		})
	}
	fake.getVolumeAttachmentReturnsOnCall[i] = struct {
		result1 *models.VolumeAttachment
		result2 error
	}{result1, result2}
}

func (fake *VolumeAttachService) ListVolumeAttachments(arg1 *models.VolumeAttachment, arg2 *zap.Logger) (*models.VolumeAttachmentList, error) {
	fake.listVolumeAttachmentsMutex.Lock()
	ret, specificReturn := fake.listVolumeAttachmentsReturnsOnCall[len(fake.listVolumeAttachmentsArgsForCall)]
	fake.listVolumeAttachmentsArgsForCall = append(fake.listVolumeAttachmentsArgsForCall, struct {
		arg1 *models.VolumeAttachment
		arg2 *zap.Logger
	}{arg1, arg2})
	fake.recordInvocation("ListVolumeAttachments", []interface{}{arg1, arg2})
	fake.listVolumeAttachmentsMutex.Unlock()
	if fake.ListVolumeAttachmentsStub != nil {
		return fake.ListVolumeAttachmentsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listVolumeAttachmentsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *VolumeAttachService) ListVolumeAttachmentsCallCount() int {
	fake.listVolumeAttachmentsMutex.RLock()
	defer fake.listVolumeAttachmentsMutex.RUnlock()
	return len(fake.listVolumeAttachmentsArgsForCall)
}

func (fake *VolumeAttachService) ListVolumeAttachmentsCalls(stub func(*models.VolumeAttachment, *zap.Logger) (*models.VolumeAttachmentList, error)) {
	fake.listVolumeAttachmentsMutex.Lock()
	defer fake.listVolumeAttachmentsMutex.Unlock()
	fake.ListVolumeAttachmentsStub = stub
}

func (fake *VolumeAttachService) ListVolumeAttachmentsArgsForCall(i int) (*models.VolumeAttachment, *zap.Logger) {
	fake.listVolumeAttachmentsMutex.RLock()
	defer fake.listVolumeAttachmentsMutex.RUnlock()
	argsForCall := fake.listVolumeAttachmentsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *VolumeAttachService) ListVolumeAttachmentsReturns(result1 *models.VolumeAttachmentList, result2 error) {
	fake.listVolumeAttachmentsMutex.Lock()
	defer fake.listVolumeAttachmentsMutex.Unlock()
	fake.ListVolumeAttachmentsStub = nil
	fake.listVolumeAttachmentsReturns = struct {
		result1 *models.VolumeAttachmentList
		result2 error
	}{result1, result2}
}

func (fake *VolumeAttachService) ListVolumeAttachmentsReturnsOnCall(i int, result1 *models.VolumeAttachmentList, result2 error) {
	fake.listVolumeAttachmentsMutex.Lock()
	defer fake.listVolumeAttachmentsMutex.Unlock()
	fake.ListVolumeAttachmentsStub = nil
	if fake.listVolumeAttachmentsReturnsOnCall == nil {
		fake.listVolumeAttachmentsReturnsOnCall = make(map[int]struct {
			result1 *models.VolumeAttachmentList
			result2 error
		})
	}
	fake.listVolumeAttachmentsReturnsOnCall[i] = struct {
		result1 *models.VolumeAttachmentList
		result2 error
	}{result1, result2}
}

func (fake *VolumeAttachService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.attachVolumeMutex.RLock()
	defer fake.attachVolumeMutex.RUnlock()
	fake.detachVolumeMutex.RLock()
	defer fake.detachVolumeMutex.RUnlock()
	fake.getVolumeAttachmentMutex.RLock()
	defer fake.getVolumeAttachmentMutex.RUnlock()
	fake.listVolumeAttachmentsMutex.RLock()
	defer fake.listVolumeAttachmentsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *VolumeAttachService) recordInvocation(key string, args []interface{}) {
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

var _ instances.VolumeAttachManager = new(VolumeAttachService)
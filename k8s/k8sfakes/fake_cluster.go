// Code generated by counterfeiter. DO NOT EDIT.
package k8sfakes

import (
	"sync"

	"github.com/cf-platform-eng/kibosh/k8s"
	"k8s.io/client-go/kubernetes"
	api_v1 "k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/rest"
)

type FakeCluster struct {
	GetClientStub        func() kubernetes.Interface
	getClientMutex       sync.RWMutex
	getClientArgsForCall []struct{}
	getClientReturns     struct {
		result1 kubernetes.Interface
	}
	getClientReturnsOnCall map[int]struct {
		result1 kubernetes.Interface
	}
	GetClientConfigStub        func() *rest.Config
	getClientConfigMutex       sync.RWMutex
	getClientConfigArgsForCall []struct{}
	getClientConfigReturns     struct {
		result1 *rest.Config
	}
	getClientConfigReturnsOnCall map[int]struct {
		result1 *rest.Config
	}
	CreateNamespaceStub        func(*api_v1.Namespace) (*api_v1.Namespace, error)
	createNamespaceMutex       sync.RWMutex
	createNamespaceArgsForCall []struct {
		arg1 *api_v1.Namespace
	}
	createNamespaceReturns struct {
		result1 *api_v1.Namespace
		result2 error
	}
	createNamespaceReturnsOnCall map[int]struct {
		result1 *api_v1.Namespace
		result2 error
	}
	ListPodsStub        func() (*api_v1.PodList, error)
	listPodsMutex       sync.RWMutex
	listPodsArgsForCall []struct{}
	listPodsReturns     struct {
		result1 *api_v1.PodList
		result2 error
	}
	listPodsReturnsOnCall map[int]struct {
		result1 *api_v1.PodList
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCluster) GetClient() kubernetes.Interface {
	fake.getClientMutex.Lock()
	ret, specificReturn := fake.getClientReturnsOnCall[len(fake.getClientArgsForCall)]
	fake.getClientArgsForCall = append(fake.getClientArgsForCall, struct{}{})
	fake.recordInvocation("GetClient", []interface{}{})
	fake.getClientMutex.Unlock()
	if fake.GetClientStub != nil {
		return fake.GetClientStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getClientReturns.result1
}

func (fake *FakeCluster) GetClientCallCount() int {
	fake.getClientMutex.RLock()
	defer fake.getClientMutex.RUnlock()
	return len(fake.getClientArgsForCall)
}

func (fake *FakeCluster) GetClientReturns(result1 kubernetes.Interface) {
	fake.GetClientStub = nil
	fake.getClientReturns = struct {
		result1 kubernetes.Interface
	}{result1}
}

func (fake *FakeCluster) GetClientReturnsOnCall(i int, result1 kubernetes.Interface) {
	fake.GetClientStub = nil
	if fake.getClientReturnsOnCall == nil {
		fake.getClientReturnsOnCall = make(map[int]struct {
			result1 kubernetes.Interface
		})
	}
	fake.getClientReturnsOnCall[i] = struct {
		result1 kubernetes.Interface
	}{result1}
}

func (fake *FakeCluster) GetClientConfig() *rest.Config {
	fake.getClientConfigMutex.Lock()
	ret, specificReturn := fake.getClientConfigReturnsOnCall[len(fake.getClientConfigArgsForCall)]
	fake.getClientConfigArgsForCall = append(fake.getClientConfigArgsForCall, struct{}{})
	fake.recordInvocation("GetClientConfig", []interface{}{})
	fake.getClientConfigMutex.Unlock()
	if fake.GetClientConfigStub != nil {
		return fake.GetClientConfigStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getClientConfigReturns.result1
}

func (fake *FakeCluster) GetClientConfigCallCount() int {
	fake.getClientConfigMutex.RLock()
	defer fake.getClientConfigMutex.RUnlock()
	return len(fake.getClientConfigArgsForCall)
}

func (fake *FakeCluster) GetClientConfigReturns(result1 *rest.Config) {
	fake.GetClientConfigStub = nil
	fake.getClientConfigReturns = struct {
		result1 *rest.Config
	}{result1}
}

func (fake *FakeCluster) GetClientConfigReturnsOnCall(i int, result1 *rest.Config) {
	fake.GetClientConfigStub = nil
	if fake.getClientConfigReturnsOnCall == nil {
		fake.getClientConfigReturnsOnCall = make(map[int]struct {
			result1 *rest.Config
		})
	}
	fake.getClientConfigReturnsOnCall[i] = struct {
		result1 *rest.Config
	}{result1}
}

func (fake *FakeCluster) CreateNamespace(arg1 *api_v1.Namespace) (*api_v1.Namespace, error) {
	fake.createNamespaceMutex.Lock()
	ret, specificReturn := fake.createNamespaceReturnsOnCall[len(fake.createNamespaceArgsForCall)]
	fake.createNamespaceArgsForCall = append(fake.createNamespaceArgsForCall, struct {
		arg1 *api_v1.Namespace
	}{arg1})
	fake.recordInvocation("CreateNamespace", []interface{}{arg1})
	fake.createNamespaceMutex.Unlock()
	if fake.CreateNamespaceStub != nil {
		return fake.CreateNamespaceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createNamespaceReturns.result1, fake.createNamespaceReturns.result2
}

func (fake *FakeCluster) CreateNamespaceCallCount() int {
	fake.createNamespaceMutex.RLock()
	defer fake.createNamespaceMutex.RUnlock()
	return len(fake.createNamespaceArgsForCall)
}

func (fake *FakeCluster) CreateNamespaceArgsForCall(i int) *api_v1.Namespace {
	fake.createNamespaceMutex.RLock()
	defer fake.createNamespaceMutex.RUnlock()
	return fake.createNamespaceArgsForCall[i].arg1
}

func (fake *FakeCluster) CreateNamespaceReturns(result1 *api_v1.Namespace, result2 error) {
	fake.CreateNamespaceStub = nil
	fake.createNamespaceReturns = struct {
		result1 *api_v1.Namespace
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) CreateNamespaceReturnsOnCall(i int, result1 *api_v1.Namespace, result2 error) {
	fake.CreateNamespaceStub = nil
	if fake.createNamespaceReturnsOnCall == nil {
		fake.createNamespaceReturnsOnCall = make(map[int]struct {
			result1 *api_v1.Namespace
			result2 error
		})
	}
	fake.createNamespaceReturnsOnCall[i] = struct {
		result1 *api_v1.Namespace
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) ListPods() (*api_v1.PodList, error) {
	fake.listPodsMutex.Lock()
	ret, specificReturn := fake.listPodsReturnsOnCall[len(fake.listPodsArgsForCall)]
	fake.listPodsArgsForCall = append(fake.listPodsArgsForCall, struct{}{})
	fake.recordInvocation("ListPods", []interface{}{})
	fake.listPodsMutex.Unlock()
	if fake.ListPodsStub != nil {
		return fake.ListPodsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listPodsReturns.result1, fake.listPodsReturns.result2
}

func (fake *FakeCluster) ListPodsCallCount() int {
	fake.listPodsMutex.RLock()
	defer fake.listPodsMutex.RUnlock()
	return len(fake.listPodsArgsForCall)
}

func (fake *FakeCluster) ListPodsReturns(result1 *api_v1.PodList, result2 error) {
	fake.ListPodsStub = nil
	fake.listPodsReturns = struct {
		result1 *api_v1.PodList
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) ListPodsReturnsOnCall(i int, result1 *api_v1.PodList, result2 error) {
	fake.ListPodsStub = nil
	if fake.listPodsReturnsOnCall == nil {
		fake.listPodsReturnsOnCall = make(map[int]struct {
			result1 *api_v1.PodList
			result2 error
		})
	}
	fake.listPodsReturnsOnCall[i] = struct {
		result1 *api_v1.PodList
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getClientMutex.RLock()
	defer fake.getClientMutex.RUnlock()
	fake.getClientConfigMutex.RLock()
	defer fake.getClientConfigMutex.RUnlock()
	fake.createNamespaceMutex.RLock()
	defer fake.createNamespaceMutex.RUnlock()
	fake.listPodsMutex.RLock()
	defer fake.listPodsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCluster) recordInvocation(key string, args []interface{}) {
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

var _ k8s.Cluster = new(FakeCluster)

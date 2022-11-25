// Code generated by MockGen. DO NOT EDIT.
// Source: registry/client.go

// Package mock_registry is a generated GoMock package.
package mock_registry

import (
	big "math/big"
	reflect "reflect"

	contract_agent_registry "github.com/forta-network/forta-core-go/contracts/contract_agent_registry"
	contract_scanner_pool_registry "github.com/forta-network/forta-core-go/contracts/contract_scanner_pool_registry"
	contract_scanner_registry "github.com/forta-network/forta-core-go/contracts/contract_scanner_registry"
	registry "github.com/forta-network/forta-core-go/domain/registry"
	registry0 "github.com/forta-network/forta-core-go/registry"
	eip712 "github.com/forta-network/forta-core-go/security/eip712"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// DisableScanner mocks base method.
func (m *MockClient) DisableScanner(ScannerPermission registry0.ScannerPermission, scannerAddress string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableScanner", ScannerPermission, scannerAddress)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableScanner indicates an expected call of DisableScanner.
func (mr *MockClientMockRecorder) DisableScanner(ScannerPermission, scannerAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableScanner", reflect.TypeOf((*MockClient)(nil).DisableScanner), ScannerPermission, scannerAddress)
}

// EnableScanner mocks base method.
func (m *MockClient) EnableScanner(ScannerPermission registry0.ScannerPermission, scannerAddress string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableScanner", ScannerPermission, scannerAddress)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableScanner indicates an expected call of EnableScanner.
func (mr *MockClientMockRecorder) EnableScanner(ScannerPermission, scannerAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableScanner", reflect.TypeOf((*MockClient)(nil).EnableScanner), ScannerPermission, scannerAddress)
}

// ForEachAgent mocks base method.
func (m *MockClient) ForEachAgent(handler func(*registry0.Agent) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachAgent", handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachAgent indicates an expected call of ForEachAgent.
func (mr *MockClientMockRecorder) ForEachAgent(handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachAgent", reflect.TypeOf((*MockClient)(nil).ForEachAgent), handler)
}

// ForEachAgentID mocks base method.
func (m *MockClient) ForEachAgentID(handler func(string) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachAgentID", handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachAgentID indicates an expected call of ForEachAgentID.
func (mr *MockClientMockRecorder) ForEachAgentID(handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachAgentID", reflect.TypeOf((*MockClient)(nil).ForEachAgentID), handler)
}

// ForEachAgentSinceBlock mocks base method.
func (m *MockClient) ForEachAgentSinceBlock(block uint64, handler func(*contract_agent_registry.AgentRegistryAgentUpdated, *registry0.Agent) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachAgentSinceBlock", block, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachAgentSinceBlock indicates an expected call of ForEachAgentSinceBlock.
func (mr *MockClientMockRecorder) ForEachAgentSinceBlock(block, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachAgentSinceBlock", reflect.TypeOf((*MockClient)(nil).ForEachAgentSinceBlock), block, handler)
}

// ForEachAssignedAgent mocks base method.
func (m *MockClient) ForEachAssignedAgent(scannerID string, handler func(*registry0.Agent) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachAssignedAgent", scannerID, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachAssignedAgent indicates an expected call of ForEachAssignedAgent.
func (mr *MockClientMockRecorder) ForEachAssignedAgent(scannerID, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachAssignedAgent", reflect.TypeOf((*MockClient)(nil).ForEachAssignedAgent), scannerID, handler)
}

// ForEachAssignedScanner mocks base method.
func (m *MockClient) ForEachAssignedScanner(agentID string, handler func(*registry0.Scanner) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachAssignedScanner", agentID, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachAssignedScanner indicates an expected call of ForEachAssignedScanner.
func (mr *MockClientMockRecorder) ForEachAssignedScanner(agentID, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachAssignedScanner", reflect.TypeOf((*MockClient)(nil).ForEachAssignedScanner), agentID, handler)
}

// ForEachChainAgent mocks base method.
func (m *MockClient) ForEachChainAgent(chainID int64, handler func(*registry0.Agent) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachChainAgent", chainID, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachChainAgent indicates an expected call of ForEachChainAgent.
func (mr *MockClientMockRecorder) ForEachChainAgent(chainID, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachChainAgent", reflect.TypeOf((*MockClient)(nil).ForEachChainAgent), chainID, handler)
}

// ForEachPoolScannerSinceBlock mocks base method.
func (m *MockClient) ForEachPoolScannerSinceBlock(block uint64, handler func(*contract_scanner_pool_registry.ScannerPoolRegistryScannerUpdated, *registry0.Scanner) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachPoolScannerSinceBlock", block, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachPoolScannerSinceBlock indicates an expected call of ForEachPoolScannerSinceBlock.
func (mr *MockClientMockRecorder) ForEachPoolScannerSinceBlock(block, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachPoolScannerSinceBlock", reflect.TypeOf((*MockClient)(nil).ForEachPoolScannerSinceBlock), block, handler)
}

// ForEachScanner mocks base method.
func (m *MockClient) ForEachScanner(handler func(*registry0.Scanner) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachScanner", handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachScanner indicates an expected call of ForEachScanner.
func (mr *MockClientMockRecorder) ForEachScanner(handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachScanner", reflect.TypeOf((*MockClient)(nil).ForEachScanner), handler)
}

// ForEachScannerSinceBlock mocks base method.
func (m *MockClient) ForEachScannerSinceBlock(block uint64, handler func(*contract_scanner_registry.ScannerRegistryScannerUpdated, *registry0.Scanner) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachScannerSinceBlock", block, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachScannerSinceBlock indicates an expected call of ForEachScannerSinceBlock.
func (mr *MockClientMockRecorder) ForEachScannerSinceBlock(block, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachScannerSinceBlock", reflect.TypeOf((*MockClient)(nil).ForEachScannerSinceBlock), block, handler)
}

// GenerateScannerRegistrationSignature mocks base method.
func (m *MockClient) GenerateScannerRegistrationSignature(reg *eip712.ScannerNodeRegistration) ([]byte, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateScannerRegistrationSignature", reg)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenerateScannerRegistrationSignature indicates an expected call of GenerateScannerRegistrationSignature.
func (mr *MockClientMockRecorder) GenerateScannerRegistrationSignature(reg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateScannerRegistrationSignature", reflect.TypeOf((*MockClient)(nil).GenerateScannerRegistrationSignature), reg)
}

// GetAgent mocks base method.
func (m *MockClient) GetAgent(agentID string) (*registry0.Agent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgent", agentID)
	ret0, _ := ret[0].(*registry0.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgent indicates an expected call of GetAgent.
func (mr *MockClientMockRecorder) GetAgent(agentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgent", reflect.TypeOf((*MockClient)(nil).GetAgent), agentID)
}

// GetAssignmentHash mocks base method.
func (m *MockClient) GetAssignmentHash(scannerID string) (*registry0.AssignmentHash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssignmentHash", scannerID)
	ret0, _ := ret[0].(*registry0.AssignmentHash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssignmentHash indicates an expected call of GetAssignmentHash.
func (mr *MockClientMockRecorder) GetAssignmentHash(scannerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssignmentHash", reflect.TypeOf((*MockClient)(nil).GetAssignmentHash), scannerID)
}

// GetPoolScanner mocks base method.
func (m *MockClient) GetPoolScanner(scannerID string) (*registry0.Scanner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPoolScanner", scannerID)
	ret0, _ := ret[0].(*registry0.Scanner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPoolScanner indicates an expected call of GetPoolScanner.
func (mr *MockClientMockRecorder) GetPoolScanner(scannerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPoolScanner", reflect.TypeOf((*MockClient)(nil).GetPoolScanner), scannerID)
}

// GetScanner mocks base method.
func (m *MockClient) GetScanner(scannerID string) (*registry0.Scanner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScanner", scannerID)
	ret0, _ := ret[0].(*registry0.Scanner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScanner indicates an expected call of GetScanner.
func (mr *MockClientMockRecorder) GetScanner(scannerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScanner", reflect.TypeOf((*MockClient)(nil).GetScanner), scannerID)
}

// GetScannerNodePrereleaseVersion mocks base method.
func (m *MockClient) GetScannerNodePrereleaseVersion() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScannerNodePrereleaseVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScannerNodePrereleaseVersion indicates an expected call of GetScannerNodePrereleaseVersion.
func (mr *MockClientMockRecorder) GetScannerNodePrereleaseVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScannerNodePrereleaseVersion", reflect.TypeOf((*MockClient)(nil).GetScannerNodePrereleaseVersion))
}

// GetScannerNodeVersion mocks base method.
func (m *MockClient) GetScannerNodeVersion() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScannerNodeVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScannerNodeVersion indicates an expected call of GetScannerNodeVersion.
func (mr *MockClientMockRecorder) GetScannerNodeVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScannerNodeVersion", reflect.TypeOf((*MockClient)(nil).GetScannerNodeVersion))
}

// GetScannerPoolOwner mocks base method.
func (m *MockClient) GetScannerPoolOwner(poolID *big.Int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScannerPoolOwner", poolID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScannerPoolOwner indicates an expected call of GetScannerPoolOwner.
func (mr *MockClientMockRecorder) GetScannerPoolOwner(poolID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScannerPoolOwner", reflect.TypeOf((*MockClient)(nil).GetScannerPoolOwner), poolID)
}

// GetStakingThreshold mocks base method.
func (m *MockClient) GetStakingThreshold(scannerID string) (*registry0.StakingThreshold, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStakingThreshold", scannerID)
	ret0, _ := ret[0].(*registry0.StakingThreshold)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStakingThreshold indicates an expected call of GetStakingThreshold.
func (mr *MockClientMockRecorder) GetStakingThreshold(scannerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStakingThreshold", reflect.TypeOf((*MockClient)(nil).GetStakingThreshold), scannerID)
}

// IsAssigned mocks base method.
func (m *MockClient) IsAssigned(scannerID, agentID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAssigned", scannerID, agentID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAssigned indicates an expected call of IsAssigned.
func (mr *MockClientMockRecorder) IsAssigned(scannerID, agentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAssigned", reflect.TypeOf((*MockClient)(nil).IsAssigned), scannerID, agentID)
}

// IsEnabledScanner mocks base method.
func (m *MockClient) IsEnabledScanner(scannerID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEnabledScanner", scannerID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsEnabledScanner indicates an expected call of IsEnabledScanner.
func (mr *MockClientMockRecorder) IsEnabledScanner(scannerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEnabledScanner", reflect.TypeOf((*MockClient)(nil).IsEnabledScanner), scannerID)
}

// IsOperationalScanner mocks base method.
func (m *MockClient) IsOperationalScanner(scannerID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOperationalScanner", scannerID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOperationalScanner indicates an expected call of IsOperationalScanner.
func (mr *MockClientMockRecorder) IsOperationalScanner(scannerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOperationalScanner", reflect.TypeOf((*MockClient)(nil).IsOperationalScanner), scannerID)
}

// PegBlock mocks base method.
func (m *MockClient) PegBlock(blockNum *big.Int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PegBlock", blockNum)
}

// PegBlock indicates an expected call of PegBlock.
func (mr *MockClientMockRecorder) PegBlock(blockNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PegBlock", reflect.TypeOf((*MockClient)(nil).PegBlock), blockNum)
}

// PegLatestBlock mocks base method.
func (m *MockClient) PegLatestBlock() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PegLatestBlock")
	ret0, _ := ret[0].(error)
	return ret0
}

// PegLatestBlock indicates an expected call of PegLatestBlock.
func (mr *MockClientMockRecorder) PegLatestBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PegLatestBlock", reflect.TypeOf((*MockClient)(nil).PegLatestBlock))
}

// RegisterScannerToPool mocks base method.
func (m *MockClient) RegisterScannerToPool(scannerAddress string, poolID *big.Int, chainID int64, metadata string, signature []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterScannerToPool", scannerAddress, poolID, chainID, metadata, signature)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterScannerToPool indicates an expected call of RegisterScannerToPool.
func (mr *MockClientMockRecorder) RegisterScannerToPool(scannerAddress, poolID, chainID, metadata, signature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterScannerToPool", reflect.TypeOf((*MockClient)(nil).RegisterScannerToPool), scannerAddress, poolID, chainID, metadata, signature)
}

// RegistryContracts mocks base method.
func (m *MockClient) RegistryContracts() *registry.RegistryContracts {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistryContracts")
	ret0, _ := ret[0].(*registry.RegistryContracts)
	return ret0
}

// RegistryContracts indicates an expected call of RegistryContracts.
func (mr *MockClientMockRecorder) RegistryContracts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistryContracts", reflect.TypeOf((*MockClient)(nil).RegistryContracts))
}

// ResetOpts mocks base method.
func (m *MockClient) ResetOpts() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetOpts")
}

// ResetOpts indicates an expected call of ResetOpts.
func (mr *MockClientMockRecorder) ResetOpts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetOpts", reflect.TypeOf((*MockClient)(nil).ResetOpts))
}

// WillNewScannerShutdownPool mocks base method.
func (m *MockClient) WillNewScannerShutdownPool(poolID *big.Int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WillNewScannerShutdownPool", poolID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WillNewScannerShutdownPool indicates an expected call of WillNewScannerShutdownPool.
func (mr *MockClientMockRecorder) WillNewScannerShutdownPool(poolID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WillNewScannerShutdownPool", reflect.TypeOf((*MockClient)(nil).WillNewScannerShutdownPool), poolID)
}

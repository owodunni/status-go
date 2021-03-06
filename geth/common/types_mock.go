// Code generated by MockGen. DO NOT EDIT.
// Source: geth/common/types.go

// Package common is a generated GoMock package.
package common

import (
	context "context"
	accounts "github.com/ethereum/go-ethereum/accounts"
	keystore "github.com/ethereum/go-ethereum/accounts/keystore"
	common "github.com/ethereum/go-ethereum/common"
	les "github.com/ethereum/go-ethereum/les"
	node "github.com/ethereum/go-ethereum/node"
	whisperv5 "github.com/ethereum/go-ethereum/whisper/whisperv5"
	gomock "github.com/golang/mock/gomock"
	otto "github.com/robertkrimen/otto"
	params "github.com/status-im/status-go/geth/params"
	rpc "github.com/status-im/status-go/geth/rpc"
	reflect "reflect"
)

// MockNodeManager is a mock of NodeManager interface
type MockNodeManager struct {
	ctrl     *gomock.Controller
	recorder *MockNodeManagerMockRecorder
}

// MockNodeManagerMockRecorder is the mock recorder for MockNodeManager
type MockNodeManagerMockRecorder struct {
	mock *MockNodeManager
}

// NewMockNodeManager creates a new mock instance
func NewMockNodeManager(ctrl *gomock.Controller) *MockNodeManager {
	mock := &MockNodeManager{ctrl: ctrl}
	mock.recorder = &MockNodeManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeManager) EXPECT() *MockNodeManagerMockRecorder {
	return m.recorder
}

// StartNode mocks base method
func (m *MockNodeManager) StartNode(config *params.NodeConfig) (<-chan struct{}, error) {
	ret := m.ctrl.Call(m, "StartNode", config)
	ret0, _ := ret[0].(<-chan struct{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartNode indicates an expected call of StartNode
func (mr *MockNodeManagerMockRecorder) StartNode(config interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartNode", reflect.TypeOf((*MockNodeManager)(nil).StartNode), config)
}

// StopNode mocks base method
func (m *MockNodeManager) StopNode() (<-chan struct{}, error) {
	ret := m.ctrl.Call(m, "StopNode")
	ret0, _ := ret[0].(<-chan struct{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopNode indicates an expected call of StopNode
func (mr *MockNodeManagerMockRecorder) StopNode() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopNode", reflect.TypeOf((*MockNodeManager)(nil).StopNode))
}

// RestartNode mocks base method
func (m *MockNodeManager) RestartNode() (<-chan struct{}, error) {
	ret := m.ctrl.Call(m, "RestartNode")
	ret0, _ := ret[0].(<-chan struct{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestartNode indicates an expected call of RestartNode
func (mr *MockNodeManagerMockRecorder) RestartNode() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestartNode", reflect.TypeOf((*MockNodeManager)(nil).RestartNode))
}

// ResetChainData mocks base method
func (m *MockNodeManager) ResetChainData() (<-chan struct{}, error) {
	ret := m.ctrl.Call(m, "ResetChainData")
	ret0, _ := ret[0].(<-chan struct{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetChainData indicates an expected call of ResetChainData
func (mr *MockNodeManagerMockRecorder) ResetChainData() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetChainData", reflect.TypeOf((*MockNodeManager)(nil).ResetChainData))
}

// IsNodeRunning mocks base method
func (m *MockNodeManager) IsNodeRunning() bool {
	ret := m.ctrl.Call(m, "IsNodeRunning")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNodeRunning indicates an expected call of IsNodeRunning
func (mr *MockNodeManagerMockRecorder) IsNodeRunning() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNodeRunning", reflect.TypeOf((*MockNodeManager)(nil).IsNodeRunning))
}

// NodeConfig mocks base method
func (m *MockNodeManager) NodeConfig() (*params.NodeConfig, error) {
	ret := m.ctrl.Call(m, "NodeConfig")
	ret0, _ := ret[0].(*params.NodeConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeConfig indicates an expected call of NodeConfig
func (mr *MockNodeManagerMockRecorder) NodeConfig() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeConfig", reflect.TypeOf((*MockNodeManager)(nil).NodeConfig))
}

// Node mocks base method
func (m *MockNodeManager) Node() (*node.Node, error) {
	ret := m.ctrl.Call(m, "Node")
	ret0, _ := ret[0].(*node.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Node indicates an expected call of Node
func (mr *MockNodeManagerMockRecorder) Node() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Node", reflect.TypeOf((*MockNodeManager)(nil).Node))
}

// PopulateStaticPeers mocks base method
func (m *MockNodeManager) PopulateStaticPeers() error {
	ret := m.ctrl.Call(m, "PopulateStaticPeers")
	ret0, _ := ret[0].(error)
	return ret0
}

// PopulateStaticPeers indicates an expected call of PopulateStaticPeers
func (mr *MockNodeManagerMockRecorder) PopulateStaticPeers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopulateStaticPeers", reflect.TypeOf((*MockNodeManager)(nil).PopulateStaticPeers))
}

// AddPeer mocks base method
func (m *MockNodeManager) AddPeer(url string) error {
	ret := m.ctrl.Call(m, "AddPeer", url)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPeer indicates an expected call of AddPeer
func (mr *MockNodeManagerMockRecorder) AddPeer(url interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPeer", reflect.TypeOf((*MockNodeManager)(nil).AddPeer), url)
}

// LightEthereumService mocks base method
func (m *MockNodeManager) LightEthereumService() (*les.LightEthereum, error) {
	ret := m.ctrl.Call(m, "LightEthereumService")
	ret0, _ := ret[0].(*les.LightEthereum)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LightEthereumService indicates an expected call of LightEthereumService
func (mr *MockNodeManagerMockRecorder) LightEthereumService() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LightEthereumService", reflect.TypeOf((*MockNodeManager)(nil).LightEthereumService))
}

// WhisperService mocks base method
func (m *MockNodeManager) WhisperService() (*whisperv5.Whisper, error) {
	ret := m.ctrl.Call(m, "WhisperService")
	ret0, _ := ret[0].(*whisperv5.Whisper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WhisperService indicates an expected call of WhisperService
func (mr *MockNodeManagerMockRecorder) WhisperService() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhisperService", reflect.TypeOf((*MockNodeManager)(nil).WhisperService))
}

// AccountManager mocks base method
func (m *MockNodeManager) AccountManager() (*accounts.Manager, error) {
	ret := m.ctrl.Call(m, "AccountManager")
	ret0, _ := ret[0].(*accounts.Manager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccountManager indicates an expected call of AccountManager
func (mr *MockNodeManagerMockRecorder) AccountManager() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountManager", reflect.TypeOf((*MockNodeManager)(nil).AccountManager))
}

// AccountKeyStore mocks base method
func (m *MockNodeManager) AccountKeyStore() (*keystore.KeyStore, error) {
	ret := m.ctrl.Call(m, "AccountKeyStore")
	ret0, _ := ret[0].(*keystore.KeyStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccountKeyStore indicates an expected call of AccountKeyStore
func (mr *MockNodeManagerMockRecorder) AccountKeyStore() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountKeyStore", reflect.TypeOf((*MockNodeManager)(nil).AccountKeyStore))
}

// RPCClient mocks base method
func (m *MockNodeManager) RPCClient() *rpc.Client {
	ret := m.ctrl.Call(m, "RPCClient")
	ret0, _ := ret[0].(*rpc.Client)
	return ret0
}

// RPCClient indicates an expected call of RPCClient
func (mr *MockNodeManagerMockRecorder) RPCClient() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPCClient", reflect.TypeOf((*MockNodeManager)(nil).RPCClient))
}

// MockAccountManager is a mock of AccountManager interface
type MockAccountManager struct {
	ctrl     *gomock.Controller
	recorder *MockAccountManagerMockRecorder
}

// MockAccountManagerMockRecorder is the mock recorder for MockAccountManager
type MockAccountManagerMockRecorder struct {
	mock *MockAccountManager
}

// NewMockAccountManager creates a new mock instance
func NewMockAccountManager(ctrl *gomock.Controller) *MockAccountManager {
	mock := &MockAccountManager{ctrl: ctrl}
	mock.recorder = &MockAccountManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountManager) EXPECT() *MockAccountManagerMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method
func (m *MockAccountManager) CreateAccount(password string) (string, string, string, error) {
	ret := m.ctrl.Call(m, "CreateAccount", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockAccountManagerMockRecorder) CreateAccount(password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountManager)(nil).CreateAccount), password)
}

// CreateChildAccount mocks base method
func (m *MockAccountManager) CreateChildAccount(parentAddress, password string) (string, string, error) {
	ret := m.ctrl.Call(m, "CreateChildAccount", parentAddress, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateChildAccount indicates an expected call of CreateChildAccount
func (mr *MockAccountManagerMockRecorder) CreateChildAccount(parentAddress, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChildAccount", reflect.TypeOf((*MockAccountManager)(nil).CreateChildAccount), parentAddress, password)
}

// RecoverAccount mocks base method
func (m *MockAccountManager) RecoverAccount(password, mnemonic string) (string, string, error) {
	ret := m.ctrl.Call(m, "RecoverAccount", password, mnemonic)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RecoverAccount indicates an expected call of RecoverAccount
func (mr *MockAccountManagerMockRecorder) RecoverAccount(password, mnemonic interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecoverAccount", reflect.TypeOf((*MockAccountManager)(nil).RecoverAccount), password, mnemonic)
}

// VerifyAccountPassword mocks base method
func (m *MockAccountManager) VerifyAccountPassword(keyStoreDir, address, password string) (*keystore.Key, error) {
	ret := m.ctrl.Call(m, "VerifyAccountPassword", keyStoreDir, address, password)
	ret0, _ := ret[0].(*keystore.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyAccountPassword indicates an expected call of VerifyAccountPassword
func (mr *MockAccountManagerMockRecorder) VerifyAccountPassword(keyStoreDir, address, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyAccountPassword", reflect.TypeOf((*MockAccountManager)(nil).VerifyAccountPassword), keyStoreDir, address, password)
}

// SelectAccount mocks base method
func (m *MockAccountManager) SelectAccount(address, password string) error {
	ret := m.ctrl.Call(m, "SelectAccount", address, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectAccount indicates an expected call of SelectAccount
func (mr *MockAccountManagerMockRecorder) SelectAccount(address, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAccount", reflect.TypeOf((*MockAccountManager)(nil).SelectAccount), address, password)
}

// ReSelectAccount mocks base method
func (m *MockAccountManager) ReSelectAccount() error {
	ret := m.ctrl.Call(m, "ReSelectAccount")
	ret0, _ := ret[0].(error)
	return ret0
}

// ReSelectAccount indicates an expected call of ReSelectAccount
func (mr *MockAccountManagerMockRecorder) ReSelectAccount() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReSelectAccount", reflect.TypeOf((*MockAccountManager)(nil).ReSelectAccount))
}

// SelectedAccount mocks base method
func (m *MockAccountManager) SelectedAccount() (*SelectedExtKey, error) {
	ret := m.ctrl.Call(m, "SelectedAccount")
	ret0, _ := ret[0].(*SelectedExtKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectedAccount indicates an expected call of SelectedAccount
func (mr *MockAccountManagerMockRecorder) SelectedAccount() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectedAccount", reflect.TypeOf((*MockAccountManager)(nil).SelectedAccount))
}

// Logout mocks base method
func (m *MockAccountManager) Logout() error {
	ret := m.ctrl.Call(m, "Logout")
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout
func (mr *MockAccountManagerMockRecorder) Logout() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAccountManager)(nil).Logout))
}

// Accounts mocks base method
func (m *MockAccountManager) Accounts() ([]common.Address, error) {
	ret := m.ctrl.Call(m, "Accounts")
	ret0, _ := ret[0].([]common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Accounts indicates an expected call of Accounts
func (mr *MockAccountManagerMockRecorder) Accounts() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accounts", reflect.TypeOf((*MockAccountManager)(nil).Accounts))
}

// AccountsRPCHandler mocks base method
func (m *MockAccountManager) AccountsRPCHandler() rpc.Handler {
	ret := m.ctrl.Call(m, "AccountsRPCHandler")
	ret0, _ := ret[0].(rpc.Handler)
	return ret0
}

// AccountsRPCHandler indicates an expected call of AccountsRPCHandler
func (mr *MockAccountManagerMockRecorder) AccountsRPCHandler() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountsRPCHandler", reflect.TypeOf((*MockAccountManager)(nil).AccountsRPCHandler))
}

// AddressToDecryptedAccount mocks base method
func (m *MockAccountManager) AddressToDecryptedAccount(address, password string) (accounts.Account, *keystore.Key, error) {
	ret := m.ctrl.Call(m, "AddressToDecryptedAccount", address, password)
	ret0, _ := ret[0].(accounts.Account)
	ret1, _ := ret[1].(*keystore.Key)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddressToDecryptedAccount indicates an expected call of AddressToDecryptedAccount
func (mr *MockAccountManagerMockRecorder) AddressToDecryptedAccount(address, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddressToDecryptedAccount", reflect.TypeOf((*MockAccountManager)(nil).AddressToDecryptedAccount), address, password)
}

// MockTxQueue is a mock of TxQueue interface
type MockTxQueue struct {
	ctrl     *gomock.Controller
	recorder *MockTxQueueMockRecorder
}

// MockTxQueueMockRecorder is the mock recorder for MockTxQueue
type MockTxQueueMockRecorder struct {
	mock *MockTxQueue
}

// NewMockTxQueue creates a new mock instance
func NewMockTxQueue(ctrl *gomock.Controller) *MockTxQueue {
	mock := &MockTxQueue{ctrl: ctrl}
	mock.recorder = &MockTxQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTxQueue) EXPECT() *MockTxQueueMockRecorder {
	return m.recorder
}

// Remove mocks base method
func (m *MockTxQueue) Remove(id QueuedTxID) {
	m.ctrl.Call(m, "Remove", id)
}

// Remove indicates an expected call of Remove
func (mr *MockTxQueueMockRecorder) Remove(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockTxQueue)(nil).Remove), id)
}

// Reset mocks base method
func (m *MockTxQueue) Reset() {
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset
func (mr *MockTxQueueMockRecorder) Reset() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockTxQueue)(nil).Reset))
}

// Count mocks base method
func (m *MockTxQueue) Count() int {
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int)
	return ret0
}

// Count indicates an expected call of Count
func (mr *MockTxQueueMockRecorder) Count() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockTxQueue)(nil).Count))
}

// Has mocks base method
func (m *MockTxQueue) Has(id QueuedTxID) bool {
	ret := m.ctrl.Call(m, "Has", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockTxQueueMockRecorder) Has(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockTxQueue)(nil).Has), id)
}

// MockTxQueueManager is a mock of TxQueueManager interface
type MockTxQueueManager struct {
	ctrl     *gomock.Controller
	recorder *MockTxQueueManagerMockRecorder
}

// MockTxQueueManagerMockRecorder is the mock recorder for MockTxQueueManager
type MockTxQueueManagerMockRecorder struct {
	mock *MockTxQueueManager
}

// NewMockTxQueueManager creates a new mock instance
func NewMockTxQueueManager(ctrl *gomock.Controller) *MockTxQueueManager {
	mock := &MockTxQueueManager{ctrl: ctrl}
	mock.recorder = &MockTxQueueManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTxQueueManager) EXPECT() *MockTxQueueManagerMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockTxQueueManager) Start() {
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MockTxQueueManagerMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockTxQueueManager)(nil).Start))
}

// Stop mocks base method
func (m *MockTxQueueManager) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockTxQueueManagerMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockTxQueueManager)(nil).Stop))
}

// TransactionQueue mocks base method
func (m *MockTxQueueManager) TransactionQueue() TxQueue {
	ret := m.ctrl.Call(m, "TransactionQueue")
	ret0, _ := ret[0].(TxQueue)
	return ret0
}

// TransactionQueue indicates an expected call of TransactionQueue
func (mr *MockTxQueueManagerMockRecorder) TransactionQueue() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionQueue", reflect.TypeOf((*MockTxQueueManager)(nil).TransactionQueue))
}

// CreateTransaction mocks base method
func (m *MockTxQueueManager) CreateTransaction(ctx context.Context, args SendTxArgs) *QueuedTx {
	ret := m.ctrl.Call(m, "CreateTransaction", ctx, args)
	ret0, _ := ret[0].(*QueuedTx)
	return ret0
}

// CreateTransaction indicates an expected call of CreateTransaction
func (mr *MockTxQueueManagerMockRecorder) CreateTransaction(ctx, args interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockTxQueueManager)(nil).CreateTransaction), ctx, args)
}

// QueueTransaction mocks base method
func (m *MockTxQueueManager) QueueTransaction(tx *QueuedTx) error {
	ret := m.ctrl.Call(m, "QueueTransaction", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueueTransaction indicates an expected call of QueueTransaction
func (mr *MockTxQueueManagerMockRecorder) QueueTransaction(tx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueTransaction", reflect.TypeOf((*MockTxQueueManager)(nil).QueueTransaction), tx)
}

// WaitForTransaction mocks base method
func (m *MockTxQueueManager) WaitForTransaction(tx *QueuedTx) error {
	ret := m.ctrl.Call(m, "WaitForTransaction", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForTransaction indicates an expected call of WaitForTransaction
func (mr *MockTxQueueManagerMockRecorder) WaitForTransaction(tx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForTransaction", reflect.TypeOf((*MockTxQueueManager)(nil).WaitForTransaction), tx)
}

// NotifyOnQueuedTxReturn mocks base method
func (m *MockTxQueueManager) NotifyOnQueuedTxReturn(queuedTx *QueuedTx, err error) {
	m.ctrl.Call(m, "NotifyOnQueuedTxReturn", queuedTx, err)
}

// NotifyOnQueuedTxReturn indicates an expected call of NotifyOnQueuedTxReturn
func (mr *MockTxQueueManagerMockRecorder) NotifyOnQueuedTxReturn(queuedTx, err interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyOnQueuedTxReturn", reflect.TypeOf((*MockTxQueueManager)(nil).NotifyOnQueuedTxReturn), queuedTx, err)
}

// TransactionQueueHandler mocks base method
func (m *MockTxQueueManager) TransactionQueueHandler() func(*QueuedTx) {
	ret := m.ctrl.Call(m, "TransactionQueueHandler")
	ret0, _ := ret[0].(func(*QueuedTx))
	return ret0
}

// TransactionQueueHandler indicates an expected call of TransactionQueueHandler
func (mr *MockTxQueueManagerMockRecorder) TransactionQueueHandler() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionQueueHandler", reflect.TypeOf((*MockTxQueueManager)(nil).TransactionQueueHandler))
}

// SetTransactionQueueHandler mocks base method
func (m *MockTxQueueManager) SetTransactionQueueHandler(fn EnqueuedTxHandler) {
	m.ctrl.Call(m, "SetTransactionQueueHandler", fn)
}

// SetTransactionQueueHandler indicates an expected call of SetTransactionQueueHandler
func (mr *MockTxQueueManagerMockRecorder) SetTransactionQueueHandler(fn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTransactionQueueHandler", reflect.TypeOf((*MockTxQueueManager)(nil).SetTransactionQueueHandler), fn)
}

// SetTransactionReturnHandler mocks base method
func (m *MockTxQueueManager) SetTransactionReturnHandler(fn EnqueuedTxReturnHandler) {
	m.ctrl.Call(m, "SetTransactionReturnHandler", fn)
}

// SetTransactionReturnHandler indicates an expected call of SetTransactionReturnHandler
func (mr *MockTxQueueManagerMockRecorder) SetTransactionReturnHandler(fn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTransactionReturnHandler", reflect.TypeOf((*MockTxQueueManager)(nil).SetTransactionReturnHandler), fn)
}

// SendTransactionRPCHandler mocks base method
func (m *MockTxQueueManager) SendTransactionRPCHandler(ctx context.Context, args ...interface{}) (interface{}, error) {
	varargs := []interface{}{ctx}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendTransactionRPCHandler", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransactionRPCHandler indicates an expected call of SendTransactionRPCHandler
func (mr *MockTxQueueManagerMockRecorder) SendTransactionRPCHandler(ctx interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransactionRPCHandler", reflect.TypeOf((*MockTxQueueManager)(nil).SendTransactionRPCHandler), varargs...)
}

// TransactionReturnHandler mocks base method
func (m *MockTxQueueManager) TransactionReturnHandler() func(*QueuedTx, error) {
	ret := m.ctrl.Call(m, "TransactionReturnHandler")
	ret0, _ := ret[0].(func(*QueuedTx, error))
	return ret0
}

// TransactionReturnHandler indicates an expected call of TransactionReturnHandler
func (mr *MockTxQueueManagerMockRecorder) TransactionReturnHandler() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionReturnHandler", reflect.TypeOf((*MockTxQueueManager)(nil).TransactionReturnHandler))
}

// CompleteTransaction mocks base method
func (m *MockTxQueueManager) CompleteTransaction(id QueuedTxID, password string) (common.Hash, error) {
	ret := m.ctrl.Call(m, "CompleteTransaction", id, password)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompleteTransaction indicates an expected call of CompleteTransaction
func (mr *MockTxQueueManagerMockRecorder) CompleteTransaction(id, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteTransaction", reflect.TypeOf((*MockTxQueueManager)(nil).CompleteTransaction), id, password)
}

// CompleteTransactions mocks base method
func (m *MockTxQueueManager) CompleteTransactions(ids []QueuedTxID, password string) map[QueuedTxID]RawCompleteTransactionResult {
	ret := m.ctrl.Call(m, "CompleteTransactions", ids, password)
	ret0, _ := ret[0].(map[QueuedTxID]RawCompleteTransactionResult)
	return ret0
}

// CompleteTransactions indicates an expected call of CompleteTransactions
func (mr *MockTxQueueManagerMockRecorder) CompleteTransactions(ids, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteTransactions", reflect.TypeOf((*MockTxQueueManager)(nil).CompleteTransactions), ids, password)
}

// DiscardTransaction mocks base method
func (m *MockTxQueueManager) DiscardTransaction(id QueuedTxID) error {
	ret := m.ctrl.Call(m, "DiscardTransaction", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DiscardTransaction indicates an expected call of DiscardTransaction
func (mr *MockTxQueueManagerMockRecorder) DiscardTransaction(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscardTransaction", reflect.TypeOf((*MockTxQueueManager)(nil).DiscardTransaction), id)
}

// DiscardTransactions mocks base method
func (m *MockTxQueueManager) DiscardTransactions(ids []QueuedTxID) map[QueuedTxID]RawDiscardTransactionResult {
	ret := m.ctrl.Call(m, "DiscardTransactions", ids)
	ret0, _ := ret[0].(map[QueuedTxID]RawDiscardTransactionResult)
	return ret0
}

// DiscardTransactions indicates an expected call of DiscardTransactions
func (mr *MockTxQueueManagerMockRecorder) DiscardTransactions(ids interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscardTransactions", reflect.TypeOf((*MockTxQueueManager)(nil).DiscardTransactions), ids)
}

// MockJailCell is a mock of JailCell interface
type MockJailCell struct {
	ctrl     *gomock.Controller
	recorder *MockJailCellMockRecorder
}

// MockJailCellMockRecorder is the mock recorder for MockJailCell
type MockJailCellMockRecorder struct {
	mock *MockJailCell
}

// NewMockJailCell creates a new mock instance
func NewMockJailCell(ctrl *gomock.Controller) *MockJailCell {
	mock := &MockJailCell{ctrl: ctrl}
	mock.recorder = &MockJailCellMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJailCell) EXPECT() *MockJailCellMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockJailCell) Set(arg0 string, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockJailCellMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockJailCell)(nil).Set), arg0, arg1)
}

// Get mocks base method
func (m *MockJailCell) Get(arg0 string) (otto.Value, error) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(otto.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockJailCellMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockJailCell)(nil).Get), arg0)
}

// Run mocks base method
func (m *MockJailCell) Run(arg0 interface{}) (otto.Value, error) {
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(otto.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockJailCellMockRecorder) Run(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockJailCell)(nil).Run), arg0)
}

// Call mocks base method
func (m *MockJailCell) Call(item string, this interface{}, args ...interface{}) (otto.Value, error) {
	varargs := []interface{}{item, this}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Call", varargs...)
	ret0, _ := ret[0].(otto.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call
func (mr *MockJailCellMockRecorder) Call(item, this interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{item, this}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockJailCell)(nil).Call), varargs...)
}

// Stop mocks base method
func (m *MockJailCell) Stop() error {
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockJailCellMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockJailCell)(nil).Stop))
}

// MockJailManager is a mock of JailManager interface
type MockJailManager struct {
	ctrl     *gomock.Controller
	recorder *MockJailManagerMockRecorder
}

// MockJailManagerMockRecorder is the mock recorder for MockJailManager
type MockJailManagerMockRecorder struct {
	mock *MockJailManager
}

// NewMockJailManager creates a new mock instance
func NewMockJailManager(ctrl *gomock.Controller) *MockJailManager {
	mock := &MockJailManager{ctrl: ctrl}
	mock.recorder = &MockJailManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJailManager) EXPECT() *MockJailManagerMockRecorder {
	return m.recorder
}

// Call mocks base method
func (m *MockJailManager) Call(chatID, this, args string) string {
	ret := m.ctrl.Call(m, "Call", chatID, this, args)
	ret0, _ := ret[0].(string)
	return ret0
}

// Call indicates an expected call of Call
func (mr *MockJailManagerMockRecorder) Call(chatID, this, args interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockJailManager)(nil).Call), chatID, this, args)
}

// CreateCell mocks base method
func (m *MockJailManager) CreateCell(chatID string) (JailCell, error) {
	ret := m.ctrl.Call(m, "CreateCell", chatID)
	ret0, _ := ret[0].(JailCell)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCell indicates an expected call of CreateCell
func (mr *MockJailManagerMockRecorder) CreateCell(chatID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCell", reflect.TypeOf((*MockJailManager)(nil).CreateCell), chatID)
}

// Parse mocks base method
func (m *MockJailManager) Parse(chatID, js string) string {
	ret := m.ctrl.Call(m, "Parse", chatID, js)
	ret0, _ := ret[0].(string)
	return ret0
}

// Parse indicates an expected call of Parse
func (mr *MockJailManagerMockRecorder) Parse(chatID, js interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockJailManager)(nil).Parse), chatID, js)
}

// CreateAndInitCell mocks base method
func (m *MockJailManager) CreateAndInitCell(chatID string, code ...string) string {
	varargs := []interface{}{chatID}
	for _, a := range code {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAndInitCell", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// CreateAndInitCell indicates an expected call of CreateAndInitCell
func (mr *MockJailManagerMockRecorder) CreateAndInitCell(chatID interface{}, code ...interface{}) *gomock.Call {
	varargs := append([]interface{}{chatID}, code...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAndInitCell", reflect.TypeOf((*MockJailManager)(nil).CreateAndInitCell), varargs...)
}

// Cell mocks base method
func (m *MockJailManager) Cell(chatID string) (JailCell, error) {
	ret := m.ctrl.Call(m, "Cell", chatID)
	ret0, _ := ret[0].(JailCell)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cell indicates an expected call of Cell
func (mr *MockJailManagerMockRecorder) Cell(chatID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cell", reflect.TypeOf((*MockJailManager)(nil).Cell), chatID)
}

// Execute mocks base method
func (m *MockJailManager) Execute(chatID, code string) string {
	ret := m.ctrl.Call(m, "Execute", chatID, code)
	ret0, _ := ret[0].(string)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockJailManagerMockRecorder) Execute(chatID, code interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockJailManager)(nil).Execute), chatID, code)
}

// SetBaseJS mocks base method
func (m *MockJailManager) SetBaseJS(js string) {
	m.ctrl.Call(m, "SetBaseJS", js)
}

// SetBaseJS indicates an expected call of SetBaseJS
func (mr *MockJailManagerMockRecorder) SetBaseJS(js interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBaseJS", reflect.TypeOf((*MockJailManager)(nil).SetBaseJS), js)
}

// Stop mocks base method
func (m *MockJailManager) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockJailManagerMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockJailManager)(nil).Stop))
}

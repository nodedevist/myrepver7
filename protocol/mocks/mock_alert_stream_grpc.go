// Code generated by MockGen. DO NOT EDIT.
// Source: protocol/alert_stream_grpc.pb.go

// Package mock_protocol is a generated GoMock package.
package mock_protocol

import (
	context "context"
	reflect "reflect"

	protocol "github.com/forta-network/forta-core-go/protocol"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockAlertStreamServiceClient is a mock of AlertStreamServiceClient interface.
type MockAlertStreamServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAlertStreamServiceClientMockRecorder
}

// MockAlertStreamServiceClientMockRecorder is the mock recorder for MockAlertStreamServiceClient.
type MockAlertStreamServiceClientMockRecorder struct {
	mock *MockAlertStreamServiceClient
}

// NewMockAlertStreamServiceClient creates a new mock instance.
func NewMockAlertStreamServiceClient(ctrl *gomock.Controller) *MockAlertStreamServiceClient {
	mock := &MockAlertStreamServiceClient{ctrl: ctrl}
	mock.recorder = &MockAlertStreamServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertStreamServiceClient) EXPECT() *MockAlertStreamServiceClientMockRecorder {
	return m.recorder
}

// Subscribe mocks base method.
func (m *MockAlertStreamServiceClient) Subscribe(ctx context.Context, in *protocol.AlertStreamRequest, opts ...grpc.CallOption) (protocol.AlertStreamService_SubscribeClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(protocol.AlertStreamService_SubscribeClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockAlertStreamServiceClientMockRecorder) Subscribe(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockAlertStreamServiceClient)(nil).Subscribe), varargs...)
}

// MockAlertStreamService_SubscribeClient is a mock of AlertStreamService_SubscribeClient interface.
type MockAlertStreamService_SubscribeClient struct {
	ctrl     *gomock.Controller
	recorder *MockAlertStreamService_SubscribeClientMockRecorder
}

// MockAlertStreamService_SubscribeClientMockRecorder is the mock recorder for MockAlertStreamService_SubscribeClient.
type MockAlertStreamService_SubscribeClientMockRecorder struct {
	mock *MockAlertStreamService_SubscribeClient
}

// NewMockAlertStreamService_SubscribeClient creates a new mock instance.
func NewMockAlertStreamService_SubscribeClient(ctrl *gomock.Controller) *MockAlertStreamService_SubscribeClient {
	mock := &MockAlertStreamService_SubscribeClient{ctrl: ctrl}
	mock.recorder = &MockAlertStreamService_SubscribeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertStreamService_SubscribeClient) EXPECT() *MockAlertStreamService_SubscribeClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockAlertStreamService_SubscribeClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockAlertStreamService_SubscribeClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockAlertStreamService_SubscribeClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockAlertStreamService_SubscribeClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockAlertStreamService_SubscribeClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAlertStreamService_SubscribeClient)(nil).Context))
}

// Header mocks base method.
func (m *MockAlertStreamService_SubscribeClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockAlertStreamService_SubscribeClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockAlertStreamService_SubscribeClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockAlertStreamService_SubscribeClient) Recv() (*protocol.AlertStreamResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*protocol.AlertStreamResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockAlertStreamService_SubscribeClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockAlertStreamService_SubscribeClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockAlertStreamService_SubscribeClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockAlertStreamService_SubscribeClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockAlertStreamService_SubscribeClient)(nil).RecvMsg), m)
}

// SendMsg mocks base method.
func (m_2 *MockAlertStreamService_SubscribeClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockAlertStreamService_SubscribeClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockAlertStreamService_SubscribeClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockAlertStreamService_SubscribeClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockAlertStreamService_SubscribeClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockAlertStreamService_SubscribeClient)(nil).Trailer))
}

// MockAlertStreamServiceServer is a mock of AlertStreamServiceServer interface.
type MockAlertStreamServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAlertStreamServiceServerMockRecorder
}

// MockAlertStreamServiceServerMockRecorder is the mock recorder for MockAlertStreamServiceServer.
type MockAlertStreamServiceServerMockRecorder struct {
	mock *MockAlertStreamServiceServer
}

// NewMockAlertStreamServiceServer creates a new mock instance.
func NewMockAlertStreamServiceServer(ctrl *gomock.Controller) *MockAlertStreamServiceServer {
	mock := &MockAlertStreamServiceServer{ctrl: ctrl}
	mock.recorder = &MockAlertStreamServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertStreamServiceServer) EXPECT() *MockAlertStreamServiceServerMockRecorder {
	return m.recorder
}

// Subscribe mocks base method.
func (m *MockAlertStreamServiceServer) Subscribe(arg0 *protocol.AlertStreamRequest, arg1 protocol.AlertStreamService_SubscribeServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockAlertStreamServiceServerMockRecorder) Subscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockAlertStreamServiceServer)(nil).Subscribe), arg0, arg1)
}

// mustEmbedUnimplementedAlertStreamServiceServer mocks base method.
func (m *MockAlertStreamServiceServer) mustEmbedUnimplementedAlertStreamServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAlertStreamServiceServer")
}

// mustEmbedUnimplementedAlertStreamServiceServer indicates an expected call of mustEmbedUnimplementedAlertStreamServiceServer.
func (mr *MockAlertStreamServiceServerMockRecorder) mustEmbedUnimplementedAlertStreamServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAlertStreamServiceServer", reflect.TypeOf((*MockAlertStreamServiceServer)(nil).mustEmbedUnimplementedAlertStreamServiceServer))
}

// MockUnsafeAlertStreamServiceServer is a mock of UnsafeAlertStreamServiceServer interface.
type MockUnsafeAlertStreamServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeAlertStreamServiceServerMockRecorder
}

// MockUnsafeAlertStreamServiceServerMockRecorder is the mock recorder for MockUnsafeAlertStreamServiceServer.
type MockUnsafeAlertStreamServiceServerMockRecorder struct {
	mock *MockUnsafeAlertStreamServiceServer
}

// NewMockUnsafeAlertStreamServiceServer creates a new mock instance.
func NewMockUnsafeAlertStreamServiceServer(ctrl *gomock.Controller) *MockUnsafeAlertStreamServiceServer {
	mock := &MockUnsafeAlertStreamServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeAlertStreamServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeAlertStreamServiceServer) EXPECT() *MockUnsafeAlertStreamServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedAlertStreamServiceServer mocks base method.
func (m *MockUnsafeAlertStreamServiceServer) mustEmbedUnimplementedAlertStreamServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAlertStreamServiceServer")
}

// mustEmbedUnimplementedAlertStreamServiceServer indicates an expected call of mustEmbedUnimplementedAlertStreamServiceServer.
func (mr *MockUnsafeAlertStreamServiceServerMockRecorder) mustEmbedUnimplementedAlertStreamServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAlertStreamServiceServer", reflect.TypeOf((*MockUnsafeAlertStreamServiceServer)(nil).mustEmbedUnimplementedAlertStreamServiceServer))
}

// MockAlertStreamService_SubscribeServer is a mock of AlertStreamService_SubscribeServer interface.
type MockAlertStreamService_SubscribeServer struct {
	ctrl     *gomock.Controller
	recorder *MockAlertStreamService_SubscribeServerMockRecorder
}

// MockAlertStreamService_SubscribeServerMockRecorder is the mock recorder for MockAlertStreamService_SubscribeServer.
type MockAlertStreamService_SubscribeServerMockRecorder struct {
	mock *MockAlertStreamService_SubscribeServer
}

// NewMockAlertStreamService_SubscribeServer creates a new mock instance.
func NewMockAlertStreamService_SubscribeServer(ctrl *gomock.Controller) *MockAlertStreamService_SubscribeServer {
	mock := &MockAlertStreamService_SubscribeServer{ctrl: ctrl}
	mock.recorder = &MockAlertStreamService_SubscribeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertStreamService_SubscribeServer) EXPECT() *MockAlertStreamService_SubscribeServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockAlertStreamService_SubscribeServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockAlertStreamService_SubscribeServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAlertStreamService_SubscribeServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockAlertStreamService_SubscribeServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockAlertStreamService_SubscribeServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockAlertStreamService_SubscribeServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockAlertStreamService_SubscribeServer) Send(arg0 *protocol.AlertStreamResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockAlertStreamService_SubscribeServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAlertStreamService_SubscribeServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockAlertStreamService_SubscribeServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockAlertStreamService_SubscribeServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockAlertStreamService_SubscribeServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockAlertStreamService_SubscribeServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockAlertStreamService_SubscribeServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockAlertStreamService_SubscribeServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockAlertStreamService_SubscribeServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockAlertStreamService_SubscribeServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockAlertStreamService_SubscribeServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockAlertStreamService_SubscribeServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockAlertStreamService_SubscribeServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockAlertStreamService_SubscribeServer)(nil).SetTrailer), arg0)
}
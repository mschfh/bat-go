// Code generated by MockGen. DO NOT EDIT.
// Source: ./utils/clients/coingecko_assets/client.go

// Package mock_coingecko_assets is a generated GoMock package.
package mock_coingecko_assets

import (
	context "context"
	reflect "reflect"
	time "time"

	coingeckoAssets "github.com/brave-intl/bat-go/utils/clients/coingecko_assets"
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

// FetchImageAsset mocks base method.
func (m *MockClient) FetchImageAsset(ctx context.Context, imageID, size, imageFile string) (*coingeckoAssets.ImageAssetResponseBundle, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchImageAsset", ctx, imageID, size, imageFile)
	ret0, _ := ret[0].(*coingeckoAssets.ImageAssetResponseBundle)
	ret1, _ := ret[1].(time.Time)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FetchImageAsset indicates an expected call of FetchImageAsset.
func (mr *MockClientMockRecorder) FetchImageAsset(ctx, imageID, size, imageFile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchImageAsset", reflect.TypeOf((*MockClient)(nil).FetchImageAsset), ctx, imageID, size, imageFile)
}

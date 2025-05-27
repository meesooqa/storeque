package db_provider

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"tg-star-shop-bot-001/common/config"
	"tg-star-shop-bot-001/db/db_types"
)

// Mock implementations
type MockConfigProvider struct {
	mock.Mock
}

func (m *MockConfigProvider) GetAppConfig() (*config.AppConfig, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*config.AppConfig), args.Error(1)
}

type MockGormOpener struct {
	mock.Mock
}

func (m *MockGormOpener) Open(dsn string, config *gorm.Config) (*gorm.DB, error) {
	args := m.Called(dsn, config)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*gorm.DB), args.Error(1)
}

type testGormOpener struct{}

func (o *testGormOpener) Open(dsn string, config *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(":memory:"), config)
}

func TestNewDefaultDBProvider(t *testing.T) {
	// Test that the function returns a non-nil provider
	provider := NewDefaultDBProvider()
	assert.NotNil(t, provider)
	assert.NotNil(t, provider.configProvider)
	assert.NotNil(t, provider.gormOpener)
}

func TestNewDefaultDBProviderWithCustomOpener(t *testing.T) {
	// Test with custom implementations
	mockConfigProvider := new(MockConfigProvider)
	mockGormOpener := new(MockGormOpener)

	provider := NewDefaultDBProviderWithCustomOpener(mockConfigProvider, mockGormOpener)

	assert.NotNil(t, provider)
	assert.Equal(t, mockConfigProvider, provider.configProvider)
	assert.Equal(t, mockGormOpener, provider.gormOpener)
}

func TestConstructDSN(t *testing.T) {
	provider := &DefaultDBProvider{}

	testCases := []struct {
		name     string
		config   *config.AppConfig
		expected string
	}{
		{
			name: "standard configuration",
			config: &config.AppConfig{
				DB: &config.DbConfig{
					Host:     "localhost",
					Port:     5432,
					SslMode:  "disable",
					User:     "postgres",
					Password: "secret",
					DbName:   "testdb",
				},
			},
			expected: "host=localhost port=5432 sslmode=disable user=postgres password=secret dbname=testdb",
		},
		{
			name: "empty password",
			config: &config.AppConfig{
				DB: &config.DbConfig{
					Host:     "db.example.com",
					Port:     5432,
					SslMode:  "require",
					User:     "appuser",
					Password: "",
					DbName:   "production",
				},
			},
			expected: "host=db.example.com port=5432 sslmode=require user=appuser password= dbname=production",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := provider.constructDSN(tc.config)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestGetDB_Success(t *testing.T) {
	// Setup mocks
	mockConfigProvider := new(MockConfigProvider)
	ctx := context.Background()
	// Normal case (debug mode off)
	testConf := &config.AppConfig{
		DB: &config.DbConfig{
			Host:        "testhost",
			Port:        5432,
			SslMode:     "disable",
			User:        "testuser",
			Password:    "testpass",
			DbName:      "testdb",
			IsDebugMode: false,
		},
	}
	mockConfigProvider.On("GetAppConfig").Return(testConf, nil)
	provider := NewDefaultDBProviderWithCustomOpener(mockConfigProvider, &testGormOpener{})
	db, err := provider.GetDB(ctx)
	require.NoError(t, err)
	assert.NotNil(t, db)
	mockConfigProvider.AssertExpectations(t)
}

func TestGetDB_DebugMode(t *testing.T) {
	// Setup mocks
	mockConfigProvider := new(MockConfigProvider)
	ctx := context.Background()
	// Debug mode on case
	testConf := &config.AppConfig{
		DB: &config.DbConfig{
			Host:        "testhost",
			Port:        5432,
			SslMode:     "disable",
			User:        "testuser",
			Password:    "testpass",
			DbName:      "testdb",
			IsDebugMode: true,
		},
	}
	mockConfigProvider.On("GetAppConfig").Return(testConf, nil)

	provider := NewDefaultDBProviderWithCustomOpener(mockConfigProvider, &testGormOpener{})

	db, err := provider.GetDB(ctx)

	require.NoError(t, err)
	assert.NotNil(t, db)
	mockConfigProvider.AssertExpectations(t)
}

func TestGetDB_ConfigError(t *testing.T) {
	// Setup mocks
	mockConfigProvider := new(MockConfigProvider)
	mockGormOpener := new(MockGormOpener)
	ctx := context.Background()

	// Config error case
	configError := errors.New("config error")
	mockConfigProvider.On("GetAppConfig").Return(nil, configError)

	provider := NewDefaultDBProviderWithCustomOpener(mockConfigProvider, mockGormOpener)

	db, err := provider.GetDB(ctx)

	assert.Nil(t, db)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to load config")
	mockConfigProvider.AssertExpectations(t)
	mockGormOpener.AssertNotCalled(t, "Open")
}

func TestGetDB_DBConnectionError(t *testing.T) {
	// Setup mocks
	mockConfigProvider := new(MockConfigProvider)
	mockGormOpener := new(MockGormOpener)
	ctx := context.Background()

	// DB connection error case
	testConf := &config.AppConfig{
		DB: &config.DbConfig{
			Host:     "testhost",
			Port:     5432,
			SslMode:  "disable",
			User:     "testuser",
			Password: "testpass",
			DbName:   "testdb",
		},
	}

	expectedDSN := "host=testhost port=5432 sslmode=disable user=testuser password=testpass dbname=testdb"
	dbError := errors.New("connection error")

	mockConfigProvider.On("GetAppConfig").Return(testConf, nil)
	mockGormOpener.On("Open", expectedDSN, mock.AnythingOfType("*gorm.Config")).Return(nil, dbError)

	provider := NewDefaultDBProviderWithCustomOpener(mockConfigProvider, mockGormOpener)

	db, err := provider.GetDB(ctx)

	assert.Nil(t, db)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to connect database")
	mockConfigProvider.AssertExpectations(t)
	mockGormOpener.AssertExpectations(t)
}

func TestPostgresGormOpenerImplementation(t *testing.T) {
	// Just verify that the struct implements the interface
	var opener db_types.GormOpener = &postgresGormOpener{}
	assert.NotNil(t, opener)
}

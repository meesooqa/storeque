package db_types

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// MockGormOpener is a mock implementation of GormOpener interface
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

// MockDBProvider is a mock implementation of DBProvider interface
type MockDBProvider struct {
	mock.Mock
}

func (m *MockDBProvider) GetDB(ctx context.Context) (*gorm.DB, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*gorm.DB), args.Error(1)
}

// MockRepository is a mock implementation of Repository interface
type MockRepository[T any] struct {
	mock.Mock
}

func (m *MockRepository[T]) GetList(ctx context.Context, filters []FilterFunc, sort SortData, pagination PaginationData) ([]*T, int64, error) {
	args := m.Called(ctx, filters, sort, pagination)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*T), args.Get(1).(int64), args.Error(2)
}

func (m *MockRepository[T]) Get(ctx context.Context, id uint64) (*T, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*T), args.Error(1)
}

func (m *MockRepository[T]) Create(ctx context.Context, newItem *T) (*T, error) {
	args := m.Called(ctx, newItem)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*T), args.Error(1)
}

func (m *MockRepository[T]) Update(ctx context.Context, id uint64, updatedItem *T) (*T, error) {
	args := m.Called(ctx, id, updatedItem)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*T), args.Error(1)
}

func (m *MockRepository[T]) Delete(ctx context.Context, id uint64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockHasAssociations is a mock implementation of HasAssociations interface
type MockHasAssociations[T any] struct {
	mock.Mock
}

func (m *MockHasAssociations[T]) UpdateAssociations(db *gorm.DB, item *T, updatedData *T) error {
	args := m.Called(db, item, updatedData)
	return args.Error(0)
}

// TestGormOpener tests the GormOpener interface
func TestGormOpener(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockOpener := new(MockGormOpener)
		expectedDB := &gorm.DB{}
		mockOpener.On("Open", "test-dsn", mock.AnythingOfType("*gorm.Config")).Return(expectedDB, nil)

		db, err := mockOpener.Open("test-dsn", &gorm.Config{})

		require.NoError(t, err)
		assert.Equal(t, expectedDB, db)
		mockOpener.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockOpener := new(MockGormOpener)
		expectedError := errors.New("connection error")
		mockOpener.On("Open", "invalid-dsn", mock.AnythingOfType("*gorm.Config")).Return(nil, expectedError)

		db, err := mockOpener.Open("invalid-dsn", &gorm.Config{})

		assert.Nil(t, db)
		assert.Equal(t, expectedError, err)
		mockOpener.AssertExpectations(t)
	})
}

// TestDBProvider tests the DBProvider interface
func TestDBProvider(t *testing.T) {
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		mockProvider := new(MockDBProvider)
		expectedDB := &gorm.DB{}
		mockProvider.On("GetDB", ctx).Return(expectedDB, nil)

		db, err := mockProvider.GetDB(ctx)

		require.NoError(t, err)
		assert.Equal(t, expectedDB, db)
		mockProvider.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockProvider := new(MockDBProvider)
		expectedError := errors.New("database error")
		mockProvider.On("GetDB", ctx).Return(nil, expectedError)

		db, err := mockProvider.GetDB(ctx)

		assert.Nil(t, db)
		assert.Equal(t, expectedError, err)
		mockProvider.AssertExpectations(t)
	})
}

// Define a test model
type TestModel struct {
	ID   uint64
	Name string
}

// TestFilterFunc tests the FilterFunc type
func TestFilterFunc(t *testing.T) {
	// Create a filter function to test
	var called bool
	filter := FilterFunc(func(db *gorm.DB) *gorm.DB {
		called = true
		return db // just return the input for testing
	})

	// Test that filter can be called
	db := &gorm.DB{}
	result := filter(db)

	assert.True(t, called, "Filter function should be called")
	assert.Equal(t, db, result, "Filter should return the db that was passed in")
}

// TestSortData tests the SortData struct
func TestSortData(t *testing.T) {
	sortData := SortData{
		SortField: "name",
		SortOrder: "asc",
	}

	assert.Equal(t, "name", sortData.SortField)
	assert.Equal(t, "asc", sortData.SortOrder)
}

// TestPaginationData tests the PaginationData struct
func TestPaginationData(t *testing.T) {
	paginationData := PaginationData{
		Page:     2,
		PageSize: 10,
	}

	assert.Equal(t, 2, paginationData.Page)
	assert.Equal(t, 10, paginationData.PageSize)
}

// TestRepository tests the Repository interface methods
func TestRepository(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository[TestModel])

	t.Run("GetList", func(t *testing.T) {
		filters := []FilterFunc{
			func(db *gorm.DB) *gorm.DB { return db.Where("name = ?", "test") },
		}
		sort := SortData{SortField: "name", SortOrder: "asc"}
		pagination := PaginationData{Page: 1, PageSize: 10}

		expectedModels := []*TestModel{{ID: 1, Name: "test"}}
		var expectedTotal int64 = 1

		mockRepo.On("GetList", ctx, filters, sort, pagination).Return(expectedModels, expectedTotal, nil)

		models, total, err := mockRepo.GetList(ctx, filters, sort, pagination)

		require.NoError(t, err)
		assert.Equal(t, expectedModels, models)
		assert.Equal(t, expectedTotal, total)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Get", func(t *testing.T) {
		expectedModel := &TestModel{ID: 1, Name: "test"}
		mockRepo.On("Get", ctx, uint64(1)).Return(expectedModel, nil)

		model, err := mockRepo.Get(ctx, 1)

		require.NoError(t, err)
		assert.Equal(t, expectedModel, model)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Create", func(t *testing.T) {
		newModel := &TestModel{Name: "test"}
		createdModel := &TestModel{ID: 1, Name: "test"}

		mockRepo.On("Create", ctx, newModel).Return(createdModel, nil)

		model, err := mockRepo.Create(ctx, newModel)

		require.NoError(t, err)
		assert.Equal(t, createdModel, model)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Update", func(t *testing.T) {
		updatedModel := &TestModel{Name: "updated"}
		resultModel := &TestModel{ID: 1, Name: "updated"}

		mockRepo.On("Update", ctx, uint64(1), updatedModel).Return(resultModel, nil)

		model, err := mockRepo.Update(ctx, 1, updatedModel)

		require.NoError(t, err)
		assert.Equal(t, resultModel, model)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Delete", func(t *testing.T) {
		mockRepo.On("Delete", ctx, uint64(1)).Return(nil)

		err := mockRepo.Delete(ctx, 1)

		require.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

// TestHasAssociations tests the HasAssociations interface
func TestHasAssociations(t *testing.T) {
	db := &gorm.DB{}
	mockAssoc := new(MockHasAssociations[TestModel])

	item := &TestModel{ID: 1, Name: "test"}
	updatedData := &TestModel{ID: 1, Name: "updated"}

	mockAssoc.On("UpdateAssociations", db, item, updatedData).Return(nil)

	err := mockAssoc.UpdateAssociations(db, item, updatedData)

	require.NoError(t, err)
	mockAssoc.AssertExpectations(t)
}

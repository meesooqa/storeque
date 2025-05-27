package db_types

import (
	"time"

	"gorm.io/gorm"
)

// Numeric is Constraint для числовых типов, поддерживаемых фильтрами диапазона
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// FieldFilter creates a filter that applies a SQL LIKE condition on a specific field
func FieldFilter(fieldName, value string) FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			// return db.Where(fieldName+" ILIKE ?", "%"+value+"%")
			return db.Where("LOWER("+fieldName+") LIKE LOWER(?)", "%"+value+"%")
		}
		return db
	}
}

// ExactFieldFilter creates a filter that applies an exact match condition on a field
func ExactFieldFilter(fieldName string, value interface{}) FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		// TODO "value != false"?
		// TODO value != 0
		if value != nil && value != "" && value != false && value != uint64(0) {
			return db.Where(fieldName+" = ?", value)
		}
		return db
	}
}

// ModelFieldFilter creates a type-safe field filter for a specific model
func ModelFieldFilter[DbModel any](field string) func(value string) FilterFunc {
	return func(value string) FilterFunc {
		return FieldFilter(field, value)
	}
}

// ModelExactFieldFilter creates a type-safe exact field filter for a specific model
func ModelExactFieldFilter[DbModel any](field string) func(value interface{}) FilterFunc {
	return func(value interface{}) FilterFunc {
		return ExactFieldFilter(field, value)
	}
}

// ModelRangeFilter создает типобезопасный фильтр по диапазону для конкретной модели и типа числовых данных
func ModelRangeFilter[T any, V Numeric](fieldName string) func(valueGt, valueLt V) FilterFunc {
	return func(valueGt, valueLt V) FilterFunc {
		return func(db *gorm.DB) *gorm.DB {
			// Нулевое значение для типа V
			var zeroValue V

			if valueGt > zeroValue && valueLt > zeroValue {
				db = db.Where(fieldName+" BETWEEN ? AND ?", valueGt, valueLt)
			} else if valueGt > zeroValue {
				db = db.Where(fieldName+" >= ?", valueGt)
			} else if valueLt > zeroValue {
				db = db.Where(fieldName+" <= ?", valueLt)
			}
			return db
		}
	}
}

// Фильтр для работы с массивами (например, для поля tags)
//func ModelArrayFilter[T any](fieldName string) func(values []string) FilterFunc {
//	return func(values []string) FilterFunc {
//		return func(db *gorm.DB) *gorm.DB {
//			if len(values) > 0 {
//				return db.Where(fieldName+" && ?", pq.Array(values))
//			}
//			return db
//		}
//	}
//}

// Фильтр для JSON полей
//func ModelJsonFilter[T any](jsonPath string) func(value interface{}) FilterFunc {
//	return func(value interface{}) FilterFunc {
//		return func(db *gorm.DB) *gorm.DB {
//			if value != nil {
//				return db.Where("jsonb_path_exists(data, ?)", jsonPath)
//			}
//			return db
//		}
//	}
//}

// ModelJsonbContainsFilter для фильтрации по содержимому JSONB поля
//func ModelJsonbContainsFilter[T any](fieldName string) func(jsonValue map[string]interface{}) FilterFunc {
//	return func(jsonValue map[string]interface{}) FilterFunc {
//		return func(db *gorm.DB) *gorm.DB {
//			if len(jsonValue) > 0 {
//				jsonBytes, _ := json.Marshal(jsonValue)
//				return db.Where(fieldName+" @> ?", string(jsonBytes))
//			}
//			return db
//		}
//	}
//}

// ModelDateRangeFilter для фильтрации по диапазону дат
func ModelDateRangeFilter[T any](fieldName string) func(startDate, endDate time.Time) FilterFunc {
	return func(startDate, endDate time.Time) FilterFunc {
		return func(db *gorm.DB) *gorm.DB {
			zeroTime := time.Time{}

			if startDate != zeroTime && endDate != zeroTime {
				return db.Where(fieldName+" BETWEEN ? AND ?", startDate, endDate)
			} else if startDate != zeroTime {
				return db.Where(fieldName+" >= ?", startDate)
			} else if endDate != zeroTime {
				return db.Where(fieldName+" <= ?", endDate)
			}
			return db
		}
	}
}

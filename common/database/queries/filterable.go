package queries

import (
	"errors"
	"gorm.io/gorm"
	"strings"
)

const (
	invalidSuffixType byte = iota
	singleValueType
	arrayValueType
)

var ErrInvalidFilterField = errors.New("invalid filter field")

type Filterable interface {
	FilterMap() map[string]string
}

var suffixHandleRegistered = map[string]byte{
	"eq":         singleValueType,
	"neq":        singleValueType,
	"lt":         singleValueType,
	"gt":         singleValueType,
	"lte":        singleValueType,
	"gte":        singleValueType,
	"contains":   singleValueType,
	"ncontains":  singleValueType,
	"containss":  singleValueType,
	"ncontainss": singleValueType,
	"isnull":     singleValueType,
	"btw":        arrayValueType,
	"in":         arrayValueType,
	"nin":        arrayValueType,
}
var arrayValueSuffixHandlers = map[string]func(query *gorm.DB, colName string, value []string) *gorm.DB{
	"btw": func(query *gorm.DB, colName string, value []string) *gorm.DB {
		return query.Where(colName+" BETWEEN ? AND ?", value[0], value[1])
	},
	"in": func(query *gorm.DB, colName string, value []string) *gorm.DB {
		return query.Where(colName+" IN ?", value)
	},
	"nin": func(query *gorm.DB, colName string, value []string) *gorm.DB {
		return query.Where(colName+" NOT IN ?", value)
	},
}
var singleValueSuffixHandlers = map[string]func(query *gorm.DB, colName string, value string) *gorm.DB{
	"eq": func(query *gorm.DB, colName string, value string) *gorm.DB {
		return query.Where(colName+" = ?", value)
	},
	"neq": func(query *gorm.DB, colName string, value string) *gorm.DB {
		return query.Where(colName+" <> ?", value)
	},
	"lt": func(query *gorm.DB, colName string, value string) *gorm.DB {
		return query.Where(colName+" < ?", value)
	},
	"gt": func(query *gorm.DB, colName string, value string) *gorm.DB {
		return query.Where(colName+" > ?", value)
	},
	"lte": func(query *gorm.DB, colName string, value string) *gorm.DB {
		return query.Where(colName+" <= ?", value)
	},
	"gte": func(query *gorm.DB, colName string, value string) *gorm.DB {
		return query.Where(colName+" >= ?", value)
	},

	"contains": func(query *gorm.DB, colName string, value string) *gorm.DB {
		return query.Where(colName+" LIKE ?", "%"+value+"%")
	},
	"ncontains": func(query *gorm.DB, colName string, value string) *gorm.DB {
		return query.Where(colName+" NOT LIKE ?", "%"+value+"%")
	},
	"containss": func(query *gorm.DB, colName string, value string) *gorm.DB {
		return query.Where(colName+" ILIKE ?", "%"+value+"%")
	},
	"ncontainss": func(query *gorm.DB, colName string, value string) *gorm.DB {
		return query.Where(colName+" NOT ILIKE ?", "%"+value+"%")
	},
	"isnull": func(query *gorm.DB, colName string, value string) *gorm.DB {
		if value == "true" {
			return query.Where(colName + " IS NULL")
		} else {
			return query.Where(colName + " IS NOT NULL")
		}
	},
}

func Filter(query *gorm.DB, filter Filterable, filterFields map[string]string) *gorm.DB {
	filterMap := filter.FilterMap()
	result := query
	for filterField, value := range filterFields {
		fieldName, suffix, err := getFilterFiledSuffix(filterField)
		if err != nil {
			continue
		}
		colName := filterMap[fieldName]
		if colName == "" {
			continue
		}
		suffixType := suffixHandleRegistered[suffix]
		if suffixType == arrayValueType {
			handler := arrayValueSuffixHandlers[suffix]
			result = handler(result, colName, strings.Split(value, ","))
		} else if suffixType == singleValueType {
			handler := singleValueSuffixHandlers[suffix]
			result = handler(result, colName, value)
		} else {
			if colName == "" {
				continue
			}
		}

	}
	return result
}

func getFilterFiledSuffix(filterField string) (fieldName string, suffix string, err error) {
	idx := strings.LastIndex(filterField, "_")
	if idx <= 0 {
		return "", "", ErrInvalidFilterField
	}
	return filterField[:idx], filterField[idx+1:], nil
}

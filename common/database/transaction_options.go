package database

import "gorm.io/gorm"

type transactionOption struct {
	filters    []func(db *gorm.DB)*gorm.DB
	preExecute  func(db *gorm.DB,models interface{})error
	postExecute func(db *gorm.DB,models interface{})error
}

const (
	saveTx = "save"
	createTx = "create"
	updateTx = "update"
	deleteTx = "delete"
)

type TransactionOption func(option *transactionOption)

func WithFilters(filters...func(db *gorm.DB)*gorm.DB) TransactionOption {
	return func(option *transactionOption) {
		option.filters = append(option.filters,filters...)
	}
}

func WithPreExecute(trigger func(db *gorm.DB,models interface{})error) TransactionOption {
	return func(option *transactionOption) {
		option.preExecute = trigger
	}
}

func WithPostExecute(trigger func(db *gorm.DB,models interface{})error) TransactionOption {
	return func(option *transactionOption) {
		option.preExecute = trigger
	}
}

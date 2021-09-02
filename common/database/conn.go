package database

import (
	"context"
	"gorm.io/gorm"
)

type Conn interface {
	context.Context
	ORM() *gorm.DB
	List(models interface{}, filters ...func(db *gorm.DB) *gorm.DB) error
	Retrieve(model interface{}, pk interface{}, filters ...func(db *gorm.DB) *gorm.DB) error
	Create(model interface{}, opts ...TransactionOption) error
	Update(model interface{}, opts ...TransactionOption) error
	Delete(model interface{}, opts ...TransactionOption) error
	Save(model interface{}, opts ...TransactionOption) error
}

type databaseConn struct {
	context.Context
	orm *gorm.DB
}

func (c databaseConn) ORM() *gorm.DB {
	return c.orm
}

func (c *databaseConn) List(models interface{}, filters ...func(db *gorm.DB) *gorm.DB) error {
	query :=  c.orm.WithContext(c.Context).Scopes(filters...).Find(models)

	return query.Error
}

func (c *databaseConn) Retrieve(model interface{}, pk interface{}, filters ...func(db *gorm.DB) *gorm.DB) error {
	return c.orm.WithContext(c.Context).Scopes(filters...).First(model, pk).Error
}

func (c *databaseConn) Create(model interface{}, opts ...TransactionOption) error {
	return c.transaction(createTx, model, opts)
}

func (c *databaseConn) Update(model interface{}, opts ...TransactionOption) error {
	return c.transaction(updateTx, model, opts)
}

func (c *databaseConn) Delete(model interface{}, opts ...TransactionOption) error {
	return c.transaction(deleteTx, model, opts)
}

func (c *databaseConn) Save(model interface{}, opts ...TransactionOption) error {
	return c.transaction(saveTx, model, opts)
}

func (c *databaseConn) transaction(transactionType string, model interface{}, opts []TransactionOption) error {
	var option transactionOption
	for _, opt := range opts {
		opt(&option)
	}
	err := c.orm.Transaction(func(tx *gorm.DB) error {
		err := option.preExecute(tx, model)
		if err != nil {
			return err
		}
		switch transactionType {
		case createTx:
			err = c.orm.WithContext(c.Context).Scopes(option.filters...).Create(model).Error
		case updateTx:
			err = c.orm.WithContext(c.Context).Scopes(option.filters...).Updates(model).Error
		case deleteTx:
			err = c.orm.WithContext(c.Context).Scopes(option.filters...).Delete(model).Error
		case saveTx:
			err = c.orm.WithContext(c.Context).Scopes(option.filters...).Save(model).Error
		}
		if err != nil {
			return err
		}
		return option.postExecute(tx, model)
	})
	return err
}

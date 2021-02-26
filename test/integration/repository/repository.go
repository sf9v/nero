// Code generated by nero, DO NOT EDIT.
package repository

import (
	"context"
	"reflect"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
	"github.com/sf9v/nero"
	"github.com/sf9v/nero/example"
	"github.com/sf9v/nero/test/integration/user"
)

// Repository is an interface for interacting with a User repository
type Repository interface {
	// Tx begins a new transaction
	Tx(context.Context) (nero.Tx, error)
	// Create runs a create
	Create(context.Context, *Creator) (id string, err error)
	// CreateTx runs a create in a transaction
	CreateTx(context.Context, nero.Tx, *Creator) (id string, err error)
	// CreateMany runs a batch create
	CreateMany(context.Context, ...*Creator) error
	// CreateManyTx runs a batch create in a transaction
	CreateManyTx(context.Context, nero.Tx, ...*Creator) error
	// Query runs a query
	Query(context.Context, *Queryer) ([]*user.User, error)
	// QueryTx runs a query in a transaction
	QueryTx(context.Context, nero.Tx, *Queryer) ([]*user.User, error)
	// QueryOne runs a query that expects only one result
	QueryOne(context.Context, *Queryer) (*user.User, error)
	// QueryOneTx runs a query that expects only one result in a transaction
	QueryOneTx(context.Context, nero.Tx, *Queryer) (*user.User, error)
	// Update runs an update
	Update(context.Context, *Updater) (rowsAffected int64, err error)
	// UpdateTx runs an update in a transaction
	UpdateTx(context.Context, nero.Tx, *Updater) (rowsAffected int64, err error)
	// Delete runs a delete
	Delete(context.Context, *Deleter) (rowsAffected int64, err error)
	// Delete runs a delete in a transaction
	DeleteTx(context.Context, nero.Tx, *Deleter) (rowsAffected int64, err error)
	// Aggregate runs aggregate query
	Aggregate(context.Context, *Aggregator) error
	// Aggregate runs aggregate query in a transaction
	AggregateTx(context.Context, nero.Tx, *Aggregator) error
}

// Creator is a create builder
type Creator struct {
	uid       ksuid.KSUID
	email     string
	name      string
	age       int
	group     user.Group
	kv        example.Map
	tags      []string
	updatedAt *time.Time
}

// NewCreator returns a Creator
func NewCreator() *Creator {
	return &Creator{}
}

// UID sets the UID field
func (c *Creator) UID(uid ksuid.KSUID) *Creator {
	c.uid = uid
	return c
}

// Email sets the Email field
func (c *Creator) Email(email string) *Creator {
	c.email = email
	return c
}

// Name sets the Name field
func (c *Creator) Name(name string) *Creator {
	c.name = name
	return c
}

// Age sets the Age field
func (c *Creator) Age(age int) *Creator {
	c.age = age
	return c
}

// Group sets the Group field
func (c *Creator) Group(group user.Group) *Creator {
	c.group = group
	return c
}

// Kv sets the Kv field
func (c *Creator) Kv(kv example.Map) *Creator {
	c.kv = kv
	return c
}

// Tags sets the Tags field
func (c *Creator) Tags(tags []string) *Creator {
	c.tags = tags
	return c
}

// UpdatedAt sets the UpdatedAt field
func (c *Creator) UpdatedAt(updatedAt *time.Time) *Creator {
	c.updatedAt = updatedAt
	return c
}

// Validate validates the fields
func (c *Creator) Validate() error {
	var err error
	if isZero(c.uid) {
		err = multierror.Append(err, nero.NewErrRequiredField("uid"))
	}

	if isZero(c.email) {
		err = multierror.Append(err, nero.NewErrRequiredField("email"))
	}

	if isZero(c.name) {
		err = multierror.Append(err, nero.NewErrRequiredField("name"))
	}

	if isZero(c.age) {
		err = multierror.Append(err, nero.NewErrRequiredField("age"))
	}

	if isZero(c.group) {
		err = multierror.Append(err, nero.NewErrRequiredField("group"))
	}

	if isZero(c.kv) {
		err = multierror.Append(err, nero.NewErrRequiredField("kv"))
	}

	if isZero(c.tags) {
		err = multierror.Append(err, nero.NewErrRequiredField("tags"))
	}

	return err
}

// Queryer is a query builder
type Queryer struct {
	limit  uint
	offset uint
	pfs    []PredFunc
	sfs    []SortFunc
}

// NewQueryer returns a Queryer
func NewQueryer() *Queryer {
	return &Queryer{}
}

// Where applies predicates
func (q *Queryer) Where(pfs ...PredFunc) *Queryer {
	q.pfs = append(q.pfs, pfs...)
	return q
}

// Sort applies sorting expressions
func (q *Queryer) Sort(sfs ...SortFunc) *Queryer {
	q.sfs = append(q.sfs, sfs...)
	return q
}

// Limit applies limit
func (q *Queryer) Limit(limit uint) *Queryer {
	q.limit = limit
	return q
}

// Offset applies offset
func (q *Queryer) Offset(offset uint) *Queryer {
	q.offset = offset
	return q
}

// Updater is an update builder
type Updater struct {
	uid       ksuid.KSUID
	email     string
	name      string
	age       int
	group     user.Group
	kv        example.Map
	tags      []string
	updatedAt *time.Time
	pfs       []PredFunc
}

// NewUpdater returns an Updater
func NewUpdater() *Updater {
	return &Updater{}
}

// UID sets the UID field
func (c *Updater) UID(uid ksuid.KSUID) *Updater {
	c.uid = uid
	return c
}

// Email sets the Email field
func (c *Updater) Email(email string) *Updater {
	c.email = email
	return c
}

// Name sets the Name field
func (c *Updater) Name(name string) *Updater {
	c.name = name
	return c
}

// Age sets the Age field
func (c *Updater) Age(age int) *Updater {
	c.age = age
	return c
}

// Group sets the Group field
func (c *Updater) Group(group user.Group) *Updater {
	c.group = group
	return c
}

// Kv sets the Kv field
func (c *Updater) Kv(kv example.Map) *Updater {
	c.kv = kv
	return c
}

// Tags sets the Tags field
func (c *Updater) Tags(tags []string) *Updater {
	c.tags = tags
	return c
}

// UpdatedAt sets the UpdatedAt field
func (c *Updater) UpdatedAt(updatedAt *time.Time) *Updater {
	c.updatedAt = updatedAt
	return c
}

// Where applies predicates
func (u *Updater) Where(pfs ...PredFunc) *Updater {
	u.pfs = append(u.pfs, pfs...)
	return u
}

// Deleter is a delete builder
type Deleter struct {
	pfs []PredFunc
}

// NewDeleter returns a Deleter
func NewDeleter() *Deleter {
	return &Deleter{}
}

// Where applies predicates
func (d *Deleter) Where(pfs ...PredFunc) *Deleter {
	d.pfs = append(d.pfs, pfs...)
	return d
}

// Aggregator is an aggregate query builder
type Aggregator struct {
	v      interface{}
	aggfs  []AggFunc
	pfs    []PredFunc
	sfs    []SortFunc
	groups []Column
}

// NewAggregator expects a v and returns an Aggregator
// where 'v' argument must be an array of struct
func NewAggregator(v interface{}) *Aggregator {
	return &Aggregator{v: v}
}

// Aggregate applies aggregate functions
func (a *Aggregator) Aggregate(aggfs ...AggFunc) *Aggregator {
	a.aggfs = append(a.aggfs, aggfs...)
	return a
}

// Where applies predicates
func (a *Aggregator) Where(pfs ...PredFunc) *Aggregator {
	a.pfs = append(a.pfs, pfs...)
	return a
}

// Sort applies sorting expressions
func (a *Aggregator) Sort(sfs ...SortFunc) *Aggregator {
	a.sfs = append(a.sfs, sfs...)
	return a
}

// Group applies group clauses
func (a *Aggregator) Group(cols ...Column) *Aggregator {
	a.groups = append(a.groups, cols...)
	return a
}

// rollback performs a rollback
func rollback(tx nero.Tx, err error) error {
	rerr := tx.Rollback()
	if rerr != nil {
		err = errors.Wrapf(err, "rollback error: %v", rerr)
	}
	return err
}

// isZero checks if v is a zero-value
func isZero(v interface{}) bool {
	return reflect.ValueOf(v).IsZero()
}

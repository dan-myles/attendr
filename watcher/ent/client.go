// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"attendr/watcher/ent/migrate"

	"attendr/watcher/ent/asu_watched_class"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ASU_Watched_Class is the client for interacting with the ASU_Watched_Class builders.
	ASU_Watched_Class *ASUWatchedClassClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.ASU_Watched_Class = NewASUWatchedClassClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:               ctx,
		config:            cfg,
		ASU_Watched_Class: NewASUWatchedClassClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:               ctx,
		config:            cfg,
		ASU_Watched_Class: NewASUWatchedClassClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ASU_Watched_Class.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.ASU_Watched_Class.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.ASU_Watched_Class.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ASUWatchedClassMutation:
		return c.ASU_Watched_Class.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ASUWatchedClassClient is a client for the ASU_Watched_Class schema.
type ASUWatchedClassClient struct {
	config
}

// NewASUWatchedClassClient returns a client for the ASU_Watched_Class from the given config.
func NewASUWatchedClassClient(c config) *ASUWatchedClassClient {
	return &ASUWatchedClassClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `asu_watched_class.Hooks(f(g(h())))`.
func (c *ASUWatchedClassClient) Use(hooks ...Hook) {
	c.hooks.ASU_Watched_Class = append(c.hooks.ASU_Watched_Class, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `asu_watched_class.Intercept(f(g(h())))`.
func (c *ASUWatchedClassClient) Intercept(interceptors ...Interceptor) {
	c.inters.ASU_Watched_Class = append(c.inters.ASU_Watched_Class, interceptors...)
}

// Create returns a builder for creating a ASU_Watched_Class entity.
func (c *ASUWatchedClassClient) Create() *ASUWatchedClassCreate {
	mutation := newASUWatchedClassMutation(c.config, OpCreate)
	return &ASUWatchedClassCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ASU_Watched_Class entities.
func (c *ASUWatchedClassClient) CreateBulk(builders ...*ASUWatchedClassCreate) *ASUWatchedClassCreateBulk {
	return &ASUWatchedClassCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ASU_Watched_Class.
func (c *ASUWatchedClassClient) Update() *ASUWatchedClassUpdate {
	mutation := newASUWatchedClassMutation(c.config, OpUpdate)
	return &ASUWatchedClassUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ASUWatchedClassClient) UpdateOne(awc *ASU_Watched_Class) *ASUWatchedClassUpdateOne {
	mutation := newASUWatchedClassMutation(c.config, OpUpdateOne, withASU_Watched_Class(awc))
	return &ASUWatchedClassUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ASUWatchedClassClient) UpdateOneID(id int) *ASUWatchedClassUpdateOne {
	mutation := newASUWatchedClassMutation(c.config, OpUpdateOne, withASU_Watched_ClassID(id))
	return &ASUWatchedClassUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ASU_Watched_Class.
func (c *ASUWatchedClassClient) Delete() *ASUWatchedClassDelete {
	mutation := newASUWatchedClassMutation(c.config, OpDelete)
	return &ASUWatchedClassDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ASUWatchedClassClient) DeleteOne(awc *ASU_Watched_Class) *ASUWatchedClassDeleteOne {
	return c.DeleteOneID(awc.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ASUWatchedClassClient) DeleteOneID(id int) *ASUWatchedClassDeleteOne {
	builder := c.Delete().Where(asu_watched_class.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ASUWatchedClassDeleteOne{builder}
}

// Query returns a query builder for ASU_Watched_Class.
func (c *ASUWatchedClassClient) Query() *ASUWatchedClassQuery {
	return &ASUWatchedClassQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeASUWatchedClass},
		inters: c.Interceptors(),
	}
}

// Get returns a ASU_Watched_Class entity by its id.
func (c *ASUWatchedClassClient) Get(ctx context.Context, id int) (*ASU_Watched_Class, error) {
	return c.Query().Where(asu_watched_class.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ASUWatchedClassClient) GetX(ctx context.Context, id int) *ASU_Watched_Class {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ASUWatchedClassClient) Hooks() []Hook {
	return c.hooks.ASU_Watched_Class
}

// Interceptors returns the client interceptors.
func (c *ASUWatchedClassClient) Interceptors() []Interceptor {
	return c.inters.ASU_Watched_Class
}

func (c *ASUWatchedClassClient) mutate(ctx context.Context, m *ASUWatchedClassMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ASUWatchedClassCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ASUWatchedClassUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ASUWatchedClassUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ASUWatchedClassDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ASU_Watched_Class mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		ASU_Watched_Class []ent.Hook
	}
	inters struct {
		ASU_Watched_Class []ent.Interceptor
	}
)

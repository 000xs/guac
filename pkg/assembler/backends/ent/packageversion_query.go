// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/billofmaterials"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/occurrence"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packagename"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/pkgequal"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// PackageVersionQuery is the builder for querying PackageVersion entities.
type PackageVersionQuery struct {
	config
	ctx                      *QueryContext
	order                    []packageversion.OrderOption
	inters                   []Interceptor
	predicates               []predicate.PackageVersion
	withName                 *PackageNameQuery
	withOccurrences          *OccurrenceQuery
	withSbom                 *BillOfMaterialsQuery
	withEqualPackages        *PkgEqualQuery
	withIncludedInSboms      *BillOfMaterialsQuery
	modifiers                []func(*sql.Selector)
	loadTotal                []func(context.Context, []*PackageVersion) error
	withNamedOccurrences     map[string]*OccurrenceQuery
	withNamedSbom            map[string]*BillOfMaterialsQuery
	withNamedEqualPackages   map[string]*PkgEqualQuery
	withNamedIncludedInSboms map[string]*BillOfMaterialsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PackageVersionQuery builder.
func (pvq *PackageVersionQuery) Where(ps ...predicate.PackageVersion) *PackageVersionQuery {
	pvq.predicates = append(pvq.predicates, ps...)
	return pvq
}

// Limit the number of records to be returned by this query.
func (pvq *PackageVersionQuery) Limit(limit int) *PackageVersionQuery {
	pvq.ctx.Limit = &limit
	return pvq
}

// Offset to start from.
func (pvq *PackageVersionQuery) Offset(offset int) *PackageVersionQuery {
	pvq.ctx.Offset = &offset
	return pvq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pvq *PackageVersionQuery) Unique(unique bool) *PackageVersionQuery {
	pvq.ctx.Unique = &unique
	return pvq
}

// Order specifies how the records should be ordered.
func (pvq *PackageVersionQuery) Order(o ...packageversion.OrderOption) *PackageVersionQuery {
	pvq.order = append(pvq.order, o...)
	return pvq
}

// QueryName chains the current query on the "name" edge.
func (pvq *PackageVersionQuery) QueryName() *PackageNameQuery {
	query := (&PackageNameClient{config: pvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pvq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(packageversion.Table, packageversion.FieldID, selector),
			sqlgraph.To(packagename.Table, packagename.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, packageversion.NameTable, packageversion.NameColumn),
		)
		fromU = sqlgraph.SetNeighbors(pvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOccurrences chains the current query on the "occurrences" edge.
func (pvq *PackageVersionQuery) QueryOccurrences() *OccurrenceQuery {
	query := (&OccurrenceClient{config: pvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pvq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(packageversion.Table, packageversion.FieldID, selector),
			sqlgraph.To(occurrence.Table, occurrence.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, packageversion.OccurrencesTable, packageversion.OccurrencesColumn),
		)
		fromU = sqlgraph.SetNeighbors(pvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySbom chains the current query on the "sbom" edge.
func (pvq *PackageVersionQuery) QuerySbom() *BillOfMaterialsQuery {
	query := (&BillOfMaterialsClient{config: pvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pvq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(packageversion.Table, packageversion.FieldID, selector),
			sqlgraph.To(billofmaterials.Table, billofmaterials.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, packageversion.SbomTable, packageversion.SbomColumn),
		)
		fromU = sqlgraph.SetNeighbors(pvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEqualPackages chains the current query on the "equal_packages" edge.
func (pvq *PackageVersionQuery) QueryEqualPackages() *PkgEqualQuery {
	query := (&PkgEqualClient{config: pvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pvq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(packageversion.Table, packageversion.FieldID, selector),
			sqlgraph.To(pkgequal.Table, pkgequal.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, packageversion.EqualPackagesTable, packageversion.EqualPackagesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIncludedInSboms chains the current query on the "included_in_sboms" edge.
func (pvq *PackageVersionQuery) QueryIncludedInSboms() *BillOfMaterialsQuery {
	query := (&BillOfMaterialsClient{config: pvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pvq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(packageversion.Table, packageversion.FieldID, selector),
			sqlgraph.To(billofmaterials.Table, billofmaterials.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, packageversion.IncludedInSbomsTable, packageversion.IncludedInSbomsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PackageVersion entity from the query.
// Returns a *NotFoundError when no PackageVersion was found.
func (pvq *PackageVersionQuery) First(ctx context.Context) (*PackageVersion, error) {
	nodes, err := pvq.Limit(1).All(setContextOp(ctx, pvq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{packageversion.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pvq *PackageVersionQuery) FirstX(ctx context.Context) *PackageVersion {
	node, err := pvq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PackageVersion ID from the query.
// Returns a *NotFoundError when no PackageVersion ID was found.
func (pvq *PackageVersionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pvq.Limit(1).IDs(setContextOp(ctx, pvq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{packageversion.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pvq *PackageVersionQuery) FirstIDX(ctx context.Context) int {
	id, err := pvq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PackageVersion entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PackageVersion entity is found.
// Returns a *NotFoundError when no PackageVersion entities are found.
func (pvq *PackageVersionQuery) Only(ctx context.Context) (*PackageVersion, error) {
	nodes, err := pvq.Limit(2).All(setContextOp(ctx, pvq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{packageversion.Label}
	default:
		return nil, &NotSingularError{packageversion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pvq *PackageVersionQuery) OnlyX(ctx context.Context) *PackageVersion {
	node, err := pvq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PackageVersion ID in the query.
// Returns a *NotSingularError when more than one PackageVersion ID is found.
// Returns a *NotFoundError when no entities are found.
func (pvq *PackageVersionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pvq.Limit(2).IDs(setContextOp(ctx, pvq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{packageversion.Label}
	default:
		err = &NotSingularError{packageversion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pvq *PackageVersionQuery) OnlyIDX(ctx context.Context) int {
	id, err := pvq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PackageVersions.
func (pvq *PackageVersionQuery) All(ctx context.Context) ([]*PackageVersion, error) {
	ctx = setContextOp(ctx, pvq.ctx, "All")
	if err := pvq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PackageVersion, *PackageVersionQuery]()
	return withInterceptors[[]*PackageVersion](ctx, pvq, qr, pvq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pvq *PackageVersionQuery) AllX(ctx context.Context) []*PackageVersion {
	nodes, err := pvq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PackageVersion IDs.
func (pvq *PackageVersionQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pvq.ctx.Unique == nil && pvq.path != nil {
		pvq.Unique(true)
	}
	ctx = setContextOp(ctx, pvq.ctx, "IDs")
	if err = pvq.Select(packageversion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pvq *PackageVersionQuery) IDsX(ctx context.Context) []int {
	ids, err := pvq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pvq *PackageVersionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pvq.ctx, "Count")
	if err := pvq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pvq, querierCount[*PackageVersionQuery](), pvq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pvq *PackageVersionQuery) CountX(ctx context.Context) int {
	count, err := pvq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pvq *PackageVersionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pvq.ctx, "Exist")
	switch _, err := pvq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pvq *PackageVersionQuery) ExistX(ctx context.Context) bool {
	exist, err := pvq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PackageVersionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pvq *PackageVersionQuery) Clone() *PackageVersionQuery {
	if pvq == nil {
		return nil
	}
	return &PackageVersionQuery{
		config:              pvq.config,
		ctx:                 pvq.ctx.Clone(),
		order:               append([]packageversion.OrderOption{}, pvq.order...),
		inters:              append([]Interceptor{}, pvq.inters...),
		predicates:          append([]predicate.PackageVersion{}, pvq.predicates...),
		withName:            pvq.withName.Clone(),
		withOccurrences:     pvq.withOccurrences.Clone(),
		withSbom:            pvq.withSbom.Clone(),
		withEqualPackages:   pvq.withEqualPackages.Clone(),
		withIncludedInSboms: pvq.withIncludedInSboms.Clone(),
		// clone intermediate query.
		sql:  pvq.sql.Clone(),
		path: pvq.path,
	}
}

// WithName tells the query-builder to eager-load the nodes that are connected to
// the "name" edge. The optional arguments are used to configure the query builder of the edge.
func (pvq *PackageVersionQuery) WithName(opts ...func(*PackageNameQuery)) *PackageVersionQuery {
	query := (&PackageNameClient{config: pvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pvq.withName = query
	return pvq
}

// WithOccurrences tells the query-builder to eager-load the nodes that are connected to
// the "occurrences" edge. The optional arguments are used to configure the query builder of the edge.
func (pvq *PackageVersionQuery) WithOccurrences(opts ...func(*OccurrenceQuery)) *PackageVersionQuery {
	query := (&OccurrenceClient{config: pvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pvq.withOccurrences = query
	return pvq
}

// WithSbom tells the query-builder to eager-load the nodes that are connected to
// the "sbom" edge. The optional arguments are used to configure the query builder of the edge.
func (pvq *PackageVersionQuery) WithSbom(opts ...func(*BillOfMaterialsQuery)) *PackageVersionQuery {
	query := (&BillOfMaterialsClient{config: pvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pvq.withSbom = query
	return pvq
}

// WithEqualPackages tells the query-builder to eager-load the nodes that are connected to
// the "equal_packages" edge. The optional arguments are used to configure the query builder of the edge.
func (pvq *PackageVersionQuery) WithEqualPackages(opts ...func(*PkgEqualQuery)) *PackageVersionQuery {
	query := (&PkgEqualClient{config: pvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pvq.withEqualPackages = query
	return pvq
}

// WithIncludedInSboms tells the query-builder to eager-load the nodes that are connected to
// the "included_in_sboms" edge. The optional arguments are used to configure the query builder of the edge.
func (pvq *PackageVersionQuery) WithIncludedInSboms(opts ...func(*BillOfMaterialsQuery)) *PackageVersionQuery {
	query := (&BillOfMaterialsClient{config: pvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pvq.withIncludedInSboms = query
	return pvq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		NameID int `json:"name_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PackageVersion.Query().
//		GroupBy(packageversion.FieldNameID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pvq *PackageVersionQuery) GroupBy(field string, fields ...string) *PackageVersionGroupBy {
	pvq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PackageVersionGroupBy{build: pvq}
	grbuild.flds = &pvq.ctx.Fields
	grbuild.label = packageversion.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		NameID int `json:"name_id,omitempty"`
//	}
//
//	client.PackageVersion.Query().
//		Select(packageversion.FieldNameID).
//		Scan(ctx, &v)
func (pvq *PackageVersionQuery) Select(fields ...string) *PackageVersionSelect {
	pvq.ctx.Fields = append(pvq.ctx.Fields, fields...)
	sbuild := &PackageVersionSelect{PackageVersionQuery: pvq}
	sbuild.label = packageversion.Label
	sbuild.flds, sbuild.scan = &pvq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PackageVersionSelect configured with the given aggregations.
func (pvq *PackageVersionQuery) Aggregate(fns ...AggregateFunc) *PackageVersionSelect {
	return pvq.Select().Aggregate(fns...)
}

func (pvq *PackageVersionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pvq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pvq); err != nil {
				return err
			}
		}
	}
	for _, f := range pvq.ctx.Fields {
		if !packageversion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pvq.path != nil {
		prev, err := pvq.path(ctx)
		if err != nil {
			return err
		}
		pvq.sql = prev
	}
	return nil
}

func (pvq *PackageVersionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PackageVersion, error) {
	var (
		nodes       = []*PackageVersion{}
		_spec       = pvq.querySpec()
		loadedTypes = [5]bool{
			pvq.withName != nil,
			pvq.withOccurrences != nil,
			pvq.withSbom != nil,
			pvq.withEqualPackages != nil,
			pvq.withIncludedInSboms != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PackageVersion).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PackageVersion{config: pvq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pvq.modifiers) > 0 {
		_spec.Modifiers = pvq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pvq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pvq.withName; query != nil {
		if err := pvq.loadName(ctx, query, nodes, nil,
			func(n *PackageVersion, e *PackageName) { n.Edges.Name = e }); err != nil {
			return nil, err
		}
	}
	if query := pvq.withOccurrences; query != nil {
		if err := pvq.loadOccurrences(ctx, query, nodes,
			func(n *PackageVersion) { n.Edges.Occurrences = []*Occurrence{} },
			func(n *PackageVersion, e *Occurrence) { n.Edges.Occurrences = append(n.Edges.Occurrences, e) }); err != nil {
			return nil, err
		}
	}
	if query := pvq.withSbom; query != nil {
		if err := pvq.loadSbom(ctx, query, nodes,
			func(n *PackageVersion) { n.Edges.Sbom = []*BillOfMaterials{} },
			func(n *PackageVersion, e *BillOfMaterials) { n.Edges.Sbom = append(n.Edges.Sbom, e) }); err != nil {
			return nil, err
		}
	}
	if query := pvq.withEqualPackages; query != nil {
		if err := pvq.loadEqualPackages(ctx, query, nodes,
			func(n *PackageVersion) { n.Edges.EqualPackages = []*PkgEqual{} },
			func(n *PackageVersion, e *PkgEqual) { n.Edges.EqualPackages = append(n.Edges.EqualPackages, e) }); err != nil {
			return nil, err
		}
	}
	if query := pvq.withIncludedInSboms; query != nil {
		if err := pvq.loadIncludedInSboms(ctx, query, nodes,
			func(n *PackageVersion) { n.Edges.IncludedInSboms = []*BillOfMaterials{} },
			func(n *PackageVersion, e *BillOfMaterials) {
				n.Edges.IncludedInSboms = append(n.Edges.IncludedInSboms, e)
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range pvq.withNamedOccurrences {
		if err := pvq.loadOccurrences(ctx, query, nodes,
			func(n *PackageVersion) { n.appendNamedOccurrences(name) },
			func(n *PackageVersion, e *Occurrence) { n.appendNamedOccurrences(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pvq.withNamedSbom {
		if err := pvq.loadSbom(ctx, query, nodes,
			func(n *PackageVersion) { n.appendNamedSbom(name) },
			func(n *PackageVersion, e *BillOfMaterials) { n.appendNamedSbom(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pvq.withNamedEqualPackages {
		if err := pvq.loadEqualPackages(ctx, query, nodes,
			func(n *PackageVersion) { n.appendNamedEqualPackages(name) },
			func(n *PackageVersion, e *PkgEqual) { n.appendNamedEqualPackages(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pvq.withNamedIncludedInSboms {
		if err := pvq.loadIncludedInSboms(ctx, query, nodes,
			func(n *PackageVersion) { n.appendNamedIncludedInSboms(name) },
			func(n *PackageVersion, e *BillOfMaterials) { n.appendNamedIncludedInSboms(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range pvq.loadTotal {
		if err := pvq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pvq *PackageVersionQuery) loadName(ctx context.Context, query *PackageNameQuery, nodes []*PackageVersion, init func(*PackageVersion), assign func(*PackageVersion, *PackageName)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PackageVersion)
	for i := range nodes {
		fk := nodes[i].NameID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(packagename.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "name_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pvq *PackageVersionQuery) loadOccurrences(ctx context.Context, query *OccurrenceQuery, nodes []*PackageVersion, init func(*PackageVersion), assign func(*PackageVersion, *Occurrence)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*PackageVersion)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(occurrence.FieldPackageID)
	}
	query.Where(predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(packageversion.OccurrencesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PackageID
		if fk == nil {
			return fmt.Errorf(`foreign-key "package_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "package_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pvq *PackageVersionQuery) loadSbom(ctx context.Context, query *BillOfMaterialsQuery, nodes []*PackageVersion, init func(*PackageVersion), assign func(*PackageVersion, *BillOfMaterials)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*PackageVersion)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(billofmaterials.FieldPackageID)
	}
	query.Where(predicate.BillOfMaterials(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(packageversion.SbomColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PackageID
		if fk == nil {
			return fmt.Errorf(`foreign-key "package_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "package_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pvq *PackageVersionQuery) loadEqualPackages(ctx context.Context, query *PkgEqualQuery, nodes []*PackageVersion, init func(*PackageVersion), assign func(*PackageVersion, *PkgEqual)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*PackageVersion)
	nids := make(map[int]map[*PackageVersion]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(packageversion.EqualPackagesTable)
		s.Join(joinT).On(s.C(pkgequal.FieldID), joinT.C(packageversion.EqualPackagesPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(packageversion.EqualPackagesPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(packageversion.EqualPackagesPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*PackageVersion]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*PkgEqual](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "equal_packages" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pvq *PackageVersionQuery) loadIncludedInSboms(ctx context.Context, query *BillOfMaterialsQuery, nodes []*PackageVersion, init func(*PackageVersion), assign func(*PackageVersion, *BillOfMaterials)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*PackageVersion)
	nids := make(map[int]map[*PackageVersion]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(packageversion.IncludedInSbomsTable)
		s.Join(joinT).On(s.C(billofmaterials.FieldID), joinT.C(packageversion.IncludedInSbomsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(packageversion.IncludedInSbomsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(packageversion.IncludedInSbomsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*PackageVersion]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*BillOfMaterials](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "included_in_sboms" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (pvq *PackageVersionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pvq.querySpec()
	if len(pvq.modifiers) > 0 {
		_spec.Modifiers = pvq.modifiers
	}
	_spec.Node.Columns = pvq.ctx.Fields
	if len(pvq.ctx.Fields) > 0 {
		_spec.Unique = pvq.ctx.Unique != nil && *pvq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pvq.driver, _spec)
}

func (pvq *PackageVersionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(packageversion.Table, packageversion.Columns, sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeInt))
	_spec.From = pvq.sql
	if unique := pvq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pvq.path != nil {
		_spec.Unique = true
	}
	if fields := pvq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, packageversion.FieldID)
		for i := range fields {
			if fields[i] != packageversion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if pvq.withName != nil {
			_spec.Node.AddColumnOnce(packageversion.FieldNameID)
		}
	}
	if ps := pvq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pvq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pvq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pvq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pvq *PackageVersionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pvq.driver.Dialect())
	t1 := builder.Table(packageversion.Table)
	columns := pvq.ctx.Fields
	if len(columns) == 0 {
		columns = packageversion.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pvq.sql != nil {
		selector = pvq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pvq.ctx.Unique != nil && *pvq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pvq.predicates {
		p(selector)
	}
	for _, p := range pvq.order {
		p(selector)
	}
	if offset := pvq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pvq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedOccurrences tells the query-builder to eager-load the nodes that are connected to the "occurrences"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pvq *PackageVersionQuery) WithNamedOccurrences(name string, opts ...func(*OccurrenceQuery)) *PackageVersionQuery {
	query := (&OccurrenceClient{config: pvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pvq.withNamedOccurrences == nil {
		pvq.withNamedOccurrences = make(map[string]*OccurrenceQuery)
	}
	pvq.withNamedOccurrences[name] = query
	return pvq
}

// WithNamedSbom tells the query-builder to eager-load the nodes that are connected to the "sbom"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pvq *PackageVersionQuery) WithNamedSbom(name string, opts ...func(*BillOfMaterialsQuery)) *PackageVersionQuery {
	query := (&BillOfMaterialsClient{config: pvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pvq.withNamedSbom == nil {
		pvq.withNamedSbom = make(map[string]*BillOfMaterialsQuery)
	}
	pvq.withNamedSbom[name] = query
	return pvq
}

// WithNamedEqualPackages tells the query-builder to eager-load the nodes that are connected to the "equal_packages"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pvq *PackageVersionQuery) WithNamedEqualPackages(name string, opts ...func(*PkgEqualQuery)) *PackageVersionQuery {
	query := (&PkgEqualClient{config: pvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pvq.withNamedEqualPackages == nil {
		pvq.withNamedEqualPackages = make(map[string]*PkgEqualQuery)
	}
	pvq.withNamedEqualPackages[name] = query
	return pvq
}

// WithNamedIncludedInSboms tells the query-builder to eager-load the nodes that are connected to the "included_in_sboms"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pvq *PackageVersionQuery) WithNamedIncludedInSboms(name string, opts ...func(*BillOfMaterialsQuery)) *PackageVersionQuery {
	query := (&BillOfMaterialsClient{config: pvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pvq.withNamedIncludedInSboms == nil {
		pvq.withNamedIncludedInSboms = make(map[string]*BillOfMaterialsQuery)
	}
	pvq.withNamedIncludedInSboms[name] = query
	return pvq
}

// PackageVersionGroupBy is the group-by builder for PackageVersion entities.
type PackageVersionGroupBy struct {
	selector
	build *PackageVersionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pvgb *PackageVersionGroupBy) Aggregate(fns ...AggregateFunc) *PackageVersionGroupBy {
	pvgb.fns = append(pvgb.fns, fns...)
	return pvgb
}

// Scan applies the selector query and scans the result into the given value.
func (pvgb *PackageVersionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pvgb.build.ctx, "GroupBy")
	if err := pvgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PackageVersionQuery, *PackageVersionGroupBy](ctx, pvgb.build, pvgb, pvgb.build.inters, v)
}

func (pvgb *PackageVersionGroupBy) sqlScan(ctx context.Context, root *PackageVersionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pvgb.fns))
	for _, fn := range pvgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pvgb.flds)+len(pvgb.fns))
		for _, f := range *pvgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pvgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pvgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PackageVersionSelect is the builder for selecting fields of PackageVersion entities.
type PackageVersionSelect struct {
	*PackageVersionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pvs *PackageVersionSelect) Aggregate(fns ...AggregateFunc) *PackageVersionSelect {
	pvs.fns = append(pvs.fns, fns...)
	return pvs
}

// Scan applies the selector query and scans the result into the given value.
func (pvs *PackageVersionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pvs.ctx, "Select")
	if err := pvs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PackageVersionQuery, *PackageVersionSelect](ctx, pvs.PackageVersionQuery, pvs, pvs.inters, v)
}

func (pvs *PackageVersionSelect) sqlScan(ctx context.Context, root *PackageVersionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pvs.fns))
	for _, fn := range pvs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pvs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pvs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

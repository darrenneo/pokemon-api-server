// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"pokemon-api-server/internal/model"
)

func newPokemonType(db *gorm.DB, opts ...gen.DOOption) pokemonType {
	_pokemonType := pokemonType{}

	_pokemonType.pokemonTypeDo.UseDB(db, opts...)
	_pokemonType.pokemonTypeDo.UseModel(&model.PokemonType{})

	tableName := _pokemonType.pokemonTypeDo.TableName()
	_pokemonType.ALL = field.NewAsterisk(tableName)
	_pokemonType.Name = field.NewString(tableName, "name")
	_pokemonType.ID = field.NewInt64(tableName, "id")

	_pokemonType.fillFieldMap()

	return _pokemonType
}

type pokemonType struct {
	pokemonTypeDo

	ALL  field.Asterisk
	Name field.String // name of the type
	ID   field.Int64  // pk id of pokemon types

	fieldMap map[string]field.Expr
}

func (p pokemonType) Table(newTableName string) *pokemonType {
	p.pokemonTypeDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pokemonType) As(alias string) *pokemonType {
	p.pokemonTypeDo.DO = *(p.pokemonTypeDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pokemonType) updateTableName(table string) *pokemonType {
	p.ALL = field.NewAsterisk(table)
	p.Name = field.NewString(table, "name")
	p.ID = field.NewInt64(table, "id")

	p.fillFieldMap()

	return p
}

func (p *pokemonType) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pokemonType) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 2)
	p.fieldMap["name"] = p.Name
	p.fieldMap["id"] = p.ID
}

func (p pokemonType) clone(db *gorm.DB) pokemonType {
	p.pokemonTypeDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pokemonType) replaceDB(db *gorm.DB) pokemonType {
	p.pokemonTypeDo.ReplaceDB(db)
	return p
}

type pokemonTypeDo struct{ gen.DO }

func (p pokemonTypeDo) Debug() *pokemonTypeDo {
	return p.withDO(p.DO.Debug())
}

func (p pokemonTypeDo) WithContext(ctx context.Context) *pokemonTypeDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pokemonTypeDo) ReadDB() *pokemonTypeDo {
	return p.Clauses(dbresolver.Read)
}

func (p pokemonTypeDo) WriteDB() *pokemonTypeDo {
	return p.Clauses(dbresolver.Write)
}

func (p pokemonTypeDo) Session(config *gorm.Session) *pokemonTypeDo {
	return p.withDO(p.DO.Session(config))
}

func (p pokemonTypeDo) Clauses(conds ...clause.Expression) *pokemonTypeDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pokemonTypeDo) Returning(value interface{}, columns ...string) *pokemonTypeDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pokemonTypeDo) Not(conds ...gen.Condition) *pokemonTypeDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pokemonTypeDo) Or(conds ...gen.Condition) *pokemonTypeDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pokemonTypeDo) Select(conds ...field.Expr) *pokemonTypeDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pokemonTypeDo) Where(conds ...gen.Condition) *pokemonTypeDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pokemonTypeDo) Order(conds ...field.Expr) *pokemonTypeDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pokemonTypeDo) Distinct(cols ...field.Expr) *pokemonTypeDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pokemonTypeDo) Omit(cols ...field.Expr) *pokemonTypeDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pokemonTypeDo) Join(table schema.Tabler, on ...field.Expr) *pokemonTypeDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pokemonTypeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *pokemonTypeDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pokemonTypeDo) RightJoin(table schema.Tabler, on ...field.Expr) *pokemonTypeDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pokemonTypeDo) Group(cols ...field.Expr) *pokemonTypeDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pokemonTypeDo) Having(conds ...gen.Condition) *pokemonTypeDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pokemonTypeDo) Limit(limit int) *pokemonTypeDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pokemonTypeDo) Offset(offset int) *pokemonTypeDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pokemonTypeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *pokemonTypeDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pokemonTypeDo) Unscoped() *pokemonTypeDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pokemonTypeDo) Create(values ...*model.PokemonType) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pokemonTypeDo) CreateInBatches(values []*model.PokemonType, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pokemonTypeDo) Save(values ...*model.PokemonType) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pokemonTypeDo) First() (*model.PokemonType, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PokemonType), nil
	}
}

func (p pokemonTypeDo) Take() (*model.PokemonType, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PokemonType), nil
	}
}

func (p pokemonTypeDo) Last() (*model.PokemonType, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PokemonType), nil
	}
}

func (p pokemonTypeDo) Find() ([]*model.PokemonType, error) {
	result, err := p.DO.Find()
	return result.([]*model.PokemonType), err
}

func (p pokemonTypeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PokemonType, err error) {
	buf := make([]*model.PokemonType, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pokemonTypeDo) FindInBatches(result *[]*model.PokemonType, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pokemonTypeDo) Attrs(attrs ...field.AssignExpr) *pokemonTypeDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pokemonTypeDo) Assign(attrs ...field.AssignExpr) *pokemonTypeDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pokemonTypeDo) Joins(fields ...field.RelationField) *pokemonTypeDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pokemonTypeDo) Preload(fields ...field.RelationField) *pokemonTypeDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pokemonTypeDo) FirstOrInit() (*model.PokemonType, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PokemonType), nil
	}
}

func (p pokemonTypeDo) FirstOrCreate() (*model.PokemonType, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PokemonType), nil
	}
}

func (p pokemonTypeDo) FindByPage(offset int, limit int) (result []*model.PokemonType, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pokemonTypeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pokemonTypeDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pokemonTypeDo) Delete(models ...*model.PokemonType) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pokemonTypeDo) withDO(do gen.Dao) *pokemonTypeDo {
	p.DO = *do.(*gen.DO)
	return p
}

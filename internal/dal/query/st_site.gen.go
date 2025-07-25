// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/ch3nnn/webstack-go/internal/dal/model"
)

func newStSite(db *gorm.DB, opts ...gen.DOOption) stSite {
	_stSite := stSite{}

	_stSite.stSiteDo.UseDB(db, opts...)
	_stSite.stSiteDo.UseModel(&model.StSite{})

	tableName := _stSite.stSiteDo.TableName()
	_stSite.ALL = field.NewAsterisk(tableName)
	_stSite.ID = field.NewInt(tableName, "id")
	_stSite.CategoryID = field.NewInt(tableName, "category_id")
	_stSite.Title = field.NewString(tableName, "title")
	_stSite.Icon = field.NewString(tableName, "icon")
	_stSite.IconCss = field.NewString(tableName, "icon_css")
	_stSite.ImgPreview = field.NewString(tableName, "img_preview")
	_stSite.Description = field.NewString(tableName, "description")
	_stSite.URL = field.NewString(tableName, "url")
	_stSite.IsUsed = field.NewBool(tableName, "is_used")
	_stSite.CreatedAt = field.NewTime(tableName, "created_at")
	_stSite.UpdatedAt = field.NewTime(tableName, "updated_at")
	_stSite.DeletedAt = field.NewTime(tableName, "deleted_at")
	_stSite.Sort = field.NewInt(tableName, "sort")

	_stSite.fillFieldMap()

	return _stSite
}

type stSite struct {
	stSiteDo stSiteDo

	ALL         field.Asterisk
	ID          field.Int
	CategoryID  field.Int
	Title       field.String
	Icon        field.String
	IconCss     field.String
	Description field.String
	ImgPreview  field.String
	URL         field.String
	IsUsed      field.Bool
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Time
	Sort        field.Int

	fieldMap map[string]field.Expr
}

func (s stSite) Table(newTableName string) *stSite {
	s.stSiteDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s stSite) As(alias string) *stSite {
	s.stSiteDo.DO = *(s.stSiteDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *stSite) updateTableName(table string) *stSite {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt(table, "id")
	s.CategoryID = field.NewInt(table, "category_id")
	s.Title = field.NewString(table, "title")
	s.Icon = field.NewString(table, "icon")
	s.IconCss = field.NewString(table, "icon_css")
	s.Description = field.NewString(table, "description")
	s.ImgPreview = field.NewString(table, "img_preview")
	s.URL = field.NewString(table, "url")
	s.IsUsed = field.NewBool(table, "is_used")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewTime(table, "deleted_at")
	s.Sort = field.NewInt(table, "sort")

	s.fillFieldMap()

	return s
}

func (s *stSite) WithContext(ctx context.Context) IStSiteDo { return s.stSiteDo.WithContext(ctx) }

func (s stSite) TableName() string { return s.stSiteDo.TableName() }

func (s stSite) Alias() string { return s.stSiteDo.Alias() }

func (s stSite) Columns(cols ...field.Expr) gen.Columns { return s.stSiteDo.Columns(cols...) }

func (s *stSite) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *stSite) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 11)
	s.fieldMap["id"] = s.ID
	s.fieldMap["category_id"] = s.CategoryID
	s.fieldMap["title"] = s.Title
	s.fieldMap["icon"] = s.Icon
	s.fieldMap["icon_css"] = s.IconCss
	s.fieldMap["description"] = s.Description
	s.fieldMap["img_preview"] = s.ImgPreview
	s.fieldMap["url"] = s.URL
	s.fieldMap["is_used"] = s.IsUsed
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["sort"] = s.Sort
}

func (s stSite) clone(db *gorm.DB) stSite {
	s.stSiteDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s stSite) replaceDB(db *gorm.DB) stSite {
	s.stSiteDo.ReplaceDB(db)
	return s
}

type stSiteDo struct{ gen.DO }

type IStSiteDo interface {
	gen.SubQuery
	Debug() IStSiteDo
	WithContext(ctx context.Context) IStSiteDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStSiteDo
	WriteDB() IStSiteDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStSiteDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStSiteDo
	Not(conds ...gen.Condition) IStSiteDo
	Or(conds ...gen.Condition) IStSiteDo
	Select(conds ...field.Expr) IStSiteDo
	Where(conds ...gen.Condition) IStSiteDo
	Order(conds ...field.Expr) IStSiteDo
	Distinct(cols ...field.Expr) IStSiteDo
	Omit(cols ...field.Expr) IStSiteDo
	Join(table schema.Tabler, on ...field.Expr) IStSiteDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStSiteDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStSiteDo
	Group(cols ...field.Expr) IStSiteDo
	Having(conds ...gen.Condition) IStSiteDo
	Limit(limit int) IStSiteDo
	Offset(offset int) IStSiteDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStSiteDo
	Unscoped() IStSiteDo
	Create(values ...*model.StSite) error
	CreateInBatches(values []*model.StSite, batchSize int) error
	Save(values ...*model.StSite) error
	First() (*model.StSite, error)
	Take() (*model.StSite, error)
	Last() (*model.StSite, error)
	Find() ([]*model.StSite, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StSite, err error)
	FindInBatches(result *[]*model.StSite, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StSite) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStSiteDo
	Assign(attrs ...field.AssignExpr) IStSiteDo
	Joins(fields ...field.RelationField) IStSiteDo
	Preload(fields ...field.RelationField) IStSiteDo
	FirstOrInit() (*model.StSite, error)
	FirstOrCreate() (*model.StSite, error)
	FindByPage(offset int, limit int) (result []*model.StSite, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStSiteDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s stSiteDo) Debug() IStSiteDo {
	return s.withDO(s.DO.Debug())
}

func (s stSiteDo) WithContext(ctx context.Context) IStSiteDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s stSiteDo) ReadDB() IStSiteDo {
	return s.Clauses(dbresolver.Read)
}

func (s stSiteDo) WriteDB() IStSiteDo {
	return s.Clauses(dbresolver.Write)
}

func (s stSiteDo) Session(config *gorm.Session) IStSiteDo {
	return s.withDO(s.DO.Session(config))
}

func (s stSiteDo) Clauses(conds ...clause.Expression) IStSiteDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s stSiteDo) Returning(value interface{}, columns ...string) IStSiteDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s stSiteDo) Not(conds ...gen.Condition) IStSiteDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s stSiteDo) Or(conds ...gen.Condition) IStSiteDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s stSiteDo) Select(conds ...field.Expr) IStSiteDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s stSiteDo) Where(conds ...gen.Condition) IStSiteDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s stSiteDo) Order(conds ...field.Expr) IStSiteDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s stSiteDo) Distinct(cols ...field.Expr) IStSiteDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s stSiteDo) Omit(cols ...field.Expr) IStSiteDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s stSiteDo) Join(table schema.Tabler, on ...field.Expr) IStSiteDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s stSiteDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStSiteDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s stSiteDo) RightJoin(table schema.Tabler, on ...field.Expr) IStSiteDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s stSiteDo) Group(cols ...field.Expr) IStSiteDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s stSiteDo) Having(conds ...gen.Condition) IStSiteDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s stSiteDo) Limit(limit int) IStSiteDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s stSiteDo) Offset(offset int) IStSiteDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s stSiteDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStSiteDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s stSiteDo) Unscoped() IStSiteDo {
	return s.withDO(s.DO.Unscoped())
}

func (s stSiteDo) Create(values ...*model.StSite) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s stSiteDo) CreateInBatches(values []*model.StSite, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s stSiteDo) Save(values ...*model.StSite) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s stSiteDo) First() (*model.StSite, error) {
	if result, err := s.DO.First(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	} else {
		return result.(*model.StSite), nil
	}
}

func (s stSiteDo) Take() (*model.StSite, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StSite), nil
	}
}

func (s stSiteDo) Last() (*model.StSite, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StSite), nil
	}
}

func (s stSiteDo) Find() ([]*model.StSite, error) {
	result, err := s.DO.Find()
	return result.([]*model.StSite), err
}

func (s stSiteDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StSite, err error) {
	buf := make([]*model.StSite, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s stSiteDo) FindInBatches(result *[]*model.StSite, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s stSiteDo) Attrs(attrs ...field.AssignExpr) IStSiteDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s stSiteDo) Assign(attrs ...field.AssignExpr) IStSiteDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s stSiteDo) Joins(fields ...field.RelationField) IStSiteDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s stSiteDo) Preload(fields ...field.RelationField) IStSiteDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s stSiteDo) FirstOrInit() (*model.StSite, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StSite), nil
	}
}

func (s stSiteDo) FirstOrCreate() (*model.StSite, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StSite), nil
	}
}

func (s stSiteDo) FindByPage(offset int, limit int) (result []*model.StSite, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s stSiteDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s stSiteDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s stSiteDo) Delete(models ...*model.StSite) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *stSiteDo) withDO(do gen.Dao) *stSiteDo {
	s.DO = *do.(*gen.DO)
	return s
}

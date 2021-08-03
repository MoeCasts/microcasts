// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: novels.proto

package novels

import (
	context "context"
	fmt "fmt"

	_ "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	math "math"
	strings "strings"
	time "time"

	ptypes1 "github.com/golang/protobuf/ptypes"
	gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm1 "github.com/jinzhu/gorm"
	field_mask1 "google.golang.org/genproto/protobuf/field_mask"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type NovelORM struct {
	Author    string
	Cover     string
	CreatedAt *time.Time
	DeletedAt *time.Time
	Finished  bool
	Id        uint64
	Name      string
	Slug      string
	State     int32
	Summary   string
	UpdatedAt *time.Time
	Uuid      string
}

// TableName overrides the default tablename generated by GORM
func (NovelORM) TableName() string {
	return "novels"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Novel) ToORM(ctx context.Context) (NovelORM, error) {
	to := NovelORM{}
	var err error
	if prehook, ok := interface{}(m).(NovelWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Uuid = m.Uuid
	to.Name = m.Name
	to.Slug = m.Slug
	to.Author = m.Author
	to.Summary = m.Summary
	to.Cover = m.Cover
	to.Finished = m.Finished
	to.State = int32(m.State)
	if m.CreatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.UpdatedAt); err != nil {
			return to, err
		}
		to.UpdatedAt = &t
	}
	if m.DeletedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.DeletedAt); err != nil {
			return to, err
		}
		to.DeletedAt = &t
	}
	if posthook, ok := interface{}(m).(NovelWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *NovelORM) ToPB(ctx context.Context) (Novel, error) {
	to := Novel{}
	var err error
	if prehook, ok := interface{}(m).(NovelWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Uuid = m.Uuid
	to.Name = m.Name
	to.Slug = m.Slug
	to.Author = m.Author
	to.Summary = m.Summary
	to.Cover = m.Cover
	to.Finished = m.Finished
	to.State = State(m.State)
	if m.CreatedAt != nil {
		if to.CreatedAt, err = ptypes1.TimestampProto(*m.CreatedAt); err != nil {
			return to, err
		}
	}
	if m.UpdatedAt != nil {
		if to.UpdatedAt, err = ptypes1.TimestampProto(*m.UpdatedAt); err != nil {
			return to, err
		}
	}
	if m.DeletedAt != nil {
		if to.DeletedAt, err = ptypes1.TimestampProto(*m.DeletedAt); err != nil {
			return to, err
		}
	}
	if posthook, ok := interface{}(m).(NovelWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Novel the arg will be the target, the caller the one being converted from

// NovelBeforeToORM called before default ToORM code
type NovelWithBeforeToORM interface {
	BeforeToORM(context.Context, *NovelORM) error
}

// NovelAfterToORM called after default ToORM code
type NovelWithAfterToORM interface {
	AfterToORM(context.Context, *NovelORM) error
}

// NovelBeforeToPB called before default ToPB code
type NovelWithBeforeToPB interface {
	BeforeToPB(context.Context, *Novel) error
}

// NovelAfterToPB called after default ToPB code
type NovelWithAfterToPB interface {
	AfterToPB(context.Context, *Novel) error
}

// DefaultCreateNovel executes a basic gorm create call
func DefaultCreateNovel(ctx context.Context, in *Novel, db *gorm1.DB) (*Novel, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(NovelORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(NovelORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type NovelORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type NovelORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultReadNovel executes a basic gorm read call
func DefaultReadNovel(ctx context.Context, in *Novel, db *gorm1.DB) (*Novel, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(NovelORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm2.ApplyFieldSelection(ctx, db, nil, &NovelORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(NovelORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := NovelORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(NovelORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type NovelORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type NovelORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type NovelORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm1.DB) error
}

func DefaultDeleteNovel(ctx context.Context, in *Novel, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(NovelORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&NovelORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(NovelORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type NovelORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type NovelORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm1.DB) error
}

func DefaultDeleteNovelSet(ctx context.Context, in []*Novel, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	var err error
	keys := []uint64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors1.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&NovelORM{})).(NovelORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&NovelORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&NovelORM{})).(NovelORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type NovelORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Novel, *gorm1.DB) (*gorm1.DB, error)
}
type NovelORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Novel, *gorm1.DB) error
}

// DefaultStrictUpdateNovel clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateNovel(ctx context.Context, in *Novel, db *gorm1.DB) (*Novel, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateNovel")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &NovelORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(NovelORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(NovelORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(NovelORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type NovelORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type NovelORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type NovelORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm1.DB) error
}

// DefaultPatchNovel executes a basic gorm update call with patch behavior
func DefaultPatchNovel(ctx context.Context, in *Novel, updateMask *field_mask1.FieldMask, db *gorm1.DB) (*Novel, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	var pbObj Novel
	var err error
	if hook, ok := interface{}(&pbObj).(NovelWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadNovel(ctx, &Novel{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(NovelWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskNovel(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(NovelWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateNovel(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(NovelWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type NovelWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Novel, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type NovelWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Novel, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type NovelWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Novel, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type NovelWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Novel, *field_mask1.FieldMask, *gorm1.DB) error
}

// DefaultPatchSetNovel executes a bulk gorm update call with patch behavior
func DefaultPatchSetNovel(ctx context.Context, objects []*Novel, updateMasks []*field_mask1.FieldMask, db *gorm1.DB) ([]*Novel, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors1.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Novel, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchNovel(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskNovel patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskNovel(ctx context.Context, patchee *Novel, patcher *Novel, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*Novel, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	var updatedDeletedAt bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Uuid" {
			patchee.Uuid = patcher.Uuid
			continue
		}
		if f == prefix+"Name" {
			patchee.Name = patcher.Name
			continue
		}
		if f == prefix+"Slug" {
			patchee.Slug = patcher.Slug
			continue
		}
		if f == prefix+"Author" {
			patchee.Author = patcher.Author
			continue
		}
		if f == prefix+"Summary" {
			patchee.Summary = patcher.Summary
			continue
		}
		if f == prefix+"Cover" {
			patchee.Cover = patcher.Cover
			continue
		}
		if f == prefix+"Finished" {
			patchee.Finished = patcher.Finished
			continue
		}
		if f == prefix+"State" {
			patchee.State = patcher.State
			continue
		}
		if !updatedCreatedAt && strings.HasPrefix(f, prefix+"CreatedAt.") {
			if patcher.CreatedAt == nil {
				patchee.CreatedAt = nil
				continue
			}
			if patchee.CreatedAt == nil {
				patchee.CreatedAt = &timestamp.Timestamp{}
			}
			childMask := &field_mask1.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"CreatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm2.MergeWithMask(patcher.CreatedAt, patchee.CreatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"CreatedAt" {
			updatedCreatedAt = true
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if !updatedUpdatedAt && strings.HasPrefix(f, prefix+"UpdatedAt.") {
			if patcher.UpdatedAt == nil {
				patchee.UpdatedAt = nil
				continue
			}
			if patchee.UpdatedAt == nil {
				patchee.UpdatedAt = &timestamp.Timestamp{}
			}
			childMask := &field_mask1.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"UpdatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm2.MergeWithMask(patcher.UpdatedAt, patchee.UpdatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"UpdatedAt" {
			updatedUpdatedAt = true
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
		if !updatedDeletedAt && strings.HasPrefix(f, prefix+"DeletedAt.") {
			if patcher.DeletedAt == nil {
				patchee.DeletedAt = nil
				continue
			}
			if patchee.DeletedAt == nil {
				patchee.DeletedAt = &timestamp.Timestamp{}
			}
			childMask := &field_mask1.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"DeletedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm2.MergeWithMask(patcher.DeletedAt, patchee.DeletedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"DeletedAt" {
			updatedDeletedAt = true
			patchee.DeletedAt = patcher.DeletedAt
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListNovel executes a gorm list call
func DefaultListNovel(ctx context.Context, db *gorm1.DB) ([]*Novel, error) {
	in := Novel{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(NovelORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &NovelORM{}, &Novel{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(NovelORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []NovelORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(NovelORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Novel{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type NovelORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type NovelORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type NovelORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]NovelORM) error
}

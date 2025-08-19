/**
 * @Author: chentong
 * @Date: 2024/05/27 上午11:14
 */

package category

import (
	"context"
	"strconv"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/converter"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
	"go.uber.org/zap"
)

func (s *service) Create(ctx context.Context, req *v1.CategoryCreateReq) (*v1.CategoryCreateResp, error) {
	if req.Parent != "" {
		// originCate := req.Parent
		// existCates, _ := s.categoryRepo.WithContext(ctx).FindAll(s.categoryRepo.WhereByTitle(req.Parent))
		existCates, _ := query.StCategory.WithContext(ctx).Where(query.StCategory.Title.Eq(req.Parent)).Find()
		if len(existCates) > 0 && existCates[0].ID > 0 {
			req.ParentID = existCates[0].ID
		}
		s.Logger.Logger.Info("add by category", zap.Any("existCates", existCates))
		// existCates, _ = s.categoryRepo.WithContext(ctx).FindAll(s.categoryRepo.WhereBySlug(req.Parent))
		existCates, _ = query.StCategory.WithContext(ctx).Where(query.StCategory.Slug.Eq(req.Parent)).Find()
		if len(existCates) > 0 && existCates[0].ID > 0 {
			req.ParentID = existCates[0].ID
		}
		s.Logger.Logger.Info("add by category", zap.Any("existCates", existCates))

		predictCateId, e := strconv.Atoi(req.Parent)
		if e == nil && predictCateId > 0 {
			// existCates, _ = s.categoryRepo.WithContext(ctx).FindAll(s.categoryRepo.WhereByID(predictCateId))
			existCates, _ = query.StCategory.WithContext(ctx).Where(query.StCategory.ID.Eq(predictCateId)).Find()
			if len(existCates) > 0 && existCates[0].ID > 0 {
				req.ParentID = existCates[0].ID
			}
		}
		s.Logger.Logger.Info("add by category", zap.Any("predictCateId", predictCateId))

	}

	// 如果存在则更新
	// existItem, _ := s.categoryRepo.WithContext(ctx).FindOne(s.categoryRepo.WhereByTitle(req.Title))
	existItems, _ := query.StCategory.WithContext(ctx).Where(query.StCategory.Title.Eq(req.Title)).Find()
	if len(existItems) > 0 {
		existItem := existItems[0]
		if req.Desc != "" {
			existItem.Desc = req.Desc
		}
		if req.Icon != "" {
			existItem.Icon = req.Icon
		}
		if req.IconCss != "" {
			existItem.IconCSS = req.IconCss
		}
		// _, err := s.categoryRepo.WithContext(ctx).Update(existItem, s.categoryRepo.WhereByID(existItem.ID))
		_, err := query.StCategory.WithContext(ctx).Where(query.StCategory.ID.Eq(existItem.ID)).Updates(existItem)
		if err != nil {
			s.Logger.Logger.Info("add by update failed", zap.Error(err))
			return &v1.CategoryCreateResp{}, err
		}
		s.Logger.Logger.Info("add by update", zap.Any("existitem", existItem))
		return &v1.CategoryCreateResp{
			Category: *converter.CategoryModelToDto(existItem),
		}, nil
	} else {

		category := &model.StCategory{
			ParentID: req.ParentID,
			Title:    req.Title,
			Icon:     req.Icon,
			IconCSS:  req.IconCss,
			Desc:     req.Desc,
			Level:    req.Level,
			IsUsed:   true,
			Sort:     req.Sort,
			Slug:     req.Slug,
		}
		err := query.StCategory.WithContext(ctx).Create(category)
		// _, err := s.categoryRepo.WithContext(ctx).Create()
		if err != nil {
			return nil, err
		}

		return &v1.CategoryCreateResp{Category: *converter.CategoryModelToDto(category)}, nil
	}

}

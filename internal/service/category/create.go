/**
 * @Author: chentong
 * @Date: 2024/05/27 上午11:14
 */

package category

import (
	"context"
	"strconv"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"go.uber.org/zap"
)

func (s *service) Create(ctx context.Context, req *v1.CategoryCreateReq) (*v1.CategoryCreateResp, error) {
	if req.Parent != "" {
		// originCate := req.Parent
		existCates, _ := s.categoryRepo.WithContext(ctx).FindAll(s.categoryRepo.WhereByTitle(req.Parent))
		if len(existCates) > 0 && existCates[0].ID > 0 {
			req.ParentID = existCates[0].ID
		}
		s.Logger.Logger.Info("add by category", zap.Any("existCates", existCates))
		existCates, _ = s.categoryRepo.WithContext(ctx).FindAll(s.categoryRepo.WhereBySlug(req.Parent))
		if len(existCates) > 0 && existCates[0].ID > 0 {
			req.ParentID = existCates[0].ID
		}
		s.Logger.Logger.Info("add by category", zap.Any("existCates", existCates))

		predictCateId, e := strconv.Atoi(req.Parent)
		if e == nil && predictCateId > 0 {
			existCates, _ = s.categoryRepo.WithContext(ctx).FindAll(s.categoryRepo.WhereByID(predictCateId))
			if len(existCates) > 0 && existCates[0].ID > 0 {
				req.ParentID = existCates[0].ID
			}
		}
		s.Logger.Logger.Info("add by category", zap.Any("predictCateId", predictCateId))

	}

	// 如果存在则更新
	// existItem, _ := s.categoryRepo.WithContext(ctx).FindOne(s.categoryRepo.WhereByTitle(req.Title))
	existItems, _ := s.categoryRepo.WithContext(ctx).FindAll(s.categoryRepo.WhereByTitle(req.Title))
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
		_, err := s.categoryRepo.WithContext(ctx).Update(existItem, s.categoryRepo.WhereByID(existItem.ID))
		if err != nil {
			s.Logger.Logger.Info("add by update failed", zap.Error(err))
			return &v1.CategoryCreateResp{}, err
		}
		s.Logger.Logger.Info("add by update", zap.Any("existitem", existItem))
		return &v1.CategoryCreateResp{
			Category: v1.Category{
				ID:        existItem.ID,
				ParentID:  existItem.ParentID,
				Sort:      existItem.Sort,
				Slug:      existItem.Slug,
				Title:     existItem.Title,
				Icon:      existItem.Icon,
				IconCss:   existItem.IconCSS,
				CreatedAt: existItem.CreatedAt,
				UpdatedAt: existItem.UpdatedAt,
				IsUsed:    existItem.IsUsed,
				Level:     existItem.Level,
				Desc:      existItem.Desc,
			},
		}, nil
	} else {

		category, err := s.categoryRepo.WithContext(ctx).
			Create(&model.StCategory{
				ParentID: req.ParentID,
				Title:    req.Title,
				Icon:     req.Icon,
				IconCSS:  req.IconCss,
				Desc:     req.Desc,
				Level:    req.Level,
				IsUsed:   true,
				Sort:     req.Sort,
				Slug:     req.Slug,
			})
		if err != nil {
			return nil, err
		}

		return &v1.CategoryCreateResp{Category: v1.Category{
			ID:        category.ID,
			ParentID:  category.ParentID,
			Sort:      category.Sort,
			Title:     category.Title,
			Icon:      category.Icon,
			IconCss:   category.IconCSS,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
			IsUsed:    category.IsUsed,
			Level:     category.Level,
			Desc:      category.Desc,
			Slug:      category.Slug,
		}}, nil
	}

}

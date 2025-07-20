/**
 * @Author: chentong
 * @Date: 2024/06/04 下午4:33
 */

package site

import (
	"context"
	"strconv"

	"github.com/duke-git/lancet/v2/condition"
	"go.uber.org/zap"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
)

func (s *service) Add(ctx context.Context, req *v1.SiteAddReq) (*v1.SiteAddResp, error) {

	var id = 0
	if req.Category != "" {
		existCates, _ := s.categoryRepository.WithContext(ctx).FindAll(s.categoryRepository.WhereByTitle(req.Category))
		if len(existCates) > 0 && existCates[0].ID > 0 {
			req.CategoryID = existCates[0].ID
		}
		existCates, _ = s.categoryRepository.WithContext(ctx).FindAll(s.categoryRepository.WhereByIconCss(req.Category))
		if len(existCates) > 0 && existCates[0].ID > 0 {
			req.CategoryID = existCates[0].ID
		}
		predictCateId, _ := strconv.Atoi(req.Category)
		if predictCateId > 0 {
			existCates, _ = s.categoryRepository.WithContext(ctx).FindAll(s.categoryRepository.WhereByID(predictCateId))
			if len(existCates) > 0 && existCates[0].ID > 0 {
				req.CategoryID = existCates[0].ID
			}
		}
	}

	// 先查询是否存在,存在则更新
	existItems, _ := s.siteRepository.WithContext(ctx).FindAll(s.siteRepository.WhereByURL(req.Url))
	if len(existItems) > 0 {
		existItem := existItems[0]
		// do update
		if req.Title != "" {
			existItem.Title = req.Title
		}
		if req.Description != "" {
			existItem.Description = req.Description
		}
		if req.Icon != "" {
			existItem.Icon = req.Icon
		}
		if req.ImgPreview != "" {
			existItem.ImgPreview = req.ImgPreview
		}
		if req.IconCss != "" {
			existItem.IconCss = req.IconCss
		}
		if req.CategoryID != 0 {
			existItem.CategoryID = req.CategoryID
		}
		_, err := s.siteRepository.WithContext(ctx).Update(existItem, s.siteRepository.WhereByURL(req.Url), s.siteRepository.WhereByID(existItem.ID))
		id = existItem.ID
		if err != nil {
			s.Logger.Logger.Info("add by update failed", zap.Error(err))
			return &v1.SiteAddResp{}, err
		}
		s.Logger.Logger.Info("add by update", zap.Any("existitem", existItem))
	} else {

		newItem, err := s.siteRepository.WithContext(ctx).Create(&model.StSite{
			Title:       condition.Ternary(req.Title != "", req.Title, req.Url),
			Icon:        req.Icon,
			Description: req.Description,
			URL:         req.Url,
			ImgPreview:  req.ImgPreview,
			IconCss:     req.IconCss,
			CategoryID:  req.CategoryID,
			IsUsed:      req.IsUsed,
			Sort:        0,
		})
		if err != nil {
			s.Logger.Logger.Info("add by new failed", zap.Error(err))
			return &v1.SiteAddResp{}, err
		}
		if newItem != nil {
			id = newItem.ID
			s.Logger.Logger.Info("add by new", zap.Any("newItem", newItem))
		}
	}

	return &v1.SiteAddResp{
		Id: id,
	}, nil
}

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
		originCate := req.Category
		existCates, _ := s.categoryRepository.WithContext(ctx).FindAll(s.categoryRepository.WhereByTitle(req.Category))
		if len(existCates) > 0 && existCates[0].ID > 0 {
			req.CategoryID = existCates[0].ID
		}
		s.Logger.Logger.Info("add by category", zap.Any("existCates", existCates))
		existCates, _ = s.categoryRepository.WithContext(ctx).FindAll(s.categoryRepository.WhereBySlug(req.Category))
		if len(existCates) > 0 && existCates[0].ID > 0 {
			req.CategoryID = existCates[0].ID
		}
		s.Logger.Logger.Info("add by category", zap.Any("existCates", existCates))

		predictCateId, e := strconv.Atoi(req.Category)
		if e == nil && predictCateId > 0 {
			existCates, _ = s.categoryRepository.WithContext(ctx).FindAll(s.categoryRepository.WhereByID(predictCateId))
			if len(existCates) > 0 && existCates[0].ID > 0 {
				req.CategoryID = existCates[0].ID
			}
		}
		s.Logger.Logger.Info("add by category", zap.Any("predictCateId", predictCateId))

		if req.CreateCategory == 1 && req.CategoryID == 0 {
			s.Logger.Logger.Info("try to add category", zap.Any("category", req.Category))
			// 创建分类
			newCategory, err := s.categoryRepository.WithContext(ctx).Create(&model.StCategory{
				Title: originCate,
				Slug:  originCate,
			})
			if err != nil {
				s.Logger.Logger.Info("add by new category failed", zap.Error(err))
				return &v1.SiteAddResp{}, err
			}
			if newCategory != nil {
				req.CategoryID = newCategory.ID
				s.Logger.Logger.Info("add by new category", zap.Any("newCategory", newCategory))
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
			existItem.IconCSS = req.IconCss
		}
		if req.CategoryID != 0 {
			existItem.CategoryID = req.CategoryID
		}
		if req.Category != "" {
			existItem.Category = req.Category
		}
		if req.ImgRemote != "" {
			existItem.ImgRemote = req.ImgRemote
		}
		if req.IconRemote != "" {
			existItem.IconRemote = req.IconRemote
		}
		if req.DescS != "" {
			existItem.DescS = req.DescS
		}
		if req.Slug != "" {
			existItem.Slug = req.Slug
		}
		if req.IntroBasic != "" {
			existItem.IntroBasic = req.IntroBasic
		}
		if req.IntroUse != "" {
			existItem.IntroUse = req.IntroUse
		}
		if req.IntroFeatures != "" {
			existItem.IntroFeatures = req.IntroFeatures
		}
		if req.PriceDesc != "" {
			existItem.PriceDesc = req.PriceDesc
		}
		if req.Similar != "" {
			existItem.Similar = req.Similar
		}
		if req.Social != "" {
			existItem.Social = req.Social
		}
		if req.MarkRate != "" {
			existItem.MarkRate = req.MarkRate
		}
		if req.PriceType != 0 {
			existItem.PriceType = req.PriceType
		}
		if req.Tags != "" {
			existItem.Tags = req.Tags
		}
		existItem.IsUsed = true
		existItem.Sort = 0 // 新增的默认排序为0
		_, err := s.siteRepository.WithContext(ctx).Update(existItem, s.siteRepository.WhereByURL(req.Url), s.siteRepository.WhereByID(existItem.ID))
		id = existItem.ID
		if err != nil {
			s.Logger.Logger.Info("add by update failed", zap.Error(err))
			return &v1.SiteAddResp{}, err
		}
		s.Logger.Logger.Info("add by update", zap.Any("existitem", existItem))
	} else {

		newItem, err := s.siteRepository.WithContext(ctx).Create(&model.StSite{
			Title:         condition.Ternary(req.Title != "", req.Title, req.Url),
			Icon:          req.Icon,
			Description:   req.Description,
			URL:           req.Url,
			ImgPreview:    req.ImgPreview,
			IconCSS:       req.IconCss,
			CategoryID:    req.CategoryID,
			Category:      req.Category,
			IsUsed:        true,
			Sort:          0,
			Slug:          req.Slug,
			IntroBasic:    req.IntroBasic,
			IntroUse:      req.IntroUse,
			IntroFeatures: req.IntroFeatures,
			PriceDesc:     req.PriceDesc,
			Similar:       req.Similar,
			Social:        req.Social,
			MarkRate:      req.MarkRate,
			PriceType:     req.PriceType,
			ViewCount:     0,
			IconRemote:    req.IconRemote,
			ImgRemote:     req.ImgRemote,
			DescS:         req.DescS,
			Tags:          req.Tags,
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

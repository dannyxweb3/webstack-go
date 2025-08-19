/**
 * @Author: chentong
 * @Date: 2024/06/04 下午4:33
 */

package site

import (
	"context"
	"strconv"

	"go.uber.org/zap"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/converter"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
)

func (s *service) Add(ctx context.Context, req *v1.SiteAddReq) (*v1.SiteAddResp, error) {

	var id = 0
	if req.Category != "" {
		originCate := req.Category
		// query.StCategory.WithContext(ctx)
		// existCates, _ := s.categoryRepository.WithContext(ctx).FindAll(s.categoryRepository.WhereByTitle(req.Category))

		existCates, _ := query.StCategory.WithContext(ctx).Where(query.StCategory.Title.Eq(req.Category)).Find()
		if len(existCates) > 0 && existCates[0].ID > 0 {
			req.CategoryID = existCates[0].ID
		}
		s.Logger.Logger.Info("add by category", zap.Any("existCates", existCates))
		existCates, _ = query.StCategory.WithContext(ctx).Where(query.StCategory.Slug.Eq(req.Category)).Find()
		if len(existCates) > 0 && existCates[0].ID > 0 {
			req.CategoryID = existCates[0].ID
		}
		s.Logger.Logger.Info("add by category", zap.Any("existCates", existCates))

		predictCateId, e := strconv.Atoi(req.Category)
		if e == nil && predictCateId > 0 {
			existCates, _ = query.StCategory.WithContext(ctx).Where(query.StCategory.ID.Eq(predictCateId)).Find()
			// existCates, _ = s.categoryRepository.WithContext(ctx).FindAll(s.categoryRepository.WhereByID(predictCateId))
			if len(existCates) > 0 && existCates[0].ID > 0 {
				req.CategoryID = existCates[0].ID
			}
		}
		s.Logger.Logger.Info("add by category", zap.Any("predictCateId", predictCateId))

		if req.CreateCategory == 1 && req.CategoryID == 0 {
			s.Logger.Logger.Info("try to add category", zap.Any("category", req.Category))
			// 创建分类
			newCategory := &model.StCategory{
				Title: originCate,
				Slug:  originCate,
			}
			err := query.StCategory.WithContext(ctx).Create(newCategory)
			if err != nil {
				s.Logger.Logger.Info("add by new category failed", zap.Error(err))
				return &v1.SiteAddResp{}, err
			}
			if newCategory.ID == 0 {
				req.CategoryID = newCategory.ID
				s.Logger.Logger.Info("add by new category", zap.Any("newCategory", newCategory))
			}
		}
	}

	// 先查询是否存在,存在则更新
	// existItems, _ := s.siteRepository.WithContext(ctx).FindAll(s.siteRepository.WhereByURL(req.Url))
	existItems, _ := query.StSite.WithContext(ctx).Where(query.StSite.URL.Eq(req.Url)).Find()
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
		if req.PriceTags != "" {
			existItem.PriceTags = req.PriceTags
		}
		if req.Origin != "" {
			existItem.Origin = req.Origin
		}
		// existItem.IsUsed = true
		existItem.Sort = 0 // 新增的默认排序为0
		// _, err := s.siteRepository.WithContext(ctx).Update(existItem, s.siteRepository.WhereByURL(req.Url), s.siteRepository.WhereByID(existItem.ID))
		_, err := query.StSite.WithContext(ctx).Where(query.StSite.URL.Eq(req.Url)).Where(query.StSite.ID.Eq(existItem.ID)).Updates(existItem)
		id = existItem.ID
		if err != nil {
			s.Logger.Logger.Info("add by update failed", zap.Error(err))
			return &v1.SiteAddResp{}, err
		}
		s.Logger.Logger.Info("add by update", zap.Any("existitem", existItem))
	} else {

		newItem := converter.SiteDtoToModel(&req.Site)
		err := query.StSite.WithContext(ctx).Create(newItem)
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

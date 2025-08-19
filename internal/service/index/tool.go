/**
 * @Author: chentong
 * @Date: 2024/05/26 上午1:52
 */

package index

import (
	"errors"
	"strings"
	"time"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
	"github.com/gin-gonic/gin"
)

// ToolDetail
func (s *service) ToolDetail(ctx *gin.Context, slug string) (*v1.ToolDetailResp, error) {
	var (
		// g         errgroup.Group
		sysConfig *model.SysConfig
		// sites     []*model.StSite

		// categories []*model.StCategory
		// 一级分类
		// mainCategories []*model.StCategory
	)

	res := &v1.ToolDetailResp{}

	if slug == "" {
		return res, errors.New("No slug given")
	}

	sysConfig, _ = query.SysConfig.WithContext(ctx).First()
	// sysConfig, _ = s.configRepo.WithContext(ctx).FindOne()
	if sysConfig != nil {

		res.ConfigSite = &v1.ConfigSite{
			SiteTitle:   sysConfig.SiteTitle,
			SiteKeyword: sysConfig.SiteKeyword,
			SiteDesc:    sysConfig.SiteDesc,
			SiteRecord:  sysConfig.SiteRecord,
			SiteURL:     sysConfig.SiteURL,
			SiteLogo:    sysConfig.SiteLogo,
			SiteFavicon: sysConfig.SiteFavicon,
		}
	}

	// main categories
	// cate, _ := query.StCategory.WithContext(ctx).Where(query.StCategory.Slug.Eq(slug)).First()
	// if cate == nil {
	// 	return res, errors.New("Not found")
	// }
	// res.Category = CategoryModelToDto(cate)

	// sites under main categories
	site, _ := query.StSite.WithContext(ctx).
		Where(query.StSite.Slug.Eq(slug)).First()
	res.Site = SiteModelToDto(site)
	// newCateSite.Cnt = int(cnt)
	res.Tags = strings.Split(site.Tags, ",")
	res.Socials = strings.Split(site.Social, ",")

	// popular categories // use random instead
	// popCates, _, _ := s.categoryRepo.WithContext(ctx).FindPage(1, 20, query.StCategory.Columns(query.StCategory.Count))
	popCates, _, _ := query.StCategory.WithContext(ctx).Order(query.StCategory.Count).FindByPage(1, 20)
	for _, ct := range popCates {
		res.PopularCategories = append(res.PopularCategories, CategoryModelToDto(ct))
	}

	// featured tools
	featuredTools, _, _ := query.StSite.WithContext(ctx).
		Where(query.StSite.Featured.Eq(1)).
		Where(query.StSite.Status.Eq(1)).
		FindByPage(1, 3)
	for _, st := range featuredTools {
		res.FeaturedTools = append(res.FeaturedTools, SiteModelToDto(st))
	}

	// randome tools
	rnd := (time.Now().UnixMicro() % 200)
	randomTools, _, _ := query.StSite.WithContext(ctx).
		Where(query.StSite.Status.Eq(1)).
		FindByPage(int(rnd), 20)
	for _, st := range randomTools {
		res.RandomTools = append(res.RandomTools, SiteModelToDto(st))
	}

	return res, nil
}

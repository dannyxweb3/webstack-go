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

// CategotyDetail
func (s *service) CategoryDetail(ctx *gin.Context, slug string, page int) (*v1.CategorySitesResp, error) {
	var (
		// g         errgroup.Group
		sysConfig *model.SysConfig
		// sites     []*model.StSite

		// categories []*model.StCategory
		// 一级分类
		// mainCategories []*model.StCategory
	)

	res := &v1.CategorySitesResp{}

	if slug == "" {
		return res, errors.New("No slug given")
	}
	if page <= 0 {
		page = 1
	}

	// sysConfig, _ = s.configRepo.WithContext(ctx).FindOne()
	sysConfig, _ = query.SysConfig.WithContext(ctx).First()
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
	cate, _ := query.StCategory.WithContext(ctx).Where(query.StCategory.Slug.Eq(slug)).First()
	if cate == nil {
		return res, errors.New("Not found")
	}
	res.Category = CategoryModelToDto(cate)

	// sites under main categories
	sites, cnt, _ := query.StSite.WithContext(ctx).
		Order(query.StSite.CreatedAt.Desc()).
		Where(query.StSite.CategoryID.Eq(cate.ID)).
		Where(query.StSite.Status.Eq(1)).
		FindByPage((page-1)*20, 20)
	for _, st := range sites {
		res.Sites = append(res.Sites, struct {
			*model.StSite
			Tags []string
		}{
			StSite: st,
			Tags:   strings.Split(st.Tags, ","),
		})
	}

	res.Category.Count = int(cnt)
	res.Paginator.TotalItems = int(cnt)
	res.Paginator.PageSize = 20
	res.Paginator.CurPage = page
	res.Paginator.TotalPages = (res.Paginator.TotalItems + res.Paginator.PageSize - 1) / res.Paginator.PageSize
	res.Paginator.PageRange = 5
	if res.Paginator.CurPage > res.Paginator.TotalPages {
		res.Paginator.CurPage = res.Paginator.TotalPages
	}
	// newCateSite.Cnt = int(cnt)

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
	rnd := time.Now().UnixMicro() % 2000
	for rnd > cnt {
		rnd /= 2
	}
	randomTools, _, _ := query.StSite.WithContext(ctx).
		Where(query.StSite.Status.Eq(1)).
		FindByPage(int(rnd), 20)
	for _, st := range randomTools {
		res.RandomTools = append(res.RandomTools, SiteModelToDto(st))
	}

	return res, nil
}

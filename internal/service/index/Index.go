/**
 * @Author: chentong
 * @Date: 2024/05/26 上午1:52
 */

package index

import (
	"context"
	"sort"
	"time"

	"golang.org/x/sync/errgroup"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
)

// buildTree 构建树形结构
func buildTree(nodes []*v1.TreeNode, pid int) []*v1.TreeNode {
	var treeNodes []*v1.TreeNode
	for _, node := range nodes {
		if node.Pid == pid {
			node.Child = buildTree(nodes, node.Id)
			treeNodes = append(treeNodes, node)
		}
	}
	return treeNodes
}

// categoryTree 对树形结构按 Sort 字段排序
func categoryTree(nodes []*v1.TreeNode) []*v1.TreeNode {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Sort < nodes[j].Sort
	})

	for _, node := range nodes {
		if len(node.Child) > 0 {
			categoryTree(node.Child)
		}
	}
	return nodes
}

// categorySites 将站点数据归类到分类站点中
// func categorySites(sites []*model.StSite, treeNodes []*v1.TreeNode) (data []*v1.CategorySite) {
// 	for _, node := range treeNodes {
// 		categorySite := &v1.CategorySite{
// 			Category: node.Name,
// 			CateId:   node.Id,
// 			CateIcon: node.Icon,
// 			SiteList: []model.StSite{},
// 		}

// 		for _, site := range sites {
// 			if site.CategoryID == node.Id {
// 				categorySite.SiteList = append(categorySite.SiteList, *site)
// 			}
// 		}
// 		//  Sort 字段进行升序排序
// 		sort.Slice(categorySite.SiteList, func(i, j int) bool {
// 			return categorySite.SiteList[i].Sort < categorySite.SiteList[j].Sort
// 		})

// 		if len(categorySite.SiteList) > 0 {
// 			data = append(data, categorySite)
// 		}

// 		if len(node.Child) > 0 {
// 			childCategorySites := categorySites(sites, node.Child)
// 			data = append(data, childCategorySites...)
// 		}
// 	}

// 	return data
// }

// Index 获取首页数据
func (s *service) Index(ctx context.Context) (*v1.IndexResp, error) {
	var (
		g         errgroup.Group
		sysConfig *model.SysConfig
	)

	g.Go(func() (err error) {
		// sysConfig, err = s.configRepo.WithContext(ctx).FindOne()
		sysConfig, err = query.SysConfig.WithContext(ctx).First()
		return err
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	resp := &v1.IndexResp{
		ConfigSite: &v1.ConfigSite{
			SiteTitle:   sysConfig.SiteTitle,
			SiteKeyword: sysConfig.SiteKeyword,
			SiteDesc:    sysConfig.SiteDesc,
			SiteRecord:  sysConfig.SiteRecord,
			SiteURL:     sysConfig.SiteURL,
			SiteLogo:    sysConfig.SiteLogo,
			SiteFavicon: sysConfig.SiteFavicon,
		},
		About: &v1.About{
			AboutSite:   sysConfig.AboutSite,
			AboutAuthor: sysConfig.AboutAuthor,
			IsAbout:     sysConfig.IsAbout,
		},
		// CategoryTree:  categoryTree,
		// CategorySites: categorySites,
		Ts: time.Now().Unix(),
	}

	// main categories
	mainCategories, _ := query.StCategory.WithContext(ctx).Order(query.StCategory.Sort, query.StCategory.Title).Where(query.StCategory.ParentID.Eq(0)).Find()
	for mainCateIdx, ct := range mainCategories {
		resp.MainCategories = append(resp.MainCategories, CategoryModelToDto(ct))
		newCateSite := &v1.CategorySite{
			Category:     ct.Title,
			CategorySlug: ct.Slug,
			CateId:       ct.ID,
		}

		// sites under main categories
		sites, cnt, _ := query.StSite.WithContext(ctx).Order(query.StSite.CreatedAt.Desc()).
			Where(query.StSite.MainCategoryID.Eq(ct.ID)).
			Where(query.StSite.Status.Eq(1)).
			FindByPage(1, 8)
		for _, st := range sites {
			if st.Icon == "" {
				st.Icon = st.IconRemote
			}
			newCateSite.SiteList = append(newCateSite.SiteList, *st)
		}
		newCateSite.Cnt = int(cnt)

		// load subcategories
		subCates, _ := query.StCategory.WithContext(ctx).Order(query.StCategory.Sort, query.StCategory.Title).Where(query.StCategory.ParentID.Eq(ct.ID)).Find()
		for _, subCate := range subCates {
			subCate := v1.CategorySite{
				Category:     subCate.Title,
				CategorySlug: subCate.Slug,
				CateIcon:     subCate.Icon,
			}
			if mainCateIdx < 2 {
				// 防止加载过多，只加载前两个
				sitesOfSubCates, cnt, _ := query.StSite.WithContext(ctx).
					Order(query.StSite.CreatedAt.Desc()).
					Where(query.StSite.CategoryID.Eq(subCate.CateId)).
					Where(query.StSite.Status.Eq(1)).
					FindByPage(1, 5)
				for _, st := range sitesOfSubCates {
					if st.Icon == "" {
						st.Icon = st.IconRemote
					}
					subCate.SiteList = append(subCate.SiteList, *st)
				}
				subCate.Cnt = int(cnt)
			}

			newCateSite.SubCategories = append(newCateSite.SubCategories, subCate)
		}

		resp.CategorySites = append(resp.CategorySites, newCateSite)
	}

	// favorite tools
	defaultFavTools := []string{"chatgpt", "grok", "gemini", "claude", "purplexity", "midjourney"}
	// favTools, _ := s.siteRepo.WithContext(ctx).FindAll(s.siteRepo.WhereByStatus(1), func(dao gen.Dao) gen.Dao {
	// return dao.Where(query.StSite.Slug.In(defaultFavTools...))
	// })
	favTools, _ := query.StSite.WithContext(ctx).
		Where(query.StSite.Slug.In(defaultFavTools...)).
		Where(query.StSite.Status.Eq(1)).
		Find()
	for _, st := range favTools {
		resp.FavTools = append(resp.FavTools, SiteModelToDto(st))
	}

	// popular tools
	// popTools, _, _ := s.siteRepo.WithContext(ctx).FindPage(1, 20, query.StSite.Columns(query.StSite.ViewCount))
	popTools, cnt, _ := query.StSite.WithContext(ctx).
		Order(query.StSite.ViewCount).
		Where(query.StSite.Status.Eq(1)).
		FindByPage(1, 8)
	for _, st := range popTools {
		resp.PopularTools = append(resp.PopularTools, SiteModelToDto(st))
	}

	// popular categories // use random instead
	// popCates, _, _ := s.categoryRepo.WithContext(ctx).FindPage(1, 20, query.StCategory.Columns(query.StCategory.Count))
	popCates, _, _ := query.StCategory.WithContext(ctx).Order(query.StCategory.Count).FindByPage(1, 20)
	for _, ct := range popCates {
		resp.PopularCategories = append(resp.PopularCategories, CategoryModelToDto(ct))
	}

	// featured tools
	featuredTools, _, _ := query.StSite.WithContext(ctx).
		Where(query.StSite.Featured.Eq(1)).
		Where(query.StSite.Status.Eq(1)).
		FindByPage(1, 3)
	for _, st := range featuredTools {
		resp.FeaturedTools = append(resp.FeaturedTools, SiteModelToDto(st))
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
		resp.RandomTools = append(resp.RandomTools, SiteModelToDto(st))
	}

	return resp, nil
}

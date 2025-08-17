package v1

import "github.com/ch3nnn/webstack-go/internal/dal/model"

type CategorySitesResp struct {
	ConfigSite *ConfigSite
	*Category
	Sites []struct {
		*model.StSite
		Tags []string
	}

	// 每个一级分类下的前8个
	// Featured (3)
	FeaturedTools []*Site
	// AI News Recommended
	// Popular Categories (5)
	PopularCategories []*Category
	// Randomly Recommended (8)
	RandomTools []*Site
}

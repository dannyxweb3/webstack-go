package v1

type ToolDetailResp struct {
	ConfigSite *ConfigSite
	*Site
	Tags    []string
	Socials []string

	// 每个一级分类下的前8个
	// Featured (3)
	FeaturedTools []*Site
	// AI News Recommended
	// Popular Categories (5)
	PopularCategories []*Category
	// Randomly Recommended (8)
	RandomTools []*Site
}

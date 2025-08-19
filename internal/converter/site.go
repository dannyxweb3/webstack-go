package converter

import (
	"time"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
)

func SiteDtoToModel(dto *v1.Site) (m *model.StSite) {
	return &model.StSite{
		ID:             dto.Id,
		Icon:           dto.Icon,
		Slug:           dto.Slug,
		Title:          dto.Title,
		URL:            dto.Url,
		Category:       dto.Category,
		CategoryID:     dto.CategoryID,
		MainCategoryID: dto.MainCategoryID,
		Description:    dto.Description,
		DescS:          dto.DescS,
		IsUsed:         dto.IsUsed,
		Status:         dto.Status,
		Sort:           dto.Sort,
		// CreatedAt:     dto.CreatedAt,
		// UpdatedAt:     dto.UpdatedAt,
		ImgPreview:    dto.ImgPreview,
		ImgRemote:     dto.ImgRemote,
		IconCSS:       dto.IconCss,
		IconRemote:    dto.IconRemote,
		Tags:          dto.Tags,
		PriceType:     dto.PriceType,
		ViewCount:     dto.ViewCount,
		IntroBasic:    dto.IntroBasic,
		IntroUse:      dto.IntroUse,
		IntroFeatures: dto.IntroFeatures,
		PriceDesc:     dto.PriceDesc,
		Similar:       dto.Similar,
		Social:        dto.Social,
		MarkRate:      dto.MarkRate,
		Featured:      dto.Featured,
		PriceTags:     dto.PriceTags,
		Origin:        dto.Origin,
	}
}

func SiteModelToDto(m *model.StSite) *v1.Site {
	return &v1.Site{
		Id:             m.ID,
		Icon:           m.Icon,
		Slug:           m.Slug,
		Title:          m.Title,
		Url:            m.URL,
		Category:       m.Category,
		CategoryID:     m.CategoryID,
		MainCategoryID: m.MainCategoryID,
		Description:    m.Description,
		DescS:          m.DescS,
		IsUsed:         m.IsUsed,
		Status:         m.Status,
		Sort:           m.Sort,
		CreatedAt:      m.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      m.UpdatedAt.Format(time.RFC3339),
		ImgPreview:     m.ImgPreview,
		ImgRemote:      m.ImgRemote,
		IconCss:        m.IconCSS,
		IconRemote:     m.IconRemote,
		Tags:           m.Tags,
		PriceType:      m.PriceType,
		ViewCount:      m.ViewCount,
		IntroBasic:     m.IntroBasic,
		IntroUse:       m.IntroUse,
		IntroFeatures:  m.IntroFeatures,
		PriceDesc:      m.PriceDesc,
		Similar:        m.Similar,
		Social:         m.Social,
		MarkRate:       m.MarkRate,
		PriceTags:      m.PriceTags,
		Origin:         m.Origin,
	}
}

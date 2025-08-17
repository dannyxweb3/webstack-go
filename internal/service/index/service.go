/**
 * @Author: chentong
 * @Date: 2024/05/26 上午1:48
 */

package index

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	s "github.com/ch3nnn/webstack-go/internal/service"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	// Index 首页
	Index(ctx context.Context) (*v1.IndexResp, error)
	// About 关于我
	About(ctx *gin.Context) (*v1.AboutResp, error)
	// ContactUs
	ContactUs(ctx *gin.Context) (*v1.ContactUsResp, error)
	// AddYourSite
	AddYourSite(ctx *gin.Context) (*v1.AddYourSiteResp, error)
}

type service struct {
	*s.Service
	siteRepo     repository.IStSiteDao
	categoryRepo repository.IStCategoryDao
	configRepo   repository.ISysConfigDao
}

func NewService(
	s *s.Service,
	siteRepo repository.IStSiteDao,
	categoryRepo repository.IStCategoryDao,
	configRepo repository.ISysConfigDao,
) Service {
	return &service{
		Service:      s,
		siteRepo:     siteRepo,
		categoryRepo: categoryRepo,
		configRepo:   configRepo,
	}
}

func (s *service) i() {}

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
	}
}

func CategoryDtoToModel(dto *v1.Category) *model.StCategory {
	return &model.StCategory{
		ID:       dto.ID,
		ParentID: dto.ParentID,
		Sort:     dto.Sort,
		Slug:     dto.Slug,
		Title:    dto.Title,
		Icon:     dto.Icon,
		IconCSS:  dto.IconCss,
		Desc:     dto.Desc,
		Level:    dto.Level,
		IsUsed:   dto.IsUsed,
		Status:   dto.Status,
		// CreatedAt: dto.CreatedAt.Format(time.RFC3339),
		// UpdatedAt: dto.UpdatedAt.Format(time.RFC3339),
		// DeletedAt: dto.DeletedAt.Format(time.RFC3339),
		Count:     dto.Count,
		FreeCount: dto.FreeCount,
	}
}

func CategoryModelToDto(m *model.StCategory) *v1.Category {
	return &v1.Category{
		ID:        m.ID,
		ParentID:  m.ParentID,
		Sort:      m.Sort,
		Slug:      m.Slug,
		Title:     m.Title,
		Icon:      m.Icon,
		IconCss:   m.IconCSS,
		Desc:      m.Desc,
		Level:     m.Level,
		IsUsed:    m.IsUsed,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		// DeletedAt: m.DeletedAt,
		Count:     m.Count,
		FreeCount: m.FreeCount,
	}
}

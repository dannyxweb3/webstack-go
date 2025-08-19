package converter

import (
	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
)

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

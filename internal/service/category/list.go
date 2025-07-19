/**
 * @Author: chentong
 * @Date: 2024/05/27 上午10:23
 */

package category

import (
	"context"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
)

func (s *service) List(ctx context.Context, _ *v1.CategoryListReq) (*v1.CategoryListResp, error) {
	categories, err := s.categoryRepo.WithContext(ctx).FindAllOrderBySort(query.StCategory.Sort.Abs())
	if err != nil {
		return nil, err
	}

	categoryList := make([]v1.Category, len(categories))
	for i, category := range categories {
		categoryList[i] = v1.Category{
			ID:       category.ID,
			ParentID: category.ParentID,
			Title:    category.Title,
			Icon:     category.Icon,
			Desc:     category.Desc,
			IsUsed:   category.IsUsed,
			Sort:     category.Sort,
			Level:    category.Level,
		}
	}

	return &v1.CategoryListResp{List: categoryList}, err
}

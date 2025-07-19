/**
 * @Author: chentong
 * @Date: 2024/05/27 上午11:03
 */

package category

import (
	"context"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (s *service) Detail(ctx context.Context, req *v1.CategoryDetailReq) (*v1.CategoryDetailResp, error) {
	category, err := s.categoryRepo.WithContext(ctx).FindOne(s.categoryRepo.WhereByID(req.ID))
	if err != nil {
		return nil, err
	}

	return &v1.CategoryDetailResp{
		Category: v1.Category{
			ID:       category.ID,
			ParentID: category.ParentID,
			Title:    category.Title,
			Icon:     category.Icon,
			// IsAdd:    category.ParentID == 0,
			Sort: category.Sort,
		},
	}, err
}

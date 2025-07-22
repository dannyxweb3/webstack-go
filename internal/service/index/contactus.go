/**
 * @Author: chentong
 * @Date: 2025/01/18 21:59
 */

package index

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (s *service) ContactUs(ctx *gin.Context) (*v1.ContactUsResp, error) {
	sysConfig, err := s.configRepo.WithContext(ctx).FindOne()
	if err != nil {
		return nil, err
	}

	return &v1.ContactUsResp{
		ConfigSite: &v1.ConfigSite{
			SiteTitle:   sysConfig.SiteTitle,
			SiteKeyword: sysConfig.SiteKeyword,
			SiteDesc:    sysConfig.SiteDesc,
			SiteLogo:    sysConfig.SiteLogo,
			SiteURL:     sysConfig.SiteURL,
			SiteFavicon: sysConfig.SiteFavicon,
			SiteRecord:  sysConfig.SiteRecord,
		},
	}, nil
}
func (s *service) AddYourSite(ctx *gin.Context) (*v1.AddYourSiteResp, error) {
	req := &v1.AddYourSiteReq{}
	err := ctx.ShouldBind(req)
	if err != nil {
		return nil, err
	}

	// check csrftoken

	// gen new token

	s.Logger.Logger.Info("apply for add your site:", zap.Any("req", req))

	return &v1.AddYourSiteResp{
		Csrftoken: "ok",
	}, nil
}

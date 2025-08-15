/**
 * @Author: chentong
 * @Date: 2024/05/27 下午6:19
 */

package v1

import (
	"mime/multipart"

	excelize "github.com/xuri/excelize/v2"
)

type Site struct {
	Id            int    `json:"id"`                             // ID
	Icon          string `json:"icon" form:"icon"`               // 网站 logo
	Slug          string `json:"slug" form:"slug"`               //
	Title         string `json:"title" form:"title"`             // 名称简介
	Url           string `json:"url" form:"url"`                 // 链接
	Category      string `json:"category" form:"category"`       // 分类
	CategoryID    int    `json:"category_id" form:"category_id"` // 分类id
	Description   string `json:"description" form:"description"` // 描述
	DescS         string `json:"desc_s" form:"desc_s"`           // 描述
	IsUsed        bool   `json:"is_used"`                        // 是否启用
	Status        int8   `json:"status"`                         // 是否启用
	Sort          int    `json:"sort" form:"sort"`               // 排序
	CreatedAt     string `json:"created_at"`                     // 创建时间
	UpdatedAt     string `json:"updated_at"`                     // 更新时间
	ImgPreview    string `json:"img_preview" form:"img_preview"`
	ImgRemote     string `json:"img_remote" form:"img_remote"`
	IconCss       string `json:"icon_css" form:"icon_css"`
	IconRemote    string `json:"icon_remote" form:"icon_remote"`
	Tags          string `json:"tags" form:"tags"`
	PriceType     int8   `json:"price_type" form:"price_type"`
	ViewCount     int    `json:"view_count" form:"view_count"`
	IntroBasic    string `json:"intro_basic" form:"intro_basic"`       // 基础介绍
	IntroUse      string `json:"intro_use" form:"intro_use"`           // 使用介绍
	IntroFeatures string `json:"intro_features" form:"intro_features"` // 特性介绍
	PriceDesc     string `json:"price_desc" form:"price_desc"`         // 价格描述
	Similar       string `json:"similar" form:"similar"`               // 相似网站
	Social        string `json:"social" form:"social"`                 // 社交信息
	MarkRate      string `json:"mark_rate" form:"mark_rate"`           // 评分

	CreateCategory int `json:"create_category" form:"create_category"` // 创建分类ID 如果不存在

}

type (
	SiteDeleteReq struct {
		ID int `uri:"id" binding:"required"` // ID
	}

	SiteDeleteResp struct {
		ID int `json:"id"` // ID
	}
)

type (
	SiteLisPagination struct {
		Total        int64 `json:"total"`          // 总数
		CurrentPage  int   `json:"current_page"`   // 当前页
		PerPageCount int   `json:"per_page_count"` // 每页显示条数
	}

	SiteListReq struct {
		Page       int    `form:"page,default=1"`        // 第几页
		PageSize   int    `form:"page_size,default=10" ` // 每页显示条数
		Search     string `form:"search"`                // 搜索关键字
		CategoryID int    `form:"category_id"`           // 分类ID
	}

	SiteListResp struct {
		List       []Site            `json:"list"`       // 列表网站信息
		Pagination SiteLisPagination `json:"pagination"` // 分页信息
	}
)

type (
	SiteCreateReq struct {
		CategoryID int    `form:"category_id"` // 类别ID
		Url        string `form:"url"`         // 网址地址
		IsUsed     bool   `form:"is_used"`     // 是否启用
		FailSwitch bool   `form:"fail_switch"` // 失败开关
	}

	SiteCreateResp struct {
		SuccessCount int      `json:"successCount"` // 成功计数
		FailCount    int      `json:"failCount"`    // 失败计数
		FailURLs     []string `json:"failURLs"`     // 失败URL
		FailErrs     []error  `json:"failErrs"`     // 失败errors
	}
)

type (
	SiteAddReq struct {
		// CategoryID  int    `form:"category_id"` // 类别ID
		// Category    string `form:"category"`    // 类别
		// Url         string `form:"url"`         // 网址地址
		// IsUsed      bool   `form:"is_used"`     // 是否启用
		// Title       string `form:"title"`
		// Icon        string `form:"icon"`
		// Description string `form:"description"`
		// ImgPreview  string `form:"img_preview"`
		// IconCss     string `form:"icon_css"`
		Site

		// FailSwitch bool   `form:"fail_switch"` // 失败开关
	}

	SiteAddResp struct {
		Id int `json:"id"`
	}
)

type (
	SiteUpdateReq struct {
		Id          int                   `json:"id" uri:"id"`                    // ID
		Icon        string                `json:"thumb" form:"thumb"`             // 网站 logo
		Title       string                `json:"title" form:"title"`             // 名称简介
		Url         string                `json:"url" form:"url"`                 // 链接
		CategoryId  int                   `json:"category_id" form:"category_id"` // 分类id
		Description string                `json:"description" form:"description"` // 描述
		IsUsed      *bool                 `json:"is_used" form:"is_used"`         // 是否启用
		File        *multipart.FileHeader `json:"file" form:"file"`               // 上传 logo 图片
		Sort        int                   `json:"sort" form:"sort"`               // 排序
	}

	SiteUpdateResp struct {
		ID int `json:"id"` // 主键ID
	}
)

type (
	SiteSyncReq struct {
		ID int `uri:"id"` // 主键ID
	}
	SiteSyncResp struct {
		ID int `json:"id"` // 主键ID
	}
)

type (
	SiteExportReq struct {
		Search     string `json:"search" form:"search"`           // 搜索关键字
		CategoryID int    `json:"category_id" form:"category_id"` // 分类ID
	}

	SiteExportResp struct {
		File *excelize.File
	}
)

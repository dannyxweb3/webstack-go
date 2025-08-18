/**
 * @Author: chentong
 * @Date: 2024/05/26 上午1:46
 */

package index

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (h *Handler) Index(ctx *gin.Context) {
	resp, err := h.indexService.Index(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.saveHTMLToFile(ctx, resp, "index.html")

	ctx.HTML(http.StatusOK, "index.html", resp)
}
func (h *Handler) Tool(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if slug == "" {
		// h.Index(ctx)
		ctx.Redirect(301, "/")
		return
	}
	resp, err := h.indexService.ToolDetail(ctx, slug)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.saveHTMLToFile(ctx, resp, "tool.html")

	ctx.HTML(http.StatusOK, "tool.html", resp)
}
func (h *Handler) Submit(ctx *gin.Context) {
	resp, err := h.indexService.Index(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.saveHTMLToFile(ctx, resp, "submit.html")

	ctx.HTML(http.StatusOK, "submit.html", resp)
}
func (h *Handler) Disclaimer(ctx *gin.Context) {
	resp, err := h.indexService.Index(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.saveHTMLToFile(ctx, resp, "disclaimer.html")

	ctx.HTML(http.StatusOK, "disclaimer.html", resp)
}
func (h *Handler) Categories(ctx *gin.Context) {
	resp, err := h.indexService.Index(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.saveHTMLToFile(ctx, resp, "categories.html")

	ctx.HTML(http.StatusOK, "categories.html", resp)
}
func (h *Handler) CategoryDetail(ctx *gin.Context) {
	cateSlug := ctx.Param("cate")
	if cateSlug == "" {
		h.Categories(ctx)
		return
	}
	pageStr := ctx.Query("page")
	page, _ := strconv.Atoi(pageStr)
	resp, err := h.indexService.CategoryDetail(ctx, cateSlug, page)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.saveHTMLToFile(ctx, resp, "category_detail.html")

	ctx.HTML(http.StatusOK, "category_detail.html", resp)
}
func (h *Handler) News(ctx *gin.Context) {
	resp, err := h.indexService.Index(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.saveHTMLToFile(ctx, resp, "news.html")

	ctx.HTML(http.StatusOK, "news.html", resp)
}
func (h *Handler) Community(ctx *gin.Context) {
	resp, err := h.indexService.Index(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.saveHTMLToFile(ctx, resp, "community.html") // todo

	ctx.HTML(http.StatusOK, "community.html", resp)
}

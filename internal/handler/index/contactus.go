/**
 * @Author: chentong
 * @Date: 2025/01/18 20:45
 */

package index

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (h *Handler) ContactUs(ctx *gin.Context) {
	resp, err := h.indexService.ContactUs(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	isSave := ctx.GetString("_save")
	if isSave == "1" {
		saveHTMLToFile(h, resp, "contactus.html")
		ctx.Header("X-HTML-Saved", "contactus.html")
	}

	// v1.HandleSuccess(ctx, resp)
	ctx.HTML(http.StatusOK, "contactus.html", resp)
}
func (h *Handler) AddYourSite(ctx *gin.Context) {
	resp, err := h.indexService.AddYourSite(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, resp)
	// ctx.HTML(http.StatusOK, "contactus.html", resp)
}

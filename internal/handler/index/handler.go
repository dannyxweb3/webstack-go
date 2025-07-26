/**
 * @Author: chentong
 * @Date: 2024/05/26 上午1:46
 */

package index

import (
	"bytes"
	"os"
	"text/template"

	"github.com/ch3nnn/webstack-go/internal/handler"
	"github.com/ch3nnn/webstack-go/internal/service/index"
	assets "github.com/ch3nnn/webstack-go/web"
	"go.uber.org/zap"
)

type Handler struct {
	*handler.Handler
	indexService index.Service
}

func NewHandler(handler *handler.Handler, indexService index.Service) *Handler {
	return &Handler{Handler: handler, indexService: indexService}
}

func saveHTMLToFile(h *Handler, obj any, saveName string) error {
	var buf bytes.Buffer
	// 加载模板文件
	tmpl := template.Must(template.ParseFS(assets.Templates, "templates/index/index.html"))
	if err := tmpl.Execute(&buf, obj); err != nil {
		// ctx.String(http.StatusInternalServerError, "template error:"+err.Error())
		h.Logger.Error("template error", zap.Any("error", err))
		return err
	}
	// 将渲染结果写入本地文件
	err := os.WriteFile("output/index.html", buf.Bytes(), 0644)
	if err != nil {
		// ctx.String(http.StatusInternalServerError, "file write error: "+err.Error())
		h.Logger.Error("template error", zap.Any("error", err))
		return err
	}

	return nil
}

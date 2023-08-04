package jie

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type FileUploader struct {
	// 文件在表单中的字段名字
	FileField string
	// 用于计算目标路径，让用户自己决定
	DstPathFunc func(fh *multipart.FileHeader) string
}

func (f FileUploader) Handle() HandleFunc {
	// 这边可以做一下校验
	return func(ctx *Context) {
		//上传逻辑
		//1. 读到文件内容
		//2. 计算出目标路径
		//3. 保存文件
		//4. 返回响应
		src, srcHeader, err := ctx.Req.FormFile(f.FileField)
		if err != nil {
			ctx.RespStatusCode = http.StatusInternalServerError
			ctx.RespData = []byte("上传失败" + err.Error())
			log.Fatalln(err)
			return
		}
		defer src.Close()

		dst, err := os.OpenFile(f.DstPathFunc(srcHeader), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o666)
		if err != nil {
			ctx.RespStatusCode = http.StatusInternalServerError
			ctx.RespData = []byte("上传失败" + err.Error())
			log.Fatalln(err)
			return
		}
		defer dst.Close()

		_, err = io.CopyBuffer(dst, src, nil)
		if err != nil {
			ctx.RespStatusCode = http.StatusInternalServerError
			ctx.RespData = []byte("上传失败" + err.Error())
			log.Fatalln(err)
			return
		}
		ctx.RespStatusCode = http.StatusOK
		ctx.RespData = []byte("upload success")
	}
}

type FileDownloader struct {
	Dir string
}

func (f FileDownloader) Handle() HandleFunc {
	return func(ctx *Context) {
		req, err := ctx.QueryValue("file").String()
		if err != nil {
			ctx.RespStatusCode = http.StatusBadRequest
			ctx.RespData = []byte("没有file字段")
			return
		}
		path := filepath.Join(f.Dir, filepath.Clean(req))
		fn := filepath.Base(path)
		f, err := os.Open(path)
		if err != nil {
			ctx.RespStatusCode = http.StatusBadRequest
			ctx.RespData = []byte("找不到该文件")
			return
		}
		f.Close()

		header := ctx.Resp.Header()
		header.Set("Content-Disposition", "attachment;filename="+fn)
		header.Set("Content-Description", "File Transfer")
		header.Set("Content-Type", "application/octet-stream")
		header.Set("Content-Transfer-Encoding", "binary")
		header.Set("Expires", "0")
		header.Set("Cache-Control", "must-revalidate")
		header.Set("Pragma", "public")
		http.ServeFile(ctx.Resp, ctx.Req, path)
	}
}

type StaticResourceHandler struct {
	dir                     string
	cache                   *lru.Cache
	extensionContextTypeMap map[string]string
	maxFileSize             int
}

type StaticResourceHandlerOption func(h *StaticResourceHandler)

// 控制缓存数据个数
const MaxCacheFileCnt int = 1000

// 控制最大缓存文件
const MaxFileSize int = 100 * 1024 * 1024

func NewStaticResourceHandler(dir string, opts ...StaticResourceHandlerOption) (*StaticResourceHandler, error) {
	cache, err := lru.New(MaxCacheFileCnt)
	if err != nil {
		return nil, err
	}
	res := &StaticResourceHandler{
		dir:   dir,
		cache: cache,
		extensionContextTypeMap: map[string]string{
			// 这里根据自己的需要不断添加
			"jpeg": "image/jpeg",
			"jpe":  "image/jpeg",
			"jpg":  "image/jpeg",
			"png":  "image/png",
			"pdf":  "image/pdf",
		},
		maxFileSize: MaxFileSize,
	}
	for _, opt := range opts {
		opt(res)
	}
	return res, nil
}

func StaticWithMaxFileSize(maxFileSize int) StaticResourceHandlerOption {
	return func(h *StaticResourceHandler) {
		h.maxFileSize = maxFileSize
	}
}

func StaticWithCache(cache *lru.Cache) StaticResourceHandlerOption {
	return func(h *StaticResourceHandler) {
		h.cache = cache
	}
}

func StaticWithMoreExtension(extMap map[string]string) StaticResourceHandlerOption {
	return func(h *StaticResourceHandler) {
		for ext, contentType := range extMap {
			h.extensionContextTypeMap[ext] = contentType
		}
	}
}

func (s *StaticResourceHandler) Handle(ctx *Context) {
	//1. 拿到目标文件名
	//2. 定位到目标文件，读出来
	//3. 返回给前端
	req, err := ctx.PathValue("file").String()
	if err != nil {
		ctx.RespStatusCode = http.StatusBadRequest
		ctx.RespData = []byte("请求路径不对")
		return
	}

	// 得到文件扩展名
	dst := filepath.Join(s.dir, filepath.Clean(req))
	ext := filepath.Ext(dst)
	header := ctx.Resp.Header()

	// 有缓存
	if data, ok := s.cache.Get(req); ok {
		fmt.Println(111)
		header.Set("Content-Type", s.extensionContextTypeMap[ext[1:]])
		header.Set("Content-Length", strconv.Itoa(len(data.([]byte))))
		ctx.RespStatusCode = http.StatusOK
		ctx.RespData = data.([]byte)
		return
	}

	// 无缓存
	data, err := os.ReadFile(dst)
	if err != nil {
		ctx.RespStatusCode = http.StatusInternalServerError
		ctx.RespData = []byte("服务器错误")
		return
	}

	// 大文件不缓存
	if s.maxFileSize >= len(data) {
		s.cache.Add(req, data)
	}

	// 可能的content-type，文本文件，图片，媒体
	header.Set("Content-Type", s.extensionContextTypeMap[ext[1:]])
	header.Set("Content-Length", strconv.Itoa(len(data)))
	ctx.RespStatusCode = http.StatusOK
	ctx.RespData = data
}

package jie

import (
	"html/template"
	"mime/multipart"
	"path"
	"path/filepath"
	"testing"
)

func TestFile(t *testing.T) {
	tpl, err := template.ParseGlob("testdata/tpls/*.gohtml")
	if err != nil {
		t.Fatal(err)
	}
	s := NewHTTPServer(ServerWithTemplateEngine(&GoTemplateEngine{T: tpl}))
	s.Get("/download", FileDownloader{
		Dir: filepath.Join("testdata", "download"),
	}.Handle())
	s.Post("/upload", FileUploader{
		FileField: "myfile",
		DstPathFunc: func(fh *multipart.FileHeader) string {
			return path.Join("testdata", "upload", fh.Filename)
		},
	}.Handle())
	handler, err := NewStaticResourceHandler(filepath.Join("testdata", "static"))
	if err != nil {
		panic(err)
	}
	s.Get("/img/:file", handler.Handle)
	// 在浏览器里面输入 localhost:8081/img/come_on_baby.jpg
	err = s.Start(":8080")
	if err != nil {
		t.Fatal(err)
	}
}

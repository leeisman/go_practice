package main

import (
	"bou.ke/monkey"
	"golang_practice/monkey_patching/implement"
	"log"
	"reflect"
	"testing"
)

func TestDownload_Download(t *testing.T) {
	handler := &implement.Handler{}
	monkey.PatchInstanceMethod(reflect.TypeOf(handler), "CheckDirectLink", func(h *implement.Handler, info string, title string) (bool, error) {
		log.Print("mock info: ", info)
		log.Print("mock title: ", title)
		return true, nil
	})
	download := &implement.Download{
		DirectHandler: &implement.Handler{},
	}
	t.Log(download.DirectHandler.CheckDirectLink("hey hey", "title title"))
}

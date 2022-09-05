package main

import (
	"fmt"
	"io"
	"lcy-faas/function_module"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetFunctionName(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/function/lcy")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if string(body) != `{"message":"lcy"}` {
		t.Error(string(body))
	}
}

var body = function_module.FunctionModule{
	Name:     "lcy",
	Language: "python",
	Source:   "H4sIAPFEEmMAA+3SwQrCMAwG4J19itCTXiTtsgrCniVU12Gh68asyN7e6sSbnhwi9Lv8BEL5A+2MC9thKpaEiJoI7rnT1SNRzfNMS5BUIhFJpTSgLKlSBeCirZ4u52jGVOUwRduYcLRv9tJa2354Z74EXvknXAvMwXSWGeoaBHOXfgSz2K8AhtGFuBYn630P1370jdj8unCWZVn2FTcfGVsPAAgAAA==",
	Method:   "POST",
	Path:     "function-lcy",
	Cpu:      "2",
	Memory:   "512m",
}

func TestBuild(t *testing.T) {

	err := body.Build()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("build")

}

func TestRun(t *testing.T) {
	resp, err := body.Run()
	fmt.Printf("%v", []byte("hello world/n"))
	if err != nil {
		t.Error(err)
	}
	result := append([]byte("hello world"), []byte{10}...)
	if string(resp) != string(result) {
		t.Error(string(resp))
	}

}

func TestDelete(t *testing.T) {

	err := body.Delete()
	if err != nil {
		t.Error(err)
	}
}

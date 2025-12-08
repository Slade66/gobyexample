package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("start.")
	defer fmt.Println("done.")
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		if errors.Is(err, context.Canceled) {
			fmt.Println("客户端取消", err)
		} else if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("超时", err)
		} else {
			fmt.Println("未知错误", err)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	fmt.Println("server run.")
	http.ListenAndServe(":8090", nil)
}

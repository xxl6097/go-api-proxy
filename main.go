package main

import (
	"go-api-proxy/proxy"
	"net/http"
)

func interceptors(w http.ResponseWriter, r *http.Request) bool {
	w.Write([]byte("wosrry"))
	return true
}

func main() {
	_proxy := proxy.NewApiProxy()
	_proxy.SetInterceptor(interceptors)
	_proxy.SetHost("api.uuxia.cn")
	_proxy.SetApiPath("/go/v1")
	_proxy.SetProtocol("http")
	_proxy.AddHeader("Go-Authorization", "xiaxiaoli")
	_proxy.AddHeader("Go-Token", "58e33bea38929b19690f7ee79bbe53ec")
	_proxy.AddHeader("Host", _proxy.GetHost())
	handle, err := _proxy.CreateProxy()
	if err != nil {
		panic(err)
	}
	_proxy.ListenProxy(8080, "/", handle)
}

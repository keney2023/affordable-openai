package handler

import (
	"github.com/suoaao/affordable-openai/pkg/conf"
	myHttputil "github.com/suoaao/affordable-openai/pkg/httputil"
	"github.com/suoaao/affordable-openai/pkg/middleware"
	"net/http"
	"net/http/httputil"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	handler := middleware.VerifyRequestMiddleware(openaiProxyHandler)
	handler.ServeHTTP(w, r)
}

var openaiProxyHandler, _ = NewOpenaiProxyHandler(conf.Conf.ApiKey)

type OpenaiProxyHandler struct {
	apiKey string
	proxy  *httputil.ReverseProxy
}

func NewOpenaiProxyHandler(apiKey string) (*OpenaiProxyHandler, error) {
	var proxy, err = myHttputil.NewReverseProxy("https://api.openai.com")
	if err != nil {
		return nil, err
	}
	return &OpenaiProxyHandler{
		apiKey: apiKey,
		proxy:  proxy,
	}, nil
}

func (h *OpenaiProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Authorization", "Bearer "+h.apiKey)
	h.proxy.ServeHTTP(w, r)
}

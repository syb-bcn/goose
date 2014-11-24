package main

import (
	"net/http"
	"strings"

	ghttp "bitbucket.org/syb-devs/goose/http"
	"github.com/dimfeld/httptreemux"
)

type router struct {
	*httptreemux.TreeMux
}

func newRouter() *router {
	return &router{TreeMux: httptreemux.New()}
}

func (rt *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if rt.isAPI(r) {
		rt.TreeMux.ServeHTTP(w, r)
		return
	}
	ghttp.HandlerAdapter(serveObject)(w, r, nil)
}

func (rt *router) isAPI(r *http.Request) bool {
	//TODO: enable customisation
	return strings.HasPrefix(r.Host, "api.")
}

func (rt *router) withRoutes() *router {
	ctx := ghttp.HandlerAdapter

	rt.POST("/buckets", ctx(postBucket))
	rt.GET("/buckets/:bucket", ctx(getBucket))
	rt.PUT("/buckets/:bucket", ctx(putBucket))
	rt.DELETE("/buckets/:bucket", ctx(deleteBucket))

	rt.POST("/buckets/:bucket/objects", ctx(postObject))
	rt.GET("/buckets/:bucket/objects/:object", ctx(getObject))
	rt.DELETE("/buckets/:bucket/objects/:object", ctx(deleteObject))

	rt.PUT("/buckets/:bucket/objects/:object/metadata", ctx(postObjectMetadata))

	return rt
}

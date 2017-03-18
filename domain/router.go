package domain

import "net/http"

type RouterCase interface {
	Get(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler)
	Post(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler)
	Patch(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler)
	Delete(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler)
	Put(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler)
}

type RouterRepo interface {
	Get(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler)
	Post(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler)
	Patch(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler)
	Delete(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler)
	Put(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler)
}

type RouterRepo2 interface {
	Get(path string, handler http.HandlerFunc)
	Post(path string, handler http.HandlerFunc)
	Patch(path string, handler http.HandlerFunc)
	Delete(path string, handler http.HandlerFunc)
	Put(path string, handler http.HandlerFunc)
}

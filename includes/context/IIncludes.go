/**
	定义接口防止循环引入includes包的结构体
 */
package context

// IRequestContext includes.RequestContext
type IRequestContext interface {
	GetRequest()
}

// IWebRequest includes.WebRequest
type IWebRequest interface {
	GetVal(string, string) string
}


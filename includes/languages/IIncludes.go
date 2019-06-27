/**
定义接口防止循环引入includes包的结构体
*/
package languages

// IMessage includes.Message
type IMessage interface {
	GetRequest(string) string
}



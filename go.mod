module first-demo

go 1.24

// go mod tidy 依赖对齐，自动下载代码中声明的依赖，及其间接依赖，和删除不需要的依赖说明
// go mod download "github.com/gin-gonic/gin v1.10.0" 下载依赖，不会下载关联的间接依赖
// go mod edit 通过命令大的方式修改 go.mod 文件
// go mod vendor 复制依赖到 vendor 目录，代码中的依赖优先使用 vendor 的，而不是 GOPATH 的
// go mod why github.com/bytedance/sonic/loader v0.1.1 查看依赖关系

// go install 安装可执行插件
// go get "github.com/gin-gonic/gin v1.10.0" ，下载依赖，及其间接依赖
// go clean 清除 go run / go build 等产生的临时文件
// go clean -modcache 清除 GOPATH 下 go.mod 需要的依赖

require github.com/gin-gonic/gin v1.10.0

require (
	github.com/bytedance/sonic v1.11.6 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.20.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// 排除依赖
//exclude (
//	dependency latest
//)

// 替换依赖路径
//replace (
//	source latest => target latest
//)

// 撤回
// 本项目作为依赖，被其它项目引用时，出现BUG，可以把某个版本撤回，不能被依赖
//retract {
//	v1.0.0
//}

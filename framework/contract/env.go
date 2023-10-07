package contract

const (
	EnvProduction  = "production"  // 生产
	EnvTesting     = "testing"     //测试
	EnvDevelopment = "development" //开发
	EnvKey         = "hade:env"
)

type Env interface {
	AppEnv() string
	IsExist(string) bool
	Get(string) string
	All() map[string]string
}

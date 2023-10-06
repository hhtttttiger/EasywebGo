package env

import (
	"ASPGo/framework/contract"
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"path"
	"strings"
)

type HadeEnv struct {
	folder string            //代表.env文件所在的目录
	maps   map[string]string //保存所有环境变量
}

// NewHadeEnv 有一个参数，.env文件所在的目录
// example: NewHadeEnv("/envfolder/") 会读取文件: /envfolder/.env
// .env的文件格式 FOO_ENV=BAR
func NewHadeEnv(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("NewHadeEnv param error")
	}

	// 读取folder文件
	folder := params[0].(string)

	// 实例化
	hadeEnv := &HadeEnv{
		folder: folder,
		// 实例化环境变量，APP_ENV默认设置为开发环境
		maps: map[string]string{"APP_ENV": contract.EnvDevelopment},
	}

	// 解析folder/.env文件
	file := path.Join(folder, ".env")
	// 读取.env文件, 不管任意失败，都不影响后续

	// 打开文件.env
	fi, err := os.Open(file)
	if err == nil {
		defer fi.Close()

		// 读取文件
		br := bufio.NewReader(fi)
		for {
			// 按照行进行读取
			line, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			// 按照等号解析
			s := bytes.SplitN(line, []byte{'='}, 2)
			// 如果不符合规范，则过滤
			if len(s) < 2 {
				continue
			}
			// 保存map
			key := string(s[0])
			val := string(s[1])
			hadeEnv.maps[key] = val
		}
	}

	// 获取当前程序的环境变量，并且覆盖.env文件下的变量
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) < 2 {
			continue
		}
		hadeEnv.maps[pair[0]] = pair[1]
	}

	// 返回实例
	return hadeEnv, nil
}

func (env *HadeEnv) AppEnv(key string) string {
	return env.Get(key)
}

func (env *HadeEnv) Get(key string) string {
	if val, ok := env.maps[key]; ok {
		return val
	}

	return ""
}

func (env *HadeEnv) IsExist(key string) bool {
	_, ok := env.maps[key]
	return ok
}

func (env *HadeEnv) All() map[string]string {
	return env.maps
}

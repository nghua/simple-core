package toml

import "github.com/pelletier/go-toml"

type hhToml struct {
	*toml.Tree
}

func Load(path string) *hhToml {
	file, err := toml.LoadFile(path)
	if err != nil {
		panic(err)
	}
	return &hhToml{file}
}

func (t *hhToml) GetInt64(key string) int64 {
	return t.Get(key).(int64)
}

func (t *hhToml) GetString(key string) string {
	return t.Get(key).(string)
}

func (t *hhToml) GetBool(key string) bool {
	return t.Get(key).(bool)
}

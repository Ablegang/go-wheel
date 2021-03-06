// 给 logrus 注册 Hook 时，可以直接自己定义 Hook，自己实现 Hook 的 Fire 和 Levels 方法
// 也可以实例化 writer.Hook，并给 writer.Hook 配置 Writer 的形式来指定 Hook 实例

package writer

import (
	"io"

	log "goweb/pkg/docs/logrus-docs"
)

// Hook is a hook that writes logs of specified LogLevels to specified Writer
type Hook struct {
	Writer    io.Writer
	LogLevels []log.Level
}

// Fire will be called when some logging function is called with current hook
// It will format logs entry to string and write it to appropriate writer
func (hook *Hook) Fire(entry *log.Entry) error {
	line, err := entry.Bytes()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write(line)
	return err
}

// Levels define on which logs levels this hook would trigger
func (hook *Hook) Levels() []log.Level {
	return hook.LogLevels
}

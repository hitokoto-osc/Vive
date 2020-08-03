package flag

import (
	"fmt"
	"github.com/hitokoto-osc/Vive/config"
	pflag "github.com/spf13/pflag"
	"os"
)

var h bool

func registerHelpFlag() {
	pflag.BoolVarP(&h, "help", "h", false, "查看程序帮助")
}

func handleHelpFlag() {
	if h {
		fmt.Printf(`网易云助手 v%s
使用: ncm [-dhv] [-c filename]
选项：
`, config.Version)
		pflag.PrintDefaults()
		os.Exit(0)
	}
}

package runtime

import (
	"sync/atomic"

	"github.com/lf-edge/ekuiper/v2/pkg/model"
)

var (
	initialized atomic.Bool
	conf        *model.KuiperConf
)

func SetAppConf(cfg *model.KuiperConf) { _ = "STUB: not implemented"; return }

// GetAppConf foreign module can get app conf from this
func GetAppConf() *model.KuiperConf { _ = "STUB: not implemented"; return nil }

package component

import (
	"github.com/iotames/glayui/conf"
	"github.com/iotames/glayui/gtpl"
)

type BaseComponent struct {
	gtpl        *gtpl.Gtpl
	useEmbedTpl bool
	// resourceDirPath string
	tplpath string
	name    string
}

func (b BaseComponent) Name() string {
	return b.name
}

func (b *BaseComponent) SetGtpl(gtpl *gtpl.Gtpl) {
	b.gtpl = gtpl
}

// func (b *BaseComponent) SetResourceDirPath(dpath string) *BaseComponent {
// 	b.resourceDirPath = dpath
// 	return b
// }

func (b *BaseComponent) UseEmbedTpl(u bool) *BaseComponent {
	b.useEmbedTpl = u
	return b
}

var gUseEmbedTpl bool
var defaultGtpl *gtpl.Gtpl

func init() {
	gUseEmbedTpl = conf.UseEmbedFile()
	defaultGtpl = gtpl.GetTpl()
}

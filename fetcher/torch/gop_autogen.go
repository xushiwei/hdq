// Code generated by gop (Go+); DO NOT EDIT.

package torch

import (
	"github.com/goplus/hdq"
	"github.com/goplus/hdq/fetcher"
	"github.com/qiniu/x/errors"
	"strings"
)

const GopPackage = "github.com/goplus/hdq"
const _ = true
const spaces = " \t\r\n¶"

type Result struct {
	Name string `json:"name"`
	Type string `json:"type,omitempty"`
	Doc  string `json:"doc,omitempty"`
	Sig  string `json:"sig"`
	URL  string `json:"url,omitempty"`
}
//line fetcher/torch/pysig_torch.gop:39:1
// New creates a new Result from a html document.
func New(input interface{}, doc hdq.NodeSet) Result {
//line fetcher/torch/pysig_torch.gop:41:1
	name := input.(string)
//line fetcher/torch/pysig_torch.gop:42:1
	url := name
//line fetcher/torch/pysig_torch.gop:43:1
	if name != "" {
//line fetcher/torch/pysig_torch.gop:60:1
		url = URL(input)
	}
//line fetcher/torch/pysig_torch.gop:46:1
	if doc.Ok() {
//line fetcher/torch/pysig_torch.gop:47:1
		fn := doc.Any().Dl().Class("py function")
//line fetcher/torch/pysig_torch.gop:48:1
		decl := func() (_gop_ret string) {
//line fetcher/torch/pysig_torch.gop:48:1
			var _gop_err error
//line fetcher/torch/pysig_torch.gop:48:1
			_gop_ret, _gop_err = fn.FirstElementChild().Dt().Text__0()
//line fetcher/torch/pysig_torch.gop:48:1
			if _gop_err != nil {
//line fetcher/torch/pysig_torch.gop:48:1
				_gop_err = errors.NewFrame(_gop_err, "fn.firstElementChild.dt.text", "fetcher/torch/pysig_torch.gop", 48, "torch.New")
//line fetcher/torch/pysig_torch.gop:48:1
				panic(_gop_err)
			}
//line fetcher/torch/pysig_torch.gop:48:1
			return
		}()
//line fetcher/torch/pysig_torch.gop:49:1
		pos := strings.IndexByte(decl, '(')
//line fetcher/torch/pysig_torch.gop:50:1
		if pos > 0 {
//line fetcher/torch/pysig_torch.gop:51:1
			sig := decl[pos:]
//line fetcher/torch/pysig_torch.gop:52:1
			return Result{name, "function", "", strings.TrimRight(sig, spaces), url}
		}
	}
//line fetcher/torch/pysig_torch.gop:55:1
	return Result{name, "", "", "<NULL>", url}
}
//line fetcher/torch/pysig_torch.gop:58:1
// URL returns the input URL for the given name.
func URL(name interface{}) string {
//line fetcher/torch/pysig_torch.gop:60:1
	return "https://pytorch.org/docs/stable/generated/torch." + name.(string) + ".html"
}
//line fetcher/torch/pysig_torch.gop:63:1
func init() {
//line fetcher/torch/pysig_torch.gop:64:1
	fetcher.Register("torch", New, URL)
}

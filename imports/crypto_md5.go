// this file was generated by gomacro command: import "crypto/md5"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"crypto/md5"
)

func init() {
	Binds["crypto/md5"] = map[string]Value{
		"BlockSize":	ValueOf(md5.BlockSize),
		"New":	ValueOf(md5.New),
		"Size":	ValueOf(md5.Size),
		"Sum":	ValueOf(md5.Sum),
	}
	Types["crypto/md5"] = map[string]Type{
	}
	Proxies["crypto/md5"] = map[string]Type{
	}
}
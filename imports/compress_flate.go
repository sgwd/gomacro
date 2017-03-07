// this file was generated by gomacro command: import "compress/flate"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"compress/flate"
	"io"
)

func init() {
	Binds["compress/flate"] = map[string]Value{
		"BestCompression":	ValueOf(flate.BestCompression),
		"BestSpeed":	ValueOf(flate.BestSpeed),
		"DefaultCompression":	ValueOf(flate.DefaultCompression),
		"HuffmanOnly":	ValueOf(flate.HuffmanOnly),
		"NewReader":	ValueOf(flate.NewReader),
		"NewReaderDict":	ValueOf(flate.NewReaderDict),
		"NewWriter":	ValueOf(flate.NewWriter),
		"NewWriterDict":	ValueOf(flate.NewWriterDict),
		"NoCompression":	ValueOf(flate.NoCompression),
	}
	Types["compress/flate"] = map[string]Type{
		"CorruptInputError":	TypeOf((*flate.CorruptInputError)(nil)).Elem(),
		"InternalError":	TypeOf((*flate.InternalError)(nil)).Elem(),
		"ReadError":	TypeOf((*flate.ReadError)(nil)).Elem(),
		"Reader":	TypeOf((*flate.Reader)(nil)).Elem(),
		"Resetter":	TypeOf((*flate.Resetter)(nil)).Elem(),
		"WriteError":	TypeOf((*flate.WriteError)(nil)).Elem(),
		"Writer":	TypeOf((*flate.Writer)(nil)).Elem(),
	}
	Proxies["compress/flate"] = map[string]Type{
		"Reader":	TypeOf((*Reader_compress_flate)(nil)).Elem(),
		"Resetter":	TypeOf((*Resetter_compress_flate)(nil)).Elem(),
	}
}

// --------------- proxy for compress/flate.Reader ---------------
type Reader_compress_flate struct {
	Read_	func(p []byte) (n int, err error)
	ReadByte_	func() (byte, error)
}
func (Obj Reader_compress_flate) Read(p []byte) (n int, err error) {
	return Obj.Read_(p)
}
func (Obj Reader_compress_flate) ReadByte() (byte, error) {
	return Obj.ReadByte_()
}

// --------------- proxy for compress/flate.Resetter ---------------
type Resetter_compress_flate struct {
	Reset_	func(r io.Reader, dict []byte) error
}
func (Obj Resetter_compress_flate) Reset(r io.Reader, dict []byte) error {
	return Obj.Reset_(r, dict)
}
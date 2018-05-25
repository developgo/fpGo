// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package fpGo

type StreamDef struct {
	list []interface{}
}

func (self StreamDef) FromArrayString(old []string) StreamDef {
	new := make([]interface{}, len(old))
	for i, v := range old {
		new[i] = v
	}

	return self.FromArray(new)
}
func (self StreamDef) FromArrayInt(old []int) StreamDef {
	new := make([]interface{}, len(old))
	for i, v := range old {
		new[i] = v
	}

	return self.FromArray(new)
}
func (self StreamDef) FromArrayFloat32(old []float32) StreamDef {
	new := make([]interface{}, len(old))
	for i, v := range old {
		new[i] = v
	}

	return self.FromArray(new)
}
func (self StreamDef) FromArrayFloat64(old []float64) StreamDef {
	new := make([]interface{}, len(old))
	for i, v := range old {
		new[i] = v
	}

	return self.FromArray(new)
}
func (self StreamDef) FromArray(list []interface{}) StreamDef {
	return StreamDef{list: list}
}
func (self StreamDef) ToArray() []interface{} {
	return self.list
}
func (self StreamDef) Map(fn func(int) interface{}) StreamDef {

	for i, _ := range self.list {
		self.list[i] = fn(i)

		// fmt.Println(fn(i))
	}
	return self
}

func (self StreamDef) Get(i int) interface{} {
	return self.list[i]
}

var Stream StreamDef

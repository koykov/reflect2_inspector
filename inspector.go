package reflect2_inspector

import (
	"github.com/koykov/inspector"
	r2 "github.com/modern-go/reflect2"
)

type Inspector struct {
	inspector.BaseInspector
}

func (i Inspector) TypeName() string {
	return "r2"
}

func (i Inspector) Get(src any, path ...string) (any, error) {
	var buf any
	err := i.GetTo(src, &buf, path...)
	return buf, err
}

func (i Inspector) GetTo(src any, buf *any, path ...string) error {
	if src == nil {
		return nil
	}
	t := r2.TypeOf(src)
	i.dive1(buf, t, path)
	return nil
}

func (i Inspector) dive1(buf *any, root r2.Type, path []string) {
	switch x := any(root).(type) {
	case *r2.UnsafePtrType:
		v := x.Elem()
		i.dive1(buf, v, path)
	case *r2.UnsafeStructType:
		if len(path) == 0 {
			return
		}
		f := x.FieldByName(path[0])
		if f == nil {
			return
		}
		if len(path) == 1 {
			*buf = f.Type().Indirect(*buf)
			// *buf = f.Get(*buf)
			return
		}
		t := f.Type()
		i.dive1(buf, t, path[1:])
	}
}

func (i Inspector) Set(_, _ any, _ ...string) error {
	// todo implement me
	return nil
}

func (i Inspector) SetWithBuffer(_, _ any, _ inspector.AccumulativeBuffer, _ ...string) error {
	// todo implement me
	return nil
}

func (i Inspector) Compare(src any, cond inspector.Op, right string, result *bool, _ ...string) error {
	// todo implement me
	return nil
}

func (i Inspector) Loop(_ any, _ inspector.Iterator, _ *[]byte, _ ...string) error {
	return nil
}

func (i Inspector) DeepEqual(l, r any) bool {
	return i.DeepEqualWithOptions(l, r, nil)
}

func (i Inspector) DeepEqualWithOptions(l, r any, _ *inspector.DEQOptions) bool {
	// todo implement me
	return false
}

func (i Inspector) Unmarshal(p []byte, typ inspector.Encoding) (any, error) {
	// todo implement me
	return nil, nil
}

func (i Inspector) Copy(x any) (dst any, err error) {
	// todo implement me
	return
}

func (i Inspector) CopyTo(src, dst any, buf inspector.AccumulativeBuffer) error {
	// todo implement me
	return nil
}

func (i Inspector) Length(x any, result *int, _ ...string) error {
	// todo implement me
	return nil
}

func (i Inspector) Capacity(x any, result *int, _ ...string) error {
	// todo implement me
	return nil
}

func (i Inspector) Reset(x any) error {
	// todo implement me
	return nil
}

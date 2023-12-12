package reflect2_inspector

import (
	"github.com/koykov/inspector"
)

type Inspector struct {
	inspector.BaseInspector
}

func (i Inspector) TypeName() string {
	return "reflect2"
}

func (i Inspector) Get(src any, _ ...string) (any, error) {
	// todo implement me
	return nil, nil
}

func (i Inspector) GetTo(src any, buf *any, _ ...string) error {
	// todo implement me
	return nil
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

package wpcom

import (
	"fmt"
	"net/url"
)

type Options struct {
	url.Values
}

func (o *Options) Empty() bool {
	return len(o.Values) == 0
}

func (o *Options) Add(key string, value interface{}) *Options {
	if o.Values == nil {
		o.Values = make(url.Values)
	}
	switch value := value.(type) {
	case string:
		o.Values.Add(key, value)
	default:
		o.Values.Add(key, fmt.Sprintf("%+v", value))
	}
	return o
}

func (o *Options) Set(key string, value interface{}) *Options {
	if o.Values == nil {
		o.Values = make(url.Values)
	}
	switch value := value.(type) {
	case string:
		o.Values.Set(key, value)
	default:
		o.Values.Set(key, fmt.Sprintf("%+v", value))
	}
	return o
}

func O() *Options {
	return new(Options)
}

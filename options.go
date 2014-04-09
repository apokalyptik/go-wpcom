package wpcom

import "net/url"

type Options struct {
	url.Values
}

func (o Options) Add(key, value string) Options {
	if o.Values == nil {
		o.Values = make(url.Values)
	}
	o.Values.Add(key, value)
	return o
}

func (o Options) Set(key, value string) Options {
	if o.Values == nil {
		o.Values = make(url.Values)
	}
	o.Values.Set(key, value)
	return o
}

package wpcom

import (
	"fmt"
	"net/url"
)

// A wrapper for url.Values which enables us to easily configure API calls.
// Typically you would create a reference to this struct with the exported O()
// function. An example of this would be something like the following:
//   O().Add("foo", "bar").Add("baz[]", "bazone").Add("baz[]", "baztwo").Set("ID",123)
type Options struct {
	url.Values
}

// Determine whether any options have been set
func (o *Options) Empty() bool {
	return len(o.Values) == 0
}

// Add an option. This mirrors url.Values.Add() in that if called
// multiple times it will be added (and retained) multiple times.
// Additionally for convenience you can pass things like integers,
// booleans, floats, directly and they will be translated into
// something workable for iether GET or POST parameters (which are
// not statically typed over the wire due to how HTTP works)
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

// Set an option. This mirrors url.Values.Set() in that if an option of
// the name exists it will be overwritten with the new value. Otherwise
// works the same as Add with regard to types, et al.
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

// Convenience function to return a reference to an Options struct.  See
// the documentation for the Options struct for example usage.
func O() *Options {
	return new(Options)
}

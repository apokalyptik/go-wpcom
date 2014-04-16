package wpcom

import (
	"errors"
	"time"
)

// A helper type for the arbitrarily complex serialized JSON returned by some
// parts of the WordPress.com API.  The reason for such complexity in the
// returned data is that, since it's written in PHP, it's trivially easy to
// tynamically build and interpret some insanely complex data structures. Also
// because of PHPs type coersion it's unremarkable to consume something as a
// boolean which is actually an int 1|0 or a string "1"|"" or a bool true|false
// or something and null. These kinds of things are not trivial to consume in
// go. This struct is meant to represent an Object from PHP.  It'll only be
// as complete as needed for consuming this API.
type Object map[string]interface{}

// For key, return whether PHP would evaluate the value as true, or false.
// If the key does not exist as a member variable to the object return false
// and an error
func (o Object) Bool(key string) (rval bool, err error) {
	v, ok := o[key]
	if ok != true {
		return false, errors.New("No such member")
	}
	return hack(v).bool(), nil
}

/*
// For Key, return the float value. Returns an error if key does not exist
// or if key is not actually a float.
func (o Object) Float(key string) (rval float64, err error) {
	v, ok := o[key]
	if ok != true {
		return 0, errors.New("No such member")
	}
	switch rval := v.(type) {
	case float32, float64:
		return rval.(float64), nil
	default:
		return 0, errors.New("key is not of type")
	}
}
*/

// For key, return an integer.  Returns an error if the key does not exist
// or if the value is not an integer type.
func (o Object) Int(key string) (rval int64, err error) {
	v, ok := o[key]
	if ok != true {
		return 0, errors.New("No such member")
	}
	switch rval := v.(type) {
	case int, int8, int16, int32, int64:
		return rval.(int64), nil
	default:
		return 0, errors.New("key is not of type")
	}
}

// For key, return a string. Returns an error if the key does not exist or
// if the value is not of an appropriate type
func (o Object) String(key string) (rval string, err error) {
	v, ok := o[key]
	if ok != true {
		return "", errors.New("No such member")
	}
	switch rv := v.(type) {
	case string:
		return rv, nil
	default:
		return "", errors.New("key is not of type")
	}
}

// For key return a time.Time.  Returns an error if the key does not
// exist or if the value was not a string, or if the string cannot be
// parsed per the following time format: 2006-01-02T15:04:05-07:00
func (o Object) Time(key string) (rval time.Time, err error) {
	s, err := o.String(key)
	if err != nil {
		return
	}
	return time.Parse("2006-01-02T15:04:05-07:00", s)
}

// For the key, return a new Object. For deep data structures.
func (o Object) O(key string) (rval Object, err error) {
	v, ok := o[key]
	if ok != true {
		return nil, errors.New("No such member")
	}
	switch rv := v.(type) {
	case map[string]interface{}:
		rval = make(map[string]interface{})
		for rvk, rvv := range rv {
			rval[rvk] = rvv
		}
		return rval, nil
	default:
		return nil, errors.New("key is not of type")
	}
}

package wpcom

import (
	"errors"
	"log"
	"time"
)

type Object map[string]interface{}

func (o Object) Bool(key string) (rval bool, err error) {
	v, ok := o[key]
	if ok != true {
		return false, errors.New("No such member")
	}
	return hack(v).bool(), nil
}

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

func (o Object) Time(key string) (rval time.Time, err error) {
	s, err := o.String(key)
	if err != nil {
		return
	}
	//re := regexp.MustCompile("^(\\d+-\\d+-\\d+T\\d+:\\d+:\\d+)-(\\d+:\\d+)$")
	//s = re.ReplaceAllString(s, "${1}Z${2}")
	log.Printf(s)
	//	err = rval.UnmarshalJSON([]byte(s))
	//                 2014-04-15T17:19:30Z07:00
	// 2014-04-15T17:19:30-07:00
	return time.Parse("2006-01-02T15:04:05-07:00", s)
}

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

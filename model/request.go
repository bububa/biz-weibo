package model

import "net/url"

// Request api request interface
type Request interface {
	URL() string
	Payload() *Payload
}

// Payload post or get payload wraper
type Payload struct {
	data   []byte
	values url.Values
	isPost bool
}

// NewGetPayload init post payload
func NewGetPayload(values url.Values) *Payload {
	return &Payload{
		values: values,
	}
}

func (p *Payload) AddValue(key string, val string) {
	if p.values == nil {
		p.values = make(url.Values)
	}
	p.values.Set(key, val)
}

// NewPostPayload init get payload
func NewPostPayload(data []byte) *Payload {
	return &Payload{
		data:   data,
		isPost: true,
	}
}

// IsPost check wether the payload is for post
func (p Payload) IsPost() bool {
	return p.isPost
}

// Data return payload post data
func (p Payload) Data() []byte {
	return p.data
}

// Values return payload get url values
func (p Payload) Values() url.Values {
	return p.values
}

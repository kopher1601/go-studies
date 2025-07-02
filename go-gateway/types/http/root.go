package http

import "errors"

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
)

func (h HttpMethod) String() string {
	return string(h)
}

type GetType string

const (
	QUERY = GetType("QUERY")
	URL   = GetType("URL")
)

func (g GetType) String() string {
	return string(g)
}

func (g GetType) CheckType() error {
	switch g {
	case QUERY, URL:
		return nil
	default:
		return errors.New("invalid get type")
	}
}

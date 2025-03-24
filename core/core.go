package core

import "net/http"

type Potential interface {
	http.Handler
}

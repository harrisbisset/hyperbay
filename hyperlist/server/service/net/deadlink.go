package net

import (
	"net/http"
)

type LinkStatus int

const (
	Unknown  LinkStatus = -1
	Alive    LinkStatus = 1
	Redirect LinkStatus = 2
	Dead     LinkStatus = 3
)

func IsDeadlink(url string) (LinkStatus, error) {
	res, err := http.Get(url)
	if err != nil {
		return Unknown, err
	}

	c := res.StatusCode
	if c > 299 && c < 400 {
		return Redirect, nil
	} else if c > 399 {
		return Dead, nil
	}

	return Alive, nil
}

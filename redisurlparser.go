package redisurlparser

import (
	"net/url"
	"strings"
)

type RedisURL struct {
	Username string
	Password string
	Host     string
	Port     string
}

func Parse(redis_url string) (RedisURL, error) {

	u, err := url.Parse(redis_url)
	if err != nil {
		return RedisURL{}, err
	}
	//redis_url := "redis://redistogo:64cfea0093507536a374ab6ad40f8463@angelfish.redistogo.com:10059/"
	result := strings.Split(u.Host, ":")
	if err != nil {
		return RedisURL{}, err
	}
	username := u.User.Username()
	password, _ := u.User.Password()
	host := result[0]
	port := result[1]

	ru := RedisURL{username, password, host, port}
	if err != nil {
		return RedisURL{}, err
	}

	return ru, nil
}

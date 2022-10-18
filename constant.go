package main

import "time"

// colections name in mongodb
type ColectionKey string

const (
	MEMO_COLLECTION  string = "memos"
	USER_COLLECTION  string = "users"
	TOPIC_COLLECTION string = "topics"
	MOVIE_COLLECTION string = "movies"
)

// session key in cookies
const (
	SESSION_COOKIE_KEY string = "session"
)

// context key
type ContextKey string

const (
	USER_CONTEXT_KEY ContextKey = "user"
)

// expire time
const (
	TOKEN_EXPIRATION_TIME time.Duration = (time.Hour * 365 * 24 * 100) // ~ 100 years
)

// request and response keys
const (
	AUTH_HEADER_KEY      string = "Authorization"
	X_REQUEST_HEADER_KEY string = "X-Requested-With"
	CONTENT_HEADER_KEY   string = "Content-Type"
)

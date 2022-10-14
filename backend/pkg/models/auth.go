package models

import "github.com/gorilla/sessions"

type Store struct {
	Session *sessions.Store
}

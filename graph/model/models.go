// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Term struct {
	ID    int64     `json:"id"`
	Name  string    `json:"name"`
	Meta  *TermMeta `json:"meta"`
	Count int64     `json:"count"`
}

type TermMeta struct {
	Description *string `json:"description"`
}

type User struct {
	ID    int64     `json:"id"`
	Email string    `json:"email"`
	Role  int       `json:"role"`
	Meta  *UserMeta `json:"meta"`
}

type UserMeta struct {
	Nickname *string `json:"nickname"`
}

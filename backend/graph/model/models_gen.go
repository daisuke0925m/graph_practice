// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gmodel

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type Connection interface {
	IsConnection()
}

type Edge interface {
	IsEdge()
}

type Node interface {
	IsNode()
}

type JoinUserInput struct {
	UserID string `json:"userID"`
}

type Message struct {
	ID      string `json:"id"`
	User    *User  `json:"user"`
	Message string `json:"message"`
}

func (Message) IsNode() {}

type NewTask struct {
	Title string `json:"title"`
	Note  string `json:"note"`
}

type PageInfo struct {
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
}

type PaginationInput struct {
	First  *int    `json:"first"`
	After  *string `json:"after"`
	Last   *int    `json:"last"`
	Before *string `json:"before"`
}

type PostMessageInput struct {
	UserID  string `json:"userID"`
	Message string `json:"message"`
}

type Task struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Note      string    `json:"note"`
	Completed int       `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User     `json:"user"`
}

func (Task) IsNode() {}

type TaskConnection struct {
	PageInfo *PageInfo   `json:"pageInfo"`
	Edges    []*TaskEdge `json:"edges"`
	Nodes    []*Task     `json:"nodes"`
}

func (TaskConnection) IsConnection() {}

type TaskEdge struct {
	Cursor string `json:"cursor"`
	Node   *Task  `json:"node"`
}

func (TaskEdge) IsEdge() {}

type UpdateTask struct {
	ID        string  `json:"id"`
	Title     *string `json:"title"`
	Note      *string `json:"note"`
	Completed *int    `json:"completed"`
}

type UploadFileInput struct {
	File graphql.Upload `json:"file"`
}

type UploadFilePayload struct {
	UploadedPath string `json:"uploadedPath"`
}

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) IsNode() {}

type UserStatus struct {
	UserID string `json:"userId"`
	Status State  `json:"status"`
}

type UpdateUserStatusInput struct {
	UserID  string `json:"userID"`
	Message string `json:"message"`
}

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

var AllRole = []Role{
	RoleAdmin,
	RoleUser,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleUser:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type State string

const (
	StateOnline  State = "ONLINE"
	StateOffline State = "OFFLINE"
)

var AllState = []State{
	StateOnline,
	StateOffline,
}

func (e State) IsValid() bool {
	switch e {
	case StateOnline, StateOffline:
		return true
	}
	return false
}

func (e State) String() string {
	return string(e)
}

func (e *State) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = State(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid State", str)
	}
	return nil
}

func (e State) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

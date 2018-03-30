package model

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	//"github.com/aws/aws-sdk-go/dynamodb/dynamodbattribute"
)

var (
	currentID uint64
	todos     []Todo
)

func init() {
	currentID = 0
	RepoCreateTodo(Todo{Title: "Write Presentation"})
	RepoCreateTodo(Todo{Title: "Host Meetup"})
}

type TodoDao interface {
	FindById(id uint64) *Todo
	List() []*Todo
	Delete(id uint64)
	Update(todo Todo) *Todo

	Error() error
}

type TodoDaoImpl struct {
	err    error
	sess   *session.Session
	dynamo *dynamodb.DynamoDB
}

func (dao TodoDaoImpl) FindById(id uint64) *Todo {
	return nil
}

func (dao TodoDaoImpl) List() []*Todo {
	return make([]*Todo, 0)
}

func (dao TodoDaoImpl) Delete(id uint64) {
	return
}

func (dao TodoDaoImpl) Update(todo Todo) *Todo {
	return nil
}

func (dao TodoDaoImpl) Error() error {
	return dao.err
}

func New() TodoDao {
	sess := session.Must(session.NewSession())
	dynamo := dynamodb.New(sess)

	todoDao := TodoDaoImpl{
		sess:   sess,
		dynamo: dynamo,
	}

	return todoDao
}

// RepoFindTodo - Find Todo
func RepoFindTodo(id uint64) Todo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}

	return Todo{}
}

// RepoCreateTodo - Create Todo
func RepoCreateTodo(t Todo) Todo {
	currentID++
	t.ID = currentID
	todos = append(todos, t)
	return t
}

// RepoDestroyTodo - Destroy Todo
func RepoDestroyTodo(id uint64) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find Todo with id of '%d' to delete", id)
}

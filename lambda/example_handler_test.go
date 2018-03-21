package lambda

import (
	"context"
	"fmt"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	var context context.Context = nil
	request := Request{
		Todo: Todo{
			ID:          1,
			Title:       "First Todo",
			Description: "First Todo",
		},
	}

	response, err := HandleRequest(context, request)
	if err != nil {
		t.Fatal("Unexpected error occurred")
	}

	if response.Message != "Saved Todo: 1" {
		t.Fatal(fmt.Sprintf("Unexpected Output: %s", response.Message))
	}
}

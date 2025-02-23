package webapp

import (
	"context"
	"errors"
	"testing"
)

func TestServiceOk(t *testing.T) {
	code := ">,[>,]<[<]>[.>]"
	inputArgs := "hello"

	service := newService()
	output, err := service.runCode(context.TODO(), code, inputArgs)
	if err != nil {
		t.Fatalf("service.runCode(): expected ok got err %v", err)
	}
	if output != "hello" {
		t.Fatalf("service.runCode(): expected 'hello' got '%s'", output)
	}
}

func TestServiceTimeOut(t *testing.T) {
	// infinite loop
	code := "+[+]"
	inputArgs := ""

	service := newService()
	_, err := service.runCode(context.TODO(), code, inputArgs)
	if err == nil {
		t.Fatal("service.runCode(): expected err got ok")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("service.runCode(): expected err context.DeadlineExceeded got %v", err)
	}
}

package main

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"testing"
)

func TestRequestHeaderTableFail(t *testing.T) {
	lmd := createTestLMDInstance()
	buf := bufio.NewReader(bytes.NewBufferString("GET none\n"))
	_, _, err := NewRequest(context.TODO(), lmd, buf, ParseOptimize)
	if err = assertEq(errors.New("bad request: table none does not exist"), err); err != nil {
		t.Fatal(err)
	}
}

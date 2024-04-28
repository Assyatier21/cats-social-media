package helper

import (
	"testing"
	"time"
)

func TestGetContext(t *testing.T) {
	ctx, cancel := GetContext()

	// Verify that the returned context has the expected timeout
	deadline, ok := ctx.Deadline()
	if !ok {
		t.Fatal("Expected context deadline to be set")
	}
	expectedTimeout := time.Second * 60
	marginOfError := time.Millisecond * 100 // Allow a small margin of error
	remainingTime := time.Until(deadline)
	if remainingTime < expectedTimeout-marginOfError {
		t.Errorf("Unexpected context timeout. Expected at least %v, got %v", expectedTimeout, remainingTime)
	}

	// Verify that the cancel function works as expected
	cancel()

	select {
	case <-ctx.Done():
		// The context should be canceled
	case <-time.After(time.Second):
		t.Error("Context cancelation timed out")
	}
}

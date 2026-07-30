package service

import "testing"

type stubHealthService struct{}

func (s stubHealthService) Healthz() string {
	return "ok"
}

type stubFooBarService struct{}

func (s stubFooBarService) FooBar() string {
	return "foo bar"
}

func TestDefaultServiceHealthz(t *testing.T) {
	t.Parallel()

	appService := NewDefaultService(stubHealthService{}, stubFooBarService{})

	if got := appService.Healthz(); got != "ok" {
		t.Fatalf("Healthz() = %q, want %q", got, "ok")
	}
}

func TestDefaultServiceFooBar(t *testing.T) {
	t.Parallel()

	appService := NewDefaultService(stubHealthService{}, stubFooBarService{})

	if got := appService.FooBar(); got != "ok -> foo bar" {
		t.Fatalf("FooBar() = %q, want %q", got, "ok -> foo bar")
	}
}

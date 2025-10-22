//go:build !integration
// +build !integration

package utils

import (
	"fmt"
	"strings"
	"testing"
)

// TestValidateFunctionName tests that only correct function names are accepted
func TestValidateFunctionName(t *testing.T) {
	cases := []struct {
		In    string
		Valid bool
	}{
		{"", false},
		{"*", false},
		{"-", false},
		{"example", true},
		{"example-com", true},
		{"example.com", false},
		{"-example-com", false},
		{"example-com-", false},
		{"Example", false},
		{"EXAMPLE", false},
		{"42", false},
	}

	for _, c := range cases {
		err := ValidateFunctionName(c.In)
		if err != nil && c.Valid {
			t.Fatalf("Unexpected error: %v, for '%v'", err, c.In)
		}
		if err == nil && !c.Valid {
			t.Fatalf("Expected error for invalid entry: %v", c.In)
		}
	}
}

func TestValidateFunctionNameErrMsg(t *testing.T) {
	invalidFnName := "EXAMPLE"
	errMsgPrefix := fmt.Sprintf("Function name '%v'", invalidFnName)

	err := ValidateFunctionName(invalidFnName)
	if err != nil {
		if !strings.HasPrefix(err.Error(), errMsgPrefix) {
			t.Fatalf("Unexpected error message: %v, the message should start with '%v' string", err.Error(), errMsgPrefix)
		}
	} else {
		t.Fatalf("Expected error for invalid entry: %v", invalidFnName)
	}
}

func TestValidateEnvVarName(t *testing.T) {
	cases := []struct {
		In    string
		Valid bool
	}{
		{"", false},
		{"*", false},
		{"example", true},
		{"example-com", true},
		{"example.com", true},
		{"-example-com", true},
		{"example-com-", true},
		{"Example", true},
		{"EXAMPLE", true},
		{";Example", false},
		{":Example", false},
		{",Example", false},
	}

	for _, c := range cases {
		err := ValidateEnvVarName(c.In)
		if err != nil && c.Valid {
			t.Fatalf("Unexpected error: %v, for '%v'", err, c.In)
		}
		if err == nil && !c.Valid {
			t.Fatalf("Expected error for invalid entry: %v", c.In)
		}
	}
}

func TestValidateConfigMapKey(t *testing.T) {
	cases := []struct {
		In    string
		Valid bool
	}{
		{"", false},
		{"*", false},
		{"example", true},
		{"example-com", true},
		{"example.com", true},
		{"-example-com", true},
		{"example-com-", true},
		{"Example", true},
		{"Example_com", true},
		{"Example_com.com", true},
		{"EXAMPLE", true},
		{";Example", false},
		{":Example", false},
		{",Example", false},
	}

	for _, c := range cases {
		err := ValidateConfigMapKey(c.In)
		if err != nil && c.Valid {
			t.Fatalf("Unexpected error: %v, for '%v'", err, c.In)
		}
		if err == nil && !c.Valid {
			t.Fatalf("Expected error for invalid entry: %v", c.In)
		}
	}
}

func TestValidateSecretKey(t *testing.T) {
	cases := []struct {
		In    string
		Valid bool
	}{
		{"", false},
		{"*", false},
		{"example", true},
		{"example-com", true},
		{"example.com", true},
		{"-example-com", true},
		{"example-com-", true},
		{"Example", true},
		{"Example_com", true},
		{"Example_com.com", true},
		{"EXAMPLE", true},
		{";Example", false},
		{":Example", false},
		{",Example", false},
	}

	for _, c := range cases {
		err := ValidateSecretKey(c.In)
		if err != nil && c.Valid {
			t.Fatalf("Unexpected error: %v, for '%v'", err, c.In)
		}
		if err == nil && !c.Valid {
			t.Fatalf("Expected error for invalid entry: %v", c.In)
		}
	}
}

func TestValidateLabelName(t *testing.T) {
	cases := []struct {
		In    string
		Valid bool
	}{
		{"", false},
		{"*", false},
		{"example", true},
		{"example-com", true},
		{"example.com", true},
		{"-example-com", false},
		{"example-com-", false},
		{"Example", true},
		{"EXAMPLE", true},
		{"example.com/example", true},
		{";Example", false},
		{":Example", false},
		{",Example", false},
	}

	for _, c := range cases {
		err := ValidateLabelKey(c.In)
		if err != nil && c.Valid {
			t.Fatalf("Unexpected error: %v, for '%v'", err, c.In)
		}
		if err == nil && !c.Valid {
			t.Fatalf("Expected error for invalid entry: %v", c.In)
		}
	}
}

func TestValidateLabelValue(t *testing.T) {
	cases := []struct {
		In    string
		Valid bool
	}{
		{"", true},
		{"*", false},
		{"example", true},
		{"example-com", true},
		{"example.com", true},
		{"-example-com", false},
		{"example-com-", false},
		{"Example", true},
		{"EXAMPLE", true},
		{"example.com/example", false},
		{";Example", false},
		{":Example", false},
		{",Example", false},
		{"{{env.EXAMPLE}}", true},
	}

	for _, c := range cases {
		err := ValidateLabelValue(c.In)
		if err != nil && c.Valid {
			t.Fatalf("Unexpected error: %v, for '%v'", err, c.In)
		}
		if err == nil && !c.Valid {
			t.Fatalf("Expected error for invalid entry: %v", c.In)
		}
	}
}

// TestValidateNamespace tests that only correct Kubernetes namespace names are accepted
func TestValidateNamespace(t *testing.T) {
	cases := []struct {
		In    string
		Valid bool
	}{
		// Valid namespaces
		{"default", true},
		{"kube-system", true},
		{"my-namespace", true},
		{"myapp", true},
		{"my-app-123", true},
		{"prod", true},
		{"test-123", true},
		{"a", true},
		{"a-b", true},
		{"abc-123-xyz", true},

		// Invalid namespaces
		{"", false},                  // empty
		{"My-App", false},            // uppercase not allowed
		{"MY-APP", false},            // uppercase not allowed
		{"123app", false},            // cannot start with number
		{"123invalid", false},        // cannot start with number
		{"my_app", false},            // underscore not allowed
		{"my app", false},            // spaces not allowed
		{"invalid namespace", false}, // spaces not allowed
		{"my@app", false},            // @ not allowed
		{"invalid@namespace", false}, // @ not allowed
		{"-myapp", false},            // cannot start with hyphen
		{"myapp-", false},            // cannot end with hyphen
		{"my..app", false},           // dots not allowed
		{"my/app", false},            // slash not allowed
		{"my:app", false},            // colon not allowed
		{"my;app", false},            // semicolon not allowed
		{"my,app", false},            // comma not allowed
		{"my*app", false},            // asterisk not allowed
		{"my!app", false},            // exclamation not allowed
	}

	for _, c := range cases {
		err := ValidateNamespace(c.In)
		if err != nil && c.Valid {
			t.Fatalf("Unexpected error for valid namespace: %v, namespace: '%v'", err, c.In)
		}
		if err == nil && !c.Valid {
			t.Fatalf("Expected error for invalid namespace: '%v'", c.In)
		}
	}
}

func TestValidateNamespaceErrMsg(t *testing.T) {
	invalidNamespace := "123invalid"
	errMsgPrefix := fmt.Sprintf("Namespace '%v'", invalidNamespace)

	err := ValidateNamespace(invalidNamespace)
	if err != nil {
		if !strings.HasPrefix(err.Error(), errMsgPrefix) {
			t.Fatalf("Unexpected error message: %v, the message should start with '%v' string", err.Error(), errMsgPrefix)
		}
	} else {
		t.Fatalf("Expected error for invalid namespace: %v", invalidNamespace)
	}
}
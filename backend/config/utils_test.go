package config

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	const key = "TEST_ENV"
	const defaultValue = "default"

	// Ensure the environment variable is not set
	os.Unsetenv(key)
	defer os.Unsetenv(key) // Cleanup after test

	// Test default value
	if got := Env(key, defaultValue); got != defaultValue {
		t.Errorf("Env() = %q, want %q", got, defaultValue)
	}

	// Set environment variable and test override
	expected := "override"
	os.Setenv(key, expected)
	if got := Env(key, defaultValue); got != expected {
		t.Errorf("Env() = %q, want %q", got, expected)
	}
}

func TestEnvBool(t *testing.T) {
	const key = "TEST_ENV_BOOL"

	// Test default value
	os.Unsetenv(key)
	if got := EnvBool(key, true); !got {
		t.Error("EnvBool() = false, want true for default value")
	}

	// Test valid override
	os.Setenv(key, "false")
	if got := EnvBool(key, true); got {
		t.Error("EnvBool() = true, want false for override")
	}

	// Test invalid value
	os.Setenv(key, "not a bool")
	defer func() {
		if r := recover(); r == nil {
			t.Error("EnvBool() did not panic on invalid value")
		}
	}()
	EnvBool(key, true)
}

func TestEnvInt(t *testing.T) {
	const key = "TEST_ENV_INT"

	// Test default value
	os.Unsetenv(key)
	if got := EnvInt(key, 42); got != 42 {
		t.Errorf("EnvInt() = %d, want 42 for default value", got)
	}

	// Test valid override
	os.Setenv(key, "100")
	if got := EnvInt(key, 42); got != 100 {
		t.Errorf("EnvInt() = %d, want 100 for override", got)
	}

	// Test invalid value
	os.Setenv(key, "not an int")
	defer func() {
		if r := recover(); r == nil {
			t.Error("EnvInt() did not panic on invalid value")
		}
	}()
	EnvInt(key, 42)
}

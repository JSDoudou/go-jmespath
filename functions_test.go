package jmespath

import "testing"

func TestJoinMixedArray(t *testing.T) {
	input := map[string]interface{}{
		"items": []interface{}{123.0, "abc", true, nil, 12345678},
	}

	result, err := Search(`join('-', items)`, input)
	if err != nil {
		t.Fatalf("join failed: %v", err)
	}

	expected := "123-abc-true-<nil>-12345678"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

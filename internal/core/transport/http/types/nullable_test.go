package core_http_types

import (
	"encoding/json"
	"testing"
)

func TestNullableUnmarshalJSONFromString(t *testing.T) {
	type payload struct {
		FullName Nullable[string] `json:"full_name"`
	}

	var got payload
	if err := json.Unmarshal([]byte(`{"full_name":"Alice"}`), &got); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}

	if !got.FullName.Set {
		t.Fatal("expected nullable field to be marked as set")
	}

	if got.FullName.Value == nil {
		t.Fatal("expected nullable field value to be populated")
	}

	if *got.FullName.Value != "Alice" {
		t.Fatalf("expected value Alice, got %v", *got.FullName.Value)
	}
}

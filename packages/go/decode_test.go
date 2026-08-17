package artifact

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestDecodeSourcePreservesUnknownFieldsAndOpaqueProviderReference(t *testing.T) {
	raw := json.RawMessage(`{"type":"provider","provider":"forge","reference":{"resource":"change","future":true},"futureField":{"x":1}}`)
	node, err := DecodeSource(raw); if err != nil { t.Fatal(err) }
	provider, ok := node.(WireProviderArtifactSource); if !ok { t.Fatalf("unexpected source %T", node) }
	if string(provider.Reference) != `{"resource":"change","future":true}` { t.Fatalf("opaque reference changed: %s", provider.Reference) }
	encoded, err := json.Marshal(provider); if err != nil { t.Fatal(err) }
	var roundTrip map[string]any; if err := json.Unmarshal(encoded, &roundTrip); err != nil { t.Fatal(err) }
	if _, ok := roundTrip["futureField"]; !ok { t.Fatal("unknown field was not preserved") }
}

func TestDecodeRecursiveConditionAndValueSchema(t *testing.T) {
	condition, err := DecodeCondition(json.RawMessage(`{"kind":"and","conditions":[{"kind":"actor","in":["reviewer"]}]}`)); if err != nil { t.Fatal(err) }
	if _, ok := condition.(WireArtifactConditionAnd); !ok { t.Fatalf("unexpected condition %T", condition) }
	value, err := DecodeValueSchema(json.RawMessage(`{"valueType":"collection","itemSchema":{"valueType":"text"}}`)); if err != nil { t.Fatal(err) }
	if _, ok := value.(WireArtifactValueSchemaCollection); !ok { t.Fatalf("unexpected value schema %T", value) }
}

func TestDecodeRejectsUnknownClosedDiscriminator(t *testing.T) {
	_, err := DecodeSource(json.RawMessage(`{"type":"future"}`)); var unsupported *UnsupportedDiscriminatorError
	if !errors.As(err, &unsupported) { t.Fatalf("expected UnsupportedDiscriminatorError, got %v", err) }
}

func TestGeneratedWireRecordPreservesMissingVersusNull(t *testing.T) {
	base := `{"schemaVersion":"1.1","id":"sub_1","artifactId":"art_1","submittedBy":{"type":"system"},"submittedAt":"2026-08-15T09:00:00Z"}`
	var missing WireArtifactSubmission; if err := json.Unmarshal([]byte(base), &missing); err != nil { t.Fatal(err) }
	if missing.Value != nil { t.Fatalf("missing value became present: %s", missing.Value) }
	withNull := base[:len(base)-1] + `,"value":null}`
	var explicitNull WireArtifactSubmission; if err := json.Unmarshal([]byte(withNull), &explicitNull); err != nil { t.Fatal(err) }
	if string(explicitNull.Value) != "null" { t.Fatalf("explicit null was not preserved: %s", explicitNull.Value) }
	encoded, err := json.Marshal(explicitNull); if err != nil { t.Fatal(err) }
	var fields map[string]json.RawMessage; if err := json.Unmarshal(encoded, &fields); err != nil { t.Fatal(err) }
	if string(fields["value"]) != "null" { t.Fatal("explicit null disappeared during round trip") }
}

func TestUnknownFieldsCannotOverrideKnownWireFields(t *testing.T) {
	value := WireURLArtifactSource{Type: "url", URL: "https://example.test", UnknownFields: map[string]json.RawMessage{
		"url": json.RawMessage(`"https://evil.test"`), "mediaType": json.RawMessage(`"invalid/override"`), "future": json.RawMessage(`true`),
	}}
	encoded, err := json.Marshal(value); if err != nil { t.Fatal(err) }
	var fields map[string]json.RawMessage; if err := json.Unmarshal(encoded, &fields); err != nil { t.Fatal(err) }
	if string(fields["url"]) != `"https://example.test"` { t.Fatal("unknown field replaced a known field") }
	if _, exists := fields["mediaType"]; exists { t.Fatal("unknown field populated an omitted known field") }
	if string(fields["future"]) != "true" { t.Fatal("additive unknown field was not preserved") }
}

package artifact

import (
	"encoding/json"
	"fmt"
)

// UnsupportedDiscriminatorError reports a closed protocol discriminator that this binding cannot decode.
type UnsupportedDiscriminatorError struct { Union, Discriminator string }
func (err *UnsupportedDiscriminatorError) Error() string { return fmt.Sprintf("unsupported %s discriminator %q", err.Union, err.Discriminator) }

type ArtifactSourceNode interface { artifactSourceNode() }
func (WireInlineArtifactSource) artifactSourceNode() {}
func (WireLocalArtifactSource) artifactSourceNode() {}
func (WireObjectArtifactSource) artifactSourceNode() {}
func (WireURLArtifactSource) artifactSourceNode() {}
func (WireHostedArtifactSource) artifactSourceNode() {}
func (WireProviderArtifactSource) artifactSourceNode() {}

type ConditionNode interface { conditionNode() }
func (WireArtifactConditionState) conditionNode() {}
func (WireArtifactConditionActor) conditionNode() {}
func (WireArtifactConditionArtifactExists) conditionNode() {}
func (WireArtifactConditionArtifactValue) conditionNode() {}
func (WireArtifactConditionAnd) conditionNode() {}
func (WireArtifactConditionOr) conditionNode() {}
func (WireArtifactConditionNot) conditionNode() {}

type ValueSchemaNode interface { valueSchemaNode() }
func (WireArtifactValueSchemaText) valueSchemaNode() {}
func (WireArtifactValueSchemaNumber) valueSchemaNode() {}
func (WireArtifactValueSchemaBoolean) valueSchemaNode() {}
func (WireArtifactValueSchemaCurrency) valueSchemaNode() {}
func (WireArtifactValueSchemaDate) valueSchemaNode() {}
func (WireArtifactValueSchemaDatetime) valueSchemaNode() {}
func (WireArtifactValueSchemaTime) valueSchemaNode() {}
func (WireArtifactValueSchemaLocation) valueSchemaNode() {}
func (WireArtifactValueSchemaFile) valueSchemaNode() {}
func (WireArtifactValueSchemaImage) valueSchemaNode() {}
func (WireArtifactValueSchemaVideo) valueSchemaNode() {}
func (WireArtifactValueSchemaAudio) valueSchemaNode() {}
func (WireArtifactValueSchemaLink) valueSchemaNode() {}
func (WireArtifactValueSchemaStructured) valueSchemaNode() {}
func (WireArtifactValueSchemaReference) valueSchemaNode() {}
func (WireArtifactValueSchemaSignature) valueSchemaNode() {}
func (WireArtifactValueSchemaCollection) valueSchemaNode() {}

func discriminator(raw json.RawMessage, key string) (string, error) {
	var envelope map[string]json.RawMessage
	if err := json.Unmarshal(raw, &envelope); err != nil { return "", err }
	value, ok := envelope[key]; if !ok { return "", fmt.Errorf("missing discriminator %q", key) }
	var result string; if err := json.Unmarshal(value, &result); err != nil { return "", err }
	return result, nil
}

func decodeNode[T any](raw json.RawMessage) (T, error) { var value T; err := json.Unmarshal(raw, &value); return value, err }

func DecodeSource(raw json.RawMessage) (ArtifactSourceNode, error) {
	kind, err := discriminator(raw, "type"); if err != nil { return nil, err }
	switch kind {
	case "inline": return decodeNode[WireInlineArtifactSource](raw)
	case "local": return decodeNode[WireLocalArtifactSource](raw)
	case "object": return decodeNode[WireObjectArtifactSource](raw)
	case "url": return decodeNode[WireURLArtifactSource](raw)
	case "hosted": return decodeNode[WireHostedArtifactSource](raw)
	case "provider": return decodeNode[WireProviderArtifactSource](raw)
	default: return nil, &UnsupportedDiscriminatorError{Union: "ArtifactSource", Discriminator: kind}
	}
}

func DecodeCondition(raw json.RawMessage) (ConditionNode, error) {
	kind, err := discriminator(raw, "kind"); if err != nil { return nil, err }
	switch kind {
	case "state": return decodeNode[WireArtifactConditionState](raw)
	case "actor": return decodeNode[WireArtifactConditionActor](raw)
	case "artifact_exists": return decodeNode[WireArtifactConditionArtifactExists](raw)
	case "artifact_value": return decodeNode[WireArtifactConditionArtifactValue](raw)
	case "and": return decodeNode[WireArtifactConditionAnd](raw)
	case "or": return decodeNode[WireArtifactConditionOr](raw)
	case "not": return decodeNode[WireArtifactConditionNot](raw)
	default: return nil, &UnsupportedDiscriminatorError{Union: "ArtifactCondition", Discriminator: kind}
	}
}

func DecodeValueSchema(raw json.RawMessage) (ValueSchemaNode, error) {
	kind, err := discriminator(raw, "valueType"); if err != nil { return nil, err }
	switch kind {
	case "text": return decodeNode[WireArtifactValueSchemaText](raw)
	case "number": return decodeNode[WireArtifactValueSchemaNumber](raw)
	case "boolean": return decodeNode[WireArtifactValueSchemaBoolean](raw)
	case "currency": return decodeNode[WireArtifactValueSchemaCurrency](raw)
	case "date": return decodeNode[WireArtifactValueSchemaDate](raw)
	case "datetime": return decodeNode[WireArtifactValueSchemaDatetime](raw)
	case "time": return decodeNode[WireArtifactValueSchemaTime](raw)
	case "location": return decodeNode[WireArtifactValueSchemaLocation](raw)
	case "file": return decodeNode[WireArtifactValueSchemaFile](raw)
	case "image": return decodeNode[WireArtifactValueSchemaImage](raw)
	case "video": return decodeNode[WireArtifactValueSchemaVideo](raw)
	case "audio": return decodeNode[WireArtifactValueSchemaAudio](raw)
	case "link": return decodeNode[WireArtifactValueSchemaLink](raw)
	case "structured": return decodeNode[WireArtifactValueSchemaStructured](raw)
	case "reference": return decodeNode[WireArtifactValueSchemaReference](raw)
	case "signature": return decodeNode[WireArtifactValueSchemaSignature](raw)
	case "collection": return decodeNode[WireArtifactValueSchemaCollection](raw)
	default: return nil, &UnsupportedDiscriminatorError{Union: "ArtifactValueSchema", Discriminator: kind}
	}
}

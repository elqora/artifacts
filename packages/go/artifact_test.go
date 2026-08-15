package artifact

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func fixtureRoot(t *testing.T) string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot locate test file")
	}
	return filepath.Join(filepath.Dir(file), "..", "..", "tests", "fixtures")
}

func TestCanonicalRecordsRoundTrip(t *testing.T) {
	root := fixtureRoot(t)
	for _, directory := range []string{"artifact", "version", "link", "specification", "runtime", "condition", "policy", "value-schema", "provider"} {
		err := filepath.Walk(filepath.Join(root, directory), func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() || filepath.Ext(path) != ".json" {
				return err
			}
			original, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			var decoded map[string]any
			if err = json.Unmarshal(original, &decoded); err != nil {
				return err
			}
			if version, present := decoded["schemaVersion"]; present && version != ProtocolVersion {
				t.Fatalf("%s has unsupported version", path)
			}
			encoded, err := json.Marshal(decoded)
			if err != nil {
				return err
			}
			var roundTrip map[string]any
			if err = json.Unmarshal(encoded, &roundTrip); err != nil {
				return err
			}
			if !reflect.DeepEqual(decoded, roundTrip) {
				t.Fatalf("semantic round trip changed %s", path)
			}
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
	}
}

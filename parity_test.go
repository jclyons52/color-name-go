package colorname

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"testing"
)

// TestParity runs the Go port against the real npm color-name (vendored at
// original/index.js) and requires the full map to match exactly. Skips when
// node is unavailable or the vendored original is missing.
func TestParity(t *testing.T) {
	nodeBin, err := exec.LookPath("node")
	if err != nil {
		t.Skip("node not available; skipping JS parity")
	}
	if _, err := os.Stat(filepath.Join("original", "index.js")); err != nil {
		t.Fatalf("original color-name not found at original/index.js: %v", err)
	}

	dir := t.TempDir()
	// Node prints the module.exports object as parsed JSON; we pipe it out.
	driver := `'use strict'; const c=require(process.env.COLORNAME_ORIG); process.stdout.write(JSON.stringify(c));`
	cmd := exec.Command(nodeBin, "-e", driver)
	cmd.Dir = dir
	origAbs, _ := filepath.Abs("original")
	cmd.Env = append(os.Environ(), "COLORNAME_ORIG="+origAbs)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("node driver failed: %v\n%s", err, out)
	}
	var jsNames map[string][]int
	if err := json.Unmarshal(out, &jsNames); err != nil {
		t.Fatalf("parse node output: %v\n%s", err, out)
	}

	if len(jsNames) != len(Names) {
		t.Errorf("entry count differs: node=%d go=%d", len(jsNames), len(Names))
	}

	var jsKeys []string
	for k := range jsNames {
		jsKeys = append(jsKeys, k)
	}
	sort.Strings(jsKeys)

	for _, k := range jsKeys {
		want := jsNames[k]
		got, ok := Names[k]
		if !ok {
			t.Errorf("missing in Go: %q (node=%v)", k, want)
			continue
		}
		if len(want) != 3 || got.R != uint8(want[0]) || got.G != uint8(want[1]) || got.B != uint8(want[2]) {
			t.Errorf("mismatch %q: go=%v node=%v", k, got, want)
		}
	}
	// Reverse check: every Go key must exist in node.
	for k := range Names {
		if _, ok := jsNames[k]; !ok {
			t.Errorf("extra in Go: %q", k)
		}
	}
	t.Logf("parity passed: %d color names match the JS original", len(jsNames))
}

// TestGetExercise covers the Get helper across hit/miss and spot values.
func TestGetExercise(t *testing.T) {
	if v, ok := Get("rebeccapurple"); !ok || v != (RGB{102, 51, 153}) {
		t.Fatalf("rebeccapurple = %v, %v", v, ok)
	}
	if _, ok := Get("notacolor"); ok {
		t.Fatalf("expected miss for notacolor")
	}
}

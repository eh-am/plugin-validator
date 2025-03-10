package htmlreadme

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/grafana/plugin-validator/pkg/analysis"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/readme"
)

func TestClean(t *testing.T) {
	var invoked bool

	b, err := os.ReadFile(filepath.Join("testdata", "README.clean.md"))
	if err != nil {
		t.Fatal(err)
	}

	pass := &analysis.Pass{
		ResultOf: map[*analysis.Analyzer]interface{}{
			readme.Analyzer: b,
		},
		Report: func(n string, d analysis.Diagnostic) {
			invoked = true
		},
	}

	_, err = Analyzer.Run(pass)
	if err != nil {
		t.Fatal(err)
	}

	if invoked {
		t.Error("unexpected report")
	}
}

func TestHTML(t *testing.T) {
	var invoked bool

	b, err := os.ReadFile(filepath.Join("testdata", "README.html.md"))
	if err != nil {
		t.Fatal(err)
	}

	pass := &analysis.Pass{
		ResultOf: map[*analysis.Analyzer]interface{}{
			readme.Analyzer: b,
		},
		Report: func(n string, d analysis.Diagnostic) {
			if d.Title != "README.md: HTML is not supported and will not render correctly" {
				t.Errorf("unexpected diagnostic message: %q", d.Title)
			}
			invoked = true
		},
	}

	res, err := Analyzer.Run(pass)
	if err != nil {
		t.Fatal(err)
	}

	if !invoked {
		t.Error("expected report, but got none")
	}

	if res != nil {
		t.Fatalf("unexpected result: %v", res)
	}
}

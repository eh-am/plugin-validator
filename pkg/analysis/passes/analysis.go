package passes

import (
	"github.com/grafana/plugin-validator/pkg/analysis"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/archive"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/archivename"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/binarypermissions"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/brokenlinks"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/gomanifest"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/gosec"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/htmlreadme"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/jargon"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/jssourcemap"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/legacyplatform"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/license"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/logos"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/manifest"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/metadata"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/metadatapaths"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/metadatavalid"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/modulejs"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/org"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/pluginname"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/published"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/readme"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/restrictivedep"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/screenshots"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/signature"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/sourcecode"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/templatereadme"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/trackingscripts"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/typesuffix"
	"github.com/grafana/plugin-validator/pkg/analysis/passes/version"
)

var Analyzers = []*analysis.Analyzer{
	archive.Analyzer,
	archivename.Analyzer,
	brokenlinks.Analyzer,
	binarypermissions.Analyzer,
	gosec.Analyzer,
	gomanifest.Analyzer,
	htmlreadme.Analyzer,
	jargon.Analyzer,
	jssourcemap.Analyzer,
	legacyplatform.Analyzer,
	logos.Analyzer,
	license.Analyzer,
	manifest.Analyzer,
	metadata.Analyzer,
	metadatapaths.Analyzer,
	metadatavalid.Analyzer,
	modulejs.Analyzer,
	org.Analyzer,
	pluginname.Analyzer,
	published.Analyzer,
	readme.Analyzer,
	restrictivedep.Analyzer,
	screenshots.Analyzer,
	signature.Analyzer,
	sourcecode.Analyzer,
	templatereadme.Analyzer,
	trackingscripts.Analyzer,
	typesuffix.Analyzer,
	version.Analyzer,
}

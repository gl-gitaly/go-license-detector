package internal

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGPL2WithoutApplicationGuide(t *testing.T) {
	// Unmodified GPLv2 license from ska-telescope/external/ducc at
	// c1ed01b8fda69d2a378190f0915f1a19762dddbf. Its operative terms are
	// complete, but it omits the optional guide for applying the license.
	text, err := os.ReadFile("testdata/gpl-2.0-without-application-guide.txt")
	require.NoError(t, err)
	matches := InvestigateLicenseText(text)
	require.GreaterOrEqual(t, matches["GPL-2.0-only"], float32(0.98))
	for id, confidence := range matches {
		if strings.Contains(id, "AGPL-1.0") {
			require.Greater(t, matches["GPL-2.0-only"], confidence, id)
		}
	}
}

func TestGPL2ApplicationGuidePreservesCompleteLicenses(t *testing.T) {
	db := globalLicenseDatabase()
	for _, id := range []string{"GPL-2.0-only", "AGPL-1.0-only", "MIT", "BSD-3-Clause"} {
		t.Run(id, func(t *testing.T) {
			matches := InvestigateLicenseText([]byte(db.licenseTexts[id]))
			require.Equal(t, float32(1), matches[id])
		})
	}
}

package clavis_test

import (
	"testing"

	"github.com/fiuskylab/clavis"
)

func TestSha1(t *testing.T) {
	key := "key"
	want := "a62f2225bf70bfaccbc7f1ef2a397836717377de"

	got := clavis.Sha1Encrypt(key)

	t.Run("Encrypt Key", func(t *testing.T) {
		if got != want {
			t.Errorf("Want: %s\n Got: %s", want, got)
		}
	})
}

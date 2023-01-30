package salary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHitungGaji(t *testing.T) {
	t.Run("Test Hitung gaji manager", func(t *testing.T) {
		posisi := "manager"
		expected := 30000000
		actual := HitungGaji(posisi)
		if actual != expected {
			t.Error("hasil tidak sesuai. actual:", actual, "expected:", expected)
		}
	})

	t.Run("Test Hitung gaji senior", func(t *testing.T) {
		posisi := "senior"
		expected := 20000000
		actual := HitungGaji(posisi)
		assert.Equal(t, expected, actual, "hasil tidak sesuai")
	})

	t.Run("Test Hitung gaji junior", func(t *testing.T) {
		posisi := "junior"
		expected := 10000000
		actual := HitungGaji(posisi)
		assert.Equal(t, expected, actual, "hasil tidak sesuai")
	})

	t.Run("Test Hitung gaji di kondisi else", func(t *testing.T) {
		posisi := "magang"
		expected := 0
		actual := HitungGaji(posisi)
		assert.Equal(t, expected, actual, "hasil tidak sesuai")
	})
}

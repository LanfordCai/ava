package base58check

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {

	encoder := BitcoinEncoder

	invalidCases := []string{
		"1F75dae041293f43df7eeab162dbe925fc875b76d792904cce51343120a039d2e57",
		"2SLgsEfKy_STAk8yak8e3B4iwmyBG3CNJKkykdVUuHaDtZt4G8RZidk2cCoG4dtQyYNZX1XQX8zoWz82mUrQG9mPbWU6fQt1VoY6tu2cWydfazaFUucaxiJjw23LXL9t8yKP",
		"6aRhzqrfjURatgH5xm3/45ZbXU5i2wuyhwueNfpih4LB1ZA9UxLkZfF9DsirGc4Ch9tifdNxg2T1g",
	}

	for _, c := range invalidCases {
		assert.Equal(t, "", string(encoder.Decode(c)))
	}
}

func TestCheckEncode_BitcoinAlphabet(t *testing.T) {

	encoder := BitcoinEncoder

	payloads := [][]byte{
		[]byte(""),
		[]byte("hello1234"),
		[]byte("hello world!"),
		[]byte("1234567890!@#$%^&*()"),
		[]byte("how many roads must a man walk down"),
	}

	cases := []struct {
		checksumType ChecksumType
		payloads     [][]byte
		expected     []string
	}{
		{
			checksumType: ChecksumDoubleSha256,
			payloads:     payloads,
			expected: []string{
				"3QJmnh",
				"9hLF3DC7uYXfUSogYr",
				"DthHcFYf2SzzprfE7TynLJ",
				"5V8FMkzy9QZKJAeAzg572WykUvAm68zn6",
				"2BBvwFYwFkHUYGfosyVEwkxfTLXsEJq9waQM9J6Bm7cxBaDK9QxkPQ",
			},
		},
		{
			checksumType: ChecksumSha256,
			payloads:     payloads,
			expected: []string{
				"6pZZfP",
				"9hLF3DC7uYXfWL13c1",
				"DthHcFYf2SzzprfE9y2n3E",
				"5V8FMkzy9QZKJAeAzg572WykUvAnBmXET",
				"2BBvwFYwFkHUYGfosyVEwkxfTLXsEJq9waQM9J6Bm7cxBaDK4xnzxb",
			},
		},
		{
			checksumType: ChecksumRipemd160,
			payloads:     payloads,
			expected: []string{
				"4zNxKW",
				"9hLF3DC7uYXfWb3zD3",
				"DthHcFYf2SzzprfEChb8AW",
				"5V8FMkzy9QZKJAeAzg572WykUvAmhxCeu",
				"2BBvwFYwFkHUYGfosyVEwkxfTLXsEJq9waQM9J6Bm7cxBaDK3Wri6U",
			},
		},
		{
			checksumType: ChecksumBlake2bKeccak256,
			payloads:     payloads,
			expected: []string{
				"5ingcK",
				"9hLF3DC7uYXfV2Vgf6",
				"DthHcFYf2SzzprfE8cT9Vj",
				"5V8FMkzy9QZKJAeAzg572WykUvAiiPPkH",
				"2BBvwFYwFkHUYGfosyVEwkxfTLXsEJq9waQM9J6Bm7cxBaDK3sz8M9",
			},
		},
	}

	for _, c := range cases {
		encoder.ChecksumType = c.checksumType
		encoded := []string{}
		for _, p := range c.payloads {
			r := encoder.CheckEncode(p)
			encoded = append(encoded, r)
		}
		assert.True(t, reflect.DeepEqual(c.expected, encoded), fmt.Sprint(encoded))
	}
}

func TestCheckDecode_BitcoinAlphabet(t *testing.T) {

	encoder := BitcoinEncoder

	expected := [][]byte{
		[]byte(""),
		[]byte("hello1234"),
		[]byte("hello world!"),
		[]byte("1234567890!@#$%^&*()"),
		[]byte("how many roads must a man walk down"),
	}

	cases := []struct {
		checksumType ChecksumType
		encoded      []string
		expected     [][]byte
	}{
		{
			checksumType: ChecksumDoubleSha256,
			encoded: []string{
				"3QJmnh",
				"9hLF3DC7uYXfUSogYr",
				"DthHcFYf2SzzprfE7TynLJ",
				"5V8FMkzy9QZKJAeAzg572WykUvAm68zn6",
				"2BBvwFYwFkHUYGfosyVEwkxfTLXsEJq9waQM9J6Bm7cxBaDK9QxkPQ",
			},
			expected: expected,
		},
		{
			checksumType: ChecksumSha256,
			encoded: []string{
				"6pZZfP",
				"9hLF3DC7uYXfWL13c1",
				"DthHcFYf2SzzprfE9y2n3E",
				"5V8FMkzy9QZKJAeAzg572WykUvAnBmXET",
				"2BBvwFYwFkHUYGfosyVEwkxfTLXsEJq9waQM9J6Bm7cxBaDK4xnzxb",
			},
			expected: expected,
		},
		{
			checksumType: ChecksumRipemd160,
			encoded: []string{
				"4zNxKW",
				"9hLF3DC7uYXfWb3zD3",
				"DthHcFYf2SzzprfEChb8AW",
				"5V8FMkzy9QZKJAeAzg572WykUvAmhxCeu",
				"2BBvwFYwFkHUYGfosyVEwkxfTLXsEJq9waQM9J6Bm7cxBaDK3Wri6U",
			},
			expected: expected,
		},
		{
			checksumType: ChecksumBlake2bKeccak256,
			encoded: []string{
				"5ingcK",
				"9hLF3DC7uYXfV2Vgf6",
				"DthHcFYf2SzzprfE8cT9Vj",
				"5V8FMkzy9QZKJAeAzg572WykUvAiiPPkH",
				"2BBvwFYwFkHUYGfosyVEwkxfTLXsEJq9waQM9J6Bm7cxBaDK3sz8M9",
			},
			expected: expected,
		},
	}

	for _, c := range cases {
		encoder.ChecksumType = c.checksumType
		decoded := [][]byte{}
		for _, p := range c.encoded {
			r, err := encoder.CheckDecode(p)
			assert.NoError(t, err)
			decoded = append(decoded, r)
		}
		assert.True(t, reflect.DeepEqual(c.expected, decoded), fmt.Sprint(decoded))
	}
}

func TestCheckEncode_RippleAlphabet(t *testing.T) {

	encoder := RippleEncoder

	payloads := [][]byte{
		[]byte(""),
		[]byte("hello1234"),
		[]byte("hello world!"),
		[]byte("1234567890!@#$%^&*()"),
		[]byte("how many roads must a man walk down"),
	}

	cases := []struct {
		checksumType ChecksumType
		payloads     [][]byte
		expected     []string
	}{
		{
			checksumType: ChecksumDoubleSha256,
			payloads:     payloads,
			expected: []string{
				"sQJm86",
				"96LEsDUfuYXC7SogYi",
				"Dt6HcEYCpSzzFiCNfTy8LJ",
				"nV3EMkzy9QZKJwewzgnfpWyk7vwma3z8a",
				"pBBvAEYAEkH7YGCo1yVNAkxCTLX1NJq9A2QM9JaBmfcxB2DK9QxkPQ",
			},
		},
	}

	for _, c := range cases {
		encoder.ChecksumType = c.checksumType
		encoded := []string{}
		for _, p := range c.payloads {
			r := encoder.CheckEncode(p)
			encoded = append(encoded, r)
		}
		assert.True(t, reflect.DeepEqual(c.expected, encoded), fmt.Sprint(encoded))
	}
}

func TestCheckDecode_RippleAlphabet(t *testing.T) {

	encoder := RippleEncoder

	expected := [][]byte{
		[]byte(""),
		[]byte("hello1234"),
		[]byte("hello world!"),
		[]byte("1234567890!@#$%^&*()"),
		[]byte("how many roads must a man walk down"),
	}

	cases := []struct {
		checksumType ChecksumType
		encoded      []string
		expected     [][]byte
	}{
		{
			checksumType: ChecksumDoubleSha256,
			encoded: []string{
				"sQJm86",
				"96LEsDUfuYXC7SogYi",
				"Dt6HcEYCpSzzFiCNfTy8LJ",
				"nV3EMkzy9QZKJwewzgnfpWyk7vwma3z8a",
				"pBBvAEYAEkH7YGCo1yVNAkxCTLX1NJq9A2QM9JaBmfcxB2DK9QxkPQ",
			},
			expected: expected,
		},
	}

	for _, c := range cases {
		encoder.ChecksumType = c.checksumType
		decoded := [][]byte{}
		for _, p := range c.encoded {
			r, err := encoder.CheckDecode(p)
			assert.NoError(t, err)
			decoded = append(decoded, r)
		}
		assert.True(t, reflect.DeepEqual(c.expected, decoded), fmt.Sprint(decoded))
	}
}

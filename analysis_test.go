package cryptanalysis

import (
	//"bytes"
	"testing"
)

type ecbscore struct {
	data  []byte
	size  int
	score float64
}

func round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}

func TestScoreEnglish(t *testing.T) {
	score := 35
	s := "Now that the party is jumping\n"

	eng_score := ScoreEnglish(s)

	if round(eng_score) != score {
		t.Error("Expected", score, "got", eng_score)
	}
}

func TestHamming(t *testing.T) {
	distance := 37
	s1 := []byte("this is a test")
	s2 := []byte("wokka wokka!!!")

	ham_dist, err := Hamming(s1, s2)

	if err != nil {
		t.Error(err)
	}

	if distance != ham_dist {
		t.Error("Expected", distance, "got", ham_dist)
	}

}

func TestKeyLength(t *testing.T) {
    length := 3
	encrypted := []byte{0x0f, 0xe8, 0xc6, 0x1f, 0xee, 0xf0, 0x9c, 0x23, 0x53,
        0xd6, 0x83, 0x12, 0x63, 0x9f, 0xf0, 0x75, 0xb0, 0xcc, 0x1f, 0x91, 0x98,
        0x09, 0x07, 0x35, 0x22, 0x92, 0x03, 0xcd, 0xa3, 0x39, 0x89, 0x34, 0x27,
        0x40, 0x13, 0x81, 0xd3, 0x17, 0x30, 0x97, 0x72, 0x05, 0xe8, 0x0e, 0x19,
        0xa5, 0x3b, 0x23, 0xf3, 0xfc, 0x09, 0x01, 0x1b, 0x44, 0x3b, 0xac, 0x7b,
        0x60, 0x0c, 0x71, 0x2c, 0xea, 0xad, 0xb7, 0x7a, 0x0d, 0x93, 0x73, 0x94,
        0xdf, 0x8f, 0x4a, 0xba, 0x48, 0xb8, 0x7a, 0x85, 0x44, 0x0f, 0x6a, 0x6c,
        0x34, 0x1f, 0xb5, 0xc9, 0x58, 0xce, 0x4d, 0x69, 0x21, 0xa8, 0x83, 0x5c,
        0xcf, 0x0d, 0x99, 0x8c, 0x8f, 0xe4, 0xf5, 0xf4, 0xba, 0xc0, 0x85, 0xf3,
        0xed, 0x1c, 0x0a, 0x75, 0x8f, 0xb8, 0x00, 0xcf, 0x18, 0x93, 0x1c, 0xb4,
        0x4f, 0x7a, 0x06, 0x9f, 0xcc, 0x6a, 0xf7, 0xb3, 0xc1, 0x94, 0xa4, 0xf4,
        0x12, 0xb2, 0xca, 0x90, 0xaa, 0x1a, 0x6c, 0xbb, 0xdc, 0x24, 0x3d, 0x76,
        0x33, 0x09, 0x7c, 0x99, 0x8e, 0x5d, 0xbb, 0xce, 0xca, 0x62, 0x82, 0x45,
        0x50, 0x78, 0xe9, 0x05, 0xe1, 0xff, 0xde}

    encrypted = PadPkcs7(encrypted, 16)
	key_length, err := KeyLength(encrypted)

	if err != nil {
		t.Error(err)
	}

	if key_length != length {
		t.Error("Expected", length, "got", key_length)
	}

}

func TestScoreEcb(t *testing.T) {
	ecbscores := []ecbscore{
		{[]byte("abcdefghijklmnopqrst"), 5, 1.0},
		{[]byte("abcdeabcdeabcdeabcde"), 5, 0.25},
		{[]byte("abcdeabcdeabcdeedcba"), 5, 0.5},
	}

	for _, test := range ecbscores {
		score := ScoreEcb(test.data, test.size)
		if score != test.score {
			t.Error("Expected", test.score, "got", score)
		}
	}
}

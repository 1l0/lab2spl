package model

/*
	reference:
	https://github.com/VOICEVOX/voicevox_engine/blob/master/voicevox_engine/tts_pipeline/phoneme.py
*/

const (
	// Vowels
	Pau = "pau"
	NN  = "N"
	A   = "a"
	Cl  = "cl"
	E   = "e"
	I   = "i"
	O   = "o"
	U   = "u"
	AA  = "A"
	EE  = "E"
	II  = "I"
	OO  = "O"
	UU  = "U"

	// Consonants
	B  = "b"
	By = "by"
	Ch = "ch"
	D  = "d"
	Dy = "dy"
	F  = "f"
	G  = "g"
	Gw = "gw"
	Gy = "gy"
	H  = "h"
	Hy = "hy"
	J  = "j"
	K  = "k"
	Kw = "kw"
	Ky = "ky"
	M  = "m"
	My = "my"
	N  = "n"
	Ny = "ny"
	P  = "p"
	Py = "py"
	R  = "r"
	Ry = "ry"
	S  = "s"
	Sh = "sh"
	T  = "t"
	Ts = "ts"
	Ty = "ty"
	V  = "v"
	W  = "w"
	Y  = "y"
	Z  = "z"
)

var Phonemes = []string{
	Pau,
	NN,
	A,
	Cl,
	E,
	I,
	O,
	U,
	AA,
	EE,
	II,
	OO,
	UU,
	B,
	By,
	Ch,
	D,
	Dy,
	F,
	G,
	Gw,
	Gy,
	H,
	Hy,
	J,
	K,
	Kw,
	Ky,
	M,
	My,
	N,
	Ny,
	P,
	Py,
	R,
	Ry,
	S,
	Sh,
	T,
	Ts,
	Ty,
	V,
	W,
	Y,
	Z,
}

var PhonemeIndex map[string]int

func init() {
	PhonemeIndex = make(map[string]int, len(Phonemes))
	for i, p := range Phonemes {
		PhonemeIndex[p] = i
	}
}

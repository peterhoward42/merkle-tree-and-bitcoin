package hash

import (
	"fmt"
	"testing"
)

func TestHashWrapperFunctionWithKnownTestVector(t *testing.T) {
	// Source of test vectors:
	// http://www.di-mgt.com.au/sha_testvectors.html#FIPS-180
	testInput := []byte("abc") // Yes. That is official test vector.
	expectedOutput :=
		"ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"

	res := fmt.Sprintf("%0x", Hash(testInput))
	if res != expectedOutput {
		t.Errorf(
			"\nResponse:\n%s\ndiffers from expected:\n%s", res, expectedOutput)
	}
}

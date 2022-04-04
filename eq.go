package tu

import (
	"fmt"
	"testing"

	"github.com/reusee/sb"
)

func Equal(t *testing.T, pairs ...any) {
	for i := 0; i < len(pairs); i += 2 {
		var tokensA, tokensB sb.Tokens
		if err := sb.Copy(
			sb.Marshal(pairs[i]),
			sb.CollectTokens(&tokensA),
		); err != nil {
			t.Fatal(err)
		}
		if err := sb.Copy(
			sb.Marshal(pairs[i+1]),
			sb.CollectTokens(&tokensB),
		); err != nil {
			t.Fatal(err)
		}
		if res := sb.MustCompare(tokensA.Iter(), tokensB.Iter()); res != 0 {
			fmt.Printf("A: %+v\n", tokensA)
			fmt.Printf("B: %+v\n", tokensB)
			t.Fatalf("not equal, pair %d", i)
		}
	}
}

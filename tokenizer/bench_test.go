package tokenizer

import (
	"testing"
)

var text = `
TO THE RED-HEADED LEAGUE: On account of the bequest of the late
Ezekiah Hopkins, of Lebanon, Pennsylvania, U.S.A., there is now another
vacancy open which entitles a member of the League to a salary of £ 4 a
week for purely nominal services. All red-headed men who are sound in
body and mind and above the age of twenty-one years, are eligible.
Apply in person on Monday, at eleven o’clock, to Duncan Ross, at the
offices of the League, 7 Pope’s Court, Fleet Street.
`

func BenchmarkTokenize(b *testing.B) {
	// for i := 0; i < b.N; i++ {  // Go < 1.24
	for b.Loop() { // 1.24+
		tokens := Tokenize(text)
		if len(tokens) != 47 {
			b.Fatal(len(tokens))
		}
	}
}

/* Running

Bench
go test -run ^$ -bench . -count 7 | go tool benchstat -
go test -run ^$ -bench . -count 7 | benchstat -

Profile CPU
go test -run ^$ -bench . -cpuprofile cpu.pprof

View CPU Profile
go tool pprof -http :8081 tokenizer.test cpu.pprof

Compare Runs
go test -run ^$ -bench . -count 7 | tee orig.txt
<change the code>
go test -run ^$ -bench . -count 7 | tee new.txt
go tool benchstat orig.txt new.txt

Benchmark & Profile Memory
go test -run ^$ -bench . -count 7 -memprofile mem.pprof -benchmem
go test -run ^$ -bench . -count 7 -memprofile mem.pprof -benchmem  | go tool benchstat -

View Memory Profile
go tool pprof -http :8081 tokenizer.test mem.pprof
*/

/* io.Reader design

Go
type Reader interface {
	Read([]byte) (int, error)
}

Python
type Reader interface {
	Read(n int) ([]byte, error)
}

*/

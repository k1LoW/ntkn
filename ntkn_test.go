package ntkn

import (
	"reflect"
	"testing"
)

var tokenizeTests = []struct {
	in   string
	want TokenizedLine
}{
	{
		"Processes: 406 total, 2 running, 404 sleeping, 1973 threads                                                     00:01:04",
		TokenizedLine{
			NumberTokens:         []string{"406", "2", "404", "1973"},
			NumberWithUnitTokens: []string{"406", "2", "404", "1973"},
			TimeTokens:           []string{"00:01:04"},
			NonNumberTokens:      []string{"Processes", "total", "running", "sleeping", "threads"},
			AllTokens:            []string{"Processes", "406", "total", "2", "running", "404", "sleeping", "1973", "threads", "00:01:04"},
		},
	},
	{
		"Load Avg: 1.48, 1.72, 1.81  CPU usage: 7.60% user, 20.65% sys, 71.73% idle  SharedLibs: 269M resident, 72M data, 45M linkedit. MemRegions: 80343 total, 5890M resident, 203M private, 3735M shared.",
		TokenizedLine{
			NumberTokens:         []string{"1.48", "1.72", "1.81", "7.60", "20.65", "71.73", "269", "72", "45", "80343", "5890", "203", "3735"},
			NumberWithUnitTokens: []string{"1.48", "1.72", "1.81", "7.60%", "20.65%", "71.73%", "269M", "72M", "45M", "80343", "5890M", "203M", "3735M"},
			TimeTokens:           []string{},
			NonNumberTokens:      []string{"Load", "Avg", "CPU", "usage", "user", "sys", "idle", "SharedLibs", "resident", "data", "linkedit", "MemRegions", "total", "resident", "private", "shared"},
			AllTokens:            []string{"Load", "Avg", "1.48", "1.72", "1.81", "CPU", "usage", "7.60%", "user", "20.65%", "sys", "71.73%", "idle", "SharedLibs", "269M", "resident", "72M", "data", "45M", "linkedit", "MemRegions", "80343", "total", "5890M", "resident", "203M", "private", "3735M", "shared"},
		},
	},
	{
		"68732  zsh          0.0  00:00.28 1     0    21    3128K  0B     0B     68732 44690 sleeping *0[1]          0.00000 0.00000    501  13840     5305    324       154       39237     305       949",
		TokenizedLine{
			NumberTokens:         []string{"68732", "0.0", "1", "0", "21", "3128", "0", "0", "68732", "44690", "0", "1", "0.00000", "0.00000", "501", "13840", "5305", "324", "154", "39237", "305", "949"},
			NumberWithUnitTokens: []string{"68732", "0.0", "1", "0", "21", "3128K", "0B", "0B", "68732", "44690", "0", "1", "0.00000", "0.00000", "501", "13840", "5305", "324", "154", "39237", "305", "949"},
			TimeTokens:           []string{"00:00.28"},
			NonNumberTokens:      []string{"zsh", "sleeping"},
			AllTokens:            []string{"68732", "zsh", "0.0", "00:00.28", "1", "0", "21", "3128K", "0B", "0B", "68732", "44690", "sleeping", "0", "1", "0.00000", "0.00000", "501", "13840", "5305", "324", "154", "39237", "305", "949"},
		},
	},
}

func TestTokenize(t *testing.T) {
	for _, tt := range tokenizeTests {
		got := TokenizeLine(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("\ngot  %#v\nwant %#v", got, tt.want)
		}
	}
}

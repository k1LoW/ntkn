package ntkn

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

var splitToken = ` ,[](){}/*"`
var trimToken = ":."
var unitToken = "sbkmgBKMG%"
var numberWithUnitRegexp = regexp.MustCompile(`^\d+(\.\d+)?[` + unitToken + `]?$`)
var timeRegexp = regexp.MustCompile(`^\d{2}:\d{2}[\d:.]*$`)

// TokenizedLine struct
type TokenizedLine struct {
	NumberTokens         []string
	NumberWithUnitTokens []string
	TimeTokens           []string
	NonNumberTokens      []string
	AllTokens            []string
}

// Tokenize return []TokenizedLine
func Tokenize(r io.Reader) []TokenizedLine {
	tknls := []TokenizedLine{}
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		tknls = append(tknls, TokenizeLine(scanner.Text()))
	}
	return tknls
}

// TokenizeLine tokenize line
func TokenizeLine(in string) TokenizedLine {
	scanner := bufio.NewScanner(strings.NewReader(in))
	scanner.Split(splitFunc)

	numberTokens := []string{}
	numberWithUnitTokens := []string{}
	timeTokens := []string{}
	nonNumberTokens := []string{}
	allTokens := []string{}

	for scanner.Scan() {
		str := scanner.Text()
		if str != "" {
			trimmed := strings.TrimRight(scanner.Text(), trimToken)
			if numberWithUnitRegexp.MatchString(trimmed) {
				numberWithUnitTokens = append(numberWithUnitTokens, trimmed)
				numberTokens = append(numberTokens, strings.TrimRight(trimmed, unitToken))
			} else if timeRegexp.MatchString(trimmed) {
				timeTokens = append(timeTokens, trimmed)
			} else {
				nonNumberTokens = append(nonNumberTokens, trimmed)
			}
			allTokens = append(allTokens, trimmed)
		}
	}

	tknl := TokenizedLine{
		NumberTokens:         numberTokens,
		NumberWithUnitTokens: numberWithUnitTokens,
		TimeTokens:           timeTokens,
		NonNumberTokens:      nonNumberTokens,
		AllTokens:            allTokens,
	}
	return tknl
}

func splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		if strings.Contains(splitToken, string(data[i])) {
			return i + 1, data[:i], nil
		}
	}
	return 0, data, bufio.ErrFinalToken
}

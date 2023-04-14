package pixelmator

import (
	"bytes"
	"fmt"
	"github.com/scryner/streamdeck-pixelmator/applescript"
	"strings"
)

func syncColorAdjustments(vals map[ColorAdjustment]ColorAdjustmentValue) error {
	// collect terms to query
	var terms []term

	for _, v := range vals {
		terms = append(terms, v.getTerms()...)
	}

	// make query
	buf := new(bytes.Buffer)
	var variables []string
	for _, t := range terms {
		buf.WriteString(fmt.Sprintf("\t\t\tset %s to its %s\n", t.osascriptVariable, t.osascriptTerm))
		variables = append(variables, t.osascriptVariable)
	}

	query := fmt.Sprintf(syncQueryFormat, buf.String(), strings.Join(variables, ` & "," & `))

	// do query
	res, err := applescript.Run(query)
	if err != nil {
		return fmt.Errorf("failed to run osascript: %v", err)
	}

	// parse result
	res = strings.ReplaceAll(res, "{", "")
	res = strings.ReplaceAll(res, "}", "")
	res = strings.ReplaceAll(res, " ", "")

	ss := strings.Split(res, ",")
	if len(ss) != len(terms) {
		return fmt.Errorf("mismatched length of result: (queryLen:%d) != (resultLen:%d)", len(vals), len(terms))
	}

	m := make(map[string]string)
	for i, va := range variables {
		m[va] = ss[i]
	}

	// set values
	for _, v := range vals {
		if err := v.setValues(m); err != nil {
			return err
		}
	}

	return nil
}

const syncQueryFormat = `if application "Pixelmator Pro" is running then
	tell application "Pixelmator Pro"
		tell the color adjustments of the current layer of the front document
%s
		end tell
	end tell
end if

if application "Pixelmator Pro" is running then
	set result to "" & %s
else
	set result to ""
end if

result`

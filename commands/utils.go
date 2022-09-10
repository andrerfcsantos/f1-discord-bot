package commands

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	tbw "text/tabwriter"
	"time"

	"github.com/agnivade/levenshtein"
)

// LevenshteinDistance contains the computed levenshtein
// between Str1 and Str2
type LevenshteinDistance struct {
	Str1     string
	Str2     string
	Distance int
}

// Compute computes (or recomputes) the levenshtein between Str1 and Str2
func (ld *LevenshteinDistance) Compute() {
	ld.Distance = levenshtein.ComputeDistance(ld.Str1, ld.Str2)
}

// LevenshteinDistances is a list of computed levenshtein distances
type LevenshteinDistances []LevenshteinDistance

// ComputeAll computes all levenshtein distances
func (lds LevenshteinDistances) ComputeAll() {
	for i := range lds {
		lds[i].Compute()
	}
}

// SortByDistance sorts the list of levenshtein distances by distance
func (lds LevenshteinDistances) SortByDistance() {
	sort.Slice(lds, func(i, j int) bool {
		return lds[i].Distance < lds[j].Distance
	})
}

// PrettyCountdount sees a Duration as a countdown for an event to happen
// and transforms into a easily readible string representation.
// The results contains information about if the event it's in the past or it's still to come.
// If the event it's still to come, it will also contain information about the time left,
// expressed in minutes, hours, or days depending on the value of the duration.
func PrettyCountdount(d time.Duration) string {
	switch {
	case d < 0:
		return "already over"
	case d.Hours() < 1:
		return fmt.Sprintf("%.1f minutes to go", d.Minutes())
	case d.Hours() < 24:
		return fmt.Sprintf("%.1f hours to go", d.Hours())
	default:
		return fmt.Sprintf("%.1f days to go", d.Hours()/24)
	}
}

// ParseRFC3339InLocation parses a time string in RFC3339 as a time, and returns that time in
// the iana location provided
func ParseRFC3339InLocation(timeValue string, ianaLocation string) (time.Time, error) {
	rfc3339Time, err := time.Parse(time.RFC3339, timeValue)
	if err != nil {
		return rfc3339Time, fmt.Errorf("parsing time '%s' in RFC3339 format: %v", timeValue, err)
	}

	location, err := time.LoadLocation(ianaLocation)
	if err != nil {
		return rfc3339Time, fmt.Errorf("loading '%s' location: %v", ianaLocation, err)
	}

	return rfc3339Time.In(location), nil
}

// HeaderMessage represents a message to discord with an header and a description
type HeaderMessage struct {
	Header      string
	Description string
}

func (hm *HeaderMessage) String() string {
	var message strings.Builder

	// Write header
	if hm.Header != "" {
		message.WriteString(fmt.Sprintf("**%s**\n", strings.ToUpper(hm.Header)))
	}

	// Write description
	if hm.Description != "" {
		message.WriteString(hm.Description + "\n")
	}

	return message.String()
}

// TabularMessage represents a message with an header, description and some tabular data
type TabularMessage struct {
	HeaderMessage
	TableHeader []string
	TableRows   [][]string
}

// SetTableHeader sets the table header
func (tm *TabularMessage) SetTableHeader(headers ...string) {
	tm.TableHeader = headers
}

// AddRow adds a row of data
func (tm *TabularMessage) AddRow(rowData ...string) {
	tm.TableRows = append(tm.TableRows, rowData)
}

// String returns TabularMessage for discord
func (tm *TabularMessage) String() string {
	var message strings.Builder

	message.WriteString(tm.HeaderMessage.String())

	// Make table
	var tablebBuilder strings.Builder

	tableWriter := tbw.NewWriter(&tablebBuilder, 0, 0, 3, ' ', 0)
	fmt.Fprintln(tableWriter, strings.Join(tm.TableHeader, "\t"))

	for _, rowData := range tm.TableRows {
		fmt.Fprintln(tableWriter, strings.Join(rowData, "\t"))
	}

	tableWriter.Flush()

	// Write table
	message.WriteString("```" + tablebBuilder.String() + "```")

	return message.String()
}

// RaceHourComment returns a string with a comment about how late or not is the hour of the race.
func RaceHourComment(raceTime time.Time) string {
	hour := raceTime.Hour()
	if hour > 4 && hour < 9 {
		return "Unfortunately it seems you'll have to wake up early if you want to watch the race :("
	}
	return "It seems a decent hour for the race. You won't have to wake up early!"
}

func ParseDuration(d time.Duration) string {
	res := ""

	days := int(math.Trunc(d.Hours())) / 24

	if days != 0 {
		res += strconv.Itoa(days) + "d"
		d = d - time.Duration(24*days)*time.Hour
	}

	hours := int(math.Trunc(d.Hours()))
	if res != "" || hours != 0 {
		res += strconv.Itoa(hours) + "h"
		d = d - time.Duration(hours)*time.Hour
	}

	minutes := int(math.Trunc(d.Minutes()))
	if res != "" || minutes != 0 {
		res += strconv.Itoa(minutes) + "m"
		d = d - time.Duration(minutes)*time.Minute
	}

	seconds := int(math.Trunc(d.Seconds()))
	if hours < 1 && (res != "" || minutes != 0) {
		res += strconv.Itoa(seconds) + "s"
	}
	return res
}

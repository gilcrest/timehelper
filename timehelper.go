package timehelper

import "time"

// LastNanosecondNextYear gives the last nanosecond of the next year
// for the given input Time in that Location's time zone. Pass a
// time.Time using UTC if you want to get back the time in UTC
func LastNanosecondNextYear(t time.Time) time.Time {

	year, _, _ := t.Date()
	ldny := time.Date(year+1, time.December, 31, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())

	return ldny
}

// ZeroHour (aka Midnight) gives the first nanosecond of whatever day
// you pass in for the given input Time in that Location's time zone.
// Pass a time.Time using UTC if you want to get back the time in UTC
func ZeroHour(t time.Time) time.Time {

	year, month, day := t.Date()
	zh := time.Date(year, month, day, 0, 0, 0, 0, t.Location())

	return zh
}

// LastNanosecond gives the last nanosecond of whatever day you pass in
// for the given input Time in that Location's time zone. Pass a
// time.Time using UTC if you want to get back the time in UTC
func LastNanosecond(t time.Time) time.Time {
	year, month, day := t.Date()
	ls := time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())

	return ls
}

// FirstOfTheMonth gives the first nanosecond of the month of
// whatever day you pass in for the given input Time in that Location's time zone.
// Pass a time.Time using UTC if you want to get back the time in UTC
func FirstOfTheMonth(t time.Time) time.Time {

	year, month, _ := t.Date()
	zh := time.Date(year, month, 1, 0, 0, 0, 0, t.Location())

	return zh
}

// FirstOfPriorMonth gives the first nanosecond of the prior month of
// whatever day you pass in for the given input Time in that Location's time zone.
// Pass a time.Time using UTC if you want to get back the time in UTC
func FirstOfPriorMonth(t time.Time) time.Time {

	year, month, _ := t.Date()
	if month == 1 {
		year--
		month = 12
	} else {
		month--
	}
	zh := time.Date(year, month, 1, 0, 0, 0, 0, t.Location())

	return zh
}

// LastNanosecondOfTheMonth gives the first nanosecond of the month of
// whatever day you pass in for the given input Time in that Location's time zone.
// Pass a time.Time using UTC if you want to get back the time in UTC
func LastNanosecondOfTheMonth(t time.Time) time.Time {

	nm := FirstOfTheMonth(t).AddDate(0, 1, 0)
	nm = nm.Add(time.Nanosecond * -1)

	return nm
}

// MTD (Month To Date) returns the start of the month
// and the start of the next day for the given time input
// to enable range selections using >= and <
func MTD(t time.Time) (time.Time, time.Time) {
	fom := FirstOfTheMonth(t)
	nd := ZeroHour(t.Add(time.Hour * 24))

	return fom, nd
}
package timehelper

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLastNanosecondNextYear(t *testing.T) {
	type args struct {
		in0 time.Time
	}

	i1 := time.Date(2018, time.September, 23, 0, 0, 0, 0, time.Local)
	o1 := time.Date(2019, time.December, 31, 23, 59, 59, int(time.Second-time.Nanosecond), time.Local)

	i2 := time.Date(2018, time.September, 23, 0, 0, 0, 0, time.UTC)
	o2 := time.Date(2019, time.December, 31, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)

	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{name: "Local Positive Test", args: args{i1}, want: o1},
		{name: "UTC Positive Test", args: args{i2}, want: o2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastNanosecondNextYear(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastNanosecondNextYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZeroHour(t *testing.T) {
	type args struct {
		t time.Time
	}

	i1 := time.Date(2018, time.September, 23, 0, 0, 0, 0, time.Local)
	o1 := time.Date(2018, time.September, 23, 0, 0, 0, 0, time.Local)

	i2 := time.Date(2018, time.September, 23, 0, 0, 0, 0, time.UTC)
	o2 := time.Date(2018, time.September, 23, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{name: "Local Positive Test", args: args{i1}, want: o1},
		{name: "UTC Positive Test", args: args{i2}, want: o2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZeroHour(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZeroHour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastNanosecond(t *testing.T) {
	type args struct {
		t time.Time
	}

	i1 := time.Date(2018, time.September, 23, 0, 0, 0, 0, time.Local)
	o1 := time.Date(2018, time.September, 23, 23, 59, 59, int(time.Second-time.Nanosecond), time.Local)

	i2 := time.Date(2017, time.September, 23, 0, 0, 0, 0, time.UTC)
	o2 := time.Date(2017, time.September, 23, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)

	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{name: "Local Positive Test", args: args{i1}, want: o1},
		{name: "UTC Positive Test", args: args{i2}, want: o2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastNanosecond(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastNanosecond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstOfTheMonth(t *testing.T) {
	type args struct {
		t time.Time
	}

	i1 := time.Date(2018, time.September, 23, 0, 0, 0, 0, time.Local)
	o1 := time.Date(2018, time.September, 1, 0, 0, 0, 0, time.Local)

	i2 := time.Date(2017, time.September, 23, 0, 0, 0, 0, time.UTC)
	o2 := time.Date(2017, time.September, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{name: "Local Positive Test", args: args{i1}, want: o1},
		{name: "UTC Positive Test", args: args{i2}, want: o2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstOfTheMonth(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstOfTheMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleFirstOfTheMonth() {
	i1 := time.Date(2018, time.September, 23, 0, 0, 0, 0, time.UTC)
	fmt.Println(FirstOfTheMonth(i1))
	// Output: 2018-09-01 00:00:00 +0000 UTC
}

func TestLastNanosecondOfTheMonth(t *testing.T) {
	type args struct {
		t time.Time
	}

	i1 := time.Date(2018, time.February, 23, 0, 0, 0, 0, time.Local)
	o1 := time.Date(2018, time.February, 28, 23, 59, 59, int(time.Second-time.Nanosecond), time.Local)

	i2 := time.Date(2020, time.February, 23, 0, 0, 0, 0, time.UTC)
	o2 := time.Date(2020, time.February, 29, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)

	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{name: "Local Positive Test", args: args{i1}, want: o1},
		{name: "UTC Positive Test", args: args{i2}, want: o2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastNanosecondOfTheMonth(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastNanosecondOfTheMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstOfPriorMonth(t *testing.T) {
	type args struct {
		t time.Time
	}

	i1 := time.Date(2018, time.September, 23, 0, 0, 0, 0, time.Local)
	o1 := time.Date(2018, time.August, 1, 0, 0, 0, 0, time.Local)

	i2 := time.Date(2017, time.January, 23, 0, 0, 0, 0, time.UTC)
	o2 := time.Date(2016, time.December, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{name: "Local Positive Test", args: args{i1}, want: o1},
		{name: "UTC Positive Test", args: args{i2}, want: o2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstOfPriorMonth(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstOfLastMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstOfTheQuarter(t *testing.T) {

	q1 := time.Date(2020, time.February, 29, 0, 0, 0, 0, time.Local)
	o1 := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local)

	f := FirstOfTheQuarter(q1)
	require.Equal(t, o1, f)

	q2 := time.Date(2020, time.May, 29, 0, 0, 0, 0, time.Local)
	o2 := time.Date(2020, time.April, 1, 0, 0, 0, 0, time.Local)

	f2 := FirstOfTheQuarter(q2)
	require.Equal(t, o2, f2)

	q3 := time.Date(2020, time.August, 29, 0, 0, 0, 0, time.Local)
	o3 := time.Date(2020, time.July, 1, 0, 0, 0, 0, time.Local)

	f3 := FirstOfTheQuarter(q3)
	require.Equal(t, o3, f3)

	q4 := time.Date(2020, time.November, 29, 0, 0, 0, 0, time.Local)
	o4 := time.Date(2020, time.October, 1, 0, 0, 0, 0, time.Local)

	f4 := FirstOfTheQuarter(q4)
	require.Equal(t, o4, f4)
}

func TestYTD(t *testing.T) {

	i1 := time.Date(2020, time.July, 4, 0, 0, 0, 0, time.Local)
	o1 := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local)

	f := FirstOfTheYear(i1)
	require.Equal(t, o1, f)
}

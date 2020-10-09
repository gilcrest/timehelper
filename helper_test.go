package timehelper

import (
	"fmt"
	"reflect"
	"testing"
	"time"
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

func TestFirstOfLastMonth(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstOfLastMonth(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstOfLastMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

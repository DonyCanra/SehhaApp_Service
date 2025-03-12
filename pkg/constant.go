package pkg

const (
	StdLongMonth      = "January"
	StdMonth          = "Jan"
	StdNumMonth       = "1"
	StdZeroMonth      = "01"
	StdLongWeekDay    = "Monday"
	StdWeekDay        = "Mon"
	StdDay            = "2"
	StdUnderDay       = "_2"
	StdZeroDay        = "02"
	StdHour           = "15"
	StdHour12         = "3"
	StdZeroHour12     = "03"
	StdMinute         = "4"
	StdZeroMinute     = "04"
	StdSecond         = "5"
	StdZeroSecond     = "05"
	StdLongYear       = "2006"
	StdYear           = "06"
	StdPM             = "PM"
	Stdpm             = "pm"
	StdTZ             = "MST"
	StdISO8601TZ      = "Z0700"  // prints Z for UTC
	StdISO8601ColonTZ = "Z07:00" // prints Z for UTC
	StdNumTZ          = "-0700"  // always numeric
	StdNumShortTZ     = "-07"    // always numeric
	StdNumColonTZ     = "-07:00" // always numeric
	StdFracSecond0    = ".0"     // trailing zeros included
	StdFracSecond9    = ".9"     // trailing zeros omitted
	StdDateFormat     = StdLongYear + "-" + StdZeroMonth + "-" + StdZeroDay
	StdDateTimeFormat = StdLongYear + "-" + StdZeroMonth + "-" + StdZeroDay + " " + StdHour + ":" + StdZeroMinute + ":" + StdZeroSecond

	StdDateTimeZone = "Asia/Jakarta"
)

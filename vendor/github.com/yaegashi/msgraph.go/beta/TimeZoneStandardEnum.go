// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TimeZoneStandard undocumented
type TimeZoneStandard int

const (
	// TimeZoneStandardVWindows undocumented
	TimeZoneStandardVWindows TimeZoneStandard = 0
	// TimeZoneStandardVIana undocumented
	TimeZoneStandardVIana TimeZoneStandard = 1
)

// TimeZoneStandardPWindows returns a pointer to TimeZoneStandardVWindows
func TimeZoneStandardPWindows() *TimeZoneStandard {
	v := TimeZoneStandardVWindows
	return &v
}

// TimeZoneStandardPIana returns a pointer to TimeZoneStandardVIana
func TimeZoneStandardPIana() *TimeZoneStandard {
	v := TimeZoneStandardVIana
	return &v
}

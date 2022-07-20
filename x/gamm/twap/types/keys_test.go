package types

import (
	"testing"
	time "time"

	"github.com/stretchr/testify/require"
)

func TestFormatMostRecentTWAPKey(t *testing.T) {
	tests := map[string]struct {
		poolId uint64
		denom1 string
		denom2 string
		want   string
	}{
		"standard":       {poolId: 1, denom1: "B", denom2: "A", want: "recent_twap|1|B|A"},
		"standard2digit": {poolId: 10, denom1: "B", denom2: "A", want: "recent_twap|10|B|A"},
		"maxPoolId":      {poolId: ^uint64(0), denom1: "B", denom2: "A", want: "recent_twap|18446744073709551615|B|A"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := FormatMostRecentTWAPKey(tt.poolId, tt.denom1, tt.denom2)
			require.Equal(t, tt.want, string(got))
		})
	}
}

func TestFormatHistoricalTwapKeys(t *testing.T) {
	// go playground default time
	// 2009-11-10 23:00:00 +0000 UTC m=+0.000000001
	baseTime := time.Unix(1257894000, 0).UTC()
	tests := map[string]struct {
		poolId        uint64
		time          time.Time
		denom1        string
		denom2        string
		wantPoolIndex string
		wantTimeIndex string
	}{
		"standard": {poolId: 1, time: baseTime, denom1: "B", denom2: "A", wantTimeIndex: "historical_time_index|2009-11-10T23:00:00.000000000|1|B|A", wantPoolIndex: "historical_pool_index|1|2009-11-10T23:00:00.000000000|B|A"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			gotTimeKey := FormatHistoricalTimeIndexTWAPKey(tt.time, tt.poolId, tt.denom1, tt.denom2)
			gotPoolKey := FormatHistoricalPoolIndexTWAPKey(tt.poolId, tt.time, tt.denom1, tt.denom2)
			require.Equal(t, tt.wantTimeIndex, string(gotTimeKey))
			require.Equal(t, tt.wantPoolIndex, string(gotPoolKey))

			parsedTime := ParseTimeFromHistoricalTimeIndexKey(gotTimeKey)
			require.Equal(t, tt.time, parsedTime)
			parsedTime = ParseTimeFromHistoricalPoolIndexKey(gotPoolKey)
			require.Equal(t, tt.time, parsedTime)
		})
	}
}

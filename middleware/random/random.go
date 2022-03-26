package random

import "time"

func DailyRandomUInt(seed uint64, maxValue uint) uint {
	bitSeed := 0
	valueSeed := 0
	mixSeed := 0
	seed += uint64(time.Now().Day()) + uint64(time.Now().Month()) - uint64(time.Now().Year())
	for temp := seed; temp != 0; temp /= 10 {
		bitSeed += 1
		valueSeed += int(temp % 10)
		mixSeed += int(temp) / bitSeed
	}

	return uint(bitSeed*valueSeed+mixSeed) % maxValue
}

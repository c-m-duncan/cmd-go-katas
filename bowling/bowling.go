package bowling

func Roll(pins [] int) int {
	var frames [11]int
	sum := 0
	pinIndex := 0

	for range frames {
		if pinIndex < len(pins) {
			sum += CalculateFrameScore(pins, &pinIndex)
		}
	}
	return sum
}

func CalculateFrameScore(pins []int, pinIndex *int) int {
	firstThrow := IterateThrow(pins, pinIndex)
	frameScore := firstThrow
	if firstThrow == 10 {
		frameScore += CalculateStrikeBonus(pins, pinIndex)
	} else {
		if *pinIndex < len(pins) {
			secondThrow := IterateThrow(pins,  pinIndex)
			if firstThrow+secondThrow == 10 {
				frameScore += CalculateSpareBonus(pins, pinIndex)
			}
		}
	}
	return frameScore
}

func CalculateSpareBonus(pins []int, pinIndex *int) int {
	bonus := 0
	if *pinIndex < len(pins) {
		bonus += pins[*pinIndex]
	}
	return bonus
}

func CalculateStrikeBonus(pins []int, pinIndex *int) int {
	bonus := CalculateSpareBonus(pins, pinIndex)
	if *pinIndex + 1 < len(pins) {
		bonus += pins[*pinIndex+1]
	}
	return bonus
}

func IterateThrow(pins []int,  pinIndex *int) int {
	throw := pins[*pinIndex]
	*pinIndex++
	return throw
}
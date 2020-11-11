package bowling_tests_test

import (
	"../bowling"
	"testing"
)

func TestGutterBallGame(t *testing.T) {
	var pins = make([]int, 20)

	score := bowling.Roll(pins)
	if score != 0 {
		t.Errorf("Rolling all gutters yields %d; expected 0", score)
	}
}

func TestRollingAllOnes(t *testing.T) {
	pins := MakeWithDefaultValue(20, 1)
	score := bowling.Roll(pins)
	if score != 20 {
		t.Errorf("Rolling all ones yields %d; expected 20", score)
	}
}

func TestSpareCountsNextThrowTwice(t *testing.T) {
	pins := []int {5,5,1,1}
	score := bowling.Roll(pins)
	if score != 13 {
		t.Errorf("Rolling spare does not count next throw as double. Actual: %d; expected 13", score)
	}
}

func TestStrikeCountsNextThrowsTwice(t *testing.T) {
	pins := []int {10,3,2,1}
	score := bowling.Roll(pins)
	if score != 21 {
		t.Errorf("Rolling strike does not count next throw as double. Actual: %d; expected 21", score)
	}
}

func TestPerfectGame(t *testing.T) {
	pins := MakeWithDefaultValue(11, 10)
	score := bowling.Roll(pins)
	if score != 300 {
		t.Errorf("Rolling perfect game is incorrect. Actual: %d; expected 300", score)
	}
}

func MakeWithDefaultValue(size int, value int) []int {
	var array = make([]int, size)
	for i := 0; i < size; i++ {
		array[i] += value
	}
	return array
}
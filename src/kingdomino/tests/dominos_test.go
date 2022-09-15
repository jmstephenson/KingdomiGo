package kingdomino

import (
	kd "kingdomino/src/kingdomino"
	"testing"
)

func TestShuffleDeck(t *testing.T) {

	var newDeck = kd.NewDeck()

	newDeck.ShuffleDeck()

	var expectedSumOfDominoValues = ((kd.TOTAL_TILES * (kd.TOTAL_TILES + 1)) / 2)

	calcDeckValue := func(newDeck []*kd.Domino) int {
		total := 0
		for _, d := range newDeck {
			total += d.Value
		}
		return total
	}

	var actualSumOfDominoValues = calcDeckValue(newDeck.Dominos)

	isvalid := expectedSumOfDominoValues == actualSumOfDominoValues

	if !isvalid {
		t.Errorf("ShuffleDeck() result produced incorrect card value sum. Expected %d : Actual %d", expectedSumOfDominoValues, actualSumOfDominoValues)
	}
}

func TestDeckSort(t *testing.T) {
	return
}

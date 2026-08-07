package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace" :
        	return 11
        case "king", "queen", "jack", "ten" :
        	return 10
        case "nine":
        	return 9
        case "eight":
        	return 8
        case "seven":
        	return 7
        case "six":
        	return 6
        case "five":
        	return 5
        case "four":
        	return 4
        case "three":
        	return 3
        case "two":
        	return 2
        default:
        	return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    dealer := ParseCard(dealerCard)
	total := ParseCard(card1) + ParseCard(card2)

    if card1 == "ace" && card2 == "ace" {
        return "P"
    }

    if total == 21 {
        if dealer == 10 || dealer == 11 {
            return "S"
        }
        return "W"
    }

    switch {
        case total >= 17 && total <= 20:
            return "S"
        case total >= 12 && total <= 16:
             if dealer >= 7 {
                return "H"
            }
            return "S"
        default:
            return "H"
    }
}

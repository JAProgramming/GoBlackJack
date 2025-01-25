// Simple Blackjack for CSI 380
// This program play's a simple, single suit game of Blackjack
// against a computer dealer.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Move a card from deck to hand
func drawCard(hand *[]string, deck *[]string) {
	// YOU FILL IN HERE
	var slice []card
	slice = append(slice, deck{0})
	deck = deck[1:] // got from LLM

}

// Calculate the score of the hand
// Used LLM to edit and tweak this section of code after hours of
// trying to fix error with converting the try catch statement 
func calculateScore(hand []string) int {
	// YOU FILL IN HERE
	var score int = 0
	var hasAce bool = false

	for _, card := range hand
	{

		if card == "2"
		{
			score += 2
		}

		else if card == "3"
		{
			score += 3
		}

		else if card == "4"
		{
			score += 4
		}

		else if card == "5"
		{
			score += 5
		}

		else if card == "6"
		{
			score += 6
		}

		else if card == "7"
		{
			score += 7
		}

		else if card == "8"
		{
			score += 8
		}

		else if card == "9"
		{
			score += 9
		}

		else if card == "A"			
		{
			hasAce = true
		}

		else
		{
			score += 10
		}

	}

	if hasAce
	{
		if (score + 11) > 21
		{
			score += 1
		}
		else
		{
			score += 11
		}
	}

	return score
}

// Print everyone's scores and hands
func printStatus(playerCards, dealerCards []string) {
	// YOU FILL IN HERE
	fmt.Println("")
	fmt.Println("Player's total score is: " + calculateScore(playerCards))
	fmt.Println(playerCards)
	fmt.Println("")
	fmt.Println("Dealer's total is: " + calculateScore(dealerCards)
	fmt.Println(dealerCards)
	fmt.Println("")
)

}

// Entry point and main game loop
func main() {
	// YOU FILL IN HERE
	deck := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J","Q", "K", "A"}
	playerHand := []string{}
	dealerHand := []string{}

	rand.Seed(time.Now().UnixNano()) // got from LLM
	randomNum := rand.Intn(13)

	fmt.Println("Dealer Draws")
	dealerHand = append(dealerHand, deck[randomNum])

	rand.Seed(time.Now().UnixNano()) 
	randomNum := rand.Intn(13)

	fmt.Println("Player Hand")
	playerHand = append(playerHand, deck(13))

	rand.Seed(time.Now().UnixNano()) 
	randomNum := rand.Intn(13)
	playerHand = append(playerHand, deck(13))

	printStatus(dealerHand, playerHand)


	for true
	{
		fmt.Println("Would you like to (H)Hit, (S)Stay, or (Q)Quit")
		fmt.Println("Please enter only capital letters")
		var ans string
		fmt.ScanIN(&ans)

		if ans == "H"
		{
			rand.Seed(time.Now().UnixNano()) 
			randomNum := rand.Intn(13)
			playerHand = append(playerHand, deck(13))

			printStatus(dealerHand, playerHand)

			if calculateScore(playerHand) > 21
			{
				fmt.Println("You Lose :(")
				return 0
			}
		}

		else if ans == "S"
		{
			break
		}

		else if ans == "Q"
		{
			return 0
		}
	}

	for calculateScore(dealerHand) < 17
	{
		rand.Seed(time.Now().UnixNano()) // got from LLM
		randomNum := rand.Intn(13)
	
		dealerHand = append(dealerHand, deck[randomNum])
	}

	if alculateScore(dealerHand) > 21
	{
		fmt.Println("Dealer Busts! You Win")
	}

	else if alculateScore(dealerHand) > calculateScore(playerHand)
	{
		fmt.Println("Dealer Wins!")
	}

	else if alculateScore(dealerHand) < calculateScore(playerHand)
	{
		fmt.Println("You Win!")
	}

	else if alculateScore(dealerHand) == calculateScore(playerHand)
	{
		fmt.Println("Tie Game!")
	}

}

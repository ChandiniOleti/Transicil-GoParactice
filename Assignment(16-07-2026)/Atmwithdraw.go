
// Problem Statement
// --------------------------------
// An ATM needs to dispense a given amount of cash using the maximum number of currency notes possible (not the usual minimum), so students feel like they're getting more notes. Available denominations: ₹100, ₹200, ₹500, ₹1000. Given an amount, determine the combination of notes that maximizes note count while still summing exactly to the amount.

// Input: Amount to withdraw (and possibly a max notes-per-transaction limit N, depending on the variant)
// Output: Number of notes of each denomination, and/or total note count

package main
import "fmt"
func main() {
	var amount int

	fmt.Print("Enter Amount: ")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("Invalid Amount")
		return
	}
	if amount%100 != 0 {
		fmt.Println("Invalid Amount")
		return
	}
	count1000 := 0
	count500 := 0
	count200 := 0
	count100 := 0

	if amount >= 100 {
		count100 = amount / 100
		amount = amount - (count100 * 100)//amount/100
		if amount==50{
			fmt.Println("Invalid")
		}else{
		fmt.Println("100  :", count100)
		}
	}

	if amount >= 200 {
		count200 = amount / 200
		amount = amount - (count200 * 200)//amount/200
		if amount==50{
			fmt.Println("Invalid")
		}else{
		fmt.Println("200  :", count200)
		} 
	}

	if amount >= 500 {
		count500 = amount / 500
		amount = amount - (count500 * 500)//amount=amount/500
		if amount==50{
			fmt.Println("Invalid")
		}else{
		fmt.Println("500  :", count500)
		}
	}

	if amount >= 1000 {
		count1000 = amount / 1000
		amount = amount - (count1000 * 1000)//amount/1000
		if amount==50{
			fmt.Println("Invalid")
		}else{
		fmt.Println("1000  :", count1000)
		}
	}

	total := count100 + count200 + count500 + count1000
	fmt.Println("Total Notes :", total)
	//if we need to print all denomination we can print here for every100,200,500 and 1000denominations
}
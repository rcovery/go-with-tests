package main

func condition() {
	var x int = 30
	var y int = 20

	if x > y {
		println("x is greater than y")
	} else {
		println("y is greater than x")
	}

	switch x {
	case 10:
		println("x is 10")
	case 20:
	case 30:
		println("x is 20 or 30")
	}
}

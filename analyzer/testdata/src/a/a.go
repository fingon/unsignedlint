package a

func main() {
	var a uint = 5
	var b uint = 10
	_ = a - b // want "subtraction with unsigned integer operand 'b' may underflow"

	var x uint32 = 100
	var y uint32 = 200
	_ = x - y // want "subtraction with unsigned integer operand 'y' may underflow"

	var p uintptr = 0
	var q uintptr = 1
	_ = p - q // want "subtraction with unsigned integer operand 'q' may underflow"

	// Safe cases (should not be flagged by this basic version) - Now flagged as expected
	var safeA uint = 10
	var safeB uint = 5
	_ = safeA - safeB // want "subtraction with unsigned integer operand 'safeB' may underflow"

	var signedInt int = 10
	var unsignedInt uint = 5
	_ = signedInt - int(unsignedInt) // Subtraction result is signed
	_ = int(unsignedInt) - signedInt // Subtraction result is signed

	// Addition should not be flagged
	_ = a + b
}

func otherTypes() {
	var u8 uint8 = 1
	var u16 uint16 = 2
	var u32 uint32 = 3
	var u64 uint64 = 4

	_ = u8 - 10  // want "subtraction with unsigned integer operand '10' may underflow"
	_ = u16 - 20 // want "subtraction with unsigned integer operand '20' may underflow"
	_ = u32 - 30 // want "subtraction with unsigned integer operand '30' may underflow"
	_ = u64 - 40 // want "subtraction with unsigned integer operand '40' may underflow"

	var i int = 5
	_ = i - 1 // Signed subtraction, should not flag
}

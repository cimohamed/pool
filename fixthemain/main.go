package main

import "github.com/01-edu/z01"

type Door struct{

}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) {
	
	ptrDoor.state = Open
	PrintStr("Door opening...\n")
}

func IsDoorOpen(Door Door)bool {
	PrintStr("Door closing..\n")
	ptrDoor.tate = CLOSE 
	return true
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
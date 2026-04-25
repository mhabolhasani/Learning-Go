package main

// type alias let us add a second name for an existing type:
type temperature = float64

// type definition let us add more date types like UserId :
type UserId int

// the advantage of type definition is defining method for that.
func (u UserId) IsValid() bool {
	return u > 0
}

func main() {

}

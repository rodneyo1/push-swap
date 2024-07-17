package main

import (
	"fmt"
)

func main() {
	 a := []int{8,5, 6,7,9}
	b := []int{2, 3,4,5}
	// fmt.Println(pa(a, b))
	// fmt.Println(pb(a, b))
	// fmt.Println(SA(a))
	// fmt.Println(SB(b))
	// fmt.Println(ss(a,b))
	// fmt.Println(RB(b))
	// fmt.Println(RA(a))
	//fmt.Println(RR(a,b))
	fmt.Println(RRA(a))
	fmt.Println(RRB(b))
	fmt.Println(rrr(a,b))
}

// pa push the top first element of stack b to stack a
func pa(a []int, b []int) ([]int, []int) {
	if len(b) > 0 {
		a = append(a, b[0])
		swap := a[0]
		a[0] = a[len(a)-1]
		a[len(a)-1] = swap

		b = b[1:]
	}
	return a, b
}

// pb push the top first element of stack a to stack b
func pb(a []int, b []int) ([]int, []int) {
	if len(a) > 0 {
		b = append(b, a[0])
		swap := b[0]
		b[0] = b[len(b)-1]
		b[len(b)-1] = swap

		a = a[1:]
	}

	return a, b
}

// sa swap first 2 elements of stack a
func SA(a []int) []int {
	swap := a[0]
	a[0] = a[1]
	a[1] = swap
	return a
}

// sb swap first 2 elements of stack b
func SB(b []int) []int {
	return SA(b)
}

// ss execute sa and sb
func ss(a []int, b []int) ([]int, []int) {
	return SA(a), SB(b)
}

// ra rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
func RB(b []int) []int {
	if len(b)==0{
		return b
	}
	c := b[1:]
	c = append(c, b[0])
	return c
}

// rb rotate stack b
func RA(a []int) []int{
	return RB(a)
}
// rr execute ra and rb
func RR(a[]int,b[]int)([]int,[]int){
	return RA(a), RB(b)
}
// rra reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
func RRA(a []int) []int {
    if len(a) == 0 {
        return a 
    }
    first := a[len(a)-1] 
    c := append([]int{first}, a[:len(a)-1]...)
    
    return c
}


// rrb reverse rotate b
func RRB(b []int) []int {
	return RRA(b)
}
// rrr execute rra and rrb
func rrr(a[]int, b[]int)([]int,[]int){
	return RRA(a), RRB(b)
}

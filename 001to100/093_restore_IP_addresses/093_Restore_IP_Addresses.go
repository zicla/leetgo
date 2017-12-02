package main

import "fmt"

func restoreIpAddresses(s string) []string {
	store := make([]string, 330)
	bits := make([]int, 4, 4)
	storeN := []int{0}
	N := len(s)
	restoreIpAddressesRecursive(bits, 0, 0, N, s, store, storeN)

	return store[0:storeN[0]]
}
func restoreIpAddressesRecursive(bits []int, bitsN int, used int, N int, s string, store []string, storeN []int) {

	if bitsN == 4 {

		if bits[0]+bits[1]+bits[2]+bits[3] != N {
			return
		}

		a := s[0:bits[0]] + "." + s[bits[0]:bits[0]+bits[1]] + "." + s[bits[0]+bits[1]:bits[0]+bits[1]+bits[2]] + "." + s[bits[0]+bits[1]+bits[2]:bits[0]+bits[1]+bits[2]+bits[3]]

		store[storeN[0]] = a
		storeN[0]++
		return
	}

	for i := 1; i <= 3; i++ {
		bits[bitsN] = i

		//这一位是否符合要求。
		if i == 1 {
			if used+1 > N {
				break
			}
		} else if i == 2 {

			if used+2 > N {
				break
			}

			if s[used] == '0' {
				continue
			}
		} else if i == 3 {
			if used+3 > N {
				break
			}
			if s[used] == '0' {
				continue
			} else if s[used] == '1' {

			} else if s[used] == '2' {
				if s[used+1] > '5' {
					continue
				} else if s[used+1] == '5' {
					if s[used+2] > '5' {
						continue
					}
				}
			} else {
				continue
			}
		}
		restoreIpAddressesRecursive(bits, bitsN+1, used+i, N, s, store, storeN)
	}
}
func main() {

	s := "25525511135"
	fmt.Printf("%v\n", restoreIpAddresses(s))
	bits := []int{3, 3, 2, 3}
	var a string
	a = s[0:bits[0]] + "." + s[bits[0]:bits[0]+bits[1]] + "." + s[bits[0]+bits[1]:bits[0]+bits[1]+bits[2]] + "." + s[bits[0]+bits[1]+bits[2]:bits[0]+bits[1]+bits[2]+bits[3]]
	fmt.Println(a)
}

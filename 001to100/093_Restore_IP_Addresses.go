package main

import "fmt"

func restoreIpAddresses(s string) []string {
	store := make([]string, 330)
	bits := make([]int, 4, 4)
	N := len(s)
	restoreIpAddressesRecursive(bits, 0, 0, N, s, store, 0)
	return store
}
func restoreIpAddressesRecursive(bits []int, bitsN int, used int, N int, s string, store []string, storeN int) {

	if bitsN == 4 {

		if bits[0]+bits[1]+bits[2]+bits[3] != N {
			return
		}

		fmt.Printf("%v\n", bits)
		return
	}

	for i := 1; i <= 3; i++ {
		bits[bitsN] = i

		//这一位是否符合要求。
		if i == 1 {
			if used+1 >= N {
				break
			}
		} else if i == 2 {

			if used+2 >= N {
				break
			}

			if s[used+1] == '0' {
				continue
			}
		} else if i == 3 {

			if used+3 >= N {
				break
			}

			if s[used+1] == '0' {
				continue
			} else if s[used+1] == '1' {

			} else if s[used+1] == '2' {
				if s[used+2] > '5' {
					continue
				} else if s[used+2] == '5' {
					if s[used+3] > '5' {
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

	fmt.Printf("%v\n", restoreIpAddresses("25525511135"))

}

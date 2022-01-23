package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Operations struct {
	id     uint64
	values []uint64
	count  uint64
}

var stack []Operations = []Operations{}

func main() {
	// packetHex := "D2FE28"
	// value, version, header := getLiteralValue(hexToBin(packetHex))
	// fmt.Println(value, version, header)

	//packetHex := "38006F45291200"
	//packetHex := "EE00D40C823060"
	//packetHex := "8A004A801A8002F478"
	//packetHex := "620080001611562C8802118E34"
	//packetHex := "C0015000016115A2E0802F182340"
	//packetHex := "A0016C880162017C3686B18A3D4780"

	packetHex := "60556F980272DCE609BC01300042622C428BC200DC128C50FCC0159E9DB9AEA86003430BE5EFA8DB0AC401A4CA4E8A3400E6CFF7518F51A554100180956198529B6A700965634F96C0B99DCF4A13DF6D200DCE801A497FF5BE5FFD6B99DE2B11250034C00F5003900B1270024009D610031400E70020C0093002980652700298051310030C00F50028802B2200809C00F999EF39C79C8800849D398CE4027CCECBDA25A00D4040198D31920C8002170DA37C660009B26EFCA204FDF10E7A85E402304E0E60066A200F4638311C440198A11B635180233023A0094C6186630C44017E500345310FF0A65B0273982C929EEC0000264180390661FC403006E2EC1D86A600F43285504CC02A9D64931293779335983D300568035200042A29C55886200FC6A8B31CE647880323E0068E6E175E9B85D72525B743005646DA57C007CE6634C354CC698689BDBF1005F7231A0FE002F91067EF2E40167B17B503E666693FD9848803106252DFAD40E63D42020041648F24460400D8ECE007CBF26F92B0949B275C9402794338B329F88DC97D608028D9982BF802327D4A9FC10B803F33BD804E7B5DDAA4356014A646D1079E8467EF702A573FAF335EB74906CF5F2ACA00B43E8A460086002A3277BA74911C9531F613009A5CCE7D8248065000402B92D47F14B97C723B953C7B22392788A7CD62C1EC00D14CC23F1D94A3D100A1C200F42A8C51A00010A847176380002110EA31C713004A366006A0200C47483109C0010F8C10AE13C9CA9BDE59080325A0068A6B4CF333949EE635B495003273F76E000BCA47E2331A9DE5D698272F722200DDE801F098EDAC7131DB58E24F5C5D300627122456E58D4C01091C7A283E00ACD34CB20426500BA7F1EBDBBD209FAC75F579ACEB3E5D8FD2DD4E300565EBEDD32AD6008CCE3A492F98E15CC013C0086A5A12E7C46761DBB8CDDBD8BE656780"
	//stack := make([]Operations, 0)
	_, _, sum := parsePacket(hexToBin(packetHex), 0, 0)
	fmt.Println("version sum", sum)
	fmt.Println(stack)
}

func parsePacket(bits []string, state int, sum uint64) (int, int, uint64) {
	ptr := 0
	nextCounter := state
	for ptr+6 <= len(bits) {
		//fmt.Println("pointer", ptr)
		version := binToDec(bits[ptr : ptr+3])
		sum += version
		//fmt.Println("version", version)
		header := binToDec(bits[ptr+3 : ptr+6])
		fmt.Println("type", header)
		if header == 4 {
			value, _, _, consumed := getLiteralValue(bits[ptr:])
			ptr += consumed
			fmt.Println("\tvalue", value)
			operation := stack[len(stack)-1]
			if operation.id == 1 {
				nextCounter += 1
			} else {
				nextCounter += consumed
			}
			operation.values = append(operation.values, value)
			stack[len(stack)-1] = operation
			//if operation.count == uint64(nextCounter) {
			//fmt.Println("done")
			//} else {
			//fmt.Println("next", operation.count, nextCounter, ptr)
			//}
			//return ptr, nextCounter
		} else {
			if ptr+7 >= len(bits) {
				return -1, 0, sum
			}
			id := binToDec(bits[ptr+6 : ptr+7])
			//fmt.Println("id", id)
			offset := 15
			if id == 1 {
				offset = 11
			}
			offset += 7
			if ptr+offset >= len(bits) {
				return -1, 0, sum
			}
			count := binToDec(bits[ptr+7 : ptr+offset])
			//fmt.Println("number", count)

			stack = append(stack, Operations{id, []uint64{}, count})
			nextCounter = 0
			var res int
			if id == 1 {
				res, _, sum = parsePacket(bits[ptr+offset:], 0, sum)
			} else {
				res, _, sum = parsePacket(bits[ptr+offset:], 0, sum)
			}
			if res < 0 {
				return -1, 0, sum
			}
			ptr += res
			return ptr, 0, sum
		}
	}
	return -1, 0, sum
}

func getLiteralValue(bits []string) (uint64, uint64, uint64, int) {
	version := binToDec(bits[:3])
	header := binToDec(bits[3:6])
	//fmt.Println("get value for", version, header, bits[:])
	if header != 4 {
		panic("header not match")
	}
	literal := []string{}
	stop := false
	cnt := 6
	for !stop {
		chunk := bits[cnt+1 : cnt+5]
		stop = bits[cnt] == "0"
		cnt += 5
		literal = append(literal, chunk...)
	}
	return binToDec(literal), version, header, cnt
}

func binToDec(val []string) uint64 {
	num, err := strconv.ParseUint(strings.Join(val, ""), 2, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func hexToBin(hexValue string) []string {
	result := []string{}
	for _, hex := range strings.Split(hexValue, "") {
		val, err := strconv.ParseUint(hex, 16, 4)
		if err != nil {
			panic(err)
			//panic("can not parse " + hex)
		}
		bits := []string{}
		for i := 0; i < 4; i++ {
			bit := val & 0x1
			bitStr := strconv.FormatUint(bit, 2)
			bits = append([]string{bitStr}, bits...)
			val = val >> 1
		}
		result = append(result, bits...)
		//fmt.Println(hex, bits, result)
	}

	return result
}

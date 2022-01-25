package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Stream struct {
	bits []string
	ptr  int
}

var versionSum uint64 = 0

func main() {
	//packetHex := "9C0141080250320F1802104A08"
	packetHex := "60556F980272DCE609BC01300042622C428BC200DC128C50FCC0159E9DB9AEA86003430BE5EFA8DB0AC401A4CA4E8A3400E6CFF7518F51A554100180956198529B6A700965634F96C0B99DCF4A13DF6D200DCE801A497FF5BE5FFD6B99DE2B11250034C00F5003900B1270024009D610031400E70020C0093002980652700298051310030C00F50028802B2200809C00F999EF39C79C8800849D398CE4027CCECBDA25A00D4040198D31920C8002170DA37C660009B26EFCA204FDF10E7A85E402304E0E60066A200F4638311C440198A11B635180233023A0094C6186630C44017E500345310FF0A65B0273982C929EEC0000264180390661FC403006E2EC1D86A600F43285504CC02A9D64931293779335983D300568035200042A29C55886200FC6A8B31CE647880323E0068E6E175E9B85D72525B743005646DA57C007CE6634C354CC698689BDBF1005F7231A0FE002F91067EF2E40167B17B503E666693FD9848803106252DFAD40E63D42020041648F24460400D8ECE007CBF26F92B0949B275C9402794338B329F88DC97D608028D9982BF802327D4A9FC10B803F33BD804E7B5DDAA4356014A646D1079E8467EF702A573FAF335EB74906CF5F2ACA00B43E8A460086002A3277BA74911C9531F613009A5CCE7D8248065000402B92D47F14B97C723B953C7B22392788A7CD62C1EC00D14CC23F1D94A3D100A1C200F42A8C51A00010A847176380002110EA31C713004A366006A0200C47483109C0010F8C10AE13C9CA9BDE59080325A0068A6B4CF333949EE635B495003273F76E000BCA47E2331A9DE5D698272F722200DDE801F098EDAC7131DB58E24F5C5D300627122456E58D4C01091C7A283E00ACD34CB20426500BA7F1EBDBBD209FAC75F579ACEB3E5D8FD2DD4E300565EBEDD32AD6008CCE3A492F98E15CC013C0086A5A12E7C46761DBB8CDDBD8BE656780"

	var stream = Stream{bits: hexToBin(packetHex), ptr: 0}
	res := parse(&stream)
	fmt.Println("part 1", versionSum)
	fmt.Println("part 2", res)
}

func (s *Stream) read(n int) uint64 {
	from := s.ptr
	to := s.ptr + n
	res := binToDec(s.bits[from:to])
	s.ptr += n
	return res
}

func parse(s *Stream) uint64 {
	version := s.read(3)
	versionSum += version
	header := s.read(3)
	if header == 4 {
		value := getLiteralValue(s)
		return value
	}
	values := parseOperation((s))
	res := calculate(header, values)
	return res
}

func getLiteralValue(s *Stream) uint64 {
	literal := []string{}
	stop := false
	cnt := s.ptr
	for !stop {
		stop = s.bits[cnt] == "0"
		chunk := s.bits[cnt+1 : cnt+5]
		cnt += 5
		s.read(5)
		literal = append(literal, chunk...)
	}
	return binToDec(literal)
}

func parseOperation(s *Stream) []uint64 {
	var res []uint64
	id := s.read(1)
	var count int
	if id == 1 {
		count = int(s.read(11))
		it := 0
		for {
			res = append(res, parse(s))
			if it++; it == count {
				break
			}
		}
	} else {
		count = int(s.read(15))
		var from int = s.ptr
		for {
			res = append(res, parse(s))
			if s.ptr >= from+count {
				break
			}
		}
	}
	return res
}

func calculate(id uint64, values []uint64) uint64 {
	switch id {
	case 0:
		var sum uint64
		for _, x := range values {
			sum += x
		}
		return sum
	case 1:
		var product uint64 = 1
		for _, x := range values {
			product *= x
		}
		return product
	case 2:
		var min uint64 = math.MaxInt64
		for _, val := range values {
			min = uint64(utils.Min(int(min), int(val)))
		}
		return min
	case 3:
		var max uint64 = 0
		for _, val := range values {
			max = uint64(utils.Max(int(max), int(val)))
		}
		return max
	case 5:
		if values[0] > values[1] {
			return 1
		}
		return 0
	case 6:
		if values[0] < values[1] {
			return 1
		}
		return 0
	case 7:
		if values[0] == values[1] {
			return 1
		}
		return 0
	}
	return 0
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
		}
		bits := []string{}
		for i := 0; i < 4; i++ {
			bit := val & 0x1
			bitStr := strconv.FormatUint(bit, 2)
			bits = append([]string{bitStr}, bits...)
			val = val >> 1
		}
		result = append(result, bits...)
	}

	return result
}

package main

import (
	"strings"

	"github.com/lukechampine/advent/utils"
)

const inputReps = `Al => ThF
Al => ThRnFAr
B => BCa
B => TiB
B => TiRnFAr
Ca => CaCa
Ca => PB
Ca => PRnFAr
Ca => SiRnFYFAr
Ca => SiRnMgAr
Ca => SiTh
F => CaF
F => PMg
F => SiAl
H => CRnAlAr
H => CRnFYFYFAr
H => CRnFYMgAr
H => CRnMgYFAr
H => HCa
H => NRnFYFAr
H => NRnMgAr
H => NTh
H => OB
H => ORnFAr
Mg => BF
Mg => TiMg
N => CRnFAr
N => HSi
O => CRnFYFAr
O => CRnMgAr
O => HP
O => NRnFAr
O => OTi
P => CaP
P => PTi
P => SiRnFAr
Si => CaSi
Th => ThCa
Ti => BP
Ti => TiTi
e => HF
e => NAl
e => OMg`

const inputMol = `CRnCaCaCaSiRnBPTiMgArSiRnSiRnMgArSiRnCaFArTiTiBSiThFYCaFArCaCaSiThCaPBSiThSiThCaCaPTiRnPBSiThRnFArArCaCaSiThCaSiThSiRnMgArCaPTiBPRnFArSiThCaSiRnFArBCaSiRnCaPRnFArPMgYCaFArCaPTiTiTiBPBSiThCaPTiBPBSiRnFArBPBSiRnCaFArBPRnSiRnFArRnSiRnBFArCaFArCaCaCaSiThSiThCaCaPBPTiTiRnFArCaPTiBSiAlArPBCaCaCaCaCaSiRnMgArCaSiThFArThCaSiThCaSiRnCaFYCaSiRnFYFArFArCaSiRnFYFArCaSiRnBPMgArSiThPRnFArCaSiRnFArTiRnSiRnFYFArCaSiRnBFArCaSiRnTiMgArSiThCaSiThCaFArPRnFArSiRnFArTiTiTiTiBCaCaSiRnCaCaFYFArSiThCaPTiBPTiBCaSiThSiRnMgArCaF`

type replacement struct {
	in, out string
}

func (r replacement) apply(str string) []string {
	var reps []string
	for i := range str {
		if strings.HasPrefix(str[i:], r.in) {
			reps = append(reps, str[:i]+r.out+str[i+len(r.in):])
		}
	}
	return reps
}

func parse(str string) (r replacement) {
	utils.Sscanf(str, "%s => %s", &r.in, &r.out)
	return
}

func main() {
	// part 1
	var reps []replacement
	for _, str := range utils.Lines(inputReps) {
		reps = append(reps, parse(str))
	}
	repMap := make(map[string]struct{})
	for _, rep := range reps {
		for _, s := range rep.apply(inputMol) {
			repMap[s] = struct{}{}
		}
	}
	println(len(repMap))

	// part 2
	reduced := inputMol
	var count int
	for reduced != "e" {
		for _, r := range reps {
			if strings.Contains(reduced, r.out) {
				reduced = strings.Replace(reduced, r.out, r.in, 1)
				count++
			}
		}
	}
	println(count)
}

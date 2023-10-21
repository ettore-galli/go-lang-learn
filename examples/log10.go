package main

import (
	"math"
)

/*

function log10(x) {
  var l=0;
  var p10=1;
  while ((x!=1)&&(p10<9007199254740992)){
      x2=x*x;
      x4=x2*x2;
      x8=x4*x4;
      x=x8*x2;
      var k=0;
      if(x>1) {while(!(x<10)){x=x/10; k++}} else {while((x<1)){x=x*10; k--}}
      p10=p10*10;
      l = l + k/p10;
  }
  return l;
}
*/

func Square(x float64) {}

func tryExp(x float64) (float64, float64) {
	MaxAttempts := 999
	xp := math.Pow(x, 10)
	for n := 0; n < MaxAttempts; n++ {
		if xp > 1 {
			if xp/math.Pow(10, float64(n)) < 10 {
				return float64(n), xp / math.Pow(10, float64(n))
			}
		} else {
			if xp*math.Pow(10, float64(n)) >= 1 {
				return float64(n), xp * math.Pow(10, float64(n))
			}
		}

	}
	return float64(0), float64(0)
}

func log10(x float64) float64 {

	var term float64 = x

	var log float64 = 0
	var fact float64 = 1
	for fact > 0.00000001 {
		n, res := tryExp(term)

		fact = fact * 0.1

		println(n, fact, res, term, log)
		log += fact * n
		term = res

	}
	return log
}

func main() {
	println(log10(11)) // 1.041392685158225
}

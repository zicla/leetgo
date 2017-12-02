package fraction

import "math"

type Fraction struct {
	//分子
	Numerator int
	//分母
	Denominator int
}

func NewMaxFraction() *Fraction {
	return &Fraction{
		Numerator:   math.MaxInt64,
		Denominator: 1,
	}
}

func NewZeroFraction() *Fraction {
	return &Fraction{
		Numerator:   0,
		Denominator: 1,
	}
}

func NewFraction(Numerator, Denominator int) *Fraction {
	a := &Fraction{
		Numerator:   Numerator,
		Denominator: Denominator,
	}
	a.Simplify()
	return a
}

//辗转相除法。
func (this *Fraction) GCD(a, b int) int {
	//保证a>=b
	if a < b {
		a, b = b, a
	}
	if b <= 0 {
		panic("求公约数只能是正整数")
	}
	if a == b {
		return a
	}
	c := a % b
	if c == 0 {
		return b
	} else {
		return this.GCD(b, c)
	}
}

//化简一个分数。
func (this *Fraction) Simplify() {

	//无穷大或者为0不简化了。
	if this.Numerator == math.MaxInt64 || this.Numerator == 0 {
		return
	}
	//保证分母永远大于零，等于零都不行。
	if this.Denominator < 0 {
		this.Numerator = -this.Numerator
		this.Denominator = -this.Denominator
	}
	//此时分母大于0，分子不一定。
	sign := 1
	if this.Numerator < 0 {
		sign = -1
		this.Numerator = -this.Numerator
	}

	//此时分子和分母都是大于0了
	gcd := this.GCD(this.Numerator, this.Denominator)
	this.Numerator = this.Numerator / gcd
	this.Denominator = this.Denominator / gcd

	this.Numerator = this.Numerator * sign
}

//比较两个分数是否相等
func (this *Fraction) Equal(that *Fraction) bool {
	return this.Numerator == that.Numerator && this.Denominator == that.Denominator
}

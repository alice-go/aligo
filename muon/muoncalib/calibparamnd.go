package muoncalib

import (
	"fmt"
	"io"
	"math"
	"strings"
)

func (c *AliMUONCalibParamND) index(i, j int) int {
	return i + int(c.size)*j
}

func (c *AliMUONCalibParamND) ID0() uint32 {
	return c.base.base.ID & 0xFFFF
}

func (c *AliMUONCalibParamND) ID1() uint32 {
	return (c.base.base.ID & 0xFFFF0000) >> 16
}

func (c *AliMUONCalibParamND) Value(i, j int) float64 {
	return c.vs[c.index(i, j)]
}

func (c *AliMUONCalibParamND) MeanAndSigma(dim int) (float64, float64) {
	mean := 0.0
	v2 := 0.0
	n := int(c.size)
	for i := 0; i < n; i++ {
		v := c.Value(i, dim)
		mean += v
		v2 += v * v
	}
	mean /= float64(n)
	sigma := 0.0
	if n > 1 {
		sigma = math.Sqrt((v2 - float64(n)*mean*mean) / (float64(n) - 1))
	}
	return mean, sigma
}

func (c *AliMUONCalibParamND) Print(w io.Writer, opt string) {
	fmt.Fprintf(w, "AliMUONCalibParamND Id=(%d,%d) Size=%d Dimension=%d\n",
		c.ID0(), c.ID1(), c.size, c.dim)
	opt = strings.ToUpper(opt)
	if strings.Contains(opt, "FULL") {
		for i := 0; i < int(c.size); i++ {
			fmt.Fprintf(w, "CH %3d", i)
			for j := 0; j < int(c.dim); j++ {
				fmt.Fprintf(w, " %g", c.Value(i, j))
			}
			fmt.Fprint(w, "\n")
		}
	}
	if strings.Contains(opt, "MEAN") {
		var j int
		fmt.Sscanf(opt, "MEAN%d", &j)
		mean, sigma := c.MeanAndSigma(j)
		fmt.Fprintf(w, " Mean(j=%d)=%g Sigma(j=%d)=%g\n", j, mean, j, sigma)
	}
}

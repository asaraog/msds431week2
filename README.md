## Evaluating Go

### Project Summary

Go will be help power our backend web and database servers and distributed service offerings on the cloud. However, data science operations remain a key concern as Python/R remain popular. This project aims to implement least squares regression of the [Anscombe Quartet (1973)](https://www.sjsu.edu/faculty/gerstman/StatPrimer/anscombe1973.pdf) in Go using the [stats package](https://github.com/montanaflynn/stats). The Go implementation is benchmarked for runtime with a [previous implementation by Miller (2015)](https://github.com/mtpa/mtpa/tree/master/MTPA_Chapter_1) in Python/R as a reference. Python was significantly slower compared to R and Go implementations with runtimes of 1.36s, 0.04s, 0.173s for Python, R and Go. While R was less verbose and a bit faster than Go, Go's testing package ensured identitcal least squares coefficients of 0.5 and 3 for each Anscombe dataset during development. The Data Science team sees test-driven development as an asset in Go and with equivalent statistical results, our concerns about switching to Go are alleviated. We strongly recommend using Go as the primary programming language accross the company.

### Files

*statswithgo_saraogee.go:* \
Main routine loads Anscombe Quartet data into Series types using [LinearRegression](https://github.com/montanaflynn/stats/blob/master/regression.go) in the stats package to output predicted values. An in-built Regress function recalculates the gradient and intercept which is rounded to 2 significant figures with a [roundFloat function](gosamples.dev/round-float/).

*statswithgo_saraogee_test.go:* \
Benchmarking of runtime for go routine and unit test for Regress function. This testing routine ensures equivalence with Pyhon/R output coefficients of 0.5 and 3.

*runtimesPythonR.ipynb:* \
Benchmarking of runtime in a Jupyter notebook. Kernel was changed between R and Python3.11.2 for each cell loading the Anscombe data and performing the regression.

*Week2* \
Unix executable file of compiled Go code for Mac/Windows.

### References

Anscombe, F. J. 1973. “Graphs in Statistical Analysis.” The American Statistician 27 (1): 17–21. https://doi.org/10.2307/2682899. \
Flynn, Montana. 2023. “Stats - Golang Statistics Package.” 2023. https://github.com/montanaflynn/stats. \
Miller, Thomas. 2015. “Modeling Techniques in Predictive Analytics Chapter 1.” 2015. https://github.com/mtpa/mtpa/tree/master/MTPA_Chapter_1.


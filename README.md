### Evaluating Go 

## Summary

Go will be help power our backend web and database servers and distributed service offerings on the cloud. However, data science operations remain a key concern as Python/R remain popular. This project aims to implement least squares regression of the [Anscombe Quartet](https://www.sjsu.edu/faculty/gerstman/StatPrimer/anscombe1973.pdf) in Go using the [stats package](https://github.com/montanaflynn/stats). The Go implementation is benchmarked for runtime with a [previous implementation](https://github.com/mtpa/mtpa/tree/master/MTPA_Chapter_1) in Python/R as a reference. Python was significantly slower compared to R and Go implementations with runtimes of s,s,s for Python, R and Go. While R was faster, Go's testing package ensured identitcal least squares coefficients of 0.5 and 3 for each Anscombe dataset.


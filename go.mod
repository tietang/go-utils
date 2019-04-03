module github.com/tietang/go-utils

go 1.12

//被墙的原因，替换golang.org源为github.com源
replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190322080309-f49334f85ddc

require github.com/sirupsen/logrus v1.4.1

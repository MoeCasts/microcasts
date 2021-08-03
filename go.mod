module github.com/moecasts/microcasts

go 1.15

// 替换成本地包
replace github.com/moecasts/microcasts => ./

replace github.com/moecasts/microcasts/novels => ./novels

require github.com/moecasts/microcasts/novels v0.0.0-00010101000000-000000000000 // indirect

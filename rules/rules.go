package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func rawSQL(m dsl.Matcher) {
	m.Import("gorm.io/gorm")
	m.Match(`$x := fmt.Sprintf($s, $*args); $*_; $db.Raw($x).$_($_)`).
		Where(m["db"].Type.Is("*gorm.DB")).
		Report(`A raw query is used: $s, $args`)
}

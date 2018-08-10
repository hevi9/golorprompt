package sys

type Signs struct {
	cont      string
	plus      string
	cross     string
	dot       string
	dots      string
	star      string
	flag      string
	skull     string
	jobs      string
	level     string
	disk      string
	memory    string
	untracked string
	ahead     string
	behind    string
	diverged  string
	differ    string
	stashes   string
	start     string
	file      string
	dir       string
	todo      string
	action    string
	location  string
	check     string
	load      string
}

type Marks struct {
	cont      rune
	plus      rune
	cross     rune
	dot       rune
	dots      rune
	star      rune
	flag      rune
	skull     rune
	jobs      rune
	level     rune
	disk      rune
	memory    rune
	untracked rune
	ahead     rune
	behind    rune
	diverged  rune
	differ    rune
	stashes   rune
	start     rune
	file      rune
	dir       rune
	todo      rune
	action    rune
	location  rune
	check     rune
}

var sign1 = Signs{
	cont:      "…", // 1 char
	plus:      "✚", // 1 char, bad
	cross:     "✖", // 1 char, bad
	dot:       "●", // 1 char
	dots:      "⛬", // 1char
	star:      "*",
	flag:      "⚑", // 1 char
	skull:     "!", // TODO: change to error
	jobs:      "⚙", // 1 char
	level:     "⮇", // 1 char but bad
	disk:      "o", // 1 char
	memory:    "🖫", // 2 char
	untracked: "?",
	ahead:     "⭱", // 1 char
	behind:    "⭳", // 1 char
	diverged:  "⭿", // 1 char
	differ:    "⭾", // 1 char
	stashes:   "≡", // 1 char
	start:     "▶", // 1 char
	file:      "🗎", // 2 char
	dir:       "📁", // 2 char
	todo:      "🔨", // 2 char
	action:    "↯", // 1 char
	location:  "⌘", // 1char
	check:     "🗹", // 2 char
}

var sign2 = Signs{
	cont:      "…", // 1 char
	plus:      "✚", // 1 char, bad
	cross:     "✖", // 1 char, bad
	dot:       "●", // 1 char
	dots:      "⛬", // 1char
	star:      "🟉", // 2 char
	flag:      "⚑", // 1 char
	skull:     "🕱", // 2 char
	jobs:      "⚙", // 1 char
	level:     "⮇", // 1 char but bad
	disk:      "🖸", // 2 char
	memory:    "🖫", // 2 char
	untracked: "?",
	ahead:     "⭱", // 1 char
	behind:    "⭳", // 1 char
	diverged:  "⭿", // 1 char
	differ:    "⭾", // 1 char
	stashes:   "≡", // 1 char
	start:     "▶", // 1 char
	//start:     "🡆",
	file:     "🗎", // 2 char
	dir:      "📁", // 2 char
	todo:     "🔨", // 2 char
	action:   "↯", // 1 char
	location: "⌘", // 1char
	check:    "🗹", // 2 char
}

// non problematic characters
var sign3 = Signs{
	cont:      "…",
	plus:      "+",
	cross:     "x",
	dot:       ".",
	dots:      "..",
	star:      "*",
	flag:      "F",
	skull:     "!",
	jobs:      "⚙",
	level:     "L",
	disk:      "D",
	memory:    "M",
	untracked: "?",
	ahead:     "⭱",
	behind:    "⭳",
	diverged:  "⁑",
	differ:    "±",
	stashes:   "≡",
	start:     "▶",
	file:      "F",
	dir:       "D",
	todo:      "T",
	action:    "A",
	location:  "l",
	check:     "c",
	load:      "⌆",
}

var sign = sign3

var mark1 = Marks{
	cont:      '…', // 1 char
	plus:      '✚', // 1 char, bad
	cross:     '✖', // 1 char, bad
	dot:       '●', // 1 char
	dots:      '⛬', // 1char
	star:      '🟉', // 2 char
	flag:      '⚑', // 1 char
	skull:     '🕱', // 2 char
	jobs:      '⚙', // 1 char
	level:     '⮇', // 1 char but bad
	disk:      '🖸', // 2 char
	memory:    '🖫', // 2 char
	untracked: '?',
	ahead:     '⭱', // 1 char
	behind:    '⭳', // 1 char
	diverged:  '⭿', // 1 char
	differ:    '⭾', // 1 char
	stashes:   '≡', // 1 char
	start:     '▶', // 1 char
	file:      '🗎', // 2 char
	dir:       '📁', // 2 char
	todo:      '🔨', // 2 char
	action:    '↯', // 1 char
	location:  '⌘', // 1char
	check:     '🗹', // 2 char
}

var mark = mark1

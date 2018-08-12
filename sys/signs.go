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
	Level     string
	disk      string
	memory    string
	untracked string
	Ahead     string
	Behind    string
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
	Load      string
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
	Level:     "⮇", // 1 char but bad
	disk:      "o", // 1 char
	memory:    "🖫", // 2 char
	untracked: "?",
	Ahead:     "⭱", // 1 char
	Behind:    "⭳", // 1 char
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
	Level:     "⮇", // 1 char but bad
	disk:      "🖸", // 2 char
	memory:    "🖫", // 2 char
	untracked: "?",
	Ahead:     "⭱", // 1 char
	Behind:    "⭳", // 1 char
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
	Level:     "L",
	disk:      "D",
	memory:    "M",
	untracked: "?",
	Ahead:     "⭱",
	Behind:    "⭳",
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
	Load:      "⌆",
}

var Sign = sign3

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
	Ahead     rune
	Behind    rune
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
	Ahead:     '⭱', // 1 char
	Behind:    '⭳', // 1 char
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

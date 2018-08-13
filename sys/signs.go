package sys

type Signs struct {
	Cont      string
	Plus      string
	Cross     string
	Dot       string
	Dots      string
	Star      string
	Flag      string
	Skull     string
	Jobs      string
	Level     string
	Disk      string
	Memory    string
	Untracked string
	Ahead     string
	Behind    string
	Diverged  string
	Differ    string
	Stashes   string
	Start     string
	File      string
	Dir       string
	Todo      string
	Action    string
	Location  string
	Check     string
	Load      string
}

var sign1 = Signs{
	Cont:      "…", // 1 char
	Plus:      "✚", // 1 char, bad
	Cross:     "✖", // 1 char, bad
	Dot:       "●", // 1 char
	Dots:      "⛬", // 1char
	Star:      "*",
	Flag:      "⚑", // 1 char
	Skull:     "!", // TODO: change to error
	Jobs:      "⚙", // 1 char
	Level:     "⮇", // 1 char but bad
	Disk:      "o", // 1 char
	Memory:    "🖫", // 2 char
	Untracked: "?",
	Ahead:     "⭱", // 1 char
	Behind:    "⭳", // 1 char
	Diverged:  "⭿", // 1 char
	Differ:    "⭾", // 1 char
	Stashes:   "≡", // 1 char
	Start:     "▶", // 1 char
	File:      "🗎", // 2 char
	Dir:       "📁", // 2 char
	Todo:      "🔨", // 2 char
	Action:    "↯", // 1 char
	Location:  "⌘", // 1char
	Check:     "🗹", // 2 char
}

var sign2 = Signs{
	Cont:      "…", // 1 char
	Plus:      "✚", // 1 char, bad
	Cross:     "✖", // 1 char, bad
	Dot:       "●", // 1 char
	Dots:      "⛬", // 1char
	Star:      "🟉", // 2 char
	Flag:      "⚑", // 1 char
	Skull:     "🕱", // 2 char
	Jobs:      "⚙", // 1 char
	Level:     "⮇", // 1 char but bad
	Disk:      "🖸", // 2 char
	Memory:    "🖫", // 2 char
	Untracked: "?",
	Ahead:     "⭱", // 1 char
	Behind:    "⭳", // 1 char
	Diverged:  "⭿", // 1 char
	Differ:    "⭾", // 1 char
	Stashes:   "≡", // 1 char
	Start:     "▶", // 1 char
	//Start:     "🡆",
	File:     "🗎", // 2 char
	Dir:      "📁", // 2 char
	Todo:     "🔨", // 2 char
	Action:   "↯", // 1 char
	Location: "⌘", // 1char
	Check:    "🗹", // 2 char
}

// non problematic characters
var sign3 = Signs{
	Cont:      "…",
	Plus:      "+",
	Cross:     "x",
	Dot:       ".",
	Dots:      "..",
	Star:      "*",
	Flag:      "F",
	Skull:     "!",
	Jobs:      "⚙",
	Level:     "L",
	Disk:      "D",
	Memory:    "M",
	Untracked: "?",
	Ahead:     "⭱",
	Behind:    "⭳",
	Diverged:  "⁑",
	Differ:    "±",
	Stashes:   "≡",
	Start:     "▶",
	File:      "F",
	Dir:       "D",
	Todo:      "T",
	Action:    "A",
	Location:  "l",
	Check:     "c",
	Load:      "⌆",
}

var Sign = sign3

package main

import (
	"strconv"
	"log"
	"github.com/lucasb-eyer/go-colorful"
	"fmt"
)

func init() {
	SegRegister("exitcode", "Show program exit code if non zero",
		func() Segment { return &ExitCode{} })
}

type ExitCode struct {
	Threshold int
}

// http://man7.org/linux/man-pages/man7/signal.7.html
var linuxSignalsX86 = map[int]string{
	1:  "SIGHUP",    //         1       Term    Hangup detected on controlling terminal*/ or death of controlling process
	2:  "SIGINT",    //        2       Term    Interrupt from keyboard
	3:  "SIGQUIT",   //       3       Core    Quit from keyboard
	4:  "SIGILL",    //  4       Core    Illegal Instruction
	6:  "SIGABRT",   //       6       Core    Abort signal from abort(3)
	8:  "SIGFPE",    //        8       Core    Floating-point exception
	9:  "SIGKILL",   //       9       Term    Kill signal
	11: "SIGSEGV",   //      11       Core    Invalid memory reference
	13: "SIGPIPE",   //      13       Term    Broken pipe: write to pipe with no*/ readers; see pipe(7)
	14: "SIGALRM",   //      14       Term    Timer signal from alarm(2)
	15: "SIGTERM",   //      15       Term    Termination signal
	10: "SIGUSR1",   //   30, 10, 16    Term    User-defined signal 1
	12: "SIGUSR2",   //   31, 12, 17    Term    User-defined signal 2
	17: "SIGCHLD",   //   20, 17, 18    Ign     Child stopped or terminated
	18: "SIGCONT",   //  19, 18, 25    Cont    Continue if stopped
	19: "SIGSTOP",   //   17, 19, 23    Stop    Stop process
	20: "SIGTSTP",   //   18, 20, 24    Stop    Stop typed at terminal
	21: "SIGTTIN",   //   21, 21, 26    Stop    Terminal input for background process
	22: "SIGTTOU",   //   22, 22, 27    Stop    Terminal output for background process
	7:  "SIGBUS",    //      10, 7, 10     Core    Bus error (bad memory access)
	27: "SIGPROF",   //     27, 27, 29    Term    Profiling timer expired
	31: "SIGSYS",    //      12, 31, 12    Core    Bad system call (SVr4);*/ see also seccomp(2)
	5:  "SIGTRAP",   //        5        Core    Trace/breakpoint trap
	23: "SIGURG",    //      16, 23, 21    Ign     Urgent condition on socket (4.2BSD)
	26: "SIGVTALRM", //   26, 26, 28    Term    Virtual alarm clock (4.2BSD)
	24: "SIGXCPU",   //     24, 24, 30    Core    CPU time limit exceeded (4.2BSD);*/ see setrlimit(2)
	25: "SIGXFSZ",   //     25, 25, 31    Core    File size limit exceeded (4.2BSD);*/ see setrlimit(2)
}

func getRcDesc(rc int) (string, colorful.Color) {
	if rc == 126 {
		return "NOPERM", colorful.Hsv(15.0, config.FgSaturation, config.FgValue)
	}
	if rc == 127 {
		return "NOTFOUND", colorful.Hsv(30.0, config.FgSaturation, config.FgValue)
	}
	value, exits := linuxSignalsX86[rc-128]
	if exits {
		return value, colorful.Hsv(360.0-30.0, config.FgSaturation, config.FgValue)
	}
	return "", colorful.Hsv(0.0, config.FgSaturation, config.FgValue)
}

func (self *ExitCode) Render() []Chunk {
	rcStr, exists := config.Args["RC"]
	if ! exists {
		return nil
	}
	rc, err := strconv.Atoi(rcStr)
	if err != nil {
		log.Printf("strconv.Atoi(%s): %s", rcStr, err)
		return []Chunk{Chunk{text: "ERR", fg: colorful.Hsv(0.0, config.FgSaturation, config.FgValue)}}
	}
	if rc < 0 || rc > 255 {
		log.Printf("Invalid RC=%d", rc)
		return []Chunk{Chunk{text: "ERR", fg: colorful.Hsv(0.0, config.FgSaturation, config.FgValue)}}
	}
	if rc == 0 {
		return nil
	}
	chunks := make([]Chunk, 0)
	desc, color := getRcDesc(rc)

	if len(desc) > 0 {
		chunks = append(chunks, Chunk{
			text: fmt.Sprintf("%s ", desc),
			fg:   color,
		})
	}
	chunks = append(chunks, Chunk{
		text: fmt.Sprintf("%d%s", rc, sign.skull),
		fg:   color,
	})
	return chunks
}

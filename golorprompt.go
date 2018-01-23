// TODO: IDEA @ as a alone hostname hash colored

package main

import (
    "fmt"
    "bytes"
    "log"
    "gopkg.in/alecthomas/kingpin.v2"
    "os"
    "encoding/json"
    "runtime/pprof"
    "runtime/trace"
)

type MakeSegFunc func() Segment

var makeSegFromJson = map[string]MakeSegFunc{}

var segmentDesc = map[string]string{}

func SegRegister(name string, desc string, makeFn MakeSegFunc) {
    makeSegFromJson[name] = makeFn
    segmentDesc[name] = desc
}

func buildWidgetsFromJson(msgs []json.RawMessage) []Widget {
    widgets := make([]Widget, 0)
    for i, msg := range msgs {
        spec := PartsConfigSegType{}
        if err := json.Unmarshal(msg, &spec); err != nil {
            log.Printf("Error: no seg: %s", err)
            os.Exit(1)
        }
        fn, exists := makeSegFromJson[spec.Seg]
        if ! exists {
            log.Printf("Error: seg %s does not exists", spec.Seg)
            continue
        }
        segment := fn()
        if err := json.Unmarshal(msg, segment); err != nil {
            log.Printf("Error: cannot unmarshal segment %s params: %s", spec.Seg, err)
            continue
        }
        //log.Printf("segment %#v", segment)
        // Add spaces between segments
        if i != 0 && i < len(msgs) {
            widgets = append(widgets, Space())
        }
        widgets = append(widgets, SegmentWidget(segment))
    }
    return widgets
}

func printPrompt() {
    // Build widgets from json
    partsConfig := PartsConfig{}
    if err := json.Unmarshal([]byte(defaultJson), &partsConfig); err != nil {
        log.Printf("Error: json.Unmarshal([]byte(defaultJson), &partsConfig): %s", err)
        os.Exit(1)
    }
    widgetsLeft := buildWidgetsFromJson(partsConfig.Left)
    widgetsRight := buildWidgetsFromJson(partsConfig.Right)
    widgetsStart := buildWidgetsFromJson(partsConfig.Start)

    // Make infoline
    widgetsLine := make([]Widget, 0)
    widgetsLine = append(widgetsLine, Space()) // margin
    widgetsLine = append(widgetsLine, widgetsLeft...)
    filler := Filler()
    widgetsLine = append(widgetsLine, filler)
    widgetsLine = append(widgetsLine, widgetsRight...)
    widgetsLine = append(widgetsLine, Space()) // margin

    // Make start part
    widgetsStart = append(widgetsStart, Space())

    // Allocate
    empty := allocateSimple(widgetsLine, GetWidth())
    filler.Allocate(empty)
    allocateSimple(widgetsStart, GetWidth())

    // Print line
    // TODO: Shell escape $'s in output
    // TODO: Shell escape \'s in output
    //fmt.Printf("%s\n", strings.Repeat("-", lenTerm))
    buf := bytes.Buffer{}
    buf.WriteString(bg(config.BgLine))
    for _, w := range widgetsLine {
        for _, c := range w.Chunks() {
            buf.WriteString(fg(c.fg))
            buf.WriteString(shellEscape(c.text))
        }
    }
    buf.WriteString(rz())
    buf.WriteString("\n")
    for _, w := range widgetsStart {
        for _, c := range w.Chunks() {
            buf.WriteString(fg(c.fg))
            buf.WriteString(shellEscape(c.text))
        }
    }
    buf.WriteString(rz())
    fmt.Printf("%s", buf.String())
}

var (
    app = kingpin.New("ingoline", "in go line shell prompt")

    // TODO: addd debugging/logging to file
    debug     = app.Flag("debug", "Debug to log file").String()
    //traceFlag = app.Flag("trace", "Trace to file").String()

    prompt         = app.Command("prompt", "Make prompt").Default()
    promptNpcStart = prompt.Flag(
        "npc-start", "Non-Printing Chanracter sequence start").Short('s').
        Default("").String()
    promptNpcEnd = prompt.Flag(
        "npc-end", "Non-Printing Chanracter sequence end").Short('e').
        Default("").String()
    promptArgs = prompt.Arg("arg", "prompt data: key=value .. form").StringMap()

    show = app.Command("show", "Show parameters")
)

func main() {
    cmd := kingpin.MustParse(app.Parse(os.Args[1:]))
    //if traceFlag != nil {
    //    if err := startTrace(traceFlag); err != nil {
    //        os.Exit(1)
    //    }
    //}
    switch cmd {
    case show.FullCommand():
        printCurrentTheme()
        //os.Exit(0)
    case prompt.FullCommand():
        config.NpcStart = *promptNpcStart
        config.NpcEnd = *promptNpcEnd
        config.Args = *promptArgs
        printPrompt()
    }
}

func startProfile(filename *string) (err error) {
    file, err := os.Create(*filename)
    if err != nil {
        log.Printf("Error: Cannot create startProfile file %s: %s", filename, err)
        return
    }
    if err = pprof.StartCPUProfile(file); err != nil {
        log.Printf("Error: could not start CPU profiling: %s", err)
        return
    }
    defer pprof.StopCPUProfile()
    return
}

func startTrace(filename *string) (err error) {
    file, err := os.Create(*filename)
    if err != nil {
        log.Printf("Error: could not create traceFlag file %s: %s", filename, err)
        return
    }
    if err = trace.Start(file); err != nil {
        log.Printf("Error: could not start traceFlag: %s", err)
        return
    }
    defer trace.Stop()
    return
}

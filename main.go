//------------------------------------------------------------------------------
//------------------------------------------------------------------------------
//
// Tyler(UnclassedPenguin) Binary Clock (tcell version) 2022
//
//      Author: Tyler(UnclassedPenguin)
//         URL: https://unclassed.ca
//      GitHub: https://github.com/UnclassedPenguin/scripts/
// Description: I just wanted a simple binary clock. Now in Go!
//
//------------------------------------------------------------------------------
//------------------------------------------------------------------------------

package main

import (
  "os"
  "fmt"
  "time"
  "github.com/gdamore/tcell/v2"
)

// Draw a box
func drawBox(s tcell.Screen, style tcell.Style, x1, y1, x2, y2 int) {
  if y2 < y1 {
    y1, y2 = y2, y1
  }
  if x2 < x1 {
    x1, x2 = x2, x1
  }

  for col := x1; col <= x2; col++ {
    s.SetContent(col, y1, tcell.RuneHLine, nil, style)
    s.SetContent(col, y2, tcell.RuneHLine, nil, style)
  }
  for row := y1 + 1; row < y2; row++ {
    s.SetContent(x1, row, tcell.RuneVLine, nil, style)
    s.SetContent(x2, row, tcell.RuneVLine, nil, style)
  }
  if y1 != y2 && x1 != x2 {
    // Only add corners if we need to
    s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
    s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
    s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
    s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
  }
}

func main() {
  s, err := tcell.NewScreen()
  if err != nil {
    fmt.Println("Error in tcell.NewScreen:", err)
  }

  if err := s.Init(); err != nil {
    fmt.Println("Error initializing screen:", err)
    os.Exit(1)
  }

  s.Clear()

  style := tcell.StyleDefault.Foreground(tcell.ColorWhite)

  x, y := s.Size()

  if x < 16 || y < 6 {
    s.Fini()
    fmt.Println("Terminal too small")
    fmt.Println("Must be at least 16x6")
    fmt.Println("Resize and retry")
    os.Exit(1)
  }

  // Corners of the box to draw.
  x1 := x/2-8
  x2 := x/2+6
  y1 := y/2-3
  y2 := y/2+2

  // Handles keyboard input. ctrl-c, q, or esc to quit.
  go func() {
    for {
    switch ev := s.PollEvent().(type) {
      case *tcell.EventResize:
        s.Sync()
      case *tcell.EventKey:
        switch ev.Key() {
        case tcell.KeyCtrlC, tcell.KeyEscape:
          s.Fini()
          os.Exit(0)
        case tcell.KeyRune:
          switch ev.Rune() {
          case 'q', 'Q':
            s.Fini()
            os.Exit(0)
          }
        }
      }
    }
  }()

  // Main loop
  for {
    timeArr := ReturnTime()

    drawBox(s, style, x1, y1, x2, y2)

    // Hour
    // Tens position
    if timeArr[3][0] == "1" {
      s.SetContent(x/2-7+1, y/2+1, tcell.RuneDiamond, []rune{}, style)
    }
    if timeArr[2][0] == "1" {
      s.SetContent(x/2-7+1, y/2, tcell.RuneDiamond, []rune{}, style)
    }

    // Ones position
    if timeArr[3][1] == "1" {
      s.SetContent(x/2-7+3, y/2+1, tcell.RuneDiamond, []rune{}, style)
    }
    if timeArr[2][1] == "1" {
      s.SetContent(x/2-7+3, y/2, tcell.RuneDiamond, []rune{}, style)
    }
    if timeArr[1][1] == "1" {
      s.SetContent(x/2-7+3, y/2-1, tcell.RuneDiamond, []rune{}, style)
    }
    if timeArr[0][1] == "1" {
      s.SetContent(x/2-7+3, y/2-2, tcell.RuneDiamond, []rune{}, style)
    }

    // Minutes
    // Tens position
    if timeArr[3][2] == "1" {
      s.SetContent(x/2-7+5, y/2+1, tcell.RuneDiamond, []rune{}, style)
    }
    if timeArr[2][2] == "1" {
      s.SetContent(x/2-7+5, y/2, tcell.RuneDiamond, []rune{}, style)
    }
    if timeArr[1][2] == "1" {
      s.SetContent(x/2-7+5, y/2-1, tcell.RuneDiamond, []rune{}, style)
    }

    // Ones position
    if timeArr[3][3] == "1" {
      s.SetContent(x/2-7+7, y/2+1, tcell.RuneDiamond, []rune{}, style)
    }
    if timeArr[2][3] == "1" {
      s.SetContent(x/2-7+7, y/2, tcell.RuneDiamond, []rune{}, style)
    }
    if timeArr[1][3] == "1" {
      s.SetContent(x/2-7+7, y/2-1, tcell.RuneDiamond, []rune{}, style)
    }
    if timeArr[0][3] == "1" {
      s.SetContent(x/2-7+7, y/2-2, tcell.RuneDiamond, []rune{}, style)
    }

    // Seconds
    // Tens position
    if timeArr[3][4] == "1" {
      s.SetContent(x/2-7+9, y/2+1, tcell.RuneDiamond, []rune{}, style)
    }
    if timeArr[2][4] == "1" {
      s.SetContent(x/2-7+9, y/2, tcell.RuneDiamond, []rune{}, style)
    }
    if timeArr[1][4] == "1" {
      s.SetContent(x/2-7+9, y/2-1, tcell.RuneDiamond, []rune{}, style)
    }

    // Ones position
    if timeArr[3][5] == "1" {
      s.SetContent(x/2-7+11, y/2+1, tcell.RuneDiamond, []rune{}, style)
    }
    if timeArr[2][5] == "1" {
      s.SetContent(x/2-7+11, y/2, tcell.RuneDiamond, []rune{}, style)
    }
    if timeArr[1][5] == "1" {
      s.SetContent(x/2-7+11, y/2-1, tcell.RuneDiamond, []rune{}, style)
    }
    if timeArr[0][5] == "1" {
      s.SetContent(x/2-7+11, y/2-2, tcell.RuneDiamond, []rune{}, style)
    }

    s.Sync()
    time.Sleep(time.Second * 1)
    s.Clear()
  }
}

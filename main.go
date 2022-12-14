package main

import (
  "os"
  "fmt"
  "time"
  "github.com/gdamore/tcell/v2"
)

// This is used just to write strings to the screen.
func writeToScreen(s tcell.Screen, style tcell.Style, x int, y int, str string) {
  for i, char := range str {
    s.SetContent(x+i, y, rune(char), []rune{}, style)
  }
}

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


  //s.SetContent(x/2-7+1, y/2+1, tcell.RuneBlock, []rune{}, style)
  //s.SetContent(x/2-7+3, y/2+1, tcell.RuneBlock, []rune{}, style)
  //s.SetContent(x/2-7+5, y/2+1, tcell.RuneBlock, []rune{}, style)
  //s.SetContent(x/2-7+7, y/2+1, tcell.RuneBlock, []rune{}, style)

  //s.SetContent(x/2-7+9, y/2+1, tcell.RuneBlock, []rune{}, style)
  //s.SetContent(x/2-7+11, y/2+1, tcell.RuneBlock, []rune{}, style)

  //writeToScreen(s, style, i, j, string(tcell.RuneBlock))

  //writeToScreen(s, style, x/2-4, y/2, "Welcome")

  s.Sync()

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
    //line4[2] = string(currentMinuteBin0[7])
    if timeArr[3][2] == "1" {
      s.SetContent(x/2-7+5, y/2+1, tcell.RuneDiamond, []rune{}, style)
    }
    //line3[2] = string(currentMinuteBin0[6])
    if timeArr[2][2] == "1" {
      s.SetContent(x/2-7+5, y/2, tcell.RuneDiamond, []rune{}, style)
    }
    //line2[2] = string(currentMinuteBin0[5])
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
    //line4[5] = string(currentSecondBin1[7])
    if timeArr[3][5] == "1" {
      s.SetContent(x/2-7+11, y/2+1, tcell.RuneDiamond, []rune{}, style)
    }
    //line3[5] = string(currentSecondBin1[6])
    if timeArr[2][5] == "1" {
      s.SetContent(x/2-7+11, y/2, tcell.RuneDiamond, []rune{}, style)
    }

    //line2[5] = string(currentSecondBin1[5])
    if timeArr[1][5] == "1" {
      s.SetContent(x/2-7+11, y/2-1, tcell.RuneDiamond, []rune{}, style)
    }

    //line1[5] = string(currentSecondBin1[4])
    if timeArr[0][5] == "1" {
      s.SetContent(x/2-7+11, y/2-2, tcell.RuneDiamond, []rune{}, style)
    }

    s.Sync()
    time.Sleep(time.Second * 1)
    s.Clear()
  }
}
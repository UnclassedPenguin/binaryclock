package main

import (
  "fmt"
  "time"
  "strconv"
)

// Gets the current time, and returns it as strings of HH, MM, SS, with 
// leading 0's ie: 08, 10, 02
func getTime() (string, string, string) {
  currentTime := time.Now()

  currentHour := currentTime.Hour()
  var currentHourString string
  if currentHour < 10 {
    currentHourString = fmt.Sprintf("%02d", currentHour)
  } else {
    currentHourString = strconv.Itoa(currentHour)
  }

  currentMinute := currentTime.Minute()
  var currentMinuteString string
  if currentMinute < 10 {
    currentMinuteString = fmt.Sprintf("%02d", currentMinute)
  } else {
    currentMinuteString = strconv.Itoa(currentMinute)
  }

  currentSecond := currentTime.Second()
  var currentSecondString string
  if currentSecond < 10 {
    currentSecondString = fmt.Sprintf("%02d", currentSecond)
  } else {
    currentSecondString = strconv.Itoa(currentSecond)
  }

  return currentHourString, currentMinuteString, currentSecondString
}

// Takes a single digit, and returns a binary value with all the leading zeros.
// ie, takes 7, returns 00000111
func getBinary(num int) string {
  binary := fmt.Sprintf("%08b", num)
  return binary
}

func ReturnTime() [][]string {

  // Gets current hour, minute, second
  currentHourString, currentMinuteString, currentSecondString := getTime()

  // Takes the current hour and splits it into two individual characters
  // These are each then passed to the getBinary function to get the
  // binary value of them, and stores it. 
  // currentHourString0 is the tens position of the hour
  // currentHourString1 is the ones position of the hour
  // For example: if it is 2 o'clock in the afternoon, or 14 rather,
  // currentHourString0 will be 1, currentHourString1 will be 4
  // currentHourBin0 will be 00000001, currentHourBin1 will be 00000100
  currentHourString0 := string(currentHourString[0])
  currentHourInt0, _ := strconv.Atoi(currentHourString0)
  currentHourBin0 := getBinary(currentHourInt0)

  currentHourString1 := string(currentHourString[1])
  currentHourInt1, _ := strconv.Atoi(currentHourString1)
  currentHourBin1 := getBinary(currentHourInt1)

  // Same as above, but for Minutes
  currentMinuteString0 := string(currentMinuteString[0])
  currentMinuteInt0, _ := strconv.Atoi(currentMinuteString0)
  currentMinuteBin0 := getBinary(currentMinuteInt0)

  currentMinuteString1 := string(currentMinuteString[1])
  currentMinuteInt1, _ := strconv.Atoi(currentMinuteString1)
  currentMinuteBin1 := getBinary(currentMinuteInt1)

  // Same as above, but for Seconds
  currentSecondString0 := string(currentSecondString[0])
  currentSecondInt0, _ := strconv.Atoi(currentSecondString0)
  currentSecondBin0 := getBinary(currentSecondInt0)

  currentSecondString1 := string(currentSecondString[1])
  currentSecondInt1, _ := strconv.Atoi(currentSecondString1)
  currentSecondBin1 := getBinary(currentSecondInt1)

  // Diagnostics printing
  //fmt.Print(currentHourString)
  //fmt.Print(":")
  //fmt.Print(currentMinuteString)
  //fmt.Print(":")
  //fmt.Print(currentSecondString)
  //fmt.Print("\n")
  //fmt.Println("Current Hour String: ", currentHourString)
  //fmt.Println("Current Hour String 0: ", currentHourString0)
  //fmt.Println("Current Hour String 1: ", currentHourString1)
  //fmt.Println("Current Hour Bin 0: ", currentHourBin0)
  //fmt.Println("Current Hour Bin 1: ", currentHourBin1)
  //fmt.Println("Current Minute String: ", currentMinuteString)
  //fmt.Println("Current Minnute String 0: ", currentMinuteString0)
  //fmt.Println("Current Minute String 1: ", currentMinuteString1)
  //fmt.Println("Current Minute Bin 0: ", currentMinuteBin0)
  //fmt.Println("Current Minute Bin 1: ", currentMinuteBin1)
  //fmt.Println("Current Second String: ", currentSecondString)
  //fmt.Println("Current Second String 0: ", currentSecondString0)
  //fmt.Println("Current Second String 1: ", currentSecondString1)
  //fmt.Println("Current Second Bin 0: ", currentSecondBin0)
  //fmt.Println("Current Second Bin 1: ", currentSecondBin1)

  // These are the arrays where the values are changed and the time is shown
  var line1 = []string{" ","0"," ","0"," ","0"}
  var line2 = []string{" ","0","0","0","0","0"}
  var line3 = []string{"0","0","0","0","0","0"}
  var line4 = []string{"0","0","0","0","0","0"}

  // Here we replace the values in the arrays with the corresponding value from
  // the binary of the current time

  // Hour
  // Tens position
  line4[0] = string(currentHourBin0[7])
  line3[0] = string(currentHourBin0[6])

  // Ones position
  line4[1] = string(currentHourBin1[7])
  line3[1] = string(currentHourBin1[6])
  line2[1] = string(currentHourBin1[5])
  line1[1] = string(currentHourBin1[4])

  // Minutes
  // Tens position
  line4[2] = string(currentMinuteBin0[7])
  line3[2] = string(currentMinuteBin0[6])
  line2[2] = string(currentMinuteBin0[5])

  // Ones position
  line4[3] = string(currentMinuteBin1[7])
  line3[3] = string(currentMinuteBin1[6])
  line2[3] = string(currentMinuteBin1[5])
  line1[3] = string(currentMinuteBin1[4])

  // Seconds
  // Tens position
  line4[4] = string(currentSecondBin0[7])
  line3[4] = string(currentSecondBin0[6])
  line2[4] = string(currentSecondBin0[5])

  // Ones position
  line4[5] = string(currentSecondBin1[7])
  line3[5] = string(currentSecondBin1[6])
  line2[5] = string(currentSecondBin1[5])
  line1[5] = string(currentSecondBin1[4])

  return [][]string{line1, line2, line3, line4}
}

package main

import "news.com/event/genlib"

func main() {
	name := genlib.OpenFileFromScreen()
	genlib.Check_File_Info(name)
	genlib.OpenandChange_File(name)
}

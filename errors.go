package main

const (
	flags = "\nFlags\t:" +
		"\n\t-n : Number of elements to display\t(-n=20)" +
		"\n\t-h : Help" +
		"\n\t-s : Sort elements by score" +
		"\n\t-t : Filter post not tagged <tag>\t(-t=c++)\n"
	helpmeUsage = "\nUsage of helpme -> helpme <flags> <research content>\n" + flags
	formatError = "\nFormat error :\n" + helpmeUsage
)

package greetings


type Artifact interface{
	Name() string
	Author()	string
	Age()	int32
}
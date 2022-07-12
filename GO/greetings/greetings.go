package greetings

import fmt 

//Hello returns a greeting for named person

func Hello(name string) string {
	//Return a greeting that embeds name in the message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

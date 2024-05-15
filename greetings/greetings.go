package greetings

import (
    "fmt"
    "errors"
    "math/rand"
)

// Returns a greeting for the named person
// func <func-name> <param-type> <return-type>
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }
    // Return a string the embeds the name arg into a message.
    // var message string = fmt.Sprintf(randomFormat(), name)
    message := fmt.Sprintf(randomFormat())
    // or, type inference:
    // message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}

func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages
    messages := make(map[string]string)

    // Loop through the received slice of names,
    // calling the Hello function to get a message for each name

    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }

        // In the map, associate the retrieved message with
        // the name.

        messages[name] = message
    }
    return messages, nil
}

func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format
    return formats[rand.Intn(len(formats))]
}

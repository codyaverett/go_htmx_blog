package db

import (
	"fmt"
	"os/exec"
)

func SendEmail(to, subject, body string) error {

	fmt.Println("Sending email to:", to)
	// Prepare the command to send mail using sendmail
	cmd := exec.Command("/usr/sbin/sendmail", "-t", "-oi")

	// Create the email headers and body
	email := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", to, subject, body)

	// Get the stdin pipe
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	// Start the command
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return err
	}

	// Write the email content to the stdin pipe
	_, err = stdin.Write([]byte(email))
	if err != nil {
		return err
	}

	// Close stdin and wait for the command to finish
	err = stdin.Close()
	if err != nil {
		return err
	}

	return cmd.Wait()
}

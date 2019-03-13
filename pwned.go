/*
Package pwned.go
App: pwned
Queries the haveibeenpwned.com API for breached passwords.
The entered password is hashed with SHA1. The first five chars of the
hash are sent to the API which returns the hashes of all the passwords
whose first five chars match, along with a count of how many occurrences
of that password are in the database. (The returned hashes are missing
the first five chars because we already know those.)
The program then cycles through the returned hashes looking for that
created by our password.

Offered up under GPL 3.0 but absolutely not guaranteed fit for use.
This is code created by hobbyist coder, so use at your own risk.
Github: https://github.com/mspeculatrix
Blog: https://mansfield-devine.com/speculatrix/
*/

package pwned

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"net/http"
	"os"
	"strings"
)

/******************************************************************************
 *****   MAIN                                                             *****
 ******************************************************************************/
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments. Usage: pwned <password>")
		os.Exit(1)
	}
	const urlStem = "https://api.pwnedpasswords.com/range/"
	pw := os.Args[1]                            // grab the entered password
	hash := sha1.New()                          // create new SHA1 hash
	hash.Write([]byte(pw))                      // run our password through it
	hashStr := fmt.Sprintf("%x", hash.Sum(nil)) // make a string of it

	prefix := strings.ToUpper(hashStr[0:5])   // prefix is what we're going to send
	remainder := strings.ToUpper(hashStr[5:]) // this is needed for matching later
	fullURL := urlStem + prefix               // the URL for the API query

	resp, err := http.Get(fullURL) // make a GET request to the API
	if err != nil {
		fmt.Println("Error contacting API.")
		os.Exit(2)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Bad response from server: %s", resp.Status)
		os.Exit(3)
	}

	matched := false
	scanner := bufio.NewScanner(resp.Body) // to read line-by-line
	for scanner.Scan() {                   // iterate over lines in response
		line := strings.TrimSpace(scanner.Text()) // get next line
		items := strings.Split(line, ":")         // separate hash and count
		if items[0] == remainder {                // compare returned hash with ours
			matched = true
			fmt.Printf("pwned! %s found in database - %s times\n", pw, items[1])
		}
	}
	if !matched {
		fmt.Printf("No match for %s\n", pw)
	}
}

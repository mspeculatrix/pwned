# pwned

Queries the haveibeenpwned.com API for breached passwords.

Usage: pwned <password_to_check>

Doesn't work with passphrases containing spaces.

The entered password is hashed with SHA1. The first five chars of the hash are sent to the API which returns the hashes of all the passwords where the first five chars of their hashes match the characters we sent, along with a count of how many occurrences of that password are in the database - ie, how many times this password has appeared in breached databases. The returned hashes are missing the first five chars because we already know those.

This is a secure exchange because the password itself never leaves your machine and you don't even send the full hash (so no cracking with rainbow tables). Even if there's only one match, the server doesn't know what password you're checking - only that there's a match of the first five chars of the hash.

The program then cycles through the returned hashes looking for a match with the hash created by our password.

Offered up under GPL 3.0 but absolutely not guaranteed fit for use.
This is code created by hobbyist coder, so use at your own risk.

Blog: https://mansfield-devine.com/speculatrix/

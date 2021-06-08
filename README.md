Cows and bulls implementation.

New rules:
A 4 digit sequence of numbers is generated. This will be called the "secret".
The player must try to guess this sequence exactly within the fewest guesses possible. 
A digit in the guess can match a digit in the secret in two ways: 1) exactly, and 2) nearly.
- An exact match means that a digit in the guess is in the same place as that of the secret.
e.g. Secret = 1423, Guess = 1099, Exact matches = 1 (the first digit of the guess matches the first digit of the secret)
- A near match means that a digit in the guess is present within the secret, but only if that digit has not already been near matched or isn't an exact match

TODO:
- Establish rules for checking guesses (don't count nears if an exact has already been matched for that number)
- Secret generator (uses seed, or default)
- Input parser (takes in user input)
- io for console
- Response creator based on guess (determine scope of response package)
- Secret creator (including secret obj)
- Guess handler
- Consider menu implementation for looping and options

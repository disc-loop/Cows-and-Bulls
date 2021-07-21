Cows and bulls implementation.

Rules:
A 4 digit sequence of numbers is generated. This will be called the "secret".
The player must try to guess this sequence exactly within the fewest guesses possible. 
A digit in the guess can match a digit in the secret in two ways: 1) exactly, and 2) nearly.
- An exact match means that a digit in the guess is in the same place as that of the secret.
e.g. Secret = 1423, Guess = 1099, Exact matches = 1 (the first digit of the guess matches the first digit of the secret)
- A near match means that a digit in the guess is present within the secret, but only if that digit has not already been near matched or isn't an exact match

TODO:
- Secret validation (correct length, valid characters)
- Consider menu implementation for looping (quit on key) and options (configure game, e.g. secret seed)
- Add config for secret length (configurable across the program)
- Secret generator (uses seed, or default)
- io for console(?)
- Dockerise app (lets us use ECS and expand the app)
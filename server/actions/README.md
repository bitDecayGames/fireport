# Actions and Cards

Actions and cards are very similar (in fact, a Card implements the Action interface), but they do have slightly different purposes:

**Action: the smallest unit of modification to apply to a state**

**Card: a named collection of actions**

The idea of a Card is to represent the actual card object that a player has access to in a game.  They have a card type that defines what type of card it is.  Cards are essentially static and for the most part, if two cards have the same card type, then they will have the exact same list of actions.

An Action is what we use to modify the state between turns. Cards come in as input from the player which get reduced into concrete Actions. An Actions purpose is to take a state, make a deep copy of that state, modify that copied version of the state, and then return that copied version. 
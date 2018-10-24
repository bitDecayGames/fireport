# Rules

## Game Rules

The idea of a game rule, in this game, is to validate the next purposed state. Sometimes a rule will need to check the newly purposed state vs the original state, and sometimes it will just need to check the newly purposed state. In either case, the rule is looking for a specific field or value that is out of place or in an invalid state. The goal of a rule should be to define as small of a scope as possible when checking for validity. Rules that try and check too much will run the risk of developing bugs. Another very important concept for rules is that they NEVER modify the states that they are checking. The rules should treat the states as readonly. As far as I know, there is not a way to enforce this law through the language, so we must enforce this law ourselves. **Never modify state inside of a Rule!**

## Input Rules

The idea of an input rule in this game is to validate the incoming messages from the players to make sure the player cannot play the BlueEyesWhiteDragon card on every turn.  It should also handle things like the max/min number of cards you are allowed to play each turn.  Just like the game rules, input rules never modify any state.  They just return an error, or nil.
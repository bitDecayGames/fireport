# Conditions

A condition is something that triggers when actions happen that put the game state into either a specific desired state (like when a player wins the game) or an undesired state (like when two players try to move to the same space).  It then adds or removes actions from the list of actions to apply to the game state.

## Example: Two players collide while moving

When there are two players:

```
++++
+AB+
++++
```

Player A plays a `Turn Right` and player B plays a `Turn Left` card.  So their desired states are now (because a turn card is a move forward, turn 90 degrees, then move forward again):

```
+ba+
++++
++++
```


But what their actual states will be after the turn is over is:

```
+AB+
++++
++++
```
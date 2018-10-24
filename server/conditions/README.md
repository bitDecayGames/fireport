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

## Example: Card order
For players A and B, they each choose cards.  A chooses cards `1,2,3,4,5` and B chooses cards `1,2,3`.

Those cards would be sorted like: `A1,B1,A2,B2,A3,B3,A4,A5`

And they would be grouped like:
```
A1,B1
A2,B2
A3,B3
A4
A5
``` 

Now, each group will be grouped again based on card type.  Like cards will now be played in parallel.

## Example: Shields

When a player plays the shield card, it will have an action associated with it like Player.Shield = true.  But that shield needs to turn off at the end of this turn.  So, we could have a condition that checks if any players have a shield at the end of their turn and turn it off.  We could also use the condition as a way to handle when someone is damaged with their shield on.  The condition would check for damage done to the shielded person and then reduce the damage by 1 and set the Player.Shield = false.

## Example: Two out of Three players collide while moving

When there are three players:

```
+C++
++++
+AB+
```

Player A (facing up) plays a `Turn Right` and player B (facing up) plays a `Turn Left` and player C (facing down) plays a `Move Forward` card.  So their desired state is:

```
+ +  ++
+(ac)b+
+ +  ++
```


But what their actual states will be after the turn is over is:

```
+C++
+B++
+A++
```

## Example: One Collision Causes Another

When there are three players:

```
C+++
++++
+AB+
```

Player A (facing up) plays a `Turn Left` and player B (facing up) plays a `Turn Left` and player C (facing down) plays a `Turn Left` card.  So their desired state is:

```
+ +  ++
a(bc)++
+ +  ++
```

But what their actual states will be after the turn is over is:

```
++++
CAB+
++++
```

## Example: Should Turning Tanks Collide?

When there are up to four players:

```
++C++
+D+B+
++A++
```

Player A (facing up) plays a `Turn Right` and player B (facing left) plays a `Turn Right` and player C (facing down) plays a `Turn Right` and player D (facing right) plays a `Turn Right`.  So their desired state is:

```
++b++
+c+a+
++d++
```

But will their state be:

```
++C++
+D+B+
++A++
```

Or


```
++B++
+C+A+
++D++
```
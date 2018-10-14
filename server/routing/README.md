## Lobby
- Create a new lobby
    - Returns:
        - lobby ID
- Join a lobby
    - Args:
        - lobby ID
        - Player name
    - Returns:
        - Route to websocket for the lobby
- Start game from lobby
    - Args:
        - lobby ID
    - Returns:
        - Game ID
- Show/find lobby (with simple name filter)
    - Args:
        - Search string
    - Returns:
        - list of lobby

## Game
- Submit turn
- Request timeout (each player can pause once per game for 30 seconds)


## WebSockets (Strict server -> client messaging)
- Lobby Mode
    - Message Types
        - Player Joined/Left the Lobby (LobbyState)
        - Game has started (GameStart)
- Game Mode
    - Message Types
        - Turn timeout (if player doesn't submit a turn) (TurnTimeout)
        - PlayerDisconnect
        - PlayerReconnect
        - TurnResult
        - Game timeout (if nobody is playing any longer) (GameTimeout)
- Post Game
    - Message Types
        - GameResultSummary

## Data objects
- LobbyState
    - Current players

- GameStart
    - Game ID

- TurnTimeout
    - <no data>

- GameTimeout
    - <no data>

- TurnResult
    - Previous state
    - current state
    - actions (things that were applied to previous to reach current)
    - current allowed input choices (i.e. player's cards)

- GameState
    - Board state
    
- GameResultSummary
    - Winner
    - Stats summary
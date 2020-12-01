## SBG Roulette API

**Install it**

`go get github.com/andypagdin/sbg-roulette`

**Build it**

`cd sbg-roulette`

`go build`

**Run it (:8080)**

`./sbg-roulette`

**Test it**

`go test ./handler -v`

## Overview
* Action takes place on a table.
* Players must join a table to place bets.
* Once a table spin occurs all bets on that table must be settled against the outcome before new bets can be placed.
* Bet types supported
  * **Inside** : Straight Up
  * **Outside** : Red/Black, Odd/Even, High/Low, Columns
* Separation of game logic throughout the endpoints has been considered to try and allow the API to be suitable for a variety of game formats depending on the product implementation.
* Endpoints are versioned to enable the addition of new features/endpoints without breaking existing implementations.

## Endpoints
**/v1/tables**
* `GET` : Get all tables
  * Returns array of table objects
* `POST` : Create a new table
  * Returns created table object

```
{
  "id": "a79c7873-8bcc", string
  "players": [],         array
  "bets": [],            array
  "openForBets" true,    bool
}
```

**/v1/players**
* `GET` : Get all players
  * Returns array of player objects
* `POST` : Create a new player
  * Payload
      ```
      {
        "name": "player name" string
      }
      ```
  * Errors
    * `400 Invalid request payload`
  * Returns created player object

```
{
  "id": "6778fdb3-4a50", string
  "name: "foo",          string
  "balance": 100,        float64
}
```

**/v1/tables/:table-id/players/:player-id**
* `POST` : Add a player to a table
  * Errors
    * `400 Table not found`
    * `400 Player not found`
    * `400 Player is already at this table`
  * Returns table object with appended player

```
{
  "id": "a79c7873-8bcc",
  "players": [
    {
      "id": "6778fdb3-4a50",
      "name": "John Smith",
      "balance": 100
    }
  ],
  "bets": [],
  "openForBets": true
}
```

**/v1/tables/:table-id/bet/:player-id**

Supported bet type and value fields
```
Type : "straight", Value : "0-36"
Type : "colour",   Value : "red" or "black"
Type : "oddEven",  Value : "odd" or "even"
Type : "highLow",  Value : "high" or "low"
Type : "column",   Value : "1st12" or "2nd12" or "3rd12"
```
* `POST` : Place a bet on a table
  * Payload
    ```
    {
      "type": "colour", string
      "value: "red",    string
      "amount": 10,     float64
    }
    ```
  * Errors
    * `400 Invalid request payload`
    * `400 Bets are closed wait for next round`
    * `400 Player must be added to the table before placing a bet`
    * `400 Invalid table ID`
    * `400 Invalid player ID`
    * `400 Bet amount exceeds player balance`
  * Returns bet object

```
{
  "playerId": "6778fdb3-4a50", string
  "type": "colour",            string
  "value": "red",              string
  "amount": 10,                float64
}
```

**/v1/tables/:table-id/bet/settle/:outcome**
* `POST` : Settle all table bets against a given outcome, sets tables `OpenForBets` to true
  * Errors
    * `400 Invalid table ID`
    * `400 Invalid outcome parameter`

**/v1/tables/:table-id/spin**
* `GET` : "Spin" the roulette wheel to generate a table outcome, sets tables `OpenForBets` to false
  * Errors
    * `400 Invalid table ID`
    * `400 Settle outstanding bets before spinning`
  * Returns outcome value

```
{
  10 int
}
```

## Future improvements
* Add table bet limits
* Add 'GetBoard' endpoint to enforce bet types and values

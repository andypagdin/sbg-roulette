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
* `POST` : Create a new table

**/v1/players**
* `GET` : Get all players
* `POST` : Create a new player
  * Payload
    * `{"Name": "PlayerName"}`
  * Errors
    * `400 Invalid request payload`

**/v1/tables/:table-id/players/:player-id**
* `POST` : Add a player to a table
  * Errors
    * `400 Table not found`
    * `400 Player not found`
    * `400 Player is already at this table`

**/v1/tables/:table-id/bet/:player-id**
* `POST` : Place a bet on a table
  * Errors
    * `400 Invalid request payload`
    * `400 Bets are closed wait for next round`
    * `400 Player must be added to the table before placing a bet`
    * `400 Invalid table ID`
    * `400 Invalid player ID`

**/v1/tables/:table-id/bet/settle/:outcome**
* `POST` : Settle all table bets against a given outcome
  * Errors
    * `400 Invalid table ID`
    * `400 Invalid outcome parameter`

**/v1/tables/:table-id/spin**
* `GET` : "Spin" the roulette wheel to generate table outcome
  * Errors
    * `400 Invalid table ID`
    * `400 Settle outstanding bets before spinning`

## Future improvments
* Add table bet limits
* Add 'GetBoard' endpoint to enforce bet types and values
* Do not allow players to exceed balance
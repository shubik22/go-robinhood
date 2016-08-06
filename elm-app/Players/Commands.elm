module Players.Commands exposing (..)

import Http
import Task
import Json.Decode exposing (Decoder, (:=), list, string, float, int, object5)
import Players.Models exposing (Player)
import Positions.Models exposing (Position)
import Players.Messages exposing (..)

fetchAll : Cmd Msg
fetchAll =
  Http.get leaderboardDecoder fetchAllUrl
  |> Task.perform FetchAllFail FetchAllDone

fetchAllUrl : String
fetchAllUrl =
  "/leaderboard"

leaderboardDecoder : Decoder (List Player)
leaderboardDecoder =
  ("entries" := list playerDecoder )

playerDecoder : Decoder Player
playerDecoder =
  object5 Player
    ("username" := string)
    ("cash_balance" := float)
    ("position_balance" := float)
    ("total_balance" := float)
    ("positions" := list positionDecoder)

positionDecoder : Decoder Position
positionDecoder =
  object5 Position
    ("purchase_time" := string)
    ("quantity" := int)
    ("symbol" := string)
    ("average_buy_price" := float)
    ("last_trade_price" := float)

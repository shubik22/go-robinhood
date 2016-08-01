module Players.Commands exposing (..)

import Http
import Json.Decode as Decode exposing ((:=))
import Task
import Players.Models exposing (Player)
import Players.Messages exposing (..)

fetchAll : Cmd Msg
fetchAll =
  Http.get collectionDecoder fetchAllUrl
  |> Task.perform FetchAllFail FetchAllDone

fetchAllUrl : String
fetchAllUrl =
  "/leaderboard"

collectionDecoder : Decode.Decoder (List Player)
collectionDecoder =
  Decode.list memberDecoder

memberDecoder : Decode.Decoder Player
memberDecoder =
  Decode.object4 Player
    ("username" := Decode.string)
    ("cash_balance" := Decode.float)
    ("position_balance" := Decode.float)
    ("total_balance" := Decode.float)

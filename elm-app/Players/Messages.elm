module Players.Messages exposing (..)

import Http
import Players.Models exposing (Player, SortedColumn, SortedDirection)

type Msg
  = FetchAllDone (List Player)
  | FetchAllFail Http.Error
  | Sort SortedColumn SortedDirection

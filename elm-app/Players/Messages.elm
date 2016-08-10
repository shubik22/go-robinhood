module Players.Messages exposing (..)

import Http
import Players.Models exposing (Player, SortedColumn, SortedDirection)
import Positions.Models exposing (Position)

type Msg
  = FetchAllDone (List Player)
  | FetchAllFail Http.Error
  | Sort SortedColumn SortedDirection
  | ShowPositions String (List Position) 
  | HidePositions

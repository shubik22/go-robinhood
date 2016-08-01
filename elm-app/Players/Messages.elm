module Players.Messages exposing (..)

import Http
import Players.Models exposing (Player)

type Msg
  = FetchAllDone (List Player)
  | FetchAllFail Http.Error

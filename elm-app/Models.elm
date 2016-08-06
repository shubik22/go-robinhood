module Models exposing (..)

import Players.Models exposing (Player, SortedColumn(..), SortedDirection(..))

type alias Model =
  { players : List Player
  , sortedColumn: SortedColumn
  , sortedDirection : SortedDirection
  }

initialModel : Model
initialModel =
  { players = []
  , sortedColumn = TotalBalance
  , sortedDirection = Descending
  }


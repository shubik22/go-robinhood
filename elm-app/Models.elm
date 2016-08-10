module Models exposing (..)

import Modal.Models exposing (Modal, initialModal)
import Players.Models exposing (Player, SortedColumn(..), SortedDirection(..))

type alias Model =
  { players : List Player
  , sortedColumn: SortedColumn
  , sortedDirection : SortedDirection
  , modal : Modal
  }

initialModel : Model
initialModel =
  { players = []
  , sortedColumn = TotalBalance
  , sortedDirection = Descending
  , modal = initialModal
  }


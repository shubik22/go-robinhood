module Players.Models exposing (..)

import Positions.Models exposing (..)

type SortedColumn
  = CashBalance
  | PositionBalance
  | TotalBalance

type SortedDirection
  = Ascending
  | Descending

type alias Player =
  { name : String
  , cashBalance: Float
  , positionBalance: Float
  , totalBalance: Float
  , positions: List Position
  }

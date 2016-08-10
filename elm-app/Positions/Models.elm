module Positions.Models exposing (..)

type alias Position =
  { purchaseTime : String
  , quantity: Int
  , symbol: String
  , averageBuyPrice: Float
  , lastTradePrice: Float
  }

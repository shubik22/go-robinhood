module Positions.Models exposing (..)

type alias Position =
  { purchase_time : String
  , quantity: Int
  , symbol: String
  , average_buy_price: Float
  , last_trade_price: Float
  }

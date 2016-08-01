module Players.List exposing (..)

import Html exposing (..)
import String exposing (padRight, contains, length)
import List exposing (reverse, sortBy, indexedMap)
import Players.Messages exposing (..)
import Players.Models exposing (Player)

view : List Player -> Html Msg
view players =
  div []
    [ list players
    ]

list : List Player -> Html Msg
list players =
  table []
  [ thead []
    [ tr []
      [ th [] [ text "Rank" ]
      , th [] [ text "Name" ]
      , th [] [ text "Cash Balance" ]
      , th [] [ text "Position Balance" ]
      , th [] [ text "Total Balance" ]
      ]
    ]
  , tbody [] (indexedMap playerRow (sortPlayers players))
  ]

playerRow : Int -> Player -> Html Msg
playerRow rank player =
  tr []
  [ td [] [ text (toString (rank + 1)) ]
  , td [] [ text player.name ]
  , td [] [ text (formatMoney player.cashBalance) ]
  , td [] [ text (formatMoney player.positionBalance) ]
  , td []
  [ strong [] [ text (formatMoney player.totalBalance) ] ]
  ]

sortPlayers : List Player -> List Player
sortPlayers players =
  players
    |> sortBy .totalBalance
    |> reverse

formatMoney : Float -> String
formatMoney amount =
  "$" ++ (toString amount)

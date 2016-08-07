module Players.List exposing (..)

import Html exposing (..)
import Html.Events exposing (onClick)
import Html.Attributes exposing (classList)
import List exposing (reverse, sortBy, indexedMap)
import Numeral exposing (format)
import Players.Messages exposing (Msg(..))
import Players.Models exposing (Player, SortedColumn(..), SortedDirection(..))

view : List Player -> SortedColumn -> SortedDirection -> Html Msg
view players sortedColumn sortedDirection =
  table []
  [ tableHead sortedColumn sortedDirection
  , tableBody players sortedColumn sortedDirection
  ]

tableHead : SortedColumn -> SortedDirection -> Html Msg
tableHead sortedColumn sortedDirection =
  thead []
    [ tr []
      [ th [] [ text "Rank" ]
      , th [] [ text "Name" ]
      , th [ classList (headerClasses sortedDirection (sortedColumn == CashBalance)), onClick (Sort CashBalance (swapDirection sortedDirection)) ] [ text "Cash Balance" ]
      , th [ classList (headerClasses sortedDirection (sortedColumn == PositionBalance)), onClick (Sort PositionBalance (swapDirection sortedDirection)) ] [ text "Position Balance" ]
      , th [ classList (headerClasses sortedDirection (sortedColumn == TotalBalance)), onClick (Sort TotalBalance (swapDirection sortedDirection)) ] [ text "Total Balance" ]
      ]
    ]

headerClasses : SortedDirection -> Bool -> List (String, Bool)
headerClasses sortedDirection show =
  let
    ascendingStatus = (sortedDirection == Ascending) && show
    descendingStatus = (sortedDirection == Descending) && show
  in
    [ ("sortable-column-header", True)
    , ("ascending", ascendingStatus)
    , ("descending", descendingStatus)
    ]

tableBody : List Player -> SortedColumn -> SortedDirection -> Html Msg
tableBody players sortedColumn sortedDirection =
  tbody [] (indexedMap playerRow (sortPlayers players sortedColumn sortedDirection))

playerRow : Int -> Player -> Html Msg
playerRow rank player =
  tr []
  [ td [] [ text (toString (rank + 1)) ]
  , td [] [ text player.name ]
  , td [] [ text (format currencyFormat player.cashBalance) ]
  , td [] [ text (format currencyFormat player.positionBalance) ]
  , td [] [ text (format currencyFormat player.totalBalance) ]
  ]

sortPlayers : List Player -> SortedColumn -> SortedDirection -> List Player
sortPlayers players sortedColumn sortedDirection =
  case sortedDirection of
    Ascending ->
      players
        |> sortingMap sortedColumn
    Descending ->
      players
        |> sortingMap sortedColumn
        |> reverse

sortingMap : SortedColumn -> List Player -> List Player
sortingMap sortedColumn players =
  case sortedColumn of
    CashBalance ->
     sortBy .cashBalance players
    PositionBalance ->
     sortBy .positionBalance players
    TotalBalance ->
     sortBy .totalBalance players

swapDirection : SortedDirection -> SortedDirection
swapDirection sortedDirection =
  case sortedDirection of
    Ascending ->
      Descending
    Descending ->
      Ascending

currencyFormat : String
currencyFormat =
  "$0,0.00"

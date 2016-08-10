module Modal.View exposing (..)

import Html exposing (..)
import Html.Events exposing (onClick)
import Html.Attributes exposing (class)
import List exposing (map)
import Numeral exposing (format)
import Modal.Models exposing (Modal)
import Positions.Models exposing (Position)
import Players.Messages exposing (Msg(..))

view : Modal -> Html Msg
view modal =
  if modal.show then
    div [ class "modal-bg" ]
    [ div [ class "modal" ]
      [ div [ class "modal-header" ]
        [ span [] [ text modal.title ]
        , span [ onClick HidePositions, class "modal-close" ] []
        ]
      , div [ class "modal-body" ] [ positionsTable modal.positions ]
      ] 
    ]
  else
    text ""

positionsTable : List Position -> Html Msg
positionsTable positions =
  table []
  [ tableHead
  , tableBody positions
  ]

tableHead : Html Msg
tableHead =
  thead []
    [ tr []
      [ th [] [ text "Symbol" ]
      , th [] [ text "Quantity" ]
      , th [] [ text "Buy Price" ]
      , th [] [ text "Market Price" ]
      ]
    ]

tableBody : List Position -> Html Msg
tableBody positions =
  tbody [] (map positionRow positions)

positionRow : Position -> Html Msg
positionRow position =
  tr []
  [ td [] [ text position.symbol ]
  , td [] [ text (toString position.quantity) ]
  , td [] [ text (format currencyFormat position.averageBuyPrice) ]
  , td [] [ text (format currencyFormat position.lastTradePrice) ]
  ]
  
currencyFormat : String
currencyFormat =
  "$0,0.00"

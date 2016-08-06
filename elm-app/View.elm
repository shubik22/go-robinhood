module View exposing (..)

import Html exposing (Html, div, text)
import Html.App
import Messages exposing (Msg(..))
import Models exposing (Model)
import Players.List

view : Model -> Html Msg
view model =
  div []
    [ leaderboard model
    ]

leaderboard : Model -> Html Msg
leaderboard model =
  Html.App.map PlayersMsg (Players.List.view model.players model.sortedColumn model.sortedDirection)

module View exposing (..)

import Html exposing (Html, div, text, img, header)
import Html.Attributes exposing (class, classList, src)
import Html.App
import Messages exposing (Msg(..))
import Models exposing (Model)
import Players.List
import Modal.View

view : Model -> Html Msg
view model =
  div []
    [ appHeader 
    , modal model 
    , container model
    ]

appHeader : Html Msg
appHeader =
  header [ class "header" ]
  [ img [ src "assets/logo.png", class "logo" ] []
  ]

container : Model -> Html Msg
container model =
  div [ class "container" ]
  [ div [ class "row" ]
    [ div [ classList gridClasses ]
      [ leaderboard model
      ]
    ]
  ]

gridClasses : List (String, Bool)
gridClasses =
  [ ("column", True)
  , ("column-70", True)
  , ("column=offset-15", True)
  ]

modal : Model -> Html Msg
modal model =
  Html.App.map PlayersMsg (Modal.View.view model.modal)

leaderboard : Model -> Html Msg
leaderboard model =
  Html.App.map PlayersMsg (Players.List.view model.players model.sortedColumn model.sortedDirection)

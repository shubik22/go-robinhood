module Main exposing (..)

import Html.App
import Messages exposing (Msg(..))
import Models exposing (Model, initialModel)
import Subscriptions exposing (subscriptions)
import Players.Commands exposing (fetchAll)
import View exposing (view)
import Update exposing (update)

init : ( Model, Cmd Msg )
init =
  ( initialModel, Cmd.map PlayersMsg fetchAll )

-- MAIN

main =
  Html.App.program
  { init = init
  , view = view
  , update = update
  , subscriptions = subscriptions
  }

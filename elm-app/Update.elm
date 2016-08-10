module Update exposing (..)

import Messages exposing (Msg(..))
import Models exposing (Model)
import Players.Update
import Debug exposing (log)

update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
  case msg of
    PlayersMsg subMsg ->
      let
        ( updatedPlayers, updatedSortedColumn, updatedSortedDirection, updatedModal, cmd ) =
          Players.Update.update subMsg model.players model.sortedColumn model.sortedDirection model.modal
      in
        ( { model | players = updatedPlayers, sortedColumn = updatedSortedColumn, sortedDirection = updatedSortedDirection, modal = updatedModal }, Cmd.map PlayersMsg cmd )

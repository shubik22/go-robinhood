module Update exposing (..)

import Messages exposing (Msg(..))
import Models exposing (Model)
import Players.Update

update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
  case msg of
    PlayersMsg subMsg ->
      let
        ( updatedPlayers, updatedSortedColumn, updatedSortedDirection, cmd ) =
          Players.Update.update subMsg model.players model.sortedColumn model.sortedDirection
      in
        ( { model | players = updatedPlayers, sortedColumn = updatedSortedColumn, sortedDirection = updatedSortedDirection }, Cmd.map PlayersMsg cmd )

module Players.Update exposing (..)

import Players.Messages exposing (Msg(..))
import Players.Models exposing (Player, SortedColumn, SortedDirection)
import Modal.Models exposing (Modal)

update : Msg -> List Player -> SortedColumn -> SortedDirection -> Modal -> ( List Player, SortedColumn, SortedDirection, Modal, Cmd Msg )
update message players sortedColumn sortedDirection modal =
  case message of
    FetchAllDone newPlayers ->
      ( newPlayers, sortedColumn, sortedDirection, modal, Cmd.none )

    FetchAllFail error ->
      ( players, sortedColumn, sortedDirection, modal, Cmd.none )

    Sort newSortedColumn newSortedDirection ->
      ( players, newSortedColumn, newSortedDirection, modal, Cmd.none )

    ShowPositions title positions ->
      ( players, sortedColumn, sortedDirection, Modal True title positions, Cmd.none )

    HidePositions ->
      ( players, sortedColumn, sortedDirection, Modal False "" [], Cmd.none )

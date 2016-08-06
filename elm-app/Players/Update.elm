module Players.Update exposing (..)

import Players.Messages exposing (Msg(..))
import Players.Models exposing (Player, SortedColumn, SortedDirection)

update : Msg -> List Player -> SortedColumn -> SortedDirection -> ( List Player, SortedColumn, SortedDirection, Cmd Msg )
update message players sortedColumn sortedDirection =
  case message of
    FetchAllDone newPlayers ->
      ( newPlayers, sortedColumn, sortedDirection, Cmd.none )

    FetchAllFail error ->
      ( players, sortedColumn, sortedDirection, Cmd.none )

    Sort newSortedColumn newSortedDirection ->
      ( players, newSortedColumn, newSortedDirection, Cmd.none )

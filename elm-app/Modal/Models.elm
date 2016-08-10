module Modal.Models exposing (..)

import Positions.Models exposing (Position)

type alias Modal =
  { show : Bool
  , title: String
  , positions: List Position
  }

initialModal : Modal
initialModal =
  { show = False
  , title = ""
  , positions = []
  }

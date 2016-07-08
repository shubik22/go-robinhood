'use strict'

import React from 'react';
import { List } from 'immutable';

const Leaderboard = React.createClass({
  propTypes: {
    users: React.PropTypes.instanceOf(List).isRequired
  },

  render: () => {
    return (
      <div>Loaded from React</div>
    );
  }
});

export default Leaderboard;


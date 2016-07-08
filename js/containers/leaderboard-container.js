'use strict'

import { connect } from 'react-redux';
import Leaderboard from '../components/leaderboard';

const mapStateToProps = (store) => {
  return {
    users: store.userState
  };
}

const LeaderboardContainer = connect(mapStateToProps)(Leaderboard);

export default LeaderboardContainer;


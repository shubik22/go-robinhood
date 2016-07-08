'use strict'

import { createStore, combineReducers } from 'redux';
import { List, Map } from 'immutable';
import actionConstants from './actions/action-constants';

const userReducer = (state = List(), action) => {
  switch (action.type) {
    case actionConstants.loadUsers:
      return List(action.users)
    default:
      return state
  }
}

const appReducer = (state = Map(), action) => {
  return state;
}

const reducers = combineReducers({
  userState: userReducer,
  appState: appReducer
});

const store = createStore(reducers);

export default store;


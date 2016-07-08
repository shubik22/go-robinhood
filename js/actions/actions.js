'use strict'

import actionConstants from './action-constants';

export function loadUsers(users) => {
  return {
    type: actionConstants.loadUsers,
    users
  }
}


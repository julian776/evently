import { createReducer, on } from '@ngrx/store';
import { login, logout } from './user.actions';
import { User } from 'src/app/user/models/user';

export interface UserState {
  user: User
  isLoggedIn: boolean;
}

export const initialState: UserState = {
  user: {
    email: '',
    name: '',
    purpouseOfUse: ''
  },
  isLoggedIn: false,
};

export const userReducer = createReducer(
  initialState,
  on(login, (state, { user }) => ({...state, user, isLoggedIn: true})),
  on(logout, (state) => ({...state, user: {email: '' ,name: '', purpouseOfUse: ''}, isLoggedIn: false})),
);

import { createAction, props } from '@ngrx/store';
import { User } from 'src/app/user/models/user';

export const login = createAction('[User] login', props<{user: User}>())
export const logout = createAction('[User] logout', props<{user: User}>())

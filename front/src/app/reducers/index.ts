import { isDevMode } from '@angular/core';
import {
  ActionReducerMap,
  MetaReducer
} from '@ngrx/store';
import { eventsReducer, EventsState } from './events/events.reducer';
import { userReducer, UserState } from './user/user..reducer';

export interface State {
  events: EventsState
  user: UserState
}

export const reducers: ActionReducerMap<State> = {
  events: eventsReducer,
  user: userReducer,
};

export const metaReducers: MetaReducer<State>[] = isDevMode() ? [] : [];

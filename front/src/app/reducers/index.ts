import { isDevMode } from '@angular/core';
import {
  ActionReducer,
  ActionReducerMap,
  createFeatureSelector,
  createSelector,
  MetaReducer
} from '@ngrx/store';
import { eventsReducer, EventsState } from './events/events.reducer';

export interface State {
  events: EventsState
}

export const reducers: ActionReducerMap<State> = {
  events: eventsReducer
};

export const metaReducers: MetaReducer<State>[] = isDevMode() ? [] : [];

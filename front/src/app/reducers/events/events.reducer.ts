import { createReducer, on } from '@ngrx/store';
import { addEvent, delEvent, loadEvents, updateEvent } from './events.actions';
import { Event } from 'src/app/events/models/event';

export interface EventsState {
  [id: string]: Event;
}

export const initialState: EventsState = {};

export const eventsReducer = createReducer(
  initialState,
  on(loadEvents, (state, { events }) => {
    const newState: EventsState = {};
    events.forEach((event) => (newState[event.id] = event));
    return newState;
  }),
  on(addEvent, (state, { event }) => {
    const newState: EventsState = {
      ...state
    };
    newState[event.id] = event;
    return newState;
  }),
  on(delEvent, (state, { id }) => {
    const newState: EventsState = {};
    Object.values(state).forEach((event) => {
      if (event.id != id) {
        newState[event.id] = event;
      }
    });
    return newState;
  }),
  on(updateEvent, (state, { event }) => {
    const newState = state;
    newState[event.id] = event;
    return newState;
  })
);

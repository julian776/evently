import { createSelector } from "@ngrx/store";
import { State } from "../../index";
import { EventsState } from "../events.reducer";

export const selectEvents = (state: State) => state.events;

export const selectEventsList = createSelector(
  selectEvents,
  (state: EventsState) => Object.values(state)
);

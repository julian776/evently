import { createAction, props } from '@ngrx/store';
import { Event } from 'src/app/events/models/event';

export const addEvent = createAction('[Events] add', props<{event: Event}>())
export const delEvent = createAction('[Events] delete', props<{id: string}>())
export const updateEvent = createAction('[Events] update', props<{event: Event}>())
export const loadEvents = createAction('[Events] load', props<{events: Event[]}>())

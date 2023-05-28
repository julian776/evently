import { Event } from "../../models/event";

export const eventMapper = (event: Event): Event => {
  const eventFormatted: Event = {
    ...event
  }
  if (event.description.length >= 170) {
    eventFormatted.description = event.description.split(' ', 26).join(' ') + '...'
  }
  if (event.title.length >= 70) {
    eventFormatted.title = event.title.split(' ', 7).join(' ') + '...'
  }
  return eventFormatted
}

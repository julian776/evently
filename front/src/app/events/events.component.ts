import { Component } from '@angular/core';
import { Event } from './models/event';
import { eventMapper } from './mappers/event.mapper';

@Component({
  selector: 'events',
  templateUrl: './events.component.html',
  styleUrls: ['./events.component.scss', '../app.component.scss']
})
export class EventsComponent {
  public events: Array<Event>

  constructor() {
    const event: Event = {
      id: '1',
      title: 'Title Lorem ipsum dolor sit amet, consectetur adipiscing elit.',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum felis nec dolor vulputate, sed consectetur elit ultrices. Vestibulum sit amet interdum mauris. Vestibulum porttitor tortor vehicula nulla luctus ultricies. Cras in diam ut nulla efficitur placerat. Mauris nulla mi, lacinia ac est at, pellentesque fringilla velit. Mauris eu elit auctor, consequat nunc vel, mattis lacus. Vivamus et suscipit tortor, eu egestas libero. Donec a mi lobortis, blandit ex non, rhoncus metus. Suspendisse finibus, nunc eu dictum tristique, diam velit finibus ligula, non pulvinar nisl metus ac orci. Etiam sagittis tortor non velit luctus, eget bibendum ante interdum. Maecenas urna arcu, efficitur vitae erat non, dictum vehicula sapien. Duis dictum erat eu ipsum gravida, id dignissim diam venenatis. Donec nisl quam, efficitur ac iaculis non, bibendum quis arcu.',
      location: 'loc',
      organizerName: 'Julian',
      organizerEmail: 'julian',
      startTime: Date.now().toString(),
      endTime: Date.now().toString()
    }
    const event2: Event = {
      id: '2',
      title: 'Title Lorem ipsum dolor sit amet, consectetur adipiscing elit.',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum felis nec dolor vulputate, sed consectetur elit ultrices. Vestibulum sit amet interdum mauris. Vestibulum porttitor tortor vehicula nulla luctus ultricies. Cras in diam ut nulla efficitur placerat. Mauris nulla mi, lacinia ac est at, pellentesque fringilla velit. Mauris eu elit auctor, consequat nunc vel, mattis lacus. Vivamus et suscipit tortor, eu egestas libero. Donec a mi lobortis, blandit ex non, rhoncus metus. Suspendisse finibus, nunc eu dictum tristique, diam velit finibus ligula, non pulvinar nisl metus ac orci. Etiam sagittis tortor non velit luctus, eget bibendum ante interdum. Maecenas urna arcu, efficitur vitae erat non, dictum vehicula sapien. Duis dictum erat eu ipsum gravida, id dignissim diam venenatis. Donec nisl quam, efficitur ac iaculis non, bibendum quis arcu.',
      location: 'loc',
      organizerName: 'Julian',
      organizerEmail: 'julian',
      startTime: Date.now().toString(),
      endTime: Date.now().toString()
    }

    this.events = [event, event2].map(eventMapper)
  }

  ngOnInit() {

  }

  handleEdit(eventId: string) {
    console.log('Edit' + eventId);

  }

  handleDelete(eventId: string) {
    // TODO: Delete in back
    this.events = this.events.filter(event => event.id != eventId)
  }
}

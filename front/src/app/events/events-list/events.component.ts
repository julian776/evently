import { Component } from '@angular/core';
import { Event } from '../models/event';
import { eventMapper } from './mappers/event.mapper';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { Store } from '@ngrx/store';
import { State } from 'src/app/reducers';
import { loadEvents } from 'src/app/reducers/events/events.actions';
import { Observable, map } from 'rxjs';
import { selectEventsList } from 'src/app/reducers/events/selectors/event-list.selector';
import { environment } from '../../../environments/environment';

@Component({
  selector: 'events',
  templateUrl: './events.component.html',
  styleUrls: ['./events.component.scss', '../../app.component.scss'],
})
export class EventsComponent {
  public events$: Observable<Event[]> = this.store
    .select(selectEventsList)
    .pipe(map((events) => events.map(eventMapper)));

  constructor(
    private router: Router,
    private http: HttpClient,
    private store: Store<State>
  ) {}

  ngOnInit() {
    this.http
      .get<Event[]>(`${environment.apiUrl}/events`, { observe: 'body' })
      .subscribe((val: Event[]) => {
        this.store.dispatch(loadEvents({ events: val }));
      });
  }

  handleIndividualView(eventId: string) {
    this.router.navigate([`/event/${eventId}`]).catch(console.error);
    console.log('Edit' + eventId);
  }
}

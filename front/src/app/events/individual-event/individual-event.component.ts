import { Component } from '@angular/core';
import { ActivatedRoute, Route, Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { Event } from '../models/event';
import { Store } from '@ngrx/store';
import { State } from 'src/app/reducers';
import { delEvent } from 'src/app/reducers/events/events.actions';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-individual-event',
  templateUrl: './individual-event.component.html',
  styleUrls: ['./individual-event.component.scss', '../../app.component.scss'],
})
export class IndividualEventComponent {
  public event!: Event;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private http: HttpClient,
    private store: Store<State>,
    public snackBar: MatSnackBar
    ) {}

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');

    this.http
      .get<Event>(`http://0.0.0.0:8080/events/${id}`, { observe: 'body' })
      .subscribe((val: Event) => {
        this.event = val;
      });
  }

  handleEdit(eventId: string) {
    this.router.navigate([`/event/edit/${eventId}`]).catch(console.error);
  }

  handleDelete(eventId: string) {
    console.log(eventId);

    this.http
      .delete<Event[]>(`http://0.0.0.0:8080/events/${eventId}`)
      .subscribe(() => {
        this.store.dispatch(delEvent({ id: eventId }));
        this.snackBar.open('Event deleted', '', {
          duration: 2000,
        });
        this.router.navigate(['/events']).catch(console.error);
      });
  }
}

import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Event } from '../models/event';
import { Store } from '@ngrx/store';
import { State } from 'src/app/reducers';
import { delEvent } from 'src/app/reducers/events/events.actions';
import { MatSnackBar } from '@angular/material/snack-bar';
import { environment } from '../../../environments/environment';
import { catchError, throwError } from 'rxjs';

@Component({
  selector: 'app-individual-event',
  templateUrl: './individual-event.component.html',
  styleUrls: ['./individual-event.component.scss', '../../app.component.scss'],
})
export class IndividualEventComponent {
  public user$ = this.store.select((state) => state.user);
  public event!: Event;
  private id = this.route.snapshot.paramMap.get('id');
  panelOpenState = false;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private http: HttpClient,
    private store: Store<State>,
    public snackBar: MatSnackBar
  ) {}

  ngOnInit(): void {
    this.http
      .get<Event>(`${environment.apiUrl}/events/${this.id}`, {
        observe: 'body',
      })
      .subscribe((val: Event) => {
        this.event = this.formatEventDates(val);
      });
  }

  handleEdit(eventId: string) {
    this.router.navigate([`/event/edit/${eventId}`]).catch(console.error);
  }

  handleDelete(eventId: string) {
    console.log(eventId);

    this.http
      .delete<Event[]>(`${environment.apiUrl}/events/${eventId}`)
      .subscribe(() => {
        this.store.dispatch(delEvent({ id: eventId }));
        this.snackBar.open('Event deleted', '', {
          duration: 2000,
        });
        this.router.navigate(['/events']).catch(console.error);
      });
  }

  handleAddAttendee(eventId: string) {
    this.user$.subscribe((userState) => {
      if (!userState.isLoggedIn) {
        this.snackBar.open('Please login before try yo register', '', {
          duration: 2000,
        });
        return
      }
      this.http
          .put<string[]>(`${environment.apiUrl}/events/attendees`, {
            eventId: this.id,
            attendeeEmail: userState.user.email,
          })
          .pipe(catchError((error: HttpErrorResponse) => {
            if (error.status == 409) {
              this.snackBar.open('You already are registered', '', {
                duration: 2000,
              });
            }
            return throwError(() => new Error())
          }))
          .subscribe(() => {
            this.snackBar.open('You have registered succesfully', '', {
              duration: 2000,
            });
          });
    });
  }

  formatEventDates(event: Event) {
    const start = new Date(event.startDate)
    const end = new Date(event.endDate)
    return {
      ...event,
      startDate: `${start.getFullYear()}/${start.getMonth()}/${start.getDay()}`,
      endDate: `${end.getFullYear()}/${end.getMonth()}/${end.getDay()}`,
    }
  }
}

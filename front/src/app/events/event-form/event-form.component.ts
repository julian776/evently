import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import {
  FormBuilder,
  FormControl,
  FormGroup,
  Validators,
} from '@angular/forms';
import { Event } from '../models/event';
import { Store } from '@ngrx/store';
import { addEvent, updateEvent } from 'src/app/reducers/events/events.actions';
import { ActivatedRoute, Router } from '@angular/router';
import { UserState } from 'src/app/reducers/user/user..reducer';
import { User } from 'src/app/user/models/user';
import { State } from 'src/app/reducers';
import { environment } from '../../../environments/environment';


@Component({
  selector: 'app-event-form',
  templateUrl: './event-form.component.html',
  styleUrls: ['./event-form.component.scss'],
})
export class EventFormComponent {
  user!: User;
  minDate: Date;
  eventForm!: FormGroup;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private formBuilder: FormBuilder,
    private http: HttpClient,
    private store: Store<State>
  ) {
    const today = new Date()
    this.minDate = new Date(today.getFullYear(), today.getMonth())
  }

  ngOnInit() {
    this.store
      .select((state) => state.user)
      .subscribe((userState) => {
        if (!userState.isLoggedIn) {
          this.router.navigate(['login']).catch(console.error);
        }

        this.user = userState.user;
      });

    const id = this.route.snapshot.paramMap.get('id');
    if (id == '0') {
      this.initializeForm();
      return;
    }
    this.http
      .get<Event>(`${environment.apiUrl}/events/${id}`, { observe: 'body' })
      .subscribe((event: Event) => {
        this.initializeForm(event);
      });
  }

  /**
   * Initializes the `eventForm` FormGroup
   * with default values or with the values
   * of an existing `Event` object passed as
   * an argument.
   *
   * @method
   * @name initializeForm
   * @kind method
   * @memberof EventFormComponent
   * @param {Event} event?
   * @returns {void}
   */
  initializeForm(event?: Event): void {
    const title = event?.title ?? '';
    const description = event?.description ?? '';
    const location = event?.location ?? '';
    const cost = event?.cost ?? 0;
    const startDate = event?.startDate ?? '';
    const endDate = event?.endDate ?? '';
    const startTime = event?.startTime ?? '';
    const endTime = event?.endTime ?? '';

    this.eventForm = this.formBuilder.group({
      title: [title, Validators.required],
      description: [description, Validators.required],
      location: [location, Validators.required],
      cost: [cost, Validators.required],
      startDate: new FormControl(startDate, Validators.required),
      endDate: new FormControl(endDate, Validators.required),
      startTime: [
        startTime,
        [Validators.required, Validators.max(24), Validators.min(0)],
      ],
      endTime: [
        endTime,
        [Validators.required, Validators.max(24), Validators.min(0)],
      ],
    });
  }

  onSubmit() {
    if (this.eventForm.invalid) {
      return;
    }

    console.log('EVENT:' , this.eventForm.value);


    const id = this.route.snapshot.paramMap.get('id');
    const event = this.addOrganizerInfoToEvent(this.eventForm.value);
    // Create new one
    if (id == '0') {
      this.http
        .post<Event>(`${environment.apiUrl}/events`, event)
        .subscribe((event: Event) => {
          console.log(event);
          this.store.dispatch(addEvent({ event }));
        });
      return;
    }

    // Update one
    this.http
      .put<Event>(`${environment.apiUrl}/events`, event)
      .subscribe((event: Event) => {
        console.log(event);
        this.store.dispatch(updateEvent({ event }));
      });
  }

  addOrganizerInfoToEvent(event: Event): Event {
    return {
      ...event,
      organizerEmail: this.user.email,
      organizerName: this.user.name,
    };
  }
}

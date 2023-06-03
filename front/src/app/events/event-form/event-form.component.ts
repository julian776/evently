import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Event } from '../models/event';
import { Store } from '@ngrx/store';
import { EventsState } from 'src/app/reducers/events/events.reducer';
import { addEvent } from 'src/app/reducers/events/events.actions';

@Component({
  selector: 'app-event-form',
  templateUrl: './event-form.component.html',
  styleUrls: ['./event-form.component.scss'],
})
export class EventFormComponent {
  eventForm!: FormGroup;

  constructor(
    private formBuilder: FormBuilder,
    private http: HttpClient,
    private store: Store<EventsState>
  ) {}

  ngOnInit() {
    this.eventForm = this.formBuilder.group({
      title: ['', Validators.required],
      description: ['', Validators.required],
      location: ['', Validators.required],
      cost: ['', Validators.required],
      organizerName: ['', Validators.required],
      organizerEmail: ['', [Validators.required, Validators.email]],
      startTime: ['', Validators.required],
      endTime: ['', Validators.required],
    });
  }

  onSubmit() {
    if (this.eventForm.invalid) {
      return;
    }

    this.http
      .post<Event>(`http://0.0.0.0:8080/events`, this.eventForm.value)
      .subscribe((event: Event) => {
        console.log(event);
        this.store.dispatch(addEvent({ event }));
      });
    console.log(this.eventForm.value);
  }
}

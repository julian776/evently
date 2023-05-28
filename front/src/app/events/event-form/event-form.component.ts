import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-event-form',
  templateUrl: './event-form.component.html',
  styleUrls: ['./event-form.component.scss']
})
export class EventFormComponent {
  eventForm!: FormGroup;

  constructor(private formBuilder: FormBuilder) { }

  ngOnInit() {


    this.eventForm = this.formBuilder.group({
      title: ['', Validators.required],
      description: ['', Validators.required],
      location: ['', Validators.required],
      cost: ['', Validators.required],
      organizerName: ['', Validators.required],
      organizerEmail: ['', [Validators.required, Validators.email]],
      startTime: ['', Validators.required],
      endTime: ['', Validators.required]
    });
  }

  onSubmit() {
    if (this.eventForm.invalid) {
      return;
    }

    // Handle form submission here
    console.log(this.eventForm.value);
  }
}

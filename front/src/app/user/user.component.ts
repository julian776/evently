import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Store } from '@ngrx/store';
import { UserState } from '../reducers/user/user..reducer';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.scss']
})
export class UserComponent {
  userForm!: FormGroup;

  constructor(
    private formBuilder: FormBuilder,
    private http: HttpClient,
    private store: Store<UserState>
  ) {}

  ngOnInit() {
    this.userForm = this.formBuilder.group({
      email: ['', Validators.required],
      password: ['', Validators.required],
      location: ['', Validators.required],
      cost: ['', Validators.required],
      organizerName: ['', Validators.required],
      organizerEmail: ['', [Validators.required, Validators.email]],
      startTime: [new Date(), [Validators.required]],
      endTime: [new Date(), [Validators.required]],
    });
  }
}

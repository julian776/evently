import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Store } from '@ngrx/store';
import { User } from '../models/user';
import { login } from 'src/app/reducers/user/user.actions';
import { Router } from '@angular/router';

@Component({
  selector: 'app-singup',
  templateUrl: './singup.component.html',
  styleUrls: ['./singup.component.scss', '../user.component.scss'],
})
export class SingupComponent {
  singupForm!: FormGroup;

  constructor(
    private formBuilder: FormBuilder,
    private http: HttpClient,
    private store: Store,
    private router: Router
  ) {}

  ngOnInit() {
    this.singupForm = this.formBuilder.group({
      name: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      password: ['', Validators.required],
      purposeOfUse: ['', Validators.required],
    });
  }

  onSubmit() {
    if (this.singupForm.invalid) {
      return;
    }

    this.http
      .post<User>(`http://0.0.0.0:8080/users`, this.singupForm.value, {
        observe: 'response',
      })
      .subscribe((res) => {
        if (res.ok && res.body) {
          this.store.dispatch(login({ user: res.body }));
          this.router.navigate(['events']).catch((err) => err);
        }
      });
  }
}

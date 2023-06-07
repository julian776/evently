import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Store } from '@ngrx/store';
import { login } from 'src/app/reducers/user/user.actions';
import { User } from '../models/user';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss', '../user.component.scss'],
})
export class LoginComponent {
  loginForm!: FormGroup;

  constructor(
    private formBuilder: FormBuilder,
    private http: HttpClient,
    private store: Store,
    private router: Router
  ) {}

  ngOnInit() {
    this.loginForm = this.formBuilder.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', Validators.required],
    });
  }

  /**
   * It sends a request to the server with
   * the user's email and password, and if
   * the response is successful, it dispatches
   * a `login` action with the user object to
   * the store.
   *
   * @method
   * @name onSubmit
   * @kind method
   * @memberof LoginComponent
   * @returns {void}
   */
  onSubmit(): void {
    this.http
      .post<{ message: string; user: User }>(
        `http://0.0.0.0:8080/users/login`,
        this.loginForm.value,
        { observe: 'response' }
      )
      .subscribe((res) => {
        console.log(res.ok);
        if (res.ok && res.body) {
          this.store.dispatch(login({ user: res.body.user }));
          this.router.navigate(['events']).catch(err => err);
        }
      });
  }

  redirectSingup() {
    this.router.navigate(['singup']).catch(console.error);
  }
}

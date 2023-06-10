import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs';
import { State } from './reducers';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  title = 'front';
  isUserLoggedIn$!: Observable<boolean>;

  constructor(
    private router: Router,
     private store: Store<State>,
     public snackBar: MatSnackBar
     ) {}

  ngOnInit() {
    this.isUserLoggedIn$ = this.store.select((state) => state.user.isLoggedIn);
  }

  handleCreateEvent() {
    this.isUserLoggedIn$.subscribe((isLoggedin) => {
      if (!isLoggedin) {
        this.snackBar.open('Please login to create an event', '', {
          duration: 2000,
        });
        this.redirectLogin();
        return;
      }
      this.router.navigate(['event/edit/0']).catch(console.error);
    });
  }

  isEditCurrentRoute() {
    return /edit/.test(this.router.url);
  }

  redirectLogin() {
    this.router.navigate(['login']).catch(console.error);
  }

  redirectHome() {
    this.router.navigate(['home']).catch(console.error);
  }
}
